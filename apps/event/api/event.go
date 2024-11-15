package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/ioc/utils"
	"github.com/qiaogy91/maudit/apps/event"
)

func (h *Handler) QueryEvents(r *restful.Request, w *restful.Response) {
	req := &event.QueryEventRequest{}
	if err := utils.Decoder.Decode(req, r.Request.URL.Query()); err != nil {
		utils.SendFailed(w, ErrEventQueryParams(err))
		return
	}

	ins, err := h.svc.QueryEvent(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrEventQuery(err))
		return
	}
	utils.SendSuccess(w, ins)
}
