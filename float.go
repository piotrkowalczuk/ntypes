package ntypes

import (
	"database/sql/driver"
	"encoding/json"
	fmt "fmt"
	"strconv"
)

// NewFloat allocates new valid Float.
func NewFloat(f float32) *Float {
	return &Float{Optional: &FloatValue{Value: f}, Valid: true}
}

func (m *Float) isNull() bool {
	return m == nil || !m.Valid || m.Optional == nil
}

// FloatOr returns given Float value if receiver/wrapper is nil or invalid.
func (m *Float) FloatOr(or float32) float32 {
	if m.isNull() {
		return or
	}

	return m.Optional.Value
}

// Value implements the driver Valuer interface.
func (m Float) Value() (driver.Value, error) {
	if m.isNull() {
		return nil, nil
	}
	return m.Optional.Value, nil
}

// Scan implements the Scanner interface.
func (m *Float) Scan(value interface{}) (err error) {
	if value == nil {
		m.Optional, m.Valid = nil, false
		return nil
	}
	m.Valid = true
	m.Optional = &FloatValue{}

	var tmp float64
	switch v := value.(type) {
	case []byte:
		tmp, err = strconv.ParseFloat(string(v), 32)
		m.Optional.Value = float32(tmp)
	case string:
		tmp, err = strconv.ParseFloat(v, 32)
		m.Optional.Value = float32(tmp)
	case float32:
		m.Optional.Value = v
	default:
		err = fmt.Errorf("ntypes: unsuported type (%T) passed to Float.Scan", value)
	}

	return
}

// MarshalJSON implements json.Marshaler interface.
func (m *Float) MarshalJSON() ([]byte, error) {
	if m.isNull() {
		return []byte("null"), nil
	}

	return json.Marshal(m.Optional.Value)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (m *Float) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		m.Valid = false
		return nil
	}

	m.Valid = true
	m.Optional = &FloatValue{}

	if err := json.Unmarshal(data, &m.Optional.Value); err != nil {
		m.Valid = false
		return err
	}

	return nil
}
