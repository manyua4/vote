package tools

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

func GetUUID() string {
	id := uuid.New() //默认V4 版本 基于一个随机数的，有的基于时间和硬件地址 如MAC
	fmt.Printf("uuid:%s,version:%s", id.String(), id.Version().String())
	return id.String()
}

var snowNode *snowflake.Node

func GetUid() int64 {
	if snowNode == nil {
		snowNode, _ = snowflake.NewNode(1)
	}
	return snowNode.Generate().Int64()

}
