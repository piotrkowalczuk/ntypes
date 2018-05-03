package ntypes

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
)

var (
	// StringArrayScan is used by StringArray Scan method.
	StringArrayScan = func(interface{}) ([]string, error) {
		return nil, errors.New("ntypes: StringArrayScan not set")
	}
	// StringArrayValue is used by StringArray Value method.
	StringArrayValue = func([]string) (driver.Value, error) {
		return nil, errors.New("ntypes: StringArrayValue not set")
	}
	// Float64ArrayScan is used by Float64Array Scan method.
	Float64ArrayScan = func(interface{}) ([]float64, error) {
		return nil, errors.New("ntypes: Float64ArrayScan not set")
	}
	// Float64ArrayValue is used by Float64Array Value method.
	Float64ArrayValue = func([]float64) (driver.Value, error) {
		return nil, errors.New("ntypes: Float64ArrayValue not set")
	}
	// Int64ArrayScan is used by Int64Array Scan method.
	Int64ArrayScan = func(interface{}) ([]int64, error) {
		return nil, errors.New("ntypes: Float64ArrayValue not set")
	}
	// Int64ArrayValue is used by Int64Array Value method.
	Int64ArrayValue = func([]int64) (driver.Value, error) {
		return nil, errors.New("ntypes: Int64ArrayValue not set")
	}
	// BoolArrayScan is used by BoolArray Scan method.
	BoolArrayScan = func(interface{}) ([]bool, error) {
		return nil, errors.New("ntypes: BoolArrayScan not set")
	}
	// BoolArrayValue is used by BoolArray Value method.
	BoolArrayValue = func([]bool) (driver.Value, error) {
		return nil, errors.New("ntypes: BoolArrayValue not set")
	}
	// BytesArrayScan is used by BytesArray Scan method.
	BytesArrayScan = func(interface{}) ([][]byte, error) {
		return nil, errors.New("ntypes: BytesArrayScan not set")
	}
	// ByteArrayValue is used by BytesArray Value method.
	BytesArrayValue = func([][]byte) (driver.Value, error) {
		return nil, errors.New("ntypes: BytesArrayValue not set")
	}
)

// MarshalJSON implements json.Marshaler interface.
func (ba *BytesArray) MarshalJSON() ([]byte, error) {
	if ba == nil || !ba.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ba.BytesArray)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (ba *BytesArray) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		ba.Valid = false
		return nil
	}
	if data == nil {
		ba.BytesArray, ba.Valid = nil, false
		return nil
	}
	ba.Valid = true
	return json.Unmarshal(data, &ba.BytesArray)
}

// Scan implements the Scanner interface.
func (ba *BytesArray) Scan(value interface{}) (err error) {
	if value == nil {
		ba.BytesArray, ba.Valid = nil, false
		return
	}
	ba.Valid = true
	ba.BytesArray, err = BytesArrayScan(value)
	return
}

// Value implements the driver Valuer interface.
func (ba BytesArray) Value() (driver.Value, error) {
	if !ba.Valid {
		return nil, nil
	}

	return BytesArrayValue(ba.BytesArray)
}

// BytesArrayOr returns given slice if receiver is nil or invalid.
func (ba *BytesArray) BytesArrayOr(or [][]byte) [][]byte {
	switch {
	case ba == nil:
		return or
	case !ba.Valid:
		return or
	default:
		return ba.BytesArray
	}
}

// NewString allocates new valid string.
func NewString(s string) *String {
	return &String{Chars: s, Valid: true}
}

// StringOr returns given string value if receiver is nil or invalid.
func (s *String) StringOr(or string) string {
	switch {
	case s == nil:
		return or
	case !s.Valid:
		return or
	default:
		return s.Chars
	}
}

// MarshalJSON implements json.Marshaler interface.
func (s *String) MarshalJSON() ([]byte, error) {
	if s == nil || !s.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(s.Chars)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (s *String) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		s.Valid = false
		return nil
	}
	if data == nil {
		s.Chars, s.Valid = "", false
		return nil
	}

	s.Valid = true

	return json.Unmarshal(data, &s.Chars)
}

// Value implements the driver Valuer interface.
func (s String) Value() (driver.Value, error) {
	if !s.Valid {
		return nil, nil
	}
	return s.Chars, nil
}

