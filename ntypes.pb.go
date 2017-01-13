// Code generated by protoc-gen-go.
// source: ntypes.proto
// DO NOT EDIT!

/*
Package ntypes is a generated protocol buffer package.

It is generated from these files:
	ntypes.proto

It has these top-level messages:
	Bytes
	BytesArray
	String
	StringArray
	Int32
	Int32Array
	Int64
	Int64Array
	Uint32
	Uint32Array
	Uint64
	Uint64Array
	Float32
	Float32Array
	Float64
	Float64Array
	Bool
	BoolArray
*/
package ntypes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Bytes represents a byte slice that may be nil.
type Bytes struct {
	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Valid bool   `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Bytes) Reset()                    { *m = Bytes{} }
func (m *Bytes) String() string            { return proto.CompactTextString(m) }
func (*Bytes) ProtoMessage()               {}
func (*Bytes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// BytesArray represents an array of byte slices that may be nil.
type BytesArray struct {
	BytesArray [][]byte `protobuf:"bytes,1,rep,name=bytes_array,json=bytesArray,proto3" json:"bytes_array,omitempty"`
	Valid      bool     `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *BytesArray) Reset()                    { *m = BytesArray{} }
func (m *BytesArray) String() string            { return proto.CompactTextString(m) }
func (*BytesArray) ProtoMessage()               {}
func (*BytesArray) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// String represents a string that may be nil.
type String struct {
	Chars string `protobuf:"bytes,1,opt,name=chars" json:"chars,omitempty"`
	Valid bool   `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *String) Reset()                    { *m = String{} }
func (m *String) String() string            { return proto.CompactTextString(m) }
func (*String) ProtoMessage()               {}
func (*String) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// StringArray represents an array of strings that may be nil.
type StringArray struct {
	StringArray []string `protobuf:"bytes,1,rep,name=string_array,json=stringArray" json:"string_array,omitempty"`
	Valid       bool     `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *StringArray) Reset()                    { *m = StringArray{} }
func (m *StringArray) String() string            { return proto.CompactTextString(m) }
func (*StringArray) ProtoMessage()               {}
func (*StringArray) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Int32 represents a int32 that may be nil.
type Int32 struct {
	Int32 int32 `protobuf:"varint,1,opt,name=int32" json:"int32,omitempty"`
	Valid bool  `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Int32) Reset()                    { *m = Int32{} }
func (m *Int32) String() string            { return proto.CompactTextString(m) }
func (*Int32) ProtoMessage()               {}
func (*Int32) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Int32Array represents an array of int32s that may be nil.
type Int32Array struct {
	Int32Array []int32 `protobuf:"varint,1,rep,packed,name=int32_array,json=int32Array" json:"int32_array,omitempty"`
	Valid      bool    `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Int32Array) Reset()                    { *m = Int32Array{} }
func (m *Int32Array) String() string            { return proto.CompactTextString(m) }
func (*Int32Array) ProtoMessage()               {}
func (*Int32Array) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// Int64 represents a int64 that may be nil.
type Int64 struct {
	Int64 int64 `protobuf:"varint,1,opt,name=int64" json:"int64,omitempty"`
	Valid bool  `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Int64) Reset()                    { *m = Int64{} }
func (m *Int64) String() string            { return proto.CompactTextString(m) }
func (*Int64) ProtoMessage()               {}
func (*Int64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// Int64Array represents an array of int64s that may be nil.
type Int64Array struct {
	Int64Array []int64 `protobuf:"varint,1,rep,packed,name=int64_array,json=int64Array" json:"int64_array,omitempty"`
	Valid      bool    `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Int64Array) Reset()                    { *m = Int64Array{} }
func (m *Int64Array) String() string            { return proto.CompactTextString(m) }
func (*Int64Array) ProtoMessage()               {}
func (*Int64Array) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// Uint32 represents a Uint32 that may be nil.
type Uint32 struct {
	Uint32 uint32 `protobuf:"varint,1,opt,name=uint32" json:"uint32,omitempty"`
	Valid  bool   `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Uint32) Reset()                    { *m = Uint32{} }
func (m *Uint32) String() string            { return proto.CompactTextString(m) }
func (*Uint32) ProtoMessage()               {}
func (*Uint32) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// Uint32Array represents an array of uint32s that may be nil.
type Uint32Array struct {
	Uint32Array []uint32 `protobuf:"varint,1,rep,packed,name=uint32_array,json=uint32Array" json:"uint32_array,omitempty"`
	Valid       bool     `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Uint32Array) Reset()                    { *m = Uint32Array{} }
