package main

import (
	"qa_go/conf"
	"qa_go/cron"
	"qa_go/routes"
)

func main() {

	conf.Init()

	cron.StartSchedule()

	r := routes.NewRouter()

	r.Run(":8000")
}
