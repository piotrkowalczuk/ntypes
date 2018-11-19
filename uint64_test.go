package ntypes_test

import (
	"encoding/json"
	"math"
	"strconv"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/piotrkowalczuk/ntypes"
)

func TestNewUInt64(t *testing.T) {
	given := uint64(124)
	got := ntypes.NewUInt64(given)
	if !got.Valid {
		t.Error("uint64 should be valid")
	}
	if got.GetOptional().GetValue() != given {
		t.Errorf("wrong uint64, expected %d but got %d", given, got.GetOptional().GetValue())
	}
}

func TestUInt64_UInt64Or(t *testing.T) {
	cases := map[string]struct {
		expected uint64
		or       uint64
		given    *ntypes.UInt64
	}{
		"valid": {
			expected: 1,
			given:    &ntypes.UInt64{Optional: &ntypes.UInt64Value{Value: 1}, Valid: true},
			or:       2,
		},
		"invalid": {
			expected: 2,
			given:    &ntypes.UInt64{Optional: &ntypes.UInt64Value{Value: 1}, Valid: false},
			or:       2,
		},
		"nil": {
			expected: 3,
			or:       3,
		},
	}

	for hUInt64, c := range cases {
		t.Run(hUInt64, func(t *testing.T) {
			got := c.given.UInt64Or(c.or)

			if got != c.expected {
				t.Fatalf("wrong output, expected %d but got %d", c.expected, got)
			}
		})
	}
}

func TestUInt64_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		given    *ntypes.UInt64
		expected string
	}{

		"nil": {
			given:    nil,
			expected: "null",
		},
		"zero value": {
			given:    &ntypes.UInt64{},
			expected: "null",
		},
		"valid": {
			given:    &ntypes.UInt64{Valid: true},
			expected: "0",
		},
		"invalid": {
			given:    &ntypes.UInt64{Valid: false},
			expected: "null",
		},
		"non zero valid value": {
			given:    &ntypes.UInt64{Optional: &ntypes.UInt64Value{Value: 123}, Valid: true},
			expected: "123",
		},
		"non zero valid max value": {
			given:    &ntypes.UInt64{Optional: &ntypes.UInt64Value{Value: 4294967295}, Valid: true},
			expected: "4294967295",
		},
		"non zero valid min value": {
			given:    &ntypes.UInt64{Optional: &ntypes.UInt64Value{Value: 0}, Valid: true},
			expected: "0",
		},
		"non zero invalid max value": {
			given:    &ntypes.UInt64{Optional: &ntypes.UInt64Value{Value: 4294967295}, Valid: false},
			expected: "null",
		},
		"non zero invalid min value": {
			given:    &ntypes.UInt64{Optional: &ntypes.UInt64Value{Value: 0}, Valid: false},
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
		ID *ntypes.UInt64 `json:"id"`
	}
}

func TestUInt64_UnmarshalJSON(t *testing.T) {
	cases := map[string]struct {
		expected ntypes.UInt64
		given    []byte
	}{
		"valid": {
			given:    []byte("1"),
			expected: ntypes.UInt64{Optional: &ntypes.UInt64Value{Value: 1}, Valid: true},
		},
		"max": {
			given:    []byte(strconv.FormatUint(math.MaxUint64, 10)),
			expected: ntypes.UInt64{Optional: &ntypes.UInt64Value{Value: math.MaxUint64}, Valid: true},
		},
		"zero": {
			given:    []byte("0"),
			expected: ntypes.UInt64{Optional: &ntypes.UInt64Value{Value: 0}, Valid: true},
		},
		"null": {
			given:    []byte("null"),
			expected: ntypes.UInt64{Optional: &ntypes.UInt64Value{Value: 0}, Valid: false},
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			var got ntypes.UInt64
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
