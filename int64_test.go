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

func TestNewInt64(t *testing.T) {
	given := int64(124)
	got := ntypes.NewInt64(given)
	if !got.Valid {
		t.Error("int64 should be valid")
	}
	if got.GetOptional().GetValue() != given {
		t.Errorf("wrong int64, expected %d but got %d", given, got.GetOptional().GetValue())
	}
}

func TestInt64_Int64Or(t *testing.T) {
	cases := map[string]struct {
		expected int64
		or       int64
		given    *ntypes.Int64
	}{
		"valid": {
			expected: 1,
			given:    &ntypes.Int64{Optional: &ntypes.Int64Value{Value: 1}, Valid: true},
			or:       2,
		},
		"invalid": {
			expected: 2,
			given:    &ntypes.Int64{Optional: &ntypes.Int64Value{Value: 1}, Valid: false},
			or:       2,
		},
		"nil": {
			expected: 3,
			or:       3,
		},
	}

	for hint64, c := range cases {
		t.Run(hint64, func(t *testing.T) {
			got := c.given.Int64Or(c.or)

			if got != c.expected {
				t.Fatalf("wrong output, expected %d but got %d", c.expected, got)
			}
		})
	}
}

func TestInt64_ProtoMessage(t *testing.T) {
	var (
		buf []byte
		err error
		tmp ntypes.Int64
	)
	cases := map[string]ntypes.Int64{
		"valid":        {Optional: &ntypes.Int64Value{Value: 1}, Valid: true},
		"zero":         {Optional: &ntypes.Int64Value{Value: 0}, Valid: true},
		"invalid-zero": {Optional: &ntypes.Int64Value{Value: 0}, Valid: false},
		"invalid":      {Optional: &ntypes.Int64Value{Value: 13123}, Valid: false},
		"negative":     {Optional: &ntypes.Int64Value{Value: -1241223}, Valid: true},
	}

	for hint, given := range cases {
		t.Run(hint, func(t *testing.T) {
			buf, err = proto.Marshal(&given)
			if err != nil {
				t.Fatalf("marshal returned unexpected error: %s", err.Error())
			}
			err = proto.Unmarshal(buf, &tmp)
			if err != nil {
				t.Fatalf("unmarshal returned unexpected error: %s", err.Error())
			}
			if tmp.GetOptional().GetValue() != given.GetOptional().GetValue() {
				t.Errorf("integers are not equal expected %d, got %d", given.GetOptional().GetValue(), tmp.GetOptional().GetValue())
			}
			if tmp.Valid != given.Valid {
				t.Errorf("booleans are not equal expected %t, got %t", given.Valid, tmp.Valid)
			}
		})
	}
}

func TestInt64_Value(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		given := ntypes.Int64{Optional: &ntypes.Int64Value{Value: 10}, Valid: true}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if _, ok := val.(int64); !ok {
			t.Fatalf("unexpected output type, expected int64 but got %T", val)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		given := ntypes.Int64{Optional: &ntypes.Int64Value{Value: 10}, Valid: false}
		val, err := given.Value()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if val != nil {
			t.Fatalf("unexpected value, expected nil but got %v", val)
		}
	})
}

func TestInt64_Scan(t *testing.T) {
	cases := map[string]struct {
		given    interface{}
		expected int64
	}{
		"bytes": {
			given:    []byte("15"),
			expected: 15,
		},
		"string": {
			given:    "16",
			expected: 16,
		},
		"int64": {
			given:    int64(18),
			expected: 18,
		},
		"nil": {
			expected: 0,
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			str := &ntypes.Int64{}
			err := str.Scan(c.given)
			if err != nil {
				t.Fatalf("unexpected error: %s", err.Error())
			}
			if c.expected != str.GetOptional().GetValue() {
				t.Fatalf("wrong output, expected %d but got %d", c.expected, str.GetOptional().GetValue())
			}

		})
	}
}

func TestInt64_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		given    *ntypes.Int64
		expected string
	}{

		"nil": {
			given:    nil,
			expected: "null",
		},
		"zero value": {
			given:    &ntypes.Int64{},
			expected: "null",
		},
		"valid": {
			given:    &ntypes.Int64{Valid: true},
			expected: "0",
		},
		"invalid": {
			given:    &ntypes.Int64{Valid: false},
			expected: "null",
		},
		"non zero valid value": {
			given:    &ntypes.Int64{Optional: &ntypes.Int64Value{Value: 123}, Valid: true},
			expected: "123",
		},
		"non zero valid max value": {
			given:    &ntypes.Int64{Optional: &ntypes.Int64Value{Value: math.MaxInt64}, Valid: true},
			expected: fmt.Sprintf("%d", math.MaxInt64),
		},
		"non zero valid min value": {
			given:    &ntypes.Int64{Optional: &ntypes.Int64Value{Value: math.MinInt64}, Valid: true},
			expected: fmt.Sprintf("%d", math.MinInt64),
		},
		"non zero invalid max value": {
			given:    &ntypes.Int64{Optional: &ntypes.Int64Value{Value: math.MaxInt64}, Valid: false},
			expected: "null",
		},
		"non zero invalid min value": {
			given:    &ntypes.Int64{Optional: &ntypes.Int64Value{Value: -9223372036854775808}, Valid: false},
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

func TestInt64_UnmarshalJSON(t *testing.T) {
	cases := map[string]struct {
		expected ntypes.Int64
		given    []byte
	}{
		"valid": {
			given:    []byte("1"),
			expected: ntypes.Int64{Optional: &ntypes.Int64Value{Value: 1}, Valid: true},
		},
		"max": {
			given:    []byte(strconv.FormatInt(math.MaxInt64, 10)),
			expected: ntypes.Int64{Optional: &ntypes.Int64Value{Value: math.MaxInt64}, Valid: true},
		},
		"min": {
			given:    []byte(strconv.FormatInt(math.MinInt64, 10)),
			expected: ntypes.Int64{Optional: &ntypes.Int64Value{Value: math.MinInt64}, Valid: true},
		},
		"zero": {
			given:    []byte("0"),
			expected: ntypes.Int64{Optional: &ntypes.Int64Value{Value: 0}, Valid: true},
		},
		"invalid-null": {
			given:    []byte("null"),
			expected: ntypes.Int64{Optional: &ntypes.Int64Value{Value: 0}, Valid: false},
		},
		"valid-null": {
			given:    []byte("null"),
			expected: ntypes.Int64{Optional: nil, Valid: true},
		},
	}

	for hint, c := range cases {
		t.Run(hint, func(t *testing.T) {
			var got ntypes.Int64
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
