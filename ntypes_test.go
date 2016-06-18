package ntypes_test

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"math"

	"github.com/golang/protobuf/proto"
	"github.com/piotrkowalczuk/ntypes"
)

func TestInt64_ProtoMessage(t *testing.T) {
	var (
		buf []byte
		err error
		tmp ntypes.Int64
	)
	success := []ntypes.Int64{
		{Int64: 1, Valid: true},
		{Int64: 0, Valid: false},
		{Int64: 13123, Valid: false},
		{Int64: -1241223, Valid: true},
	}

	for _, given := range success {
		buf, err = proto.Marshal(&given)
		if err != nil {
			t.Errorf("marshal returned unexpected error: %s", err.Error())
			continue
		}

		err = proto.Unmarshal(buf, &tmp)
		if err != nil {
			t.Errorf("unmarshal returned unexpected error: %s", err.Error())
			continue
		}

		if tmp.Int64 != given.Int64 {
			t.Errorf("integers are not equal expected %d, got %d", given.Int64, tmp.Int64)
		}

		if tmp.Valid != given.Valid {
			t.Errorf("booleans are not equal expected %t, got %t", given.Valid, tmp.Valid)
		}
	}
}

func TestInt64_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		given    *ntypes.Int64
		expected string
	}{

		"nil": {
			given:    nil,
			expected: "null",
		},
		"zero value": {
			given:    &ntypes.Int64{},
			expected: "null",
		},
		"valid": {
			given:    &ntypes.Int64{Valid: true},
			expected: "0",
		},
		"invalid": {
			given:    &ntypes.Int64{Valid: false},
			expected: "null",
		},
		"non zero valid value": {
			given:    &ntypes.Int64{Int64: 123, Valid: true},
			expected: "123",
		},
		"non zero valid max value": {
			given:    &ntypes.Int64{Int64: math.MaxInt64, Valid: true},
			expected: fmt.Sprintf("%d", math.MaxInt64),
		},
		"non zero valid min value": {
			given:    &ntypes.Int64{Int64: math.MinInt64, Valid: true},
			expected: fmt.Sprintf("%d", math.MinInt64),
		},
		"non zero invalid max value": {
			given:    &ntypes.Int64{Int64: math.MaxInt64, Valid: false},
			expected: "null",
		},
		"non zero invalid min value": {
			given:    &ntypes.Int64{Int64: -9223372036854775808, Valid: false},
			expected: "null",
		},
	}

	for d, c := range cases {
		t.Run(d, func(t *testing.T) {
			b, err := json.Marshal(c.given)
			if err != nil {
				t.Fatalf("%s: unexpected error: %s", d, err.Error())
			}
			if string(b) != c.expected {
				t.Errorf("%s: wrong output, expected %s but got %s", d, c.expected, string(b))
			}
		})
	}

	type within struct {
		ID *ntypes.Int64 `json:"id"`
	}
}

func TestInt32_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		given    *ntypes.Int32
		expected string
	}{

		"nil": {
			given:    nil,
			expected: "null",
		},
		"zero value": {
			given:    &ntypes.Int32{},
			expected: "null",
		},
		"valid": {
			given:    &ntypes.Int32{Valid: true},
			expected: "0",
		},
		"invalid": {
			given:    &ntypes.Int32{Valid: false},
			expected: "null",
		},
		"non zero valid value": {
			given:    &ntypes.Int32{Int32: 123, Valid: true},
			expected: "123",
		},
		"non zero valid max value": {
			given:    &ntypes.Int32{Int32: math.MaxInt32, Valid: true},
			expected: fmt.Sprintf("%d", math.MaxInt32),
		},
		"non zero valid min value": {
			given:    &ntypes.Int32{Int32: math.MinInt32, Valid: true},
			expected: fmt.Sprintf("%d", math.MinInt32),
		},
		"non zero invalid max value": {
			given:    &ntypes.Int32{Int32: math.MaxInt32, Valid: false},
			expected: "null",
		},
		"non zero invalid min value": {
			given:    &ntypes.Int32{Int32: math.MinInt32, Valid: false},
			expected: "null",
		},
	}

	for d, c := range cases {
		t.Run(d, func(t *testing.T) {
			b, err := json.Marshal(c.given)
			if err != nil {
				t.Fatalf("%s: unexpected error: %s", d, err.Error())
			}
			if string(b) != c.expected {
				t.Errorf("%s: wrong output, expected %s but got %s", d, c.expected, string(b))
			}
		})
	}

	type within struct {
		ID *ntypes.Int32 `json:"id"`
	}
}

