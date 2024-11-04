package impl_test

import (
	"context"
	"github.com/qiaogy91/ioc"
	_ "github.com/qiaogy91/ioc/config/otlp"
	_ "github.com/qiaogy91/maudit/apps"
	"github.com/qiaogy91/maudit/apps/bridge"
	"github.com/qiaogy91/maudit/apps/event"
	"testing"
	"time"
)

var (
	ctx = context.Background()
	c   bridge.Service
)

func init() {
	cf := "/Users/qiaogy/GolandProjects/projects/github/CloudManager/maudit/etc/application.yaml"
	if err := ioc.ConfigIocObject(cf); err != nil {
		panic(err)
	}
	c = bridge.GetSvc()
}

func TestImpl_Write(t *testing.T) {
	req := &event.Event{
		Time:         time.Now().Unix(),
		User:         "qiaogy3",
		Source:       "10.0.0.1",
		Location:     "中国",
		Agent:        "测试",
		Service:      "svc01",
		Resource:     "cmdb",
		Action:       "get",
		ResponseCode: 200,
	}
	if _, err := c.Write(ctx, req); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_Read(t *testing.T) {
	inst, err := c.Read(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", inst)
}
