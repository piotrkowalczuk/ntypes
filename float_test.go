package ntypes_test

import (
	"encoding/json"
	"math"
	"strconv"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/piotrkowalczuk/ntypes"
)

func TestNewFloat(t *testing.T) {
	given := float32(124.124)
	got := ntypes.NewFloat(given)
	if !got.Valid {
		t.Error("float32 should be valid")
	}
	if got.GetOptional().GetValue() != given {
		t.Errorf("wrong float32, expected %f but got %f", given, got.GetOptional().GetValue())
	}
}

func TestFloat_FloatOr(t *testing.T) {
	cases := map[string]struct {
		expected float32
		or       float32
		given    *ntypes.Float
	}{
		"valid": {
			expected: 0.1,
			given:    &ntypes.Float{Optional: &ntypes.FloatValue{Value: 0.1}, Valid: true},
			or:       0.2,
		},
		"invalid": {
			expected: 0.2,
			given:    &ntypes.Float{Optional: &ntypes.FloatValue{Value: 0.1}, Valid: false},
			or:       0.2,
		},
		"nil": {
			expected: 0.3,
			or:       0.3,
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			got := c.given.FloatOr(c.or)

			if got != c.expected {
				t.Fatalf("wrong output, expected %f but got %f", c.expected, got)
			}
		})
	}
}

func TestFloat_Value(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		given := ntypes.Float{Optional: &ntypes.FloatValue{Value: 10}, Valid: true}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if _, ok := val.(float32); !ok {
			t.Fatalf("unexpected output type, expected Float but got %T", val)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		given := ntypes.Float{Optional: &ntypes.FloatValue{Value: 10.12}, Valid: false}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if val != nil {
			t.Fatalf("unexpected value, expected nil but got %v", val)
		}
	})
}

func TestFloat_Scan(t *testing.T) {
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
			str := &ntypes.Float{}
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

func TestFloat_UnmarshalJSON(t *testing.T) {
	cases := map[string]struct {
		expected ntypes.Float
		given    []byte
	}{
		"valid": {
			given:    []byte("1"),
			expected: ntypes.Float{Optional: &ntypes.FloatValue{Value: 1}, Valid: true},
		},
		"max": {
			given:    []byte(strconv.FormatFloat(math.MaxFloat32, 'e', 10, 32)),
			expected: ntypes.Float{Optional: &ntypes.FloatValue{Value: math.MaxFloat32}, Valid: true},
		},
		"min": {
			given:    []byte(strconv.FormatFloat(math.SmallestNonzeroFloat32, 'e', 10, 32)),
			expected: ntypes.Float{Optional: &ntypes.FloatValue{Value: math.SmallestNonzeroFloat32}, Valid: true},
		},
		"zero": {
			given:    []byte("0"),
			expected: ntypes.Float{Optional: &ntypes.FloatValue{Value: 0}, Valid: true},
		},
		"null": {
			given:    []byte("null"),
			expected: ntypes.Float{Optional: &ntypes.FloatValue{Value: 0}, Valid: false},
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			var got ntypes.Float
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

func TestFloat_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		given    *ntypes.Float
		expected string
	}{

		"nil": {
			given:    nil,
			expected: "null",
		},
		"zero value": {
			given:    &ntypes.Float{},
			expected: "null",
		},
		"valid": {
			given:    &ntypes.Float{Valid: true},
			expected: "0",
		},
		"invalid": {
			given:    &ntypes.Float{Valid: false},
			expected: "null",
		},
		"non zero valid value": {
			given:    &ntypes.Float{Optional: &ntypes.FloatValue{Value: 123}, Valid: true},
			expected: "123",
		},
		"non zero valid max value": {
			given:    &ntypes.Float{Optional: &ntypes.FloatValue{Value: math.MaxFloat32}, Valid: true},
			expected: strconv.FormatFloat(float64(math.MaxFloat32), 'e', 7, 32),
		},
		"non zero invalid max value": {
			given:    &ntypes.Float{Optional: &ntypes.FloatValue{Value: math.MaxFloat32}, Valid: false},
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