func (m *Uint32Array) String() string            { return proto.CompactTextString(m) }
func (*Uint32Array) ProtoMessage()               {}
func (*Uint32Array) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// Uint64 represents a Uint64 that may be nil.
type Uint64 struct {
	Uint64 uint64 `protobuf:"varint,1,opt,name=uint64" json:"uint64,omitempty"`
	Valid  bool   `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Uint64) Reset()                    { *m = Uint64{} }
func (m *Uint64) String() string            { return proto.CompactTextString(m) }
func (*Uint64) ProtoMessage()               {}
func (*Uint64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// Uint64Array represents an array of uint64s that may be nil.
type Uint64Array struct {
	Uint64Array []uint64 `protobuf:"varint,1,rep,packed,name=uint64_array,json=uint64Array" json:"uint64_array,omitempty"`
	Valid       bool     `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Uint64Array) Reset()                    { *m = Uint64Array{} }
func (m *Uint64Array) String() string            { return proto.CompactTextString(m) }
func (*Uint64Array) ProtoMessage()               {}
func (*Uint64Array) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

// Float32 represents a Float32 that may be nil.
type Float32 struct {
	Float32 float32 `protobuf:"fixed32,1,opt,name=float32" json:"float32,omitempty"`
	Valid   bool    `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Float32) Reset()                    { *m = Float32{} }
func (m *Float32) String() string            { return proto.CompactTextString(m) }
func (*Float32) ProtoMessage()               {}
func (*Float32) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

// Float32Array represents an array of float32s that may be nil.
type Float32Array struct {
	Float32Array []float32 `protobuf:"fixed32,1,rep,packed,name=float32_array,json=float32Array" json:"float32_array,omitempty"`
	Valid        bool      `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Float32Array) Reset()                    { *m = Float32Array{} }
func (m *Float32Array) String() string            { return proto.CompactTextString(m) }
func (*Float32Array) ProtoMessage()               {}
func (*Float32Array) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

// Float64 represents a Float64 that may be nil.
type Float64 struct {
	Float64 float64 `protobuf:"fixed64,1,opt,name=float64" json:"float64,omitempty"`
	Valid   bool    `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Float64) Reset()                    { *m = Float64{} }
func (m *Float64) String() string            { return proto.CompactTextString(m) }
func (*Float64) ProtoMessage()               {}
func (*Float64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

// Float64Array represents an array of float64s that may be nil.
type Float64Array struct {
	Float64Array []float64 `protobuf:"fixed64,1,rep,packed,name=float64_array,json=float64Array" json:"float64_array,omitempty"`
	Valid        bool      `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Float64Array) Reset()                    { *m = Float64Array{} }
func (m *Float64Array) String() string            { return proto.CompactTextString(m) }
func (*Float64Array) ProtoMessage()               {}
func (*Float64Array) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

// Bool represents a bool that may be nil.
type Bool struct {
	Bool  bool `protobuf:"varint,1,opt,name=bool" json:"bool,omitempty"`
	Valid bool `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Bool) Reset()                    { *m = Bool{} }
func (m *Bool) String() string            { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()               {}
func (*Bool) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

// BoolArray represents an array of booleans that may be nil.
type BoolArray struct {
	BoolArray []bool `protobuf:"varint,1,rep,packed,name=bool_array,json=boolArray" json:"bool_array,omitempty"`
	Valid     bool   `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *BoolArray) Reset()                    { *m = BoolArray{} }
func (m *BoolArray) String() string            { return proto.CompactTextString(m) }
func (*BoolArray) ProtoMessage()               {}
func (*BoolArray) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func init() {
	proto.RegisterType((*Bytes)(nil), "ntypes.Bytes")
	proto.RegisterType((*BytesArray)(nil), "ntypes.BytesArray")
	proto.RegisterType((*String)(nil), "ntypes.String")
	proto.RegisterType((*StringArray)(nil), "ntypes.StringArray")
	proto.RegisterType((*Int32)(nil), "ntypes.Int32")
	proto.RegisterType((*Int32Array)(nil), "ntypes.Int32Array")
	proto.RegisterType((*Int64)(nil), "ntypes.Int64")
	proto.RegisterType((*Int64Array)(nil), "ntypes.Int64Array")
	proto.RegisterType((*Uint32)(nil), "ntypes.Uint32")
	proto.RegisterType((*Uint32Array)(nil), "ntypes.Uint32Array")
	proto.RegisterType((*Uint64)(nil), "ntypes.Uint64")
	proto.RegisterType((*Uint64Array)(nil), "ntypes.Uint64Array")
	proto.RegisterType((*Float32)(nil), "ntypes.Float32")
	proto.RegisterType((*Float32Array)(nil), "ntypes.Float32Array")
	proto.RegisterType((*Float64)(nil), "ntypes.Float64")
	proto.RegisterType((*Float64Array)(nil), "ntypes.Float64Array")
	proto.RegisterType((*Bool)(nil), "ntypes.Bool")
	proto.RegisterType((*BoolArray)(nil), "ntypes.BoolArray")
}

