package ntypes

import (
	"encoding/json"
)

// NewUInt64 allocates new valid UInt64.
func NewUInt64(u uint64) *UInt64 {
	return &UInt64{Optional: &UInt64Value{Value: u}, Valid: true}
}

func (m *UInt64) isNull() bool {
	return m == nil || !m.Valid || m.Optional == nil
}

// UInt64Or returns given uint64 value if receiver is nil or invalid.
func (m *UInt64) UInt64Or(or uint64) uint64 {
	if m.isNull() {
		return or
	}

	return m.Optional.Value
}

// MarshalJSON implements json.Marshaler interface.
func (m *UInt64) MarshalJSON() ([]byte, error) {
	if m.isNull() {
		return []byte("null"), nil
	}

	return json.Marshal(m.Optional.Value)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (m *UInt64) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		m.Valid = false
		return nil
	}

	m.Valid = true
	m.Optional = &UInt64Value{}

	if err := json.Unmarshal(data, &m.Optional.Value); err != nil {
		m.Valid = false
		return err
	}
	return nil
}
