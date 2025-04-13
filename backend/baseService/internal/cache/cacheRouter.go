package cache

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"

	"DouTok-example/backend/baseService/internal/cache/config"
	"DouTok-example/backend/baseService/internal/cache/constants"
	"DouTok-example/backend/baseService/internal/cache/types"
)

var (
	once sync.Once
)

const (
	selfGroup = "test"
	selfNode  = "127.0.0.1:9001"
)

func init() {
	once.Do(func() {
		NewCacheGroupPool("test", config.WithCacheSize(1000))
		RegisterNode()
	})
}

func RegisterNode() {
	addrs := []string{
		"127.0.0.1:9000",
		"127.0.0.1:9001",
	}
	cacheNodePool := NewCacheNodesPool()
	cacheNodePool.AddNodes(addrs...)
	GetCacheGroup(selfGroup).RegisterNode(cacheNodePool)
	cacheNodePool.startRPCServer()
}

func Get(c context.Context, ctx *app.RequestContext) {
	group := ctx.Param(constants.Group_param)
	key := ctx.Param(constants.Key_param)
	data, ok := GetCacheGroup(group).Get(key)
	if !ok {
		ctx.WriteString(fmt.Sprintf("key:%s not found", key))
		return
	}
	res := &types.Entry{
		Key:   key,
		Value: string(data.Value),
		TTL:   int64(data.LifeTime.Seconds()),
	}
	ctx.JSON(http.StatusOK, res)
}

func Store(c context.Context, ctx *app.RequestContext) {
	var (
		body []byte
		err  error
	)
	if body, err = ctx.Body(); err != nil {
		hlog.Errorf("failed to read body, err:%s", err.Error())
		ctx.WriteString("failed to read body")
		return
	}
	hlog.Infof("body:%s", string(body))

	var data types.Entry
	_ = json.Unmarshal(body, &data)
	GetCacheGroup(data.Group).Set(types.NewKvStore(data.Key, types.ByteView(data.Value), time.Duration(data.TTL)*time.Second))
	ctx.JSON(http.StatusOK, data)
}
