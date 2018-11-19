package ntypes_test

import (
	"reflect"
	"testing"

	"github.com/piotrkowalczuk/ntypes"
)

func TestStringArray_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		json   string
		object *ntypes.StringArray
	}{
		"valid": {
			json:   `["test"]`,
			object: &ntypes.StringArray{Array: []string{"test"}, Valid: true},
		},
		"invalid": {
			json:   `null`,
			object: &ntypes.StringArray{Array: []string{"test"}, Valid: false},
		},
		"nil": {
			json: `null`,
		},
	}
	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			got, err := c.object.MarshalJSON()
			if err != nil {
				t.Fatalf("unexpected error: %s", err.Error())
			}
			if string(got) != c.json {
				t.Fatalf("wrong output, expected %s but got %s", c.json, string(got))
			}
		})
	}
}

func TestStringArray_UnmarshalJSON(t *testing.T) {
	cases := map[string]struct {
		json   string
		object *ntypes.StringArray
	}{
		"valid": {
			json:   `["test"]`,
			object: &ntypes.StringArray{Array: []string{"test"}, Valid: true},
		},
		"invalid": {
			json:   `null`,
			object: &ntypes.StringArray{Valid: false},
		},
	}
	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			var got ntypes.StringArray
			err := got.UnmarshalJSON([]byte(c.json))
			if err != nil {
				t.Fatalf("unexpected error: %s", err.Error())
			}
			if c.object != nil {
				if !reflect.DeepEqual(got, *c.object) {
					t.Fatalf("wrong output, expected %s but got %s", c.object, &got)
				}
			}
		})
	}
}
