package producer

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/application"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/ioc/default/ip2region"
	iockafka "github.com/qiaogy91/ioc/default/kafka"
	"github.com/segmentio/kafka-go"
	"log/slog"
)

const AppName = "mauditProducer"

type Impl struct {
	ioc.ObjectImpl
	svc     *kafka.Writer
	log     *slog.Logger
	Topic   string `json:"topic" yaml:"topic"`
	c       *ip2region.Ip2Region
	appName string // 微服务 svc01 加载了这个中间件，那么这个 app 就是 svc01
}

func (i *Impl) Name() string  { return AppName }
func (i *Impl) Priority() int { return 202 }

func (i *Impl) Init() {
	i.svc = iockafka.GetClient().Producer(i.Topic)
	i.log = log.Sub(AppName)
	i.c = ip2region.Get()
	i.appName = application.Get().AppName

	container := gorestful.RootContainer()
	container.Filter(i.Filter())
	i.log.Info("mauditFilter added")
}

func init() {
	ioc.Default().Registry(&Impl{})
}
