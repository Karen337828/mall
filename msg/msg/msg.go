// Code generated by goctl. DO NOT EDIT.
// Source: msg.proto

package msg

import (
	"context"

	"mall/msg/proto"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SendRequest  = proto.SendRequest
	SendResponse = proto.SendResponse

	Msg interface {
		SendMsg(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	}

	defaultMsg struct {
		cli zrpc.Client
	}
)

func NewMsg(cli zrpc.Client) Msg {
	return &defaultMsg{
		cli: cli,
	}
}

func (m *defaultMsg) SendMsg(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	client := proto.NewMsgClient(m.cli.Conn())
	return client.SendMsg(ctx, in, opts...)
}
