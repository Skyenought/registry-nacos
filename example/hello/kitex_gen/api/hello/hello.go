// Code generated by Kitex v0.5.1. DO NOT EDIT.

package hello

import (
	"context"

	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	api "github.com/kitex-contrib/registry-nacos/example/hello/kitex_gen/api"
)

func serviceInfo() *kitex.ServiceInfo {
	return helloServiceInfo
}

var helloServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Hello"
	handlerType := (*api.Hello)(nil)
	methods := map[string]kitex.MethodInfo{
		"echo": kitex.NewMethodInfo(echoHandler, newHelloEchoArgs, newHelloEchoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.1",
		Extra:           extra,
	}
	return svcInfo
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.HelloEchoArgs)
	realResult := result.(*api.HelloEchoResult)
	success, err := handler.(api.Hello).Echo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloEchoArgs() interface{} {
	return api.NewHelloEchoArgs()
}

func newHelloEchoResult() interface{} {
	return api.NewHelloEchoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Echo(ctx context.Context, req *api.Request) (r *api.Response, err error) {
	var _args api.HelloEchoArgs
	_args.Req = req
	var _result api.HelloEchoResult
	if err = p.c.Call(ctx, "echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