func TestFloat32_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		given    *ntypes.Float32
		expected string
	}{

		"nil": {
			given:    nil,
			expected: "null",
		},
		"zero value": {
			given:    &ntypes.Float32{},
			expected: "null",
		},
		"valid": {
			given:    &ntypes.Float32{Valid: true},
			expected: "0",
		},
		"invalid": {
			given:    &ntypes.Float32{Valid: false},
			expected: "null",
		},
		"non zero valid value": {
			given:    &ntypes.Float32{Float32: 123, Valid: true},
			expected: "123",
		},
		"non zero valid max value": {
			given:    &ntypes.Float32{Float32: math.MaxFloat32, Valid: true},
			expected: strconv.FormatFloat(float64(math.MaxFloat32), 'e', 7, 32),
		},
		"non zero invalid max value": {
			given:    &ntypes.Float32{Float32: math.MaxFloat32, Valid: false},
			expected: "null",
		},
	}

	for d, c := range cases {
		t.Run(d, func(t *testing.T) {
			b, err := json.Marshal(c.given)
			if err != nil {
				t.Fatalf("%s: unexpected error: %s", d, err.Error())
			}
			if string(b) != c.expected {
				t.Errorf("%s: wrong output, expected %s but got %s", d, c.expected, string(b))
			}
		})
	}

	type within struct {
		ID *ntypes.Float32 `json:"id"`
	}
}

func TestInt_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		given    *ntypes.Int
		expected string
	}{

		"nil": {
			given:    nil,
			expected: "null",
		},
		"zero value": {
			given:    &ntypes.Int{},
			expected: "null",
		},
		"valid": {
			given:    &ntypes.Int{Valid: true},
			expected: "0",
		},
		"invalid": {
			given:    &ntypes.Int{Valid: false},
			expected: "null",
		},
		"non zero valid value": {
			given:    &ntypes.Int{Int: 123, Valid: true},
			expected: "123",
		},
	}

	for d, c := range cases {
		t.Run(d, func(t *testing.T) {
			b, err := json.Marshal(c.given)
			if err != nil {
				t.Fatalf("%s: unexpected error: %s", d, err.Error())
			}
			if string(b) != c.expected {
				t.Errorf("%s: wrong output, expected %s but got %s", d, c.expected, string(b))
			}
		})
	}

	type within struct {
		ID *ntypes.Int `json:"id"`
	}
}

func TestUint32_Scan(t *testing.T) {
	testUint32_Scan_success(t, nil, 0, false)
}

func TestUint32_Scan_string(t *testing.T) {
	success := map[uint32]string{
		100:        "100",
		4294967295: "4294967295",
		0:          "0",
	}

	for expected, given := range success {
		testUint32_Scan_success(t, given, expected, true)
	}
}

func TestUint32_Scan_int64(t *testing.T) {
	success := map[uint32]int64{
		100:        100,
		4294967295: 4294967295,
		0:          0,
	}

	for expected, given := range success {
		testUint32_Scan_success(t, given, expected, true)
	}
}

func TestUint32_Scan_bytes(t *testing.T) {
	success := map[uint32][]byte{
		100:        []byte("100"),
		4294967295: []byte("4294967295"),
		0:          []byte("0"),
	}

	for expected, given := range success {
		testUint32_Scan_success(t, given, expected, true)
	}
}

func testUint32_Scan_success(t *testing.T, given interface{}, expectedValue uint32, expectedValid bool) {
	var u ntypes.Uint32
	if err := u.Scan(given); err != nil {
		t.Errorf("unexpected error: %s", err.Error())
		return
	}

	if u.Valid != expectedValid {
		t.Errorf("wrong valid property value, got %t but expected %t", u.Valid, expectedValid)
	}

	if u.Uint32 != expectedValue {
		t.Errorf("wrong uint32 property value, got %d but expected %d", u.Uint32, expectedValue)
	}
}

