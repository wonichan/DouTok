package launcher

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/logger/accesslog"

	"DouTok-example/backend/baseService/internal/logger"
	"DouTok-example/backend/baseService/launcher/router"
)

type Launcher struct {
	serverHertz *server.Hertz
}

func New() *Launcher {
	launcher := new(Launcher)
	launcher.serverHertz = server.Default(
		server.WithHostPorts("127.0.0.1:8000"),
		server.WithKeepAlive(true),
		server.WithReadTimeout(60*time.Second),
		server.WithWriteTimeout(60*time.Second),
	)
	logger.InitLogger()
	launcher.serverHertz.Engine.OnRun = append(launcher.serverHertz.Engine.OnRun, beginStart)
	return launcher
}

func beginStart(ctx context.Context) error {
	logger.GetLogger().Infof("here we go...")
	return nil
}

func (l *Launcher) Run() {
	// 注册路由
	l.serverHertz.Use(accesslog.New(accesslog.WithFormat("[${time}] ${status} - ${latency} ${method} ${path} ${queryParams}")))
	router.RegisterRouter(l.serverHertz)
	l.serverHertz.Spin()
}
