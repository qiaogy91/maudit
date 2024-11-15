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