// Scan implements the Scanner interface.
func (s *String) Scan(value interface{}) error {
	if value == nil {
		s.Chars, s.Valid = "", false
		return nil
	}
	s.Valid = true

	switch v := value.(type) {
	case []byte:
		s.Chars = string(v)
	case string:
		s.Chars = v
	default:
		return fmt.Errorf("ntypes: unsuported type (%T) passed to String.Scan", value)
	}

	return nil
}

// MarshalJSON implements json.Marshaler interface.
func (sa *StringArray) MarshalJSON() ([]byte, error) {
	if sa == nil || !sa.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(sa.StringArray)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (sa *StringArray) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		sa.Valid = false
		return nil
	}
	if data == nil {
		sa.StringArray, sa.Valid = nil, false
		return nil
	}

	sa.Valid = true

	return json.Unmarshal(data, &sa.StringArray)
}

// Scan implements the Scanner interface.
func (sa *StringArray) Scan(value interface{}) (err error) {
	if value == nil {
		sa.StringArray, sa.Valid = nil, false
		return
	}
	sa.Valid = true
	sa.StringArray, err = StringArrayScan(value)
	return
}

// Value implements the driver Valuer interface.
func (sa StringArray) Value() (driver.Value, error) {
	if !sa.Valid {
		return nil, nil
	}

	return StringArrayValue(sa.StringArray)
}

// StringArrayOr returns given slice if receiver is nil or invalid.
func (sa *StringArray) StringArrayOr(or []string) []string {
	switch {
	case sa == nil:
		return or
	case !sa.Valid:
		return or
	default:
		return sa.StringArray
	}
}

// NewInt64 allocates new valid Int64.
func NewInt64(i int64) *Int64 {
	return &Int64{Int64: i, Valid: true}
}

// Int64Or returns given int64 value if receiver is nil or invalid.
func (i *Int64) Int64Or(or int64) int64 {
	if i == nil {
		return or
	}
	if !i.Valid {
		return or
	}

	return i.Int64
}

// Value implements the driver Valuer interface.
func (i Int64) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return i.Int64, nil
}

// Scan implements the Scanner interface.
func (i *Int64) Scan(value interface{}) (err error) {
	if value == nil {
		i.Int64, i.Valid = 0, false
		return nil
	}
	i.Valid = true

	switch v := value.(type) {
	case []byte:
		i.Int64, err = strconv.ParseInt(string(v), 10, 64)
	case string:
		i.Int64, err = strconv.ParseInt(v, 10, 64)
	case int64:
		i.Int64 = v
	default:
		err = fmt.Errorf("ntypes: unsuported type (%T) passed to Int64.Scan", value)
	}

	return
}

// MarshalJSON implements json.Marshaler interface.
func (i *Int64) MarshalJSON() ([]byte, error) {
	if i == nil || !i.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(i.Int64)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (i *Int64) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		i.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &i.Int64); err != nil {
		return err
	}
	i.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler interface.
func (ia *Int64Array) MarshalJSON() ([]byte, error) {
	if ia == nil || !ia.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ia.Int64Array)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (ia *Int64Array) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		ia.Valid = false
		return nil
	}
	if data == nil {
		ia.Int64Array, ia.Valid = nil, false
		return nil
	}
	ia.Valid = true
	return json.Unmarshal(data, &ia.Int64Array)
}

// Scan implements the Scanner interface.
func (ia *Int64Array) Scan(value interface{}) (err error) {
	if value == nil {
		ia.Int64Array, ia.Valid = nil, false
		return
	}
	ia.Valid = true
	ia.Int64Array, err = Int64ArrayScan(value)
	return
}

// Value implements the driver Valuer interface.
func (ia Int64Array) Value() (driver.Value, error) {
	if !ia.Valid {
		return nil, nil
	}

	return Int64ArrayValue(ia.Int64Array)
}

// Int64ArrayOr returns given slice if receiver is nil or invalid.
func (ia *Int64Array) Int64ArrayOr(or []int64) []int64 {
	switch {
	case ia == nil:
		return or
	case !ia.Valid:
		return or
	default:
		return ia.Int64Array
	}
}

// NewInt32 allocates new valid Int32.
func NewInt32(i int32) *Int32 {
	return &Int32{Int32: i, Valid: true}
}

// Int32Or returns given int32 value if receiver is nil or invalid.
func (i *Int32) Int32Or(or int32) int32 {
	if i == nil {
		return or
	}
	if !i.Valid {
		return or
	}

	return i.Int32
}

// Value implements the driver Valuer interface.
func (i Int32) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return int64(i.Int32), nil
}

// Scan implements the Scanner interface.
func (i *Int32) Scan(value interface{}) (err error) {
	if value == nil {
		i.Int32, i.Valid = 0, false
		return nil
	}
	i.Valid = true

	var tmp int64
	switch v := value.(type) {
	case []byte:
		tmp, err = strconv.ParseInt(string(v), 10, 32)
		i.Int32 = int32(tmp)
	case string:
		tmp, err = strconv.ParseInt(v, 10, 32)
		i.Int32 = int32(tmp)
	case int64:
		i.Int32 = int32(v)
	default:
		err = fmt.Errorf("ntypes: unsuported type (%T) passed to Int32.Scan", value)
	}

	return
}

// MarshalJSON implements json.Marshaler interface.
func (i *Int32) MarshalJSON() ([]byte, error) {
	if i == nil || !i.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(i.Int32)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (i *Int32) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		i.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &i.Int32); err != nil {
		return err
	}
	i.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler interface.
func (ia *Int32Array) MarshalJSON() ([]byte, error) {
	if ia == nil || !ia.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ia.Int32Array)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (ia *Int32Array) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		ia.Valid = false
		return nil
	}
	if data == nil {
		ia.Int32Array, ia.Valid = nil, false
		return nil
	}
	ia.Valid = true
	return json.Unmarshal(data, &ia.Int32Array)
}

// Int32ArrayOr returns given slice if receiver is nil or invalid.
func (ia *Int32Array) Int32ArrayOr(or []int32) []int32 {
	switch {
	case ia == nil:
		return or
	case !ia.Valid:
		return or
	default:
		return ia.Int32Array
	}
}

// NewUint32 allocates new valid Uint32.
func NewUint32(u uint32) *Uint32 {
	return &Uint32{Uint32: u, Valid: true}
}

// Uint32Or returns given uint32 value if receiver is nil or invalid.
func (u *Uint32) Uint32Or(or uint32) uint32 {
	if u == nil {
		return or
	}
	if !u.Valid {
		return or
	}

	return u.Uint32
}

// Value implements the driver Valuer interface.
func (u Uint32) Value() (driver.Value, error) {
	if !u.Valid {
		return nil, nil
	}
	return int64(u.Uint32), nil
}

// Scan implements the Scanner interface.
func (u *Uint32) Scan(value interface{}) (err error) {
	if value == nil {
		u.Uint32, u.Valid = 0, false
		return nil
	}
	u.Valid = true

	var tmp uint64
	switch v := value.(type) {
	case []byte:
		tmp, err = strconv.ParseUint(string(v), 10, 32)
		if err != nil {
			return fmt.Errorf("ntypes: Uint32.Scan failure: %s", err)
		}
		u.Uint32 = uint32(tmp)
	case string:
		tmp, err = strconv.ParseUint(v, 10, 32)
		if err != nil {
			return fmt.Errorf("ntypes: Uint32.Scan failure: %s", err)
		}
		u.Uint32 = uint32(tmp)
	case int64:
		if v > math.MaxUint32 {
			return errors.New("ntypes: Uint32.Scan failure: value out of range")
		}
		u.Uint32 = uint32(v)
	default:
		err = fmt.Errorf("ntypes: unsuported type (%T) passed to Uint32.Scan", value)
	}

	return
}

// MarshalJSON implements json.Marshaler interface.
func (u *Uint32) MarshalJSON() ([]byte, error) {
	if u == nil || !u.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(u.Uint32)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (u *Uint32) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		u.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &u.Uint32); err != nil {
		return err
	}
	u.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler interface.
func (ua *Uint32Array) MarshalJSON() ([]byte, error) {
	if ua == nil || !ua.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ua.Uint32Array)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (ua *Uint32Array) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		ua.Valid = false
		return nil
	}
	if data == nil {
		ua.Uint32Array, ua.Valid = nil, false
		return nil
	}
	ua.Valid = true
	return json.Unmarshal(data, &ua.Uint32Array)
}

// Uint32ArrayOr returns given slice if receiver is nil or invalid.
func (ua *Uint32Array) Uint32ArrayOr(or []uint32) []uint32 {
	switch {
	case ua == nil:
		return or
	case !ua.Valid:
		return or
	default:
		return ua.Uint32Array
	}
}

// NewUint64 allocates new valid Uint64.
func NewUint64(u uint64) *Uint64 {
	return &Uint64{Uint64: u, Valid: true}
}

// Uint64Or returns given uint64 value if receiver is nil or invalid.
func (u *Uint64) Uint64Or(or uint64) uint64 {
	if u == nil {
		return or
	}
	if !u.Valid {
		return or
	}

	return u.Uint64
}

// MarshalJSON implements json.Marshaler interface.
func (u *Uint64) MarshalJSON() ([]byte, error) {
	if u == nil || !u.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(u.Uint64)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (u *Uint64) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		u.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &u.Uint64); err != nil {
		return err
	}
	u.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler interface.
func (ua *Uint64Array) MarshalJSON() ([]byte, error) {
	if ua == nil || !ua.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ua.Uint64Array)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (ua *Uint64Array) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		ua.Valid = false
		return nil
	}
	if data == nil {
		ua.Uint64Array, ua.Valid = nil, false
		return nil
	}
	ua.Valid = true
	return json.Unmarshal(data, &ua.Uint64Array)
}

// Uint64ArrayOr returns given slice if receiver is nil or invalid.
func (ua *Uint64Array) Uint64ArrayOr(or []uint64) []uint64 {
	switch {
	case ua == nil:
		return or
	case !ua.Valid:
		return or
	default:
		return ua.Uint64Array
	}
}

// NewFloat32 allocates new valid Float32.
func NewFloat32(f float32) *Float32 {
	return &Float32{Float32: f, Valid: true}
}

// Float32Or returns given Float32 value if receiver is nil or invalid.
func (f *Float32) Float32Or(or float32) float32 {
	if f == nil {
		return or
	}
	if !f.Valid {
		return or
	}

	return f.Float32
}

// Value implements the driver Valuer interface.
func (f Float32) Value() (driver.Value, error) {
	if !f.Valid {
		return nil, nil
	}
	return f.Float32, nil
}

// Scan implements the Scanner interface.
func (f *Float32) Scan(value interface{}) (err error) {
	if value == nil {
		f.Float32, f.Valid = 0.0, false
		return nil
	}
	f.Valid = true

	var tmp float64
	switch v := value.(type) {
	case []byte:
		tmp, err = strconv.ParseFloat(string(v), 32)
		f.Float32 = float32(tmp)
	case string:
		tmp, err = strconv.ParseFloat(v, 32)
		f.Float32 = float32(tmp)
	case float32:
		f.Float32 = v
	default:
		err = fmt.Errorf("ntypes: unsuported type (%T) passed to Float32.Scan", value)
	}

	return
}

// MarshalJSON implements json.Marshaler interface.
func (f *Float32) MarshalJSON() ([]byte, error) {
	if f == nil || !f.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(f.Float32)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (f *Float32) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		f.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &f.Float32); err != nil {
		return err
	}
	f.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler interface.
func (fa *Float32Array) MarshalJSON() ([]byte, error) {
	if fa == nil || !fa.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(fa.Float32Array)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (fa *Float32Array) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		fa.Valid = false
		return nil
	}
	if data == nil {
		fa.Float32Array, fa.Valid = nil, false
		return nil
	}
	fa.Valid = true
	return json.Unmarshal(data, &fa.Float32Array)
}

// Float32ArrayOr returns given slice if receiver is nil or invalid.
func (fa *Float32Array) Float32ArrayOr(or []float32) []float32 {
	switch {
	case fa == nil:
		return or
	case !fa.Valid:
		return or
	default:
		return fa.Float32Array
	}
}

// NewFloat64 allocates new valid Float64.
func NewFloat64(f float64) *Float64 {
	return &Float64{Float64: f, Valid: true}
}

// Float64Or returns given float64 value if receiver is nil or invalid.
func (f *Float64) Float64Or(or float64) float64 {
	if f == nil {
		return or
	}
	if !f.Valid {
		return or
	}

	return f.Float64
}

// Value implements the driver Valuer interface.
func (f Float64) Value() (driver.Value, error) {
	if !f.Valid {
		return nil, nil
	}
	return f.Float64, nil
}

// Scan implements the Scanner interface.
func (f *Float64) Scan(value interface{}) (err error) {
	if value == nil {
		f.Float64, f.Valid = 0.0, false
		return nil
	}
	f.Valid = true

	switch v := value.(type) {
	case []byte:
		f.Float64, err = strconv.ParseFloat(string(v), 64)
	case string:
		f.Float64, err = strconv.ParseFloat(v, 64)
	case float64:
		f.Float64 = v
	default:
		err = fmt.Errorf("ntypes: unsuported type (%T) passed to Float64.Scan", value)
	}

	return
}

