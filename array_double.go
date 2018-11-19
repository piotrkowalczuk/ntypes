package ntypes

import (
	"database/sql/driver"
	"encoding/json"
)

// MarshalJSON implements json.Marshaler interface.
func (fa *DoubleArray) MarshalJSON() ([]byte, error) {
	if fa == nil || !fa.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(fa.Array)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (fa *DoubleArray) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		fa.Valid = false
		return nil
	}
	if data == nil {
		fa.Array, fa.Valid = nil, false
		return nil
	}
	fa.Valid = true
	return json.Unmarshal(data, &fa.Array)
}

// Scan implements the Scanner interface.
func (fa *DoubleArray) Scan(value interface{}) (err error) {
	if value == nil {
		fa.Array, fa.Valid = nil, false
		return
	}
	fa.Valid = true
	fa.Array, err = DoubleArrayScan(value)
	return
}

// Value implements the driver Valuer interface.
func (fa DoubleArray) Value() (driver.Value, error) {
	if !fa.Valid {
		return nil, nil
	}

	return DoubleArrayValue(fa.Array)
}

// DoubleArrayOr returns given slice if receiver is nil or invalid.
func (fa *DoubleArray) DoubleArrayOr(or []float64) []float64 {
	switch {
	case fa == nil:
		return or
	case !fa.Valid:
		return or
	default:
		return fa.Array
	}
}
