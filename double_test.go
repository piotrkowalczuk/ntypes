package ntypes_test

import (
	"encoding/json"
	"math"
	"strconv"
	"testing"

	"github.com/piotrkowalczuk/ntypes"

	"github.com/golang/protobuf/proto"
)

func TestNewDouble(t *testing.T) {
	given := float64(124.124)
	got := ntypes.NewDouble(given)
	if !got.Valid {
		t.Error("float64 should be valid")
	}
	if got.GetOptional().GetValue() != given {
		t.Errorf("wrong float64, expected %f but got %f", given, got.GetOptional().GetValue())
	}
}

func TestDouble_Value(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		given := ntypes.Double{Optional: &ntypes.DoubleValue{Value: 10}, Valid: true}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if _, ok := val.(float64); !ok {
			t.Fatalf("unexpected output type, expected Float64 but got %T", val)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		given := ntypes.Double{Optional: &ntypes.DoubleValue{Value: 10.12}, Valid: false}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if val != nil {
			t.Fatalf("unexpected value, expected nil but got %v", val)
		}
	})
}

func TestDouble_Scan(t *testing.T) {
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
			str := &ntypes.Double{}
			err := str.Scan(c.given)
			if err != nil {
				t.Fatalf("unexpected error: %s", err.Error())
			}
			if c.expected != str.GetOptional().GetValue() {
				t.Fatalf("wrong output, expected %f but got %f", c.expected, str.GetOptional().GetValue())
			}

		})
	}
}

func TestDouble_Float64Or(t *testing.T) {
	cases := map[string]struct {
		expected float64
		or       float64
		given    *ntypes.Double
	}{
		"valid": {
			expected: 0.1,
			given:    &ntypes.Double{Optional: &ntypes.DoubleValue{Value: 0.1}, Valid: true},
			or:       0.2,
		},
		"invalid": {
			expected: 0.2,
			given:    &ntypes.Double{Optional: &ntypes.DoubleValue{Value: 0.1}, Valid: false},
			or:       0.2,
		},
		"nil": {
			expected: 0.3,
			or:       0.3,
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			got := c.given.DoubleOr(c.or)

			if got != c.expected {
				t.Fatalf("wrong output, expected %f but got %f", c.expected, got)
			}
		})
	}
}

func TestDouble_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		given    *ntypes.Double
		expected string
	}{

		"nil": {
			given:    nil,
			expected: "null",
		},
		"zero value": {
			given:    &ntypes.Double{},
			expected: "null",
		},
		"valid": {
			given:    &ntypes.Double{Valid: true},
			expected: "0",
		},
		"invalid": {
			given:    &ntypes.Double{Valid: false},
			expected: "null",
		},
		"non zero valid value": {
			given:    &ntypes.Double{Optional: &ntypes.DoubleValue{Value: 123}, Valid: true},
			expected: "123",
		},
		"non zero valid max value": {
			given:    &ntypes.Double{Optional: &ntypes.DoubleValue{Value: math.MaxFloat64}, Valid: true},
			expected: strconv.FormatFloat(float64(math.MaxFloat64), 'e', 16, 64),
		},
		"non zero invalid max value": {
			given:    &ntypes.Double{Optional: &ntypes.DoubleValue{Value: math.MaxFloat64}, Valid: false},
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

	type within struct {
		ID *ntypes.Double `json:"id"`
	}
}

func TestDouble_UnmarshalJSON(t *testing.T) {
	cases := map[string]struct {
		expected ntypes.Double
		given    []byte
	}{
		"valid": {
			given:    []byte("1"),
			expected: ntypes.Double{Optional: &ntypes.DoubleValue{Value: 1}, Valid: true},
		},
		"max": {
			given:    []byte(strconv.FormatFloat(math.MaxFloat64, 'g', -1, 64)),
			expected: ntypes.Double{Optional: &ntypes.DoubleValue{Value: math.MaxFloat64}, Valid: true},
		},
		"min": {
			given:    []byte(strconv.FormatFloat(math.SmallestNonzeroFloat64, 'g', -1, 64)),
			expected: ntypes.Double{Optional: &ntypes.DoubleValue{Value: math.SmallestNonzeroFloat64}, Valid: true},
		},
		"zero": {
			given:    []byte("0"),
			expected: ntypes.Double{Optional: &ntypes.DoubleValue{Value: 0}, Valid: true},
		},
		"null": {
			given:    []byte("null"),
			expected: ntypes.Double{Optional: &ntypes.DoubleValue{Value: 0}, Valid: false},
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			var got ntypes.Double
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
