package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Cs1799205202/API-Gateway/biz/handler/gateway"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGreetings(t *testing.T) {
	h := server.Default()
	h.GET("/agw", gateway.Greeting)
	w := ut.PerformRequest(h.Engine, "GET", "/agw", nil,
		ut.Header{Key: "Connection", Value: "close"})
	resp := w.Result()
	assert.DeepEqual(t, http.StatusOK, resp.StatusCode())
	assert.DeepEqual(t, "Welcome to API Gateway!", string(resp.Body()))
}

func TestListAndUpdate1(t *testing.T) {
	h := server.Default()

	register(h)

	w := ut.PerformRequest(h.Engine, "GET", "/idl/list", nil)
	resp := w.Result()
	assert.DeepEqual(t, http.StatusOK, resp.StatusCode())
	assert.DeepEqual(t, "[]", string(resp.Body()))

	w = ut.PerformRequest(h.Engine, "POST", "/idl/update/lower", nil)
	resp = w.Result()
	fmt.Print(string(resp.Body()))
	assert.DeepEqual(t, http.StatusOK, resp.StatusCode())
	assert.DeepEqual(t, "idl lower.thrift found, add lower service to API gateway!", string(resp.Body()))

	_, ok := gateway.ClientMap.Load("lower")
	assert.Assert(t, ok)

	w = ut.PerformRequest(h.Engine, "POST", "/idl/update/lower", nil)
	resp = w.Result()
	fmt.Print(string(resp.Body()))
	assert.DeepEqual(t, http.StatusOK, resp.StatusCode())
	assert.DeepEqual(t, "idl lower.thrift found, idl management platform updated!", string(resp.Body()))
	_, ok = gateway.ClientMap.Load("lower")
	assert.Assert(t, ok)
}

func TestListAndUpdate2(t *testing.T) {
	h := server.Default()

	register(h)

	gateway.ClientMap.Delete("lower")
	gateway.ClientMap.Delete("upper")
	w := ut.PerformRequest(h.Engine, "GET", "/idl/list", nil)
	resp := w.Result()
	assert.DeepEqual(t, http.StatusOK, resp.StatusCode())
	assert.DeepEqual(t, "[]", string(resp.Body()))

	w = ut.PerformRequest(h.Engine, "POST", "/idl/update/upper", nil)
	resp = w.Result()
	fmt.Print(string(resp.Body()))
	assert.DeepEqual(t, http.StatusOK, resp.StatusCode())
	assert.DeepEqual(t, "idl upper.thrift found, add upper service to API gateway!", string(resp.Body()))

	_, ok := gateway.ClientMap.Load("upper")
	assert.Assert(t, ok)

	w = ut.PerformRequest(h.Engine, "POST", "/idl/update/upper", nil)
	resp = w.Result()
	fmt.Print(string(resp.Body()))
	assert.DeepEqual(t, http.StatusOK, resp.StatusCode())
	assert.DeepEqual(t, "idl upper.thrift found, idl management platform updated!", string(resp.Body()))

	_, ok = gateway.ClientMap.Load("upper")
	assert.Assert(t, ok)
}

func TestDelete(t *testing.T) {
	h := server.Default()

	register(h)
	
	gateway.ClientMap.Delete("lower")
	gateway.ClientMap.Delete("upper")

	w := ut.PerformRequest(h.Engine, "GET", "/idl/list", nil)
	resp := w.Result()
	assert.DeepEqual(t, http.StatusOK, resp.StatusCode())
	assert.DeepEqual(t, "[]", string(resp.Body()))

	w = ut.PerformRequest(h.Engine, "POST", "/idl/update/upper", nil)
	resp = w.Result()
	fmt.Print(string(resp.Body()))
	assert.DeepEqual(t, http.StatusOK, resp.StatusCode())
	assert.DeepEqual(t, "idl upper.thrift found, add upper service to API gateway!", string(resp.Body()))

}