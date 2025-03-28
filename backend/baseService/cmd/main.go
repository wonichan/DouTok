package main

import (
	"DouTok-example/backend/baseService/internal/conf"
	"DouTok-example/backend/common/launcher"
)

func main() {
	c := &conf.Config{}
	launcher := launcher.New(
		launcher.WithConfig(c),
	)
	launcher.Run()
}
