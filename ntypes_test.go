package ntypes_test

import (
	"encoding/json"
	"testing"

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
			given:    &ntypes.Int64{Int64: 9223372036854775807, Valid: true},
			expected: "9223372036854775807",
		},
		"non zero valid min value": {
			given:    &ntypes.Int64{Int64: -9223372036854775808, Valid: true},
			expected: "-9223372036854775808",
		},
		"non zero invalid max value": {
			given:    &ntypes.Int64{Int64: 9223372036854775807, Valid: false},
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
	t.Run("valid", func(t *testing.T) {
		expected := "test"
		given := &ntypes.String{String: expected, Valid: true}
		got := given.StringOr("alternative")

		if got != expected {
			t.Fatalf("wrong output, expected %s but got %s", expected, got)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		expected := "alternative"
		given := &ntypes.String{String: "test", Valid: false}
		got := given.StringOr(expected)

		if got != expected {
			t.Fatalf("wrong output, expected %s but got %s", expected, got)
		}
	})
	t.Run("nil", func(t *testing.T) {
		expected := "alternative"
		var given *ntypes.String
		got := given.StringOr(expected)

		if got != expected {
			t.Fatalf("wrong output, expected %s but got %s", expected, got)
		}
	})
}
