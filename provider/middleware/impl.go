package middleware

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/application"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/ioc/default/ip2region"
	"github.com/qiaogy91/maudit/apps/bridge"
	"github.com/qiaogy91/maudit/provider/client"
	"log/slog"
)

const AppName = "mauditFilter"

type Impl struct {
	ioc.ObjectImpl
	c       bridge.RpcClient
	log     *slog.Logger
	s       *ip2region.Ip2Region
	appName string // 微服务 svc01 加载了这个中间件，那么这个 app 就是 svc01
}

func (i *Impl) Name() string  { return AppName }
func (i *Impl) Priority() int { return 204 }
func (i *Impl) Init() {
	i.log = log.Sub(AppName)
	i.c = client.GetSvc().MauditClient()
	i.s = ip2region.Get()
	i.appName = application.Get().AppName

	ws := gorestful.RootContainer()
	ws.Filter(i.Filter())
	i.log.Info("MauditFilter added")
}

func init() {
	ioc.Default().Registry(&Impl{})
}
