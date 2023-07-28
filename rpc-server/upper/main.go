package main

import (
	"log"
	"net"
	upper "upper/kitex_gen/upper/upperservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal("注册etcd失败")
	}
	addr, _ := net.ResolveTCPAddr("tcp", ":8890")
	svr := upper.NewServer(new(UpperServiceImpl),
						   server.WithRegistry(r),
						   server.WithServiceAddr(addr),
						   server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
								ServiceName: "upper",
						   }))
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
