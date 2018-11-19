package ntypes_test

import (
	"encoding/json"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/piotrkowalczuk/ntypes"
)

func TestFalse(t *testing.T) {
	got := ntypes.False()
	if !got.Valid {
		t.Error("bool should be valid")
	}
	if got.GetOptional().GetValue() {
		t.Error("bool should be false")
	}
}

func TestTrue(t *testing.T) {
	got := ntypes.True()
	if !got.Valid {
		t.Error("bool should be valid")
	}
	if !got.GetOptional().GetValue() {
		t.Error("bool should be true")
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
			given:    &ntypes.Bool{Optional: &ntypes.BoolValue{Value: true}, Valid: true},
			or:       false,
		},
		"invalid": {
			expected: false,
			given:    &ntypes.Bool{Optional: &ntypes.BoolValue{Value: true}, Valid: false},
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
				t.Fatalf("wrong output, expected %t but got %t", c.expected, got)
			}
		})
	}
}

func TestBool_Value(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		given := ntypes.Bool{Optional: &ntypes.BoolValue{Value: true}, Valid: true}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if _, ok := val.(bool); !ok {
			t.Fatalf("unexpected output type, expected bool but got %T", val)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		given := ntypes.Bool{Optional: &ntypes.BoolValue{Value: true}, Valid: false}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if val != nil {
			t.Fatalf("unexpected value, expected nil but got %v", val)
		}
	})
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
		"bool": {
			given:    true,
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
				t.Fatalf("unexpected error: %s", err.Error())
			}
			if c.expected != str.GetOptional().GetValue() {
				t.Fatalf("wrong output, expected %t but got %t", c.expected, str.GetOptional().GetValue())
			}

		})
	}
}

func TestBool_JSON(t *testing.T) {
	cases := []*ntypes.Bool{
		{},
		nil,
	}

	for _, given := range cases {
		t.Run(given.String(), func(t *testing.T) {
			buf, err := json.Marshal(given)
			if err != nil {
				t.Fatalf("json marshal failure: %s", err.Error())
			}
			var got ntypes.Bool
			if err =json.Unmarshal(buf,&got); err != nil {
				t.Fatalf("json unmarshal failure: %s", err.Error())
			}

			if proto.Equal(given, got)
		})
	}
}
func TestBool_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		given    *ntypes.Bool
		expected string
	}{
		"nil": {
			given:    nil,
			expected: "",
		},
		"zero-value": {
			given:    &ntypes.Bool{},
			expected: "",
		},
		"valid": {
			given:    &ntypes.Bool{Optional: &ntypes.BoolValue{Value: false}, Valid: true},
			expected: "false",
		},
		"nil-invalid": {given: &ntypes.Bool{Valid: false}, expected: ""},
		"true-true":   {given: &ntypes.Bool{Optional: &ntypes.BoolValue{Value: true}, Valid: true}, expected: "true"},
		"true-false":  {given: &ntypes.Bool{Optional: &ntypes.BoolValue{Value: true}, Valid: false}, expected: "null"},
		"false-false": {given: &ntypes.Bool{Optional: &ntypes.BoolValue{Value: false}, Valid: false}, expected: "null"},
		"false-true":  {given: &ntypes.Bool{Optional: &ntypes.BoolValue{Value: false}, Valid: true}, expected: "false"},
	}

	for d, c := range cases {
		t.Run(d, func(t *testing.T) {
			got, err := json.Marshal(c.given)
			if err != nil {
				t.Fatalf("simple: %s: unexpected error: %s", d, err.Error())
			}
			if string(got) != c.expected {
				t.Errorf("wrong output, expected %s but got %s", c.expected, string(got))
			} else {
				t.Logf("correct output: %s", string(got))
			}

		})
	}

	type within struct {
		Exists *ntypes.Bool `json:"exists"`
	}

	for d, c := range cases {
		t.Run(d, func(t *testing.T) {
			w := within{Exists: c.given}
			_, err := json.Marshal(w)
			if err != nil {
				t.Errorf("within: %s: unexpected error: %s", d, err.Error())
			}
		})
	}
}

func TestBool_UnmarshalJSON(t *testing.T) {
	cases := map[string]struct {
		expected *ntypes.Bool
		given    []byte
	}{
		"valid-true": {
			given:    []byte("true"),
			expected: &ntypes.Bool{Optional: &ntypes.BoolValue{Value: true}, Valid: true},
		},
		"valid-false": {
			given:    []byte("false"),
			expected: &ntypes.Bool{Optional: &ntypes.BoolValue{Value: false}, Valid: true},
		},
		"invalid": {
			given:    []byte(""),
			expected: &ntypes.Bool{Valid: true},
		},
		"nil": {
			given: []byte(""),
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			var got ntypes.Bool
			err := json.Unmarshal(c.given, &got)
			if err != nil {
				t.Fatalf("unexpected error for %s: %s", string(c.given), err.Error())
			}
			if !proto.Equal(c.expected, &got) {
				t.Errorf("values are not equal, expected %v, got %v", *c.expected, got)
			}
		})
	}
}