func TestUint32_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		given    *ntypes.Uint32
		expected string
	}{

		"nil": {
			given:    nil,
			expected: "null",
		},
		"zero value": {
			given:    &ntypes.Uint32{},
			expected: "null",
		},
		"valid": {
			given:    &ntypes.Uint32{Valid: true},
			expected: "0",
		},
		"invalid": {
			given:    &ntypes.Uint32{Valid: false},
			expected: "null",
		},
		"non zero valid value": {
			given:    &ntypes.Uint32{Uint32: 123, Valid: true},
			expected: "123",
		},
		"non zero valid max value": {
			given:    &ntypes.Uint32{Uint32: 4294967295, Valid: true},
			expected: "4294967295",
		},
		"non zero valid min value": {
			given:    &ntypes.Uint32{Uint32: 0, Valid: true},
			expected: "0",
		},
		"non zero invalid max value": {
			given:    &ntypes.Uint32{Uint32: 4294967295, Valid: false},
			expected: "null",
		},
		"non zero invalid min value": {
			given:    &ntypes.Uint32{Uint32: 0, Valid: false},
			expected: "null",
		},
	}

	for d, c := range cases {
		t.Run(d, func(t *testing.T) {
			b, err := json.Marshal(c.given)
			if err != nil {
				t.Fatalf("%s: unexpected error: %s", d, err.Error())
			}
			if string(b) != c.expected {
				t.Errorf("%s: wrong output, expected %s but got %s", d, c.expected, string(b))
			}
		})
	}

	type within struct {
		ID *ntypes.Uint32 `json:"id"`
	}
}

func TestBool_MarshalJSON(t *testing.T) {
	cases := map[string]*ntypes.Bool{
		"nil":         nil,
		"zero value":  {},
		"valid":       {Valid: true},
		"invalid":     {Valid: false},
		"true true":   {Bool: true, Valid: true},
		"true false":  {Bool: true, Valid: false},
		"false false": {Bool: false, Valid: false},
		"false true":  {Bool: false, Valid: true},
	}

	for d, c := range cases {
		t.Run(d, func(t *testing.T) {
			_, err := json.Marshal(c)
			if err != nil {
				t.Fatalf("simple: %s: unexpected error: %s", d, err.Error())
			}
		})
	}

	type within struct {
		Exists *ntypes.Bool `json:"exists"`
	}

	for d, c := range cases {
		t.Run(d, func(t *testing.T) {
			w := within{Exists: c}
			_, err := json.Marshal(w)
			if err != nil {
				t.Errorf("within: %s: unexpected error: %s", d, err.Error())
			}
		})
	}
}

func TestString_Value(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		given := ntypes.String{String: "something", Valid: true}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if _, ok := val.(string); !ok {
			t.Fatalf("unexpected output type, expected string but got %T", val)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		given := ntypes.String{String: "something", Valid: false}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if val != nil {
			t.Fatalf("unexpected value, expected nil but got %v", val)
		}
	})
}

func TestString_MarshalJSON(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		expected := `"test"`
		given := &ntypes.String{String: "test", Valid: true}
		got, err := given.MarshalJSON()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if string(got) != expected {
			t.Fatalf("wrong output, expected %s but got %s", expected, string(got))
		}
	})
	t.Run("invalid", func(t *testing.T) {
		expected := `null`
		given := &ntypes.String{String: "test", Valid: false}
		got, err := given.MarshalJSON()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if string(got) != expected {
			t.Fatalf("wrong output, expected %s but got %s", expected, string(got))
		}
	})
	t.Run("nil", func(t *testing.T) {
		expected := `null`
		var given *ntypes.String
		got, err := given.MarshalJSON()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if string(got) != expected {
			t.Fatalf("wrong output, expected %s but got %s", expected, string(got))
		}
	})
}

func TestString_Scan(t *testing.T) {
	cases := map[string]struct {
		given    interface{}
		expected string
	}{
		"bytes": {
			given:    []byte("some byte slice"),
			expected: "some byte slice",
		},
		"string": {
			given:    "some string",
			expected: "some string",
		},
		"nil": {
			expected: "",
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			str := &ntypes.String{}
			err := str.Scan(c.given)
			if err != nil {
				t.Fatalf("unexpected error: %s")
			}
			if c.expected != str.String {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, str.String)
			}

		})
	}
}

