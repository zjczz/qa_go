package main

import (
	"likezh/conf"
	"likezh/routes"
)

func main(){

	conf.Init()

	r:=routes.NewRouter()

	r.Run(":8000")
}