func init() { proto.RegisterFile("ntypes.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x94, 0x41, 0x6b, 0xab, 0x40,
	0x14, 0x85, 0x31, 0x51, 0xa3, 0x57, 0xb3, 0x91, 0xc7, 0xc3, 0xcd, 0xe3, 0x25, 0xe9, 0x26, 0xab,
	0x52, 0xaa, 0x0c, 0x74, 0xd7, 0xa6, 0x10, 0xc8, 0xf6, 0x96, 0x6c, 0xba, 0x29, 0xda, 0x26, 0xad,
	0x20, 0x99, 0xa0, 0xa6, 0x90, 0x7f, 0x5f, 0x66, 0xee, 0x4c, 0x32, 0x29, 0x6a, 0x76, 0x73, 0x9c,
	0x73, 0xee, 0x77, 0xcf, 0x2c, 0x84, 0x70, 0xd7, 0x1c, 0xf7, 0x9b, 0xfa, 0x76, 0x5f, 0xf1, 0x86,
	0x47, 0x2e, 0xa9, 0x59, 0x02, 0xce, 0xe2, 0xd8, 0x6c, 0xea, 0xe8, 0x0f, 0x38, 0xb9, 0x38, 0xc4,
	0xd6, 0xc4, 0x9a, 0x87, 0x48, 0x42, 0x7c, 0xfd, 0xce, 0xca, 0xe2, 0x23, 0x1e, 0x4c, 0xac, 0xb9,
	0x87, 0x24, 0x66, 0xcf, 0x00, 0x32, 0xf4, 0x54, 0x55, 0xd9, 0x31, 0xfa, 0x0f, 0x81, 0x34, 0xbf,
	0x65, 0x42, 0xc6, 0xd6, 0x64, 0x38, 0x0f, 0x11, 0xf2, 0xb3, 0xa1, 0x7d, 0x48, 0x0a, 0xee, 0x4b,
	0x53, 0x15, 0xbb, 0x4f, 0x71, 0xff, 0xfe, 0x95, 0x55, 0x84, 0xf6, 0x91, 0x44, 0x47, 0x6a, 0x09,
	0x01, 0xa5, 0x68, 0xf4, 0x14, 0xc2, 0x5a, 0x4a, 0x03, 0xee, 0x63, 0x50, 0x1b, 0x96, 0xf6, 0x39,
	0x09, 0x38, 0xab, 0x5d, 0x93, 0xdc, 0x8b, 0xeb, 0x42, 0x1c, 0x24, 0xdc, 0x41, 0x12, 0xdd, 0xbd,
	0x65, 0xe8, 0xd4, 0x5b, 0x9a, 0x0d, 0xb4, 0x83, 0x50, 0x9c, 0x0d, 0x7d, 0x64, 0x96, 0x2a, 0x32,
	0x4b, 0x25, 0x79, 0x88, 0x24, 0x7a, 0xc9, 0x2c, 0x35, 0xc9, 0x2c, 0x35, 0xc8, 0x43, 0x49, 0xd6,
	0x86, 0xf6, 0x21, 0x0c, 0xdc, 0x35, 0xd5, 0xfb, 0x0b, 0xee, 0xe1, 0xdc, 0x7a, 0x8c, 0x4a, 0x75,
	0xbf, 0xf9, 0xda, 0xa8, 0x35, 0x85, 0xf0, 0xf0, 0xbb, 0xf8, 0x18, 0x83, 0xc3, 0xd5, 0xe6, 0x8a,
	0xcf, 0x52, 0xcd, 0x57, 0xdd, 0x6d, 0x54, 0xaa, 0x9f, 0xaf, 0xcb, 0x29, 0xfe, 0x45, 0x7d, 0x9b,
	0xf8, 0xfd, 0xfd, 0x1f, 0x60, 0xb4, 0x2c, 0x79, 0x26, 0x8a, 0xc6, 0x30, 0xda, 0xd2, 0x51, 0x6e,
	0x30, 0x40, 0x2d, 0x3b, 0xa2, 0x2b, 0x08, 0x55, 0x94, 0x00, 0x37, 0x30, 0x56, 0x01, 0x63, 0x89,
	0x01, 0x86, 0x5b, 0xd3, 0xd4, 0xbf, 0x05, 0x4b, 0x4f, 0x5b, 0xa8, 0x77, 0xb0, 0x50, 0xcb, 0x2b,
	0x5b, 0xe8, 0x9a, 0x7a, 0x8b, 0x8b, 0xa7, 0xb0, 0xd4, 0x16, 0xfd, 0x6f, 0x71, 0x07, 0xf6, 0x82,
	0xf3, 0x32, 0x8a, 0xc0, 0xce, 0x39, 0x2f, 0x25, 0xdf, 0x43, 0x79, 0xee, 0x48, 0x3c, 0x82, 0x2f,
	0x12, 0x34, 0xf4, 0x1f, 0x80, 0xb0, 0x1a, 0x58, 0x0f, 0xfd, 0xfc, 0x74, 0xdd, 0x3a, 0x61, 0xe1,
	0xbd, 0xaa, 0xbf, 0x4e, 0xee, 0xca, 0x9f, 0x50, 0xf2, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x9a,
	0x97, 0x76, 0x94, 0x04, 0x00, 0x00,
}
