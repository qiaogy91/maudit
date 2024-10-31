package impl

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/datasource"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/maudit/apps/event"
	"gorm.io/gorm"
	"log/slog"
)

var _ event.Service = &Impl{}

type Impl struct {
	ioc.ObjectImpl
	event.UnimplementedServiceServer
	log *slog.Logger
	db  *gorm.DB
}

func (i *Impl) Name() string  { return event.AppName }
func (i *Impl) Priority() int { return 398 }
func (i *Impl) Init() {
	i.log = log.Sub(event.AppName)
	i.db = datasource.DB()
}

func init() {
	ioc.Controller().Registry(&Impl{})
}
