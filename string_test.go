package ntypes_test

import (
	"reflect"
	"testing"

	"github.com/piotrkowalczuk/ntypes"
)

func TestNewString(t *testing.T) {
	given := "test"
	got := ntypes.NewString(given)
	if !got.Valid {
		t.Error("string should be valid")
	}
	if got.GetOptional().GetValue() != given {
		t.Errorf("wrong string, expected %s but got %s", given, got.GetOptional().GetValue())
	}
}

func TestString_Value(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		given := ntypes.String{Optional: &ntypes.StringValue{Value: "something"}, Valid: true}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if _, ok := val.(string); !ok {
			t.Fatalf("unexpected output type, expected string but got %T", val)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		given := ntypes.String{Optional: &ntypes.StringValue{Value: "something"}, Valid: false}
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
		given := &ntypes.String{Optional: &ntypes.StringValue{Value: "test"}, Valid: true}
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
		given := &ntypes.String{Optional: &ntypes.StringValue{Value: "test"}, Valid: false}
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

func TestString_UnmarshalJSON(t *testing.T) {
	cases := map[string]struct {
		json   string
		object *ntypes.String
	}{
		"valid": {
			json:   `"test"`,
			object: &ntypes.String{Optional: &ntypes.StringValue{Value: "test"}, Valid: true},
		},
		"invalid": {
			json:   `null`,
			object: &ntypes.String{Valid: false},
		},
	}
	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			var got ntypes.String
			err := got.UnmarshalJSON([]byte(c.json))
			if err != nil {
				t.Fatalf("unexpected error: %s", err.Error())
			}
			if c.object != nil {
				if !reflect.DeepEqual(got, *c.object) {
					t.Fatalf("wrong output, expected %v but got %v", c.object, &got)
				}
			}
		})
	}
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
				t.Fatalf("unexpected error: %s", err.Error())
			}

			if str.GetOptional().GetValue() != c.expected {
				t.Fatalf("wrong output, expected %s but got %s", c.expected, str.GetOptional().GetValue())
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
			given:    &ntypes.String{Optional: &ntypes.StringValue{Value: "test"}, Valid: true},
			or:       "alternative",
		},
		"invalid": {
			expected: "alternative",
			given:    &ntypes.String{Optional: &ntypes.StringValue{Value: "test"}, Valid: false},
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