func TestString_StringOr(t *testing.T) {
	cases := map[string]struct {
		expected string
		or       string
		given    *ntypes.String
	}{
		"valid": {
			expected: "test",
			given:    &ntypes.String{String: "test", Valid: true},
			or:       "alternative",
		},
		"invalid": {
			expected: "alternative",
			given:    &ntypes.String{String: "test", Valid: false},
			or:       "alternative",
		},
		"nil": {
			expected: "alternative",
			or:       "alternative",
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			got := c.given.StringOr(c.or)

			if got != c.expected {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, got)
			}
		})
	}
}

func TestFloat64_Float64Or(t *testing.T) {
	cases := map[string]struct {
		expected float64
		or       float64
		given    *ntypes.Float64
	}{
		"valid": {
			expected: 0.1,
			given:    &ntypes.Float64{Float64: 0.1, Valid: true},
			or:       0.2,
		},
		"invalid": {
			expected: 0.2,
			given:    &ntypes.Float64{Float64: 0.1, Valid: false},
			or:       0.2,
		},
		"nil": {
			expected: 0.3,
			or:       0.3,
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			got := c.given.Float64Or(c.or)

			if got != c.expected {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, got)
			}
		})
	}
}

func TestInt_IntOr(t *testing.T) {
	cases := map[string]struct {
		expected int
		or       int
		given    *ntypes.Int
	}{
		"valid": {
			expected: 1,
			given:    &ntypes.Int{Int: 1, Valid: true},
			or:       2,
		},
		"invalid": {
			expected: 2,
			given:    &ntypes.Int{Int: 1, Valid: false},
			or:       2,
		},
		"nil": {
			expected: 3,
			or:       3,
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			got := c.given.IntOr(c.or)

			if got != c.expected {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, got)
			}
		})
	}
}

func TestUint32_Uint32Or(t *testing.T) {
	cases := map[string]struct {
		expected uint32
		or       uint32
		given    *ntypes.Uint32
	}{
		"valid": {
			expected: 1,
			given:    &ntypes.Uint32{Uint32: 1, Valid: true},
			or:       2,
		},
		"invalid": {
			expected: 2,
			given:    &ntypes.Uint32{Uint32: 1, Valid: false},
			or:       2,
		},
		"nil": {
			expected: 3,
			or:       3,
		},
	}

	for huint32, c := range cases {
		t.Run(huint32, func(t *testing.T) {
			got := c.given.Uint32Or(c.or)

			if got != c.expected {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, got)
			}
		})
	}
}

func TestInt32_Int32Or(t *testing.T) {
	cases := map[string]struct {
		expected int32
		or       int32
		given    *ntypes.Int32
	}{
		"valid": {
			expected: 1,
			given:    &ntypes.Int32{Int32: 1, Valid: true},
			or:       2,
		},
		"invalid": {
			expected: 2,
			given:    &ntypes.Int32{Int32: 1, Valid: false},
			or:       2,
		},
		"nil": {
			expected: 3,
			or:       3,
		},
	}

	for hint32, c := range cases {
		t.Run(hint32, func(t *testing.T) {
			got := c.given.Int32Or(c.or)

			if got != c.expected {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, got)
			}
		})
	}
}

func TestInt64_Int64Or(t *testing.T) {
	cases := map[string]struct {
		expected int64
		or       int64
		given    *ntypes.Int64
	}{
		"valid": {
			expected: 1,
			given:    &ntypes.Int64{Int64: 1, Valid: true},
			or:       2,
		},
		"invalid": {
			expected: 2,
			given:    &ntypes.Int64{Int64: 1, Valid: false},
			or:       2,
		},
		"nil": {
			expected: 3,
			or:       3,
		},
	}

	for hint64, c := range cases {
		t.Run(hint64, func(t *testing.T) {
			got := c.given.Int64Or(c.or)

			if got != c.expected {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, got)
			}
		})
	}
}

func TestFloat32_Float32Or(t *testing.T) {
	cases := map[string]struct {
		expected float32
		or       float32
		given    *ntypes.Float32
	}{
		"valid": {
			expected: 0.1,
			given:    &ntypes.Float32{Float32: 0.1, Valid: true},
			or:       0.2,
		},
		"invalid": {
			expected: 0.2,
			given:    &ntypes.Float32{Float32: 0.1, Valid: false},
			or:       0.2,
		},
		"nil": {
			expected: 0.3,
			or:       0.3,
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			got := c.given.Float32Or(c.or)

			if got != c.expected {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, got)
			}
		})
	}
}

