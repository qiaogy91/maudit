package client

import "github.com/qiaogy91/ioc"

const AppName = "mauditClient"

func GetSvc() *MauditClient { return ioc.Default().Get(AppName).(*MauditClient) }
