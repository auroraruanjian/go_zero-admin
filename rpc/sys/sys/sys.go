// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package sys

import (
	"context"

	"go-zero-demo/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	InfoReq   = sysclient.InfoReq
	InfoResp  = sysclient.InfoResp
	LoginReq  = sysclient.LoginReq
	LoginResp = sysclient.LoginResp

	Sys interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error)
	}

	defaultSys struct {
		cli zrpc.Client
	}
)

func NewSys(cli zrpc.Client) Sys {
	return &defaultSys{
		cli: cli,
	}
}

func (m *defaultSys) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultSys) UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}