// MarshalJSON implements json.Marshaler interface.
func (f *Float64) MarshalJSON() ([]byte, error) {
	if f == nil || !f.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(f.Float64)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (f *Float64) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		f.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &f.Float64); err != nil {
		return err
	}
	f.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler interface.
func (fa *Float64Array) MarshalJSON() ([]byte, error) {
	if fa == nil || !fa.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(fa.Float64Array)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (fa *Float64Array) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		fa.Valid = false
		return nil
	}
	if data == nil {
		fa.Float64Array, fa.Valid = nil, false
		return nil
	}
	fa.Valid = true
	return json.Unmarshal(data, &fa.Float64Array)
}

// Scan implements the Scanner interface.
func (fa *Float64Array) Scan(value interface{}) (err error) {
	if value == nil {
		fa.Float64Array, fa.Valid = nil, false
		return
	}
	fa.Valid = true
	fa.Float64Array, err = Float64ArrayScan(value)
	return
}

// Value implements the driver Valuer interface.
func (fa Float64Array) Value() (driver.Value, error) {
	if !fa.Valid {
		return nil, nil
	}

	return Float64ArrayValue(fa.Float64Array)
}

// Float64ArrayOr returns given slice if receiver is nil or invalid.
func (fa *Float64Array) Float64ArrayOr(or []float64) []float64 {
	switch {
	case fa == nil:
		return or
	case !fa.Valid:
		return or
	default:
		return fa.Float64Array
	}
}

// True allocate new valid Bool object that holds true.
func True() *Bool {
	return &Bool{Bool: true, Valid: true}
}

// False allocate new valid Bool object that holds false.
func False() *Bool {
	return &Bool{Bool: false, Valid: true}
}

// BoolOr returns given bool value if receiver is nil or invalid.
func (b *Bool) BoolOr(or bool) bool {
	if b == nil {
		return or
	}
	if !b.Valid {
		return or
	}

	return b.Bool
}

// Value implements the driver Valuer interface.
func (b Bool) Value() (driver.Value, error) {
	if !b.Valid {
		return nil, nil
	}
	return b.Bool, nil
}

// Scan implements the Scanner interface.
func (b *Bool) Scan(value interface{}) (err error) {
	if value == nil {
		b.Bool, b.Valid = false, false
		return nil
	}
	b.Valid = true

	switch v := value.(type) {
	case []byte:
		b.Bool, err = strconv.ParseBool(string(v))
	case string:
		b.Bool, err = strconv.ParseBool(v)
	case bool:
		b.Bool = v
	default:
		err = fmt.Errorf("ntypes: unsuported type (%T) passed to Bool.Scan", value)
	}

	return
}

// MarshalJSON implements json.Marshaler interface.
func (b *Bool) MarshalJSON() ([]byte, error) {
	if b == nil || !b.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(b.Bool)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (b *Bool) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		b.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &b.Bool); err != nil {
		return err
	}
	b.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler interface.
func (ba *BoolArray) MarshalJSON() ([]byte, error) {
	if ba == nil || !ba.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ba.BoolArray)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (ba *BoolArray) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		ba.Valid = false
		return nil
	}
	if data == nil {
		ba.BoolArray, ba.Valid = nil, false
		return nil
	}
	ba.Valid = true
	return json.Unmarshal(data, &ba.BoolArray)
}

// Scan implements the Scanner interface.
func (ba *BoolArray) Scan(value interface{}) (err error) {
	if value == nil {
		ba.BoolArray, ba.Valid = nil, false
		return
	}
	ba.Valid = true
	ba.BoolArray, err = BoolArrayScan(value)
	return
}

// Value implements the driver Valuer interface.
func (ba BoolArray) Value() (driver.Value, error) {
	if !ba.Valid {
		return nil, nil
	}

	return BoolArrayValue(ba.BoolArray)
}

// BoolArrayOr returns given slice if receiver is nil or invalid.
func (ba *BoolArray) BoolArrayOr(or []bool) []bool {
	switch {
	case ba == nil:
		return or
	case !ba.Valid:
		return or
	default:
		return ba.BoolArray
	}
}

func isNull(data []byte) bool {
	return bytes.Equal(data, []byte("null"))
}
