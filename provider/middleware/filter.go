package middleware

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/maudit/apps/event"
	"github.com/qiaogy91/mcenter/apps/token"
	"log/slog"
	"strings"
	"time"
)

func (i *Impl) Filter() restful.FilterFunction {
	return func(r *restful.Request, w *restful.Response, chain *restful.FilterChain) {
		// 获取访问到的路由（返回值为nil 表示没有匹配到已经注册的路由）
		sr := r.SelectedRoute()
		if sr == nil {
			chain.ProcessFilter(r, w)
			return
		}

		m := NewMeta(sr.Metadata())
		if m.Bool("audit") {
			defer func() {
				addr := strings.Split(r.Request.RemoteAddr, ":")
				location, err := i.s.SearchByStr(addr[0])
				if err != nil {
					location = "Unknown"
				}

				// 获取用户信息
				username := "Anonymous"
				tk := r.Attribute("token")
				if tk != nil {
					username = tk.(*token.Token).Username
				}
				// 构建消息
				msg := &event.Event{
					Time:         time.Now().Unix(),
					User:         username,
					Source:       r.Request.RemoteAddr,
					Location:     location,
					Agent:        r.Request.UserAgent(),
					Service:      i.appName,
					Resource:     m.String("resource"),
					Action:       m.String("action"),
					ResponseCode: int64(w.StatusCode()),
				}

				if _, err := i.c.Write(r.Request.Context(), msg); err != nil {
					i.log.Error("kafka writeMessage err", slog.Any("err", err))
				}
			}()
		}

		chain.ProcessFilter(r, w)
	}
}
