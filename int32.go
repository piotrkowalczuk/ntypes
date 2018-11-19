package ntypes

import (
	"database/sql/driver"
	"encoding/json"
	fmt "fmt"
	"strconv"
)

// NewInt32 allocates new valid Int32.
func NewInt32(i int32) *Int32 {
	return &Int32{Optional: &Int32Value{Value: i}, Valid: true}
}

func (m *Int32) isNull() bool {
	return m == nil || m.Optional == nil || !m.Valid
}

// Int32Or returns given int32 value if receiver is nil or invalid.
func (m *Int32) Int32Or(or int32) int32 {
	if m.isNull() {
		return or
	}

	return m.Optional.Value
}

// Value implements the driver Valuer interface.
func (m Int32) Value() (driver.Value, error) {
	if m.isNull() {
		return nil, nil
	}
	return int64(m.Optional.Value), nil
}

// Scan implements the Scanner interface.
func (m *Int32) Scan(value interface{}) (err error) {
	if value == nil {
		m.Optional, m.Valid = nil, false
		return nil
	}
	m.Valid = true
	m.Optional = &Int32Value{}

	var tmp int64
	switch v := value.(type) {
	case []byte:
		tmp, err = strconv.ParseInt(string(v), 10, 32)
		m.Optional.Value = int32(tmp)
	case string:
		tmp, err = strconv.ParseInt(v, 10, 32)
		m.Optional.Value = int32(tmp)
	case int64:
		m.Optional.Value = int32(v)
	default:
		err = fmt.Errorf("ntypes: unsuported type (%T) passed to Int32.Scan", value)
	}

	return
}

// MarshalJSON implements json.Marshaler interface.
func (m *Int32) MarshalJSON() ([]byte, error) {
	if m.isNull() {
		return []byte("null"), nil
	}

	return json.Marshal(m.Optional.Value)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (m *Int32) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		m.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &m.Optional); err != nil {
		return err
	}
	m.Valid = true
	return nil
}
