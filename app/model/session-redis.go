package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rbcervilla/redisstore/v9"
)

// session 存在本机上，然后将session-name 通过cookie传给前端
var store *redisstore.RedisStore
var sessionName = "session-name"

// GetSession 从session中获取值
func GetSession(c *gin.Context) map[interface{}]interface{} {
	session, _ := store.Get(c.Request, sessionName)
	fmt.Printf("session:%+v\n", session.Values)
	return session.Values
}

// SetSession 在session中创建值
func SetSession(c *gin.Context, name string, id int64) error {
	session, _ := store.Get(c.Request, sessionName)
	session.Values["name"] = name
	session.Values["id"] = id
	return session.Save(c.Request, c.Writer)
}

// FlushSession 清楚session中的值
func FlushSession(c *gin.Context) error {
	session, _ := store.Get(c.Request, sessionName)
	fmt.Printf("session : %+v\n", session.Values)
	session.Values["name"] = ""
	session.Values["id"] = int64(0)
	return session.Save(c.Request, c.Writer)
}
