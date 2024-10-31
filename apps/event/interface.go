package event

import (
	"context"
	"github.com/qiaogy91/ioc"
)

const AppName = "event"

func Get() Service { return ioc.Controller().Get(AppName).(Service) }

type Service interface {
	ServiceServer
	CreateTable(ctx context.Context) error
}
