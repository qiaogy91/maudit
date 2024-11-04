package impl

import (
	"context"
	"encoding/json"
	"github.com/qiaogy91/maudit/apps/bridge"
	"github.com/qiaogy91/maudit/apps/event"
	"github.com/segmentio/kafka-go"
	"log/slog"
	"time"
)

func (i *Impl) Write(ctx context.Context, event *event.Event) (*bridge.WriteResponse, error) {
	i.log.Debug("向kafka 开始写入")

	// 序列化
	bs, err := json.Marshal(event)
	if err != nil {
		i.log.Error("Marshal Message err", slog.Any("err", err))
		return nil, err
	}

	// 写入kafka
	if err := i.w.WriteMessages(ctx, kafka.Message{Value: bs}); err != nil {
		i.log.Error("bridge writeMessage err", slog.Any("err", err))
		return nil, err
	}
	i.log.Debug("向kafka 写入完毕")
	return nil, err
}

func (i *Impl) Read(ctx context.Context, request *bridge.ReadRequest) (*event.Event, error) {
	for {
		// 从kafka 获取消息
		message, err := i.r.FetchMessage(ctx)
		if err != nil {
			i.log.Error("fetch message err,", slog.Any("err", err))
			time.Sleep(time.Duration(i.Duration) * time.Second)
		}

		// 反序列化为event
		e := event.NewEvent()
		if err := json.Unmarshal(message.Value, e); err != nil {
			i.log.Error("unmarshal message err,", slog.Any("err", err))
			time.Sleep(time.Duration(i.Duration) * time.Second)
		}

		// 写入mysql
		if _, err := i.s.CreateEvent(ctx, e); err != nil {
			i.log.Error("insert message err,", slog.Any("err", err))
			time.Sleep(time.Duration(i.Duration) * time.Second)
		}

		// 提交消息
		if err := i.r.CommitMessages(ctx, message); err != nil {
			i.log.Error("commit message err,", slog.Any("err", err))
			time.Sleep(time.Duration(i.Duration) * time.Second)
		}
		i.log.Debug("转换了一条消息")
	}
}
