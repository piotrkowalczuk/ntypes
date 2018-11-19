package ntypes_test

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/piotrkowalczuk/ntypes"
)

func TestNewInt32(t *testing.T) {
	given := int32(124)
	got := ntypes.NewInt32(given)
	if !got.Valid {
		t.Error("int32 should be valid")
	}
	if got.GetOptional().GetValue() != given {
		t.Errorf("wrong int32, expected %d but got %d", given, got.GetOptional().GetValue())
	}
}

func TestInt32_Value(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		given := ntypes.Int32{Optional: &ntypes.Int32Value{Value: 10}, Valid: true}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if _, ok := val.(int64); !ok {
			t.Fatalf("unexpected output type, expected int64 but got %T", val)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		given := ntypes.Int32{Optional: &ntypes.Int32Value{Value: 10}, Valid: false}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if val != nil {
			t.Fatalf("unexpected value, expected nil but got %v", val)
		}
	})
}

func TestInt32_Int32Or(t *testing.T) {
	cases := map[string]struct {
		expected int32
		or       int32
		given    *ntypes.Int32
	}{
		"valid": {
			expected: 1,
			given:    &ntypes.Int32{Optional: &ntypes.Int32Value{Value: 10}, Valid: true},
			or:       2,
		},
		"invalid": {
			expected: 2,
			given:    &ntypes.Int32{Optional: &ntypes.Int32Value{Value: 10}, Valid: false},
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
				t.Fatalf("wrong output, expected %d but got %d", c.expected, got)
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
				t.Fatalf("unexpected error: %s", err.Error())
			}
			if str.GetOptional().GetValue() != c.expected {
				t.Fatalf("wrong output, expected %d but got %d", c.expected, str.GetOptional().GetValue())
			}

		})
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
			given:    &ntypes.Int32{Optional: &ntypes.Int32Value{Value: 123}, Valid: true},
			expected: "123",
		},
		"non zero valid max value": {
			given:    &ntypes.Int32{Optional: &ntypes.Int32Value{Value: math.MaxInt32}, Valid: true},
			expected: fmt.Sprintf("%d", math.MaxInt32),
		},
		"non zero valid min value": {
			given:    &ntypes.Int32{Optional: &ntypes.Int32Value{Value: math.MinInt32}, Valid: true},
			expected: fmt.Sprintf("%d", math.MinInt32),
		},
		"non zero invalid max value": {
			given:    &ntypes.Int32{Optional: &ntypes.Int32Value{Value: math.MaxInt32}, Valid: false},
			expected: "null",
		},
		"non zero invalid min value": {
			given:    &ntypes.Int32{Optional: &ntypes.Int32Value{Value: math.MinInt32}, Valid: false},
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

func TestInt32_UnmarshalJSON(t *testing.T) {
	cases := map[string]struct {
		expected ntypes.Int32
		given    []byte
	}{
		"valid": {
			given:    []byte("1"),
			expected: ntypes.Int32{Optional: &ntypes.Int32Value{Value: 1}, Valid: true},
		},
		"max": {
			given:    []byte(strconv.FormatInt(math.MaxInt32, 10)),
			expected: ntypes.Int32{Optional: &ntypes.Int32Value{Value: math.MaxInt32}, Valid: true},
		},
		"min": {
			given:    []byte(strconv.FormatInt(math.MinInt32, 10)),
			expected: ntypes.Int32{Optional: &ntypes.Int32Value{Value: math.MinInt32}, Valid: true},
		},
		"zero": {
			given:    []byte("0"),
			expected: ntypes.Int32{Optional: &ntypes.Int32Value{Value: 0}, Valid: true},
		},
		"invalid-null": {
			given:    []byte("null"),
			expected: ntypes.Int32{Optional: &ntypes.Int32Value{Value: 1}, Valid: false},
		},
		"valid-null": {
			given:    []byte("null"),
			expected: ntypes.Int32{Optional: nil, Valid: true},
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			var got ntypes.Int32
			err := json.Unmarshal(c.given, &got)
			if err != nil {
				t.Fatalf("marshal returned unexpected error: %s", err.Error())
			}
			if !proto.Equal(&c.expected, &got) {
				t.Errorf("values are not equal expected %v, got %v", c.expected, got)
			}
		})
	}
}
