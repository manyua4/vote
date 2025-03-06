package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"mymodule/app/logic"
	"mymodule/app/model"
	"mymodule/app/tools"
	"net/http"

	_ "mymodule/docs"
)

func New() {
	r := gin.Default()
	r.LoadHTMLGlob("app/view/*")
	//相关的路径放在这

	r.GET("/redis", func(context *gin.Context) {
		s := model.GetVoteCache(context, 1)
		fmt.Printf("redis:%+v\n", s)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	index := r.Group("")
	//index.Use(chaeckUser)
	{
		//vote
		index.GET("/index", logic.Index)

		//index.POST("/vote", logic.Dovote)

		index.POST("/vote/add", logic.AddVote)
		index.POST("/vote/update", logic.UpdateVote)
		index.POST("/vote/del", logic.DelVote)

		index.GET("/result", logic.ResultInfo)
		index.GET("/result/info", logic.ResultVote)

	}
	//Restful 风格接口
	{
		//读
		index.GET("/votes", logic.GetVotes)   //获取投票列表
		index.GET("/vote", logic.GetVoteInfo) //获取投票详情

		index.POST("/vote", logic.AddVote)
		index.PUT("/vote", logic.UpdateVote)
		index.DELETE("/vote", logic.DelVote)

		index.GET("/vote/result", logic.ResultVote)

		index.POST("/do_vote", logic.Dovote)
	}
	r.GET("/", logic.Index)
	{
		//login
		r.GET("/login", logic.Getlogin)
		r.POST("/login", logic.DoLogin)
		r.GET("/logout", logic.Logout)

		//user
		r.POST("/user/create", logic.CreatUSer)

	}

	//验证码
	r.GET("/captcha", logic.GetCatcha)

	r.POST("/captcha/verify", func(context *gin.Context) {
		var param tools.CaptchaData
		if err := context.ShouldBind(&param); err != nil {
			context.JSON(http.StatusOK, tools.ParamErr)
			return
		}

		fmt.Printf("参数为：%+v", param)
		if !tools.CaptchaVerify(param) {
			context.JSON(http.StatusOK, tools.ECode{
				Code:    10008,
				Message: "验证失败",
			})
			return
		}
		context.JSON(http.StatusOK, tools.OK)
	})
	if err := r.Run(":8080"); err != nil {
		fmt.Println("启动失败")

	}

}

func chaeckUser(context *gin.Context) {
	var name string
	var id int64
	values := model.GetSession(context)

	if v, ok := values["name"]; ok {
		name = v.(string)
	}
	if v, ok := values["id"]; ok {
		id = v.(int64)
	}
	if name == "" || id <= 0 {
		context.JSON(http.StatusUnauthorized, tools.NotLogin)
		context.Abort()
	}
	context.Next()

}
