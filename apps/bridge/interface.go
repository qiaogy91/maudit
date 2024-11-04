package bridge

import "github.com/qiaogy91/ioc"

const AppName = "bridge"

func GetSvc() Service { return ioc.Controller().Get(AppName).(Service) }

type Service interface {
	RpcServer
}
