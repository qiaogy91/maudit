package comsumer

import (
	"context"
	"encoding/json"
	"github.com/qiaogy91/maudit/apps/event"
	"log/slog"
	"time"
)

func (i *Impl) Syncer(ctx context.Context) {
	i.log.Info("consumer started")
	for {
		message, err := i.c.FetchMessage(ctx)
		if err != nil {
			i.log.Error("fetch message err,", slog.Any("err", err))
			time.Sleep(time.Duration(i.Duration) * time.Second)
			continue
		}
		e := event.NewEvent()
		if err := json.Unmarshal(message.Value, e); err != nil {
			i.log.Error("unmarshal message err,", slog.Any("err", err))
			time.Sleep(time.Duration(i.Duration) * time.Second)
			continue
		}

		if _, err := i.svc.CreateEvent(ctx, e); err != nil {
			i.log.Error("insert message err,", slog.Any("err", err))
			time.Sleep(time.Duration(i.Duration) * time.Second)
			continue
		}

		if err := i.c.CommitMessages(ctx, message); err != nil {
			i.log.Error("commit message err,", slog.Any("err", err))
			time.Sleep(time.Duration(i.Duration) * time.Second)
		}
		i.log.Debug("转换了一条消息")
	}
}
