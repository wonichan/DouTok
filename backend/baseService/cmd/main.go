package main

import (
	"DouTok-example/backend/baseService/internal/config"
	"DouTok-example/backend/baseService/launcher"
)

func main() {
	config.Init()
	launcher := launcher.New()
	launcher.Run()
}
