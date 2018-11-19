package ntypes_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/piotrkowalczuk/ntypes"
)

func TestNewUInt32(t *testing.T) {
	given := uint32(124)
	got := ntypes.NewUInt32(given)
	if !got.Valid {
		t.Error("uint32 should be valid")
	}
	if got.GetOptional().GetValue() != given {
		t.Errorf("wrong uint32, expected %d but got %d", given, got.GetOptional().GetValue())
	}
}

func TestUInt32_Value(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		given := ntypes.UInt32{Optional: &ntypes.UInt32Value{Value: 10}, Valid: true}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if _, ok := val.(int64); !ok {
			t.Fatalf("unexpected output type, expected int64 but got %T", val)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		given := ntypes.UInt32{Optional: &ntypes.UInt32Value{Value: 10}, Valid: false}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if val != nil {
			t.Fatalf("unexpected value, expected nil but got %v", val)
		}
	})
}

func TestUInt32_Scan(t *testing.T) {
	test := func(t *testing.T, given interface{}, expectedValue uint32, expectedValid bool, expectedError error) {
		t.Helper()

		var u ntypes.UInt32
		if err := u.Scan(given); err != nil {
			if expectedError != nil {
				if expectedError.Error() != err.Error() {
					t.Errorf("wrong error, expected %s but got %s", expectedError.Error(), err.Error())
				}
			} else {
				t.Errorf("unexpected error: %s", err.Error())
			}
			return
		}

		if u.Valid != expectedValid {
			t.Errorf("wrong valid property value, got %t but expected %t", u.Valid, expectedValid)
		}

		if u.GetOptional().GetValue() != expectedValue {
			t.Errorf("wrong uint32 property value, got %d but expected %d", u.GetOptional().GetValue(), expectedValue)
		}
	}

	test(t, nil, 0, false, nil)
	test(t, int64(math.MaxUint32+1), 0, false, errors.New("ntypes: Uint32.Scan failure: value out of range"))
	test(t, fmt.Sprint(math.MaxUint32+1), 0, false, errors.New(`ntypes: Uint32.Scan failure: strconv.ParseUint: parsing "4294967296": value out of range`))
	test(t, []byte(fmt.Sprint(math.MaxUint32+1)), 0, false, errors.New(`ntypes: Uint32.Scan failure: strconv.ParseUint: parsing "4294967296": value out of range`))

	dataInt64 := map[uint32]int64{
		100:        100,
		4294967295: 4294967295,
		0:          0,
	}

	for expected, given := range dataInt64 {
		test(t, given, expected, true, nil)
	}

	dataByte := map[uint32][]byte{
		100:        []byte("100"),
		4294967295: []byte("4294967295"),
		0:          []byte("0"),
	}

	for expected, given := range dataByte {
		test(t, given, expected, true, nil)
	}

	dataString := map[uint32]string{
		100:        "100",
		4294967295: "4294967295",
		0:          "0",
	}

	for expected, given := range dataString {
		test(t, given, expected, true, nil)
	}
}

func TestUInt32_Uint32Or(t *testing.T) {
	cases := map[string]struct {
		expected uint32
		or       uint32
		given    *ntypes.UInt32
	}{
		"valid": {
			expected: 1,
			given:    &ntypes.UInt32{Optional: &ntypes.UInt32Value{Value: 1}, Valid: true},
			or:       2,
		},
		"invalid": {
			expected: 2,
			given:    &ntypes.UInt32{Optional: &ntypes.UInt32Value{Value: 1}, Valid: false},
			or:       2,
		},
		"nil": {
			expected: 3,
			or:       3,
		},
	}

	for huint32, c := range cases {
		t.Run(huint32, func(t *testing.T) {
			got := c.given.UInt32Or(c.or)

			if got != c.expected {
				t.Fatalf("wrong output, expected %d but got %d", c.expected, got)
			}
		})
	}
}

func TestUInt32_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		given    *ntypes.UInt32
		expected string
	}{

		"nil": {
			given:    nil,
			expected: "null",
		},
		"zero value": {
			given:    &ntypes.UInt32{},
			expected: "null",
		},
		"valid": {
			given:    &ntypes.UInt32{Valid: true},
			expected: "0",
		},
		"invalid": {
			given:    &ntypes.UInt32{Valid: false},
			expected: "null",
		},
		"non zero valid value": {
			given:    &ntypes.UInt32{Optional: &ntypes.UInt32Value{Value: 123}, Valid: true},
			expected: "123",
		},
		"non zero valid max value": {
			given:    &ntypes.UInt32{Optional: &ntypes.UInt32Value{Value: 4294967295}, Valid: true},
			expected: "4294967295",
		},
		"non zero valid min value": {
			given:    &ntypes.UInt32{Optional: &ntypes.UInt32Value{Value: 0}, Valid: true},
			expected: "0",
		},
		"non zero invalid max value": {
			given:    &ntypes.UInt32{Optional: &ntypes.UInt32Value{Value: 4294967295}, Valid: false},
			expected: "null",
		},
		"non zero invalid min value": {
			given:    &ntypes.UInt32{Optional: &ntypes.UInt32Value{Value: 0}, Valid: false},
			expected: "null",
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			b, err := json.Marshal(c.given)
			if err != nil {
				t.Fatalf("%s: unexpected error: %s", hint, err.Error())
			}
			if string(b) != c.expected {
				t.Errorf("%s: wrong output, expected %s but got %s", hint, c.expected, string(b))
			}
		})
	}
}

func TestUInt32_UnmarshalJSON(t *testing.T) {
	cases := map[string]struct {
		expected ntypes.UInt32
		given    []byte
	}{
		"valid": {
			given:    []byte("1"),
			expected: ntypes.UInt32{Optional: &ntypes.UInt32Value{Value: 1}, Valid: true},
		},
		"max": {
			given:    []byte(strconv.FormatUint(math.MaxUint32, 10)),
			expected: ntypes.UInt32{Optional: &ntypes.UInt32Value{Value: math.MaxUint32}, Valid: true},
		},
		"zero": {
			given:    []byte("0"),
			expected: ntypes.UInt32{Optional: &ntypes.UInt32Value{Value: 0}, Valid: true},
		},
		"null": {
			given:    []byte("null"),
			expected: ntypes.UInt32{Optional: &ntypes.UInt32Value{Value: 0}, Valid: false},
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			var got ntypes.UInt32
			err := json.Unmarshal(c.given, &got)
			if err != nil {
				t.Fatalf("unexpected error for %s: %s", string(c.given), err.Error())
			}
			if !proto.Equal(&c.expected, &got) {
				t.Errorf("values are not equal, expected %v, got %v", c.expected, got)
			}
		})
	}
}
