package cache

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"

	"DouTok-example/backend/baseService/internal/cache/constants"
	"DouTok-example/backend/baseService/internal/cache/types"
)

var (
	cacheCluster *CacheCluster
	once         sync.Once
)

func init() {
	once.Do(func() {
		cacheCluster = NewCacheCluster()
		RegisterNode()
	})
}

func RegisterNode() {
	var nodes []*CacheNode
	// TODO 从配置文件中读取节点信息
	nodes = append(nodes,
		NewCacheNode("node1", defaultReplicas, 1000, nil),
		NewCacheNode("node2", defaultReplicas, 1000, nil),
		NewCacheNode("node3", defaultReplicas, 1000, nil),
	)
	cacheCluster.AddNodes(nodes...)
}

func Get(c context.Context, ctx *app.RequestContext) {
	key := ctx.Param(constants.Key_param)
	cacheNode := cacheCluster.GetNode(key)

	data, ok := cacheNode.Get(key)
	if !ok {
		hlog.Infof("key:%s not found", key)
		ctx.WriteString(fmt.Sprintf("key[%s] not found", key))
		return
	}
	ctx.JSON(http.StatusOK, &types.Entry{Key: key, Value: data.String()})
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
	cacheNode := cacheCluster.GetNode(data.Key)
	cacheNode.Set(data.Key, types.ByteView(data.Value))
	ctx.JSON(http.StatusOK, data)
}
