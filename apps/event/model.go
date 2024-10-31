package event

func NewEvent() *Event {
	return &Event{}
}
func NewEventSet() *EventSet {
	return &EventSet{
		Total: 0,
		Items: []*Event{},
	}
}

func NewQueryEventRequest() *QueryEventRequest {
	return &QueryEventRequest{
		PageNum:   0,
		PageSize:  0,
		QueryType: 0,
		Keyword:   "",
	}
}
