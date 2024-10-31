package impl_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/qiaogy91/ioc"
	_ "github.com/qiaogy91/ioc/config/otlp"
	"github.com/qiaogy91/ioc/utils"
	_ "github.com/qiaogy91/maudit/apps"
	"github.com/qiaogy91/maudit/apps/event"
	"testing"
	"time"
)

var (
	ctx = context.Background()
	c   event.Service
)

func init() {
	if err := ioc.ConfigIocObject("/Users/qiaogy/GolandProjects/projects/github/CloudManager/maudit/etc/application.yaml"); err != nil {
		panic(err)
	}
	c = event.Get()
}

func TestImpl_CreateTable(t *testing.T) {
	if err := c.CreateTable(ctx); err != nil {
		t.Error(err)
	}
}

func TestImpl_CreateEvent(t *testing.T) {
	for i := 1; i <= 1000; i++ {
		req := &event.Event{
			Time:         time.Now().Unix(),
			User:         fmt.Sprintf("user_%d", i),
			Source:       fmt.Sprintf("10.0.0.%d", i),
			Agent:        "chrome",
			Service:      "svc01",
			Resource:     "ecs",
			Action:       "create",
			ResponseCode: 200,
		}
		inst, err := c.CreateEvent(ctx, req)
		if err != nil {
			var v *utils.ApiException
			if errors.As(err, &v) {
				t.Fatal(v.Message, v.Code)
			}
		}

		t.Logf("%+v", inst)
	}

}

func TestImpl_QueryEvent(t *testing.T) {
	req := &event.QueryEventRequest{
		PageNum:   4,
		PageSize:  20,
		QueryType: event.QUERY_TYPE_QUERY_TYPE_SERVICE,
		Keyword:   "svc01",
	}

	inst, err := c.QueryEvent(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(inst.Total)
	for _, item := range inst.Items {
		t.Logf("%+v", item)
	}
}
