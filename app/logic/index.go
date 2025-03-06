package logic

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"mymodule/app/model"
	"mymodule/app/tools"
	"net/http"
	"strconv"
	"time"
)

func Index(context *gin.Context) {
	ret := model.GetVotes()
	context.HTML(200, "index.tmpl", gin.H{"vote": ret})

}
func GetVotes(context *gin.Context) {
	ret := model.GetVotes()
	context.JSON(200, tools.ECode{

		Data: ret,
	})

}

// GetVoteInfo godoc
// @Summary      获取投票信息
// @Description  获取投票信息
// @Tags         vote
// @Accept       json
// @Produce      json
// @Param        id   query    int true "vote Id"
// @Success      200  {object}  tools.ECode
// @Router       /vote [get]
func GetVoteInfo(context *gin.Context) {
	var id int64
	idStr := context.Query("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	ret := model.GetVote(id)
	//官方日志包
	//log.Printf("[print]ret:%+v\n",ret)
	//log.Panicf("[panic]ret:%+v\n",ret)
	//log.Fatalf("[fatal]ret:%+v\n",ret)
	//logrus 包
	//logrus.Errorf("[error]ret:%+v\n", ret)
	//自制函数日志包
	tools.Logger.Error("[error]ret:%+v\n", ret)
	if ret.Vote.Id <= 0 {
		context.JSON(http.StatusNotFound, tools.ECode{})
		return
	}
	context.JSON(200, tools.ECode{
		Data: ret,
	})
}

func Dovote(context *gin.Context) {
	userIDStr, _ := context.Cookie("Id")
	voteIdStr, _ := context.GetPostForm("vote_id")
	optStr, _ := context.GetPostFormArray("opt[]")

	userID, _ := strconv.ParseInt(userIDStr, 10, 64)
	voteId, _ := strconv.ParseInt(voteIdStr, 10, 64)

	//前置查询
	old := model.GetVoteHistory(userID, voteId)
	if len(old) > 0 {
		context.JSON(200, tools.ECode{
			Code:    10010,
			Message: "您已投过票",
		})

	}

	opt := make([]int64, 0)
	for _, v := range optStr {
		optId, _ := strconv.ParseInt(v, 10, 64)
		opt = append(opt, optId)
	}
	model.DoVote(userID, voteId, opt)
	context.JSON(200, tools.ECode{
		Message: "投票完成",
	})
}

func CheckXYZ(context *gin.Context) bool {
	//拿到ip+ua
	ip := context.ClientIP()
	ua := context.GetHeader("user-agent")
	fmt.Printf("ip:%s\nua:%s\n", ip, ua)

	//转为MD5
	hash := md5.New()
	hash.Write([]byte(ip + ua))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	flag, _ := model.Rdb.Get(context, "ban-"+hashString).Bool()
	if flag {
		return false
	}

	i, _ := model.Rdb.Get(context, "xyz-"+hashString).Int()
	if i > 5 {
		model.Rdb.SetEx(context, "ban-"+hashString, true, 3*time.Second)
		return false
	}
	model.Rdb.Incr(context, "xyz-"+hashString)
	model.Rdb.Expire(context, "ban-"+hashString, 50*time.Second)
	return true
}

func GetCatcha(context *gin.Context) {
	if !CheckXYZ(context) {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10005,
			Message: "您点击的太快了！",
		})
		return

	}

	captcha, err := tools.CaptchaGenerate()
	if err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10005,
			Message: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, tools.ECode{
		Data: captcha,
	})

}
