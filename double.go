package ntypes

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// NewDouble allocates new valid Double.
func NewDouble(d float64) *Double {
	return &Double{Optional: &DoubleValue{Value: d}, Valid: true}
}

func (m *Double) isNull() bool {
	return m == nil || !m.Valid || m.Optional == nil
}

// DoubleOr returns given float64 value if receiver is nil or invalid.
func (m *Double) DoubleOr(or float64) float64 {
	if m.isNull() {
		return or
	}

	return m.Optional.Value
}

// Value implements the driver Valuer interface.
func (m Double) Value() (driver.Value, error) {
	if m.isNull() {
		return nil, nil
	}
	return m.Optional.Value, nil
}

// Scan implements the Scanner interface.
func (m *Double) Scan(value interface{}) (err error) {
	if value == nil {
		m.Optional, m.Valid = nil, false
		return nil
	}
	m.Valid = true
	m.Optional = &DoubleValue{}

	switch v := value.(type) {
	case []byte:
		m.Optional.Value, err = strconv.ParseFloat(string(v), 64)
	case string:
		m.Optional.Value, err = strconv.ParseFloat(v, 64)
	case float64:
		m.Optional.Value = v
	default:
		err = fmt.Errorf("ntypes: unsuported type (%T) passed to Double.Scan", value)
	}

	return
}

// MarshalJSON implements json.Marshaler interface.
func (m *Double) MarshalJSON() ([]byte, error) {
	if m.isNull() {
		return []byte("null"), nil
	}

	return json.Marshal(m.Optional.Value)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (m *Double) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		m.Optional, m.Valid = nil, true
		return nil
	}

	m.Optional, m.Valid = &DoubleValue{}, true

	if err := json.Unmarshal(data, &m.Optional.Value); err != nil {
		m.Valid = false
		return err
	}
	return nil
}
