package ntypes

import (
	"bytes"

	"github.com/golang/protobuf/ptypes/wrappers"
)

// StringValue is a handy type alias for wrappers StringValue.
type StringValue = wrappers.StringValue

// BoolValue is a handy type alias for wrappers BoolValue.
type BoolValue = wrappers.BoolValue

// BytesValue is a handy type alias for wrappers BytesValue.
type BytesValue = wrappers.BytesValue

// Int64Value is a handy type alias for wrappers Int64Value.
type Int64Value = wrappers.Int64Value

// Int32Value is a handy type alias for wrappers Int32Value.
type Int32Value = wrappers.Int32Value

// UInt32Value is a handy type alias for wrappers UInt32Value.
type UInt32Value = wrappers.UInt32Value

// UInt64Value is a handy type alias for wrappers UInt64Value.
type UInt64Value = wrappers.UInt64Value

// FloatValue is a handy type alias for wrappers FloatValue.
type FloatValue = wrappers.FloatValue

// DoubleValue is a handy type alias for wrappers DoubleValue.
type DoubleValue = wrappers.DoubleValue

func isNull(data []byte) bool {
	return bytes.Equal(data, []byte("null"))
}
