package ntypes

import "encoding/json"

// MarshalJSON implements json.Marshaler interface.
func (m *FloatArray) MarshalJSON() ([]byte, error) {
	if m == nil || !m.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(m.Array)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (m *FloatArray) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		m.Valid = false
		return nil
	}
	if data == nil {
		m.Array, m.Valid = nil, false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Array)
}

// FloatArrayOr returns given slice if receiver is nil or invalid.
func (m *FloatArray) FloatArrayOr(or []float32) []float32 {
	switch {
	case m == nil:
		return or
	case !m.Valid:
		return or
	default:
		return m.Array
	}
}
