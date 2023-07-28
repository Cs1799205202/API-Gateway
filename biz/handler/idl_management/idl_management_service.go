package idl_management

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Cs1799205202/API-Gateway/biz/handler/gateway"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	etcd "github.com/kitex-contrib/registry-etcd"
)

// GET /idl/list/
func ListService(ctx context.Context, c *app.RequestContext) {
	cnt := 0
	gateway.ClientMap.Range(func(key, value interface{}) bool {
		cnt++
		return true
	})
	services := make([]string, cnt)
	gateway.ClientMap.Range(func(key, value interface{}) bool {
		cnt--
		services[cnt] = key.(string)
		return true
	})
	c.JSON(http.StatusOK, services)
}

// POST /idl/update/:servicename
func UpdateService(ctx context.Context, c *app.RequestContext) {
	servicename := c.Param("servicename")
	idlPath := "./idl/"
	etcdResolver, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		hlog.Warn("new etcd resolver failed", err)
	}

	_, err = os.Open(idlPath + servicename + ".thrift")
	if err != nil {
		hlog.Warn("open idl file failed", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("%s.thrift not found, check your idl", servicename))
		return
	}

	provider, err := generic.NewThriftFileProvider(servicename+".thrift", idlPath)
	if err != nil {
		hlog.Warn("new thrift file provider failed", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("%s.thrift not found, check your idl", servicename))
		return
	}

	g, err := generic.HTTPThriftGeneric(provider)
	if err != nil {
		hlog.Fatal("new HTTPThriftGeneric error", err)
	}
	cli, err := genericclient.NewClient(
		servicename,
		g,
		client.WithResolver(etcdResolver),
	)
	if err != nil {
		hlog.Fatal("error creating genericclient")
	}

	_, ok := gateway.ClientMap.Load(servicename)
	gateway.ClientMap.Store(servicename, cli)
	if ok {
		c.String(http.StatusOK, fmt.Sprintf("idl %s.thrift found, idl management platform updated!", servicename))
	} else {
		c.String(http.StatusOK, fmt.Sprintf("idl %s.thrift found, add %s service to API gateway!", servicename, servicename))
	}

}

// DELETE /idl/delete/:servicename
func DeleteService(ctx context.Context, c *app.RequestContext) {
	svcname := c.Param("servicename")
	_, ok := gateway.ClientMap.Load(svcname)
	if ok {
		gateway.ClientMap.Delete(svcname)
		c.String(http.StatusOK, fmt.Sprintf("删除泛化调用客户端成功,请同步删除idl文件夹下的idl文件!"))
	} else {
		c.String(http.StatusBadRequest, fmt.Sprintf("没有此泛化调用客户端!"))
	}
}
