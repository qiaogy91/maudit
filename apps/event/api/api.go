package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/maudit/apps/event"
	"log/slog"
)

type Handler struct {
	ioc.ObjectImpl
	log *slog.Logger
	svc event.Service
}

func (h *Handler) Name() string  { return event.AppName }
func (h *Handler) Priority() int { return 499 }

func (h *Handler) Init() {
	h.log = log.Sub(event.AppName)
	h.svc = event.Get()

	ws := gorestful.ModuleWebservice(h)
	ws.Route(ws.GET("").
		To(h.QueryEvents).
		Doc("查询审计事件").
		Metadata("audit", true))

	for _, c := range restful.DefaultContainer.RegisteredWebServices() {
		for _, s := range c.Routes() {
			h.log.Debug("RouteInfo", slog.String("Method", s.Method), slog.String("Path", s.Path))
		}
	}

}

func init() {
	ioc.Api().Registry(&Handler{})
}
