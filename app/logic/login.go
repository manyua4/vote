package logic

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"mymodule/app/model"
	"mymodule/app/tools"
	"net/http"
	"regexp"
	"time"
)

type User struct {
	Name         string `json:"name" form:"name"`
	Password     string `json:"password" form:"password"`
	CaptchaId    string `json:"captcha_id" form:"captcha_id"`
	CaptchaValue string `json:"captcha_value"form:"captcha_value"`
}

func Getlogin(context *gin.Context) {
	context.HTML(200, "login.html", nil)

}

// DoLogin godoc
// @Summary      执行用户登录
// @Description  执行用户登录
// @Tags         login
// @Accept       json
// @Produce      json
// @Param        name   body      User true	"login User"
// @Success      200  {object}  tools.ECode
// @Router       /login [post]
func DoLogin(context *gin.Context) {
	var user User
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(200, tools.ECode{
			Message: err.Error(), //有风险
		})
		return
	}
	if !tools.CaptchaVerify(tools.CaptchaData{
		CaptchaId: user.CaptchaId,
		Data:      user.CaptchaValue,
	}) {
		context.JSON(200, tools.ECode{
			Code:    10010,
			Message: "验证码校验失败！", //有风险
		})
		return

	}

	ret := model.GetUser(user.Name)
	fmt.Printf("ret:%+v\n", ret) //打印出数据 便于查找错误
	if ret.Id < 1 || ret.Password != encryptV1(user.Password) {
		context.JSON(200, tools.UserErr)
		return
	}
	//context.SetCookie("name", user.Name, 3600, "/", "", true, false)
	//context.SetCookie("Id", fmt.Sprint(ret.Id), 3600, "/", "", true, false)
	_ = model.FlushSession(context)
	_ = model.SetSession(context, user.Name, ret.Id)
	context.JSON(200, tools.ECode{
		Message: "登陆成功",
	})
	return

}

// Logout godoc
// @Summary      执行用户退出
// @Description  执行用户退出
// @Tags         login
// @Accept       json
// @Produce      json
// @Success      200  {object}  tools.ECode
// @Router       /logout [get]
func Logout(context *gin.Context) {
	context.SetCookie("name", "", 3600, "/", "", true, false)
	context.SetCookie("Id", "", 3600, "/", "", true, false)
	context.Redirect(http.StatusFound, "/login")

}

type CUser struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

func CreatUSer(context *gin.Context) {
	var user CUser
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(200, tools.ECode{
			Code:    10001,
			Message: err.Error(), //有风险
		})
		return
	}
	//参数校验
	if user.Name == "" || user.Password == "" || user.Password2 == "" {
		context.JSON(200, tools.ECode{
			Code:    10002,
			Message: "账号或者密码不能为空", //这里有风险
		})
		return
	}

	//校验密码
	if user.Password != user.Password2 {
		context.JSON(200, tools.ECode{
			Code:    10003,
			Message: "两次密码不同！", //这里有风险
		})
		return
	}

	//校验用户是否存在，这种写法非常不安全。有严重的并发风险
	if oldUser := model.GetUser(user.Name); oldUser.Id > 0 {
		context.JSON(200, tools.ECode{
			Code:    10004,
			Message: "用户名已存在", //这里有风险
		})
		return
	}

	//判断位数
	lenName := len(user.Name)
	lenPwd := len(user.Password)
	if lenName < 8 || lenName > 16 || lenPwd < 8 || lenPwd > 16 {
		context.JSON(200, tools.ECode{
			Code:    10005,
			Message: "用户名或者密码要大于等于8，小于等于16！", //这里有风险
		})
		return
	}

	//密码不能是纯数字
	regex := regexp.MustCompile(`^[0-9]+$`)
	if regex.MatchString(user.Password) {
		context.JSON(200, tools.ECode{
			Code:    10006,
			Message: "密码不能为纯数字", //这里有风险
		})
		return
	}
	//开始添加用户
	newUser := model.User{
		Name:        user.Name,
		Password:    encryptV1(user.Password),
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
		Uuid:        tools.GetUUID(),
	}
	if err := model.CreateUser(&newUser); err != nil {
		context.JSON(200, tools.ECode{
			Code:    10007,
			Message: "用户创建失败", //这里有风险
		})
		return
	}

	//返回添加成功
	context.JSON(200, tools.OK)

	return
}

// 最基础的版本
func encrypt(pwd string) string {
	hash := md5.New()
	hash.Write([]byte(pwd))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	fmt.Printf("加密后的密码：%s\n", hashString)
	return hashString
}

// 比较常用的
func encryptV1(pwd string) string {
	newPwd := pwd + "香香编程喵喵喵" //不能随便起，且不能暴露
	hash := md5.New()
	hash.Write([]byte(newPwd))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	fmt.Printf("加密后的密码：%s\n", hashString)

	return hashString
}

// 高级加密
func encryptV2(pwd string) string {
	//基于Blowfish 实现加密。简单快速，但有安全风险
	//golang.org/x/crypto/ 中有大量的加密算法
	newPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("密码加密失败：", err)
		return ""
	}
	newPwdStr := string(newPwd)
	fmt.Printf("加密后的密码：%s\n", newPwdStr)
	return newPwdStr
}
