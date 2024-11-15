package api

import (
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
	h.svc = event.GetSvc()

	ws := gorestful.ModuleWebservice(h)
	ws.Route(ws.GET("").To(h.QueryEvents).
		Doc("审计列表").
		Metadata("doc", "审计列表"))

}

func init() {
	ioc.Api().Registry(&Handler{})
}
