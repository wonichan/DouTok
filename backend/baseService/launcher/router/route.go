package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	"DouTok-example/backend/baseService/internal/cache"
)

func RegisterRouter(h *server.Hertz) {
	// 注册路由
	v1 := h.Group("/v1")
	{
		v1.POST("/store", cache.Store)
		v1.GET("/get/:group/:key", cache.Get)
	}
}
