package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"qa_go/model"
)

func StartSchedule() {
	c := cron.New()

	// 每30分钟将redis数据同步到mysql
	addCronFunc(c, "@every 30m", func() {
		model.SyncUserLikeRecord()
		model.SyncAnswerLikeCount()
	})

	c.Start()
}

func addCronFunc(c *cron.Cron, sepc string, cmd func()) {
	err := c.AddFunc(sepc, cmd)
	if err != nil {
		panic(fmt.Sprintf("定时任务异常: %v", err))
	}
}
