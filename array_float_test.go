package ntypes_test

import (
	"reflect"
	"testing"

	"github.com/piotrkowalczuk/ntypes"
)

// -----------------------
// ---- FLOAT32 ARRAY ----
// -----------------------

func TestFloatArray_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		json   string
		object *ntypes.FloatArray
	}{
		"valid": {
			json:   `[1,2,3,4,5]`,
			object: &ntypes.FloatArray{Array: []float32{1, 2, 3, 4, 5}, Valid: true},
		},
		"invalid": {
			json:   `null`,
			object: &ntypes.FloatArray{Array: []float32{1, 2, 3, 4, 5}, Valid: false},
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

func TestFloatArray_UnmarshalJSON(t *testing.T) {
	cases := map[string]struct {
		json   string
		object *ntypes.FloatArray
	}{
		"valid": {
			json:   `[1,2,3,4,5]`,
			object: &ntypes.FloatArray{Array: []float32{1, 2, 3, 4, 5}, Valid: true},
		},
		"invalid": {
			json:   `null`,
			object: &ntypes.FloatArray{Valid: false},
		},
	}
	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			var got ntypes.FloatArray
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
