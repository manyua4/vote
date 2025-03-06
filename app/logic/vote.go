package logic

import (
	"github.com/gin-gonic/gin"
	"mymodule/app/model"
	"mymodule/app/tools"
	"net/http"
	"strconv"
	"time"
)

func AddVote(context *gin.Context) {
	idStr := context.Query("title")
	optStr, _ := context.GetPostFormArray("opt_name[]")
	//构建结构体
	vote := model.Vote{
		Title:       idStr,
		Type:        0,
		Status:      0,
		CreatedTime: time.Now(),
	}
	if vote.Title == "" {
		context.JSON(http.StatusBadRequest, tools.ParamErr)
		return
	}

	oldVote := model.GetVoteByName(idStr)
	if oldVote.Id > 0 {
		//context.JSON(http.StatusCreated, tools.OK)
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: "投票已存在",
		})
		return
	}

	opt := make([]model.VoteOpt, 0)
	for _, v := range optStr {
		opt = append(opt, model.VoteOpt{
			Name:        v,
			CreatedTime: time.Now(),
		})
	}

	if err := model.AddVote(vote, opt); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, tools.OK)
	return
}

func UpdateVote(context *gin.Context) {

}

// DelVote 删除一个投票
func DelVote(context *gin.Context) {
	var id int64
	idStr := context.Query("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	//实现删除的幂等性
	vote := model.GetVote(id)
	if vote.Vote.Id <= 0 {
		context.JSON(http.StatusNoContent, tools.OK)
		return

	}
	if err := model.DelVote(id); err != true {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: "删除失败",
		})
		return
	}

	context.JSON(http.StatusNoContent, tools.OK)
	return
}

func ResultInfo(context *gin.Context) {
	context.HTML(200, "result.html", nil)

}

// ResultData 新定义返回结构
type ResultData struct {
	Title string
	Count int64
	Opt   []*ResultVoteOpt
}

type ResultVoteOpt struct {
	Name  string
	Count int64
}

// ResultVote 返回一个投票结果
func ResultVote(context *gin.Context) {
	var id int64
	idStr := context.Query("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	ret := model.GetVote(id)
	data := ResultData{
		Title: ret.Vote.Title,
	}

	for _, v := range ret.Opt {
		data.Count = data.Count + v.Count
		tmp := ResultVoteOpt{
			Name:  v.Name,
			Count: v.Count,
		}
		data.Opt = append(data.Opt, &tmp)
	}

	context.JSON(http.StatusOK, tools.ECode{
		Data: data,
	})
}
