package client

import (
	"context"
	"fmt"
	"github.com/qiaogy91/ioc"
	iocgrpc "github.com/qiaogy91/ioc/config/grpc"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/maudit/apps/bridge"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

// Credential 为每个RPC 请求提供凭据
type Credential struct{ token string }

func (cr *Credential) RequireTransportSecurity() bool { return false }
func (cr *Credential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"Token": cr.token}, nil
}

type MauditClient struct {
	ioc.ObjectImpl
	log   *slog.Logger
	conn  *grpc.ClientConn
	Addr  string `json:"addr" yaml:"addr"`
	Port  int    `json:"port" yaml:"port"`
	Token string `json:"token" yaml:"token"`
}

func (mc *MauditClient) Name() string  { return AppName }
func (mc *MauditClient) Priority() int { return 203 }

func (mc *MauditClient) Init() {
	mc.log = log.Sub(AppName)
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", mc.Addr, mc.Port),
		mc.Options()...,
	)
	if err != nil {
		panic(err)
	}
	mc.conn = conn
}

func (mc *MauditClient) Options() (opt []grpc.DialOption) {
	opt = append(opt,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // 非TLS
		grpc.WithPerRPCCredentials(&Credential{token: mc.Token}), // 凭据
	)
	if iocgrpc.Get().Otlp {
		ioc.OtlpMustEnabled()
		opt = append(opt, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	}
	return
}

func (mc *MauditClient) MauditClient() bridge.RpcClient {
	return bridge.NewRpcClient(mc.conn)
}

func init() {
	ioc.Default().Registry(&MauditClient{})
}
