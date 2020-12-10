package main

import (
	"qa_go/conf"
	"qa_go/routes"
)

func main() {

	conf.Init()

	r := routes.NewRouter()

	r.Run(":8000")
}
