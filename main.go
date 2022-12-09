package main

import (
	"nothing/config"
	"nothing/service"
)

func main() {
	config.Init()
	config.InitDb()
	service.Run()
}
