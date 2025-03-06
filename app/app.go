package app

import (
	"mymodule/app/model"
	"mymodule/app/router"
	"mymodule/app/tools"
)

// Start 启动器方法
func Start() {
	model.NewMysql()
	model.NewRdb()
	defer func() {
		model.Close()
	}()
	//schedule.Start()

	tools.NewLoggeer()

	router.New()
}
