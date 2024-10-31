package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/ioc/utils"
	"github.com/qiaogy91/maudit/apps/event"
)

func (h *Handler) QueryEvents(r *restful.Request, w *restful.Response) {
	req := event.NewQueryEventRequest()
	if err := utils.Decoder.Decode(req, r.Request.URL.Query()); err != nil {
		utils.SendFailed(w, err)
		return
	}

	ins, err := h.svc.QueryEvent(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, err)
		return
	}
	utils.SendSuccess(w, ins)
}
