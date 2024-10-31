package client

import (
	"encoding/json"
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	iockafka "github.com/qiaogy91/ioc/default/kafka"
	"github.com/qiaogy91/maudit/apps/event"
	"github.com/segmentio/kafka-go"
	"log/slog"
	"time"
)

const AppName = "auditProducer"

type Impl struct {
	ioc.ObjectImpl
	svc   *kafka.Writer
	log   *slog.Logger
	Topic string `json:"topic" yaml:"topic"`
}

func (i *Impl) Name() string  { return AppName }
func (i *Impl) Priority() int { return 0 }

func (i *Impl) Init() {
	i.svc = iockafka.GetClient().Producer(i.Topic)
	i.log = log.Sub(AppName)

	container := gorestful.RootContainer()
	container.Filter(i.Filter())
}

func (i *Impl) Filter() restful.FilterFunction {
	return func(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {

		// 获取访问到的路由（返回值为nil 表示没有匹配到已经注册的路由）
		r := request.SelectedRoute()
		if r == nil {
			chain.ProcessFilter(request, response)
			return
		}

		m := NewMeta(r.Metadata())
		if m.Bool("audit") {
			defer func() {
				msg := &event.Event{
					Time:         time.Now().Unix(),
					User:         m.String("user"),
					Source:       m.String(request.Request.RemoteAddr),
					Agent:        m.String(request.Request.UserAgent()),
					Service:      m.String("service"),
					Resource:     m.String("resource"),
					Action:       m.String("action"),
					ResponseCode: int64(response.StatusCode()),
				}

				bs, _ := json.Marshal(msg)
				if err := i.svc.WriteMessages(request.Request.Context(), kafka.Message{Value: bs}); err != nil {
					i.log.Error("kafka writeMessage err", slog.Any("err", err))
				}
			}()
		}

		chain.ProcessFilter(request, response)
	}
}

func NewMeta(d map[string]any) *Meta { return &Meta{data: d} }

type Meta struct {
	data map[string]any
}

func (m *Meta) Bool(k string) bool {
	v, ok := m.data[k]
	if !ok {
		return false
	}
	return v.(bool)
}

func (m *Meta) String(k string) string {
	v, ok := m.data[k]
	if !ok {
		return ""
	}
	return v.(string)
}

func init() {
	ioc.Controller().Registry(&Impl{})
}
