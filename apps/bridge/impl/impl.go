package impl

import (
	"context"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/grpc"
	"github.com/qiaogy91/ioc/config/log"
	iockafka "github.com/qiaogy91/ioc/default/kafka"
	"github.com/qiaogy91/maudit/apps/bridge"
	"github.com/qiaogy91/maudit/apps/event"
	"github.com/segmentio/kafka-go"
	"log/slog"
)

var _ bridge.Service = &Impl{}

type Impl struct {
	ioc.ObjectImpl
	bridge.UnimplementedRpcServer
	log      *slog.Logger
	w        *kafka.Writer
	r        *kafka.Reader
	s        event.Service
	Topic    string `json:"topic" yaml:"topic"`
	Duration int64  `json:"duration" yaml:"duration"` // 将数据从kafka 读出、而后塞入Mysql 的过程中，如果发生失败，需要等待多长时间再次处理
}

func (i *Impl) Name() string  { return bridge.AppName }
func (i *Impl) Priority() int { return 399 }

func (i *Impl) Init() {
	i.log = log.Sub(bridge.AppName)
	i.w = iockafka.GetClient().Producer(i.Topic)
	i.r = iockafka.GetClient().Consumer(i.Topic, bridge.AppName)
	i.s = event.GetSvc()

	i.grpcRegistry()

	i.log.Info("start consumer")
	go func() {
		if _, err := i.Read(context.Background(), nil); err != nil {
			panic(err)
		}
	}()
}

func (i *Impl) grpcRegistry() {
	s := grpc.Get().Server()
	bridge.RegisterRpcServer(s, i)
}

func init() {
	ioc.Controller().Registry(&Impl{})
}
