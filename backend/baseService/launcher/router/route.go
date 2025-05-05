package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	"DouTok-example/backend/baseService/internal/cache"
	cacheCfg "DouTok-example/backend/baseService/internal/cache/config"
)

func RegisterRouter(h *server.Hertz) {
	// 注册路由
	v1 := h.Group("/v1")
	cacheRoute := v1.Group("/cache")
	{
		cache.Init([]string{"127.0.0.1:9000", "127.0.0.1:9001"}, "test", cacheCfg.WithCacheSize(1000))

		cacheRoute.POST("/store", cache.Store)
		cacheRoute.GET("/get/:group/:key", cache.Get)
	}
}