func TestBool_BoolOr(t *testing.T) {
	cases := map[string]struct {
		expected bool
		or       bool
		given    *ntypes.Bool
	}{
		"valid": {
			expected: true,
			given:    &ntypes.Bool{Bool: true, Valid: true},
			or:       false,
		},
		"invalid": {
			expected: false,
			given:    &ntypes.Bool{Bool: true, Valid: false},
			or:       false,
		},
		"nil": {
			expected: true,
			or:       true,
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			got := c.given.BoolOr(c.or)

			if got != c.expected {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, got)
			}
		})
	}
}

func TestInt64_Scan(t *testing.T) {
	cases := map[string]struct {
		given    interface{}
		expected int64
	}{
		"bytes": {
			given:    []byte("15"),
			expected: 15,
		},
		"string": {
			given:    "16",
			expected: 16,
		},
		"int64": {
			given:    int64(18),
			expected: 18,
		},
		"nil": {
			expected: 0,
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			str := &ntypes.Int64{}
			err := str.Scan(c.given)
			if err != nil {
				t.Fatalf("unexpected error: %s")
			}
			if c.expected != str.Int64 {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, str.Int64)
			}

		})
	}
}

func TestFloat64_Scan(t *testing.T) {
	cases := map[string]struct {
		given    interface{}
		expected float64
	}{
		"bytes": {
			given:    []byte("15.15"),
			expected: 15.15,
		},
		"string": {
			given:    "16.16",
			expected: 16.16,
		},
		"flaot64": {
			given:    19.19,
			expected: 19.19,
		},
		"nil": {
			expected: 0.0,
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			str := &ntypes.Float64{}
			err := str.Scan(c.given)
			if err != nil {
				t.Fatalf("unexpected error: %s")
			}
			if c.expected != str.Float64 {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, str.Float64)
			}

		})
	}
}

func TestFloat32_Scan(t *testing.T) {
	cases := map[string]struct {
		given    interface{}
		expected float32
	}{
		"bytes": {
			given:    []byte("15.15"),
			expected: 15.15,
		},
		"string": {
			given:    "16.16",
			expected: 16.16,
		},
		"float32": {
			given:    float32(19.19),
			expected: 19.19,
		},
		"nil": {
			expected: 0.0,
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			str := &ntypes.Float32{}
			err := str.Scan(c.given)
			if err != nil {
				t.Fatalf("unexpected error: %s")
			}
			if c.expected != str.Float32 {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, str.Float32)
			}

		})
	}
}

func TestBool_Scan(t *testing.T) {
	cases := map[string]struct {
		given    interface{}
		expected bool
	}{
		"bytes": {
			given:    []byte("true"),
			expected: true,
		},
		"string": {
			given:    "true",
			expected: true,
		},
		"nil": {
			expected: false,
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			str := &ntypes.Bool{}
			err := str.Scan(c.given)
			if err != nil {
				t.Fatalf("unexpected error: %s")
			}
			if c.expected != str.Bool {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, str.Bool)
			}

		})
	}
}

func TestInt_Scan(t *testing.T) {
	cases := map[string]struct {
		given    interface{}
		expected int
	}{
		"bytes": {
			given:    []byte("31"),
			expected: 31,
		},
		"string": {
			given:    "14",
			expected: 14,
		},
		"int64": {
			given:    int64(18),
			expected: 18,
		},
		"nil": {
			expected: 0,
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			str := &ntypes.Int{}
			err := str.Scan(c.given)
			if err != nil {
				t.Fatalf("unexpected error: %s")
			}
			if c.expected != str.Int {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, str.Int)
			}

		})
	}
}

func TestInt32_Scan(t *testing.T) {
	cases := map[string]struct {
		given    interface{}
		expected int32
	}{
		"bytes": {
			given:    []byte("31"),
			expected: 31,
		},
		"string": {
			given:    "14",
			expected: 14,
		},
		"int64": {
			given:    int64(18),
			expected: 18,
		},
		"nil": {
			expected: 0,
		},
	}

	for hint32, c := range cases {
		t.Run(hint32, func(t *testing.T) {
			str := &ntypes.Int32{}
			err := str.Scan(c.given)
			if err != nil {
				t.Fatalf("unexpected error: %s")
			}
			if c.expected != str.Int32 {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, str.Int32)
			}

		})
	}
}
