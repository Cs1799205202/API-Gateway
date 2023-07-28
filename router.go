// Code generated by hertz generator.

package main

import (
	"context"
	"net/http"

	handler "github.com/Cs1799205202/API-Gateway/biz/handler"
	"github.com/Cs1799205202/API-Gateway/biz/handler/idl_management"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	r.GET("/agw/", func(ctx context.Context, req *app.RequestContext) {
		req.JSON(http.StatusOK, "Welcome to API Gateway!")
	})

	r.GET("/idl/list", idl_management.ListService)

	r.POST("/idl/update/:servicename", idl_management.UpdateService)
	r.DELETE("/idl/delete/:servicename", idl_management.DeleteService)

	// your code ...
}
