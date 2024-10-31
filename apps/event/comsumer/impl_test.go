package comsumer_test

import (
	"context"
	"encoding/json"
	"github.com/qiaogy91/ioc"
	_ "github.com/qiaogy91/ioc/config/otlp"
	iockafka "github.com/qiaogy91/ioc/default/kafka"
	_ "github.com/qiaogy91/maudit/apps"
	"github.com/qiaogy91/maudit/apps/event"
	"github.com/segmentio/kafka-go"
	"testing"
	"time"
)

var (
	ctx = context.Background()
	c   *kafka.Writer
)

func init() {
	if err := ioc.ConfigIocObject("/Users/qiaogy/GolandProjects/projects/github/CloudManager/maudit/etc/application.yaml"); err != nil {
		panic(err)
	}
	c = iockafka.GetClient().Producer("maudit")
}

func TestImpl_Syncer(t *testing.T) {
	e := &event.Event{
		Time:         time.Now().Unix(),
		User:         "qiaogy",
		Source:       "10.0.0.1",
		Agent:        "chrome",
		Service:      "service01",
		Resource:     "SLB",
		Action:       "create",
		ResponseCode: 200,
	}
	bs, _ := json.Marshal(e)
	if err := c.WriteMessages(ctx, kafka.Message{Value: bs}); err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 12)
}
