package impl

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/qiaogy91/maudit/apps/event"
)

func (i *Impl) CreateTable(ctx context.Context) error {
	return i.db.WithContext(ctx).AutoMigrate(&event.Event{})
}

func (i *Impl) CreateEvent(ctx context.Context, req *event.Event) (*event.Event, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	if err := i.db.WithContext(ctx).Model(&event.Event{}).Create(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}

func (i *Impl) QueryEvent(ctx context.Context, req *event.QueryEventRequest) (*event.EventSet, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	inst := event.NewEventSet()

	offset := int((req.PageNum - 1) * req.PageSize)
	limit := int(req.PageSize)
	sql := i.db.WithContext(ctx).Model(&event.Event{})

	switch req.QueryType {
	case event.QUERY_TYPE_QUERY_TYPE_USERNAME:
		sql = sql.Where("user like ?", "%"+req.Keyword+"%")
	case event.QUERY_TYPE_QUERY_TYPE_SERVICE:
		sql = sql.Where("service = ?", req.Keyword)
	}

	if err := sql.Count(&inst.Total).Offset(offset).Limit(limit).Find(&inst.Items).Error; err != nil {
		return nil, err
	}

	return inst, nil
}
