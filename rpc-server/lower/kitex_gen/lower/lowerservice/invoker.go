// Code generated by Kitex v0.6.1. DO NOT EDIT.

package lowerservice

import (
	server "github.com/cloudwego/kitex/server"
	lower "lower/kitex_gen/lower"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler lower.LowerService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
