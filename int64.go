package ntypes

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// NewInt64 allocates new valid Int64.
func NewInt64(i int64) *Int64 {
	return &Int64{
		Optional: &Int64Value{
			Value: i,
		},
		Valid: true,
	}
}

func (m *Int64) isNull() bool {
	return m == nil || !m.Valid || m.Optional == nil
}

// Int64Or returns given int64 value if receiver is nil or invalid.
func (m *Int64) Int64Or(or int64) int64 {
	if m.isNull() {
		return or
	}

	return m.Optional.Value
}

// Value implements the driver Valuer interface.
func (m Int64) Value() (driver.Value, error) {
	if m.isNull() {
		return nil, nil
	}
	return m.Optional.Value, nil
}

// Scan implements the Scanner interface.
func (m *Int64) Scan(value interface{}) (err error) {
	if value == nil {
		m.Optional, m.Valid = nil, false
		return nil
	}

	m.Valid = true
	m.Optional = &Int64Value{}

	switch v := value.(type) {
	case []byte:
		m.Optional.Value, err = strconv.ParseInt(string(v), 10, 64)
	case string:
		m.Optional.Value, err = strconv.ParseInt(v, 10, 64)
	case int64:
		m.Optional.Value = v
	default:
		err = fmt.Errorf("ntypes: unsuported type (%T) passed to Int64.Scan", value)
	}

	return
}

// MarshalJSON implements json.Marshaler interface.
func (m *Int64) MarshalJSON() ([]byte, error) {
	if m.isNull() {
		return []byte("null"), nil
	}

	return json.Marshal(m.Optional.Value)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (m *Int64) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		m.Valid = false
		return nil
	}

	m.Optional = &Int64Value{}
	m.Valid = true

	if err := json.Unmarshal(data, &m.Optional.Value); err != nil {
		m.Valid = false
		return err
	}
	return nil
}
