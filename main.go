package main

import (
	_ "github.com/qiaogy91/ioc/config/otlp"
	"github.com/qiaogy91/ioc/server"
	_ "github.com/qiaogy91/maudit/apps"
)

func main() {
	server.Execute()
}
