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
	cacheGroup *CacheGroup
	once       sync.Once
)

const (
	selfNode = "node1"
)

func init() {
	once.Do(func() {
		cacheGroup = NewCacheGroup("test", config.WithCacheSize(1000))
		RegisterNode()
	})
}

func RegisterNode() {
	addrs := []string{
		"node1",
		"node2",
		"node3",
	}
	cacheGroup.RegisterNode(NewCacheNodesPool().AddNodes(addrs...))
}

func Get(c context.Context, ctx *app.RequestContext) {
	key := ctx.Param(constants.Key_param)
	data, ok := cacheGroup.Get(key)
	if !ok {
		ctx.WriteString(fmt.Sprintf("key:%s not found", key))
		return
	}
	ctx.JSON(http.StatusOK, data)
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
	cacheGroup.Set(types.NewKvStore(data.Key, types.ByteView(data.Value), time.Duration(data.TTL)*time.Second))
	ctx.JSON(http.StatusOK, data)
}
