package ntypes

import (
	"database/sql/driver"
	"encoding/json"
	fmt "fmt"
	"strconv"
)

// True allocate new valid Bool object that holds true.
func True() *Bool {
	return &Bool{Optional: &BoolValue{Value: true}, Valid: true}
}

// False allocate new valid Bool object that holds false.
func False() *Bool {
	return &Bool{Optional: &BoolValue{Value: false}, Valid: true}
}

func (m *Bool) isNull() bool {
	return m == nil || !m.Valid || m.Optional == nil
}

// BoolOr returns given bool value if receiver is nil or invalid.
func (m *Bool) BoolOr(or bool) bool {
	if m.isNull() {
		return or
	}

	return m.Optional.Value
}

// Value implements the driver Valuer interface.
func (m Bool) Value() (driver.Value, error) {
	if m.isNull() {
		return nil, nil
	}
	return m.Optional.Value, nil
}

// Scan implements the Scanner interface.
func (m *Bool) Scan(value interface{}) (err error) {
	if value == nil {
		m.Optional, m.Valid = nil, false
		return nil
	}

	m.Valid = true
	m.Optional = &BoolValue{}

	switch v := value.(type) {
	case []byte:
		m.Optional.Value, err = strconv.ParseBool(string(v))
	case string:
		m.Optional.Value, err = strconv.ParseBool(v)
	case bool:
		m.Optional.Value = v
	default:
		err = fmt.Errorf("ntypes: unsuported type (%T) passed to Bool.Scan", value)
	}

	return
}

// MarshalJSON implements json.Marshaler interface.
func (m *Bool) MarshalJSON() ([]byte, error) {
	switch {
	case m == nil:
		return []byte{}, nil
	case !m.Valid:
		// TODO: Field with field Valid set to false should be skipped, but it is not possible right now.
		// LINK: https://github.com/golang/go/issues/11939
		return []byte{}, nil
	case m.Optional == nil:
		return []byte("null"), nil
	}

	return json.Marshal(m.Optional.Value)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (m *Bool) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		m.Optional, m.Valid = nil, true
		return nil
	}

	m.Optional, m.Valid = &BoolValue{}, true

	if err := json.Unmarshal(data, &m.Optional.Value); err != nil {
		m.Valid = false
		return err
	}
	return nil
}
