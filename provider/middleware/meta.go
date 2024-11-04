package middleware

func NewMeta(d map[string]any) *Meta { return &Meta{data: d} }

type Meta struct {
	data map[string]any
}

func (m *Meta) Bool(k string) bool {
	v, ok := m.data[k]
	if !ok {
		return false
	}
	return v.(bool)
}

func (m *Meta) String(k string) string {
	v, ok := m.data[k]
	if !ok {
		return ""
	}
	return v.(string)
}
