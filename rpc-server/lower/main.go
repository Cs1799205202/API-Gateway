package main

import (
	"context"
	"log"
	"net"
	"strings"

	"lower/kitex_gen/lower"
	"lower/kitex_gen/lower/lowerservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type LowerImpl struct{}

func (l *LowerImpl) Tolower(ctx context.Context, req *lower.NormalRequest) (r *lower.LowerResponse, err error) {
	r = &lower.LowerResponse{
		Result_: strings.ToLower(req.Message),
	}
	return
}

func main() {
	// mainly from https://github.com/kitex-contrib/registry-etcd/blob/main/example/server/main.go
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"});
	if err != nil {
		log.Fatal("注册etcd失败")
	}
	addr, _ := net.ResolveTCPAddr("tcp", ":8889")
	server := lowerservice.NewServer(new(LowerImpl),
									 server.WithRegistry(r),
									 server.WithServiceAddr(addr),
									 server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
										ServiceName: "lower",
									 }))
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}