package ntypes

import (
	"database/sql/driver"
	"encoding/json"
	fmt "fmt"
)

// NewString allocates new valid string.
func NewString(s string) *String {
	return &String{Optional: &StringValue{Value: s}, Valid: true}
}

func (m *String) isNull() bool {
	return m == nil || m.Optional == nil || !m.Valid
}

// StringOr returns given string value if receiver is nil or invalid.
func (m *String) StringOr(or string) string {
	if m.isNull() {
		return or
	}
	return m.Optional.Value
}

// MarshalJSON implements json Marshaler interface.
func (m *String) MarshalJSON() ([]byte, error) {
	if m.isNull() {
		return []byte("null"), nil
	}

	return json.Marshal(m.Optional)
}

// UnmarshalJSON implements json Unmarshaler interface.
func (m *String) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		m.Valid = false
		return nil
	}
	if data == nil {
		m.Optional, m.Valid = nil, false
		return nil
	}

	m.Valid = true
	m.Optional = &StringValue{}

	return json.Unmarshal(data, &m.Optional.Value)
}

// Value implements the driver Valuer interface.
func (m String) Value() (driver.Value, error) {
	if m.isNull() {
		return nil, nil
	}
	return m.Optional.Value, nil
}

// Scan implements the Scanner interface.
func (m *String) Scan(value interface{}) error {
	if value == nil {
		m.Optional, m.Valid = nil, false
		return nil
	}
	m.Valid = true

	switch v := value.(type) {
	case []byte:
		m.Optional = &StringValue{Value: string(v)}
	case string:
		m.Optional = &StringValue{Value: v}
	default:
		return fmt.Errorf("ntypes: unsuported type (%T) passed to String.Scan", value)
	}

	return nil
}
