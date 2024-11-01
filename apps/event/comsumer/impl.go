package comsumer

import (
	"context"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/log"
	iockafka "github.com/qiaogy91/ioc/default/kafka"
	"github.com/qiaogy91/maudit/apps/event"
	"github.com/segmentio/kafka-go"
	"log/slog"
)

const AppName = "mauditConsumer"

type Impl struct {
	ioc.ObjectImpl
	log      *slog.Logger
	svc      event.Service
	c        *kafka.Reader
	Topic    string `json:"topic" yaml:"topic"`
	Duration int64  `json:"duration" yaml:"duration"`
}

func (i *Impl) Name() string  { return AppName }
func (i *Impl) Priority() int { return 399 }
func (i *Impl) Init() {
	i.log = log.Sub(AppName)
	i.svc = event.Get()
	i.c = iockafka.GetClient().Consumer(i.Topic, AppName)

	go func() {
		i.Syncer(context.Background())
	}()
}

func init() {
	ioc.Controller().Registry(&Impl{})
}
