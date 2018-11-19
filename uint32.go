package ntypes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
)

// NewUInt32 allocates new valid UInt32.
func NewUInt32(u uint32) *UInt32 {
	return &UInt32{Optional: &UInt32Value{Value: u}, Valid: true}
}

// UInt32Or returns given uint32 value if receiver is nil or invalid.
func (m *UInt32) UInt32Or(or uint32) uint32 {
	if m == nil {
		return or
	}
	if !m.Valid {
		return or
	}

	return m.Optional.Value
}

// Value implements the driver Valuer interface.
func (m UInt32) Value() (driver.Value, error) {
	if !m.Valid {
		return nil, nil
	}
	return int64(m.Optional.Value), nil
}

// Scan implements the Scanner interface.
func (m *UInt32) Scan(value interface{}) (err error) {
	if value == nil {
		m.Optional, m.Valid = nil, false
		return nil
	}

	m.Valid = true
	m.Optional = &UInt32Value{}

	var tmp uint64
	switch v := value.(type) {
	case []byte:
		tmp, err = strconv.ParseUint(string(v), 10, 32)
		if err != nil {
			return fmt.Errorf("ntypes: UInt32.Scan failure: %s", err)
		}
		m.Optional.Value = uint32(tmp)
	case string:
		tmp, err = strconv.ParseUint(v, 10, 32)
		if err != nil {
			return fmt.Errorf("ntypes: UInt32.Scan failure: %s", err)
		}
		m.Optional.Value = uint32(tmp)
	case int64:
		if v > math.MaxUint32 {
			return errors.New("ntypes: UInt32.Scan failure: value out of range")
		}
		m.Optional.Value = uint32(v)
	default:
		err = fmt.Errorf("ntypes: unsuported type (%T) passed to UInt32.Scan", value)
	}

	return
}

// MarshalJSON implements json.Marshaler interface.
func (m *UInt32) MarshalJSON() ([]byte, error) {
	if m == nil || !m.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(m.Optional)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (m *UInt32) UnmarshalJSON(data []byte) error {
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
