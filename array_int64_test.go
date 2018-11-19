package ntypes_test

import (
	"reflect"
	"testing"

	"github.com/piotrkowalczuk/ntypes"
)

func TestInt64Array_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		json   string
		object *ntypes.Int64Array
	}{
		"valid": {
			json:   `[1,2,3,4,5]`,
			object: &ntypes.Int64Array{Array: []int64{1, 2, 3, 4, 5}, Valid: true},
		},
		"invalid": {
			json:   `null`,
			object: &ntypes.Int64Array{Array: []int64{1, 2, 3, 4, 5}, Valid: false},
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

func TestInt64Array_UnmarshalJSON(t *testing.T) {
	cases := map[string]struct {
		json   string
		object *ntypes.Int64Array
	}{
		"valid": {
			json:   `[1,2,3,4,5]`,
			object: &ntypes.Int64Array{Array: []int64{1, 2, 3, 4, 5}, Valid: true},
		},
		"invalid": {
			json:   `null`,
			object: &ntypes.Int64Array{Valid: false},
		},
	}
	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			var got ntypes.Int64Array
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
