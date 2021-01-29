package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"qa_go/model"
)

func StartSchedule() {
	c := cron.New()

	addCronFunc(c, "@every 1m", func() {
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
