package schedule

import (
	"fmt"
	"mymodule/app/model"
	"time"
)

func Start() {
	go EndVote()
}

func EndVote() {
	t := time.NewTicker(1 * time.Second)
	defer func() {
		t.Stop()
	}()
	for {
		select {
		case <-t.C:
			fmt.Println("EndVote 启动")
			//执行函数
			model.EndVote()
			fmt.Println("EndVote 运行完毕")
		}
	}
}
