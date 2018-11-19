package ntypes

import (
	"database/sql/driver"
	"encoding/json"
)

// MarshalJSON implements json.Marshaler interface.
func (m *Int64Array) MarshalJSON() ([]byte, error) {
	if m == nil || !m.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(m.Array)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (m *Int64Array) UnmarshalJSON(data []byte) error {
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

// Scan implements the Scanner interface.
func (m *Int64Array) Scan(value interface{}) (err error) {
	if value == nil {
		m.Array, m.Valid = nil, false
		return
	}
	m.Valid = true
	m.Array, err = Int64ArrayScan(value)
	return
}

// Value implements the driver Valuer interface.
func (m Int64Array) Value() (driver.Value, error) {
	if !m.Valid {
		return nil, nil
	}

	return Int64ArrayValue(m.Array)
}

// Int64ArrayOr returns given slice if receiver is nil or invalid.
func (m *Int64Array) Int64ArrayOr(or []int64) []int64 {
	switch {
	case m == nil:
		return or
	case !m.Valid:
		return or
	default:
		return m.Array
	}
}
