# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ntypes.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='ntypes.proto',
  package='ntypes',
  syntax='proto3',
  serialized_pb=_b('\n\x0cntypes.proto\x12\x06ntypes\"%\n\x05\x42ytes\x12\r\n\x05\x62ytes\x18\x01 \x01(\x0c\x12\r\n\x05valid\x18\x02 \x01(\x08\"0\n\nBytesArray\x12\x13\n\x0b\x62ytes_array\x18\x01 \x03(\x0c\x12\r\n\x05valid\x18\x02 \x01(\x08\"&\n\x06String\x12\r\n\x05\x63hars\x18\x01 \x01(\t\x12\r\n\x05valid\x18\x02 \x01(\x08\"2\n\x0bStringArray\x12\x14\n\x0cstring_array\x18\x01 \x03(\t\x12\r\n\x05valid\x18\x02 \x01(\x08\"%\n\x05Int32\x12\r\n\x05int32\x18\x01 \x01(\x05\x12\r\n\x05valid\x18\x02 \x01(\x08\"0\n\nInt32Array\x12\x13\n\x0bint32_array\x18\x01 \x03(\x05\x12\r\n\x05valid\x18\x02 \x01(\x08\"%\n\x05Int64\x12\r\n\x05int64\x18\x01 \x01(\x03\x12\r\n\x05valid\x18\x02 \x01(\x08\"0\n\nInt64Array\x12\x13\n\x0bint64_array\x18\x01 \x03(\x03\x12\r\n\x05valid\x18\x02 \x01(\x08\"\'\n\x06Uint32\x12\x0e\n\x06uint32\x18\x01 \x01(\r\x12\r\n\x05valid\x18\x02 \x01(\x08\"2\n\x0bUint32Array\x12\x14\n\x0cuint32_array\x18\x01 \x03(\r\x12\r\n\x05valid\x18\x02 \x01(\x08\"\'\n\x06Uint64\x12\x0e\n\x06uint64\x18\x01 \x01(\x04\x12\r\n\x05valid\x18\x02 \x01(\x08\"2\n\x0bUint64Array\x12\x14\n\x0cuint64_array\x18\x01 \x03(\x04\x12\r\n\x05valid\x18\x02 \x01(\x08\")\n\x07\x46loat32\x12\x0f\n\x07\x66loat32\x18\x01 \x01(\x02\x12\r\n\x05valid\x18\x02 \x01(\x08\"4\n\x0c\x46loat32Array\x12\x15\n\rfloat32_array\x18\x01 \x03(\x02\x12\r\n\x05valid\x18\x02 \x01(\x08\")\n\x07\x46loat64\x12\x0f\n\x07\x66loat64\x18\x01 \x01(\x01\x12\r\n\x05valid\x18\x02 \x01(\x08\"4\n\x0c\x46loat64Array\x12\x15\n\rfloat64_array\x18\x01 \x03(\x01\x12\r\n\x05valid\x18\x02 \x01(\x08\"#\n\x04\x42ool\x12\x0c\n\x04\x62ool\x18\x01 \x01(\x08\x12\r\n\x05valid\x18\x02 \x01(\x08\".\n\tBoolArray\x12\x12\n\nbool_array\x18\x01 \x03(\x08\x12\r\n\x05valid\x18\x02 \x01(\x08\x42\x08Z\x06ntypesb\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_BYTES = _descriptor.Descriptor(
  name='Bytes',
  full_name='ntypes.Bytes',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='bytes', full_name='ntypes.Bytes.bytes', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.Bytes.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=24,
  serialized_end=61,
)


_BYTESARRAY = _descriptor.Descriptor(
  name='BytesArray',
  full_name='ntypes.BytesArray',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='bytes_array', full_name='ntypes.BytesArray.bytes_array', index=0,
      number=1, type=12, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.BytesArray.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=63,
  serialized_end=111,
)


_STRING = _descriptor.Descriptor(
  name='String',
  full_name='ntypes.String',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='chars', full_name='ntypes.String.chars', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.String.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=113,
  serialized_end=151,
)


_STRINGARRAY = _descriptor.Descriptor(
  name='StringArray',
  full_name='ntypes.StringArray',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='string_array', full_name='ntypes.StringArray.string_array', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.StringArray.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=153,
  serialized_end=203,
)


_INT32 = _descriptor.Descriptor(
  name='Int32',
  full_name='ntypes.Int32',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='int32', full_name='ntypes.Int32.int32', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.Int32.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=205,
  serialized_end=242,
)


_INT32ARRAY = _descriptor.Descriptor(
  name='Int32Array',
  full_name='ntypes.Int32Array',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='int32_array', full_name='ntypes.Int32Array.int32_array', index=0,
      number=1, type=5, cpp_type=1, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.Int32Array.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=244,
  serialized_end=292,
)


_INT64 = _descriptor.Descriptor(
  name='Int64',
  full_name='ntypes.Int64',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='int64', full_name='ntypes.Int64.int64', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.Int64.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=294,
  serialized_end=331,
)


_INT64ARRAY = _descriptor.Descriptor(
  name='Int64Array',
  full_name='ntypes.Int64Array',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='int64_array', full_name='ntypes.Int64Array.int64_array', index=0,
      number=1, type=3, cpp_type=2, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.Int64Array.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=333,
  serialized_end=381,
)


_UINT32 = _descriptor.Descriptor(
  name='Uint32',
  full_name='ntypes.Uint32',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uint32', full_name='ntypes.Uint32.uint32', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.Uint32.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=383,
  serialized_end=422,
)


_UINT32ARRAY = _descriptor.Descriptor(
  name='Uint32Array',
  full_name='ntypes.Uint32Array',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uint32_array', full_name='ntypes.Uint32Array.uint32_array', index=0,
      number=1, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.Uint32Array.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=424,
  serialized_end=474,
)


_UINT64 = _descriptor.Descriptor(
  name='Uint64',
  full_name='ntypes.Uint64',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uint64', full_name='ntypes.Uint64.uint64', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.Uint64.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=476,
  serialized_end=515,
)


_UINT64ARRAY = _descriptor.Descriptor(
  name='Uint64Array',
  full_name='ntypes.Uint64Array',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uint64_array', full_name='ntypes.Uint64Array.uint64_array', index=0,
      number=1, type=4, cpp_type=4, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.Uint64Array.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=517,
  serialized_end=567,
)


_FLOAT32 = _descriptor.Descriptor(
  name='Float32',
  full_name='ntypes.Float32',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='float32', full_name='ntypes.Float32.float32', index=0,
      number=1, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.Float32.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=569,
  serialized_end=610,
)


_FLOAT32ARRAY = _descriptor.Descriptor(
  name='Float32Array',
  full_name='ntypes.Float32Array',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='float32_array', full_name='ntypes.Float32Array.float32_array', index=0,
      number=1, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.Float32Array.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=612,
  serialized_end=664,
)


_FLOAT64 = _descriptor.Descriptor(
  name='Float64',
  full_name='ntypes.Float64',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='float64', full_name='ntypes.Float64.float64', index=0,
      number=1, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.Float64.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=666,
  serialized_end=707,
)


_FLOAT64ARRAY = _descriptor.Descriptor(
  name='Float64Array',
  full_name='ntypes.Float64Array',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='float64_array', full_name='ntypes.Float64Array.float64_array', index=0,
      number=1, type=1, cpp_type=5, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.Float64Array.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=709,
  serialized_end=761,
)


_BOOL = _descriptor.Descriptor(
  name='Bool',
  full_name='ntypes.Bool',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='bool', full_name='ntypes.Bool.bool', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.Bool.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=763,
  serialized_end=798,
)


_BOOLARRAY = _descriptor.Descriptor(
  name='BoolArray',
  full_name='ntypes.BoolArray',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='bool_array', full_name='ntypes.BoolArray.bool_array', index=0,
      number=1, type=8, cpp_type=7, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='ntypes.BoolArray.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=800,
  serialized_end=846,
)

DESCRIPTOR.message_types_by_name['Bytes'] = _BYTES
DESCRIPTOR.message_types_by_name['BytesArray'] = _BYTESARRAY
DESCRIPTOR.message_types_by_name['String'] = _STRING
DESCRIPTOR.message_types_by_name['StringArray'] = _STRINGARRAY
DESCRIPTOR.message_types_by_name['Int32'] = _INT32
DESCRIPTOR.message_types_by_name['Int32Array'] = _INT32ARRAY
DESCRIPTOR.message_types_by_name['Int64'] = _INT64
DESCRIPTOR.message_types_by_name['Int64Array'] = _INT64ARRAY
DESCRIPTOR.message_types_by_name['Uint32'] = _UINT32
DESCRIPTOR.message_types_by_name['Uint32Array'] = _UINT32ARRAY
DESCRIPTOR.message_types_by_name['Uint64'] = _UINT64
DESCRIPTOR.message_types_by_name['Uint64Array'] = _UINT64ARRAY
DESCRIPTOR.message_types_by_name['Float32'] = _FLOAT32
DESCRIPTOR.message_types_by_name['Float32Array'] = _FLOAT32ARRAY
DESCRIPTOR.message_types_by_name['Float64'] = _FLOAT64
DESCRIPTOR.message_types_by_name['Float64Array'] = _FLOAT64ARRAY
DESCRIPTOR.message_types_by_name['Bool'] = _BOOL
DESCRIPTOR.message_types_by_name['BoolArray'] = _BOOLARRAY

Bytes = _reflection.GeneratedProtocolMessageType('Bytes', (_message.Message,), dict(
  DESCRIPTOR = _BYTES,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.Bytes)
  ))
_sym_db.RegisterMessage(Bytes)

BytesArray = _reflection.GeneratedProtocolMessageType('BytesArray', (_message.Message,), dict(
  DESCRIPTOR = _BYTESARRAY,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.BytesArray)
  ))
_sym_db.RegisterMessage(BytesArray)

String = _reflection.GeneratedProtocolMessageType('String', (_message.Message,), dict(
  DESCRIPTOR = _STRING,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.String)
  ))
_sym_db.RegisterMessage(String)

StringArray = _reflection.GeneratedProtocolMessageType('StringArray', (_message.Message,), dict(
  DESCRIPTOR = _STRINGARRAY,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.StringArray)
  ))
_sym_db.RegisterMessage(StringArray)

Int32 = _reflection.GeneratedProtocolMessageType('Int32', (_message.Message,), dict(
  DESCRIPTOR = _INT32,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.Int32)
  ))
_sym_db.RegisterMessage(Int32)

Int32Array = _reflection.GeneratedProtocolMessageType('Int32Array', (_message.Message,), dict(
  DESCRIPTOR = _INT32ARRAY,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.Int32Array)
  ))
_sym_db.RegisterMessage(Int32Array)

Int64 = _reflection.GeneratedProtocolMessageType('Int64', (_message.Message,), dict(
  DESCRIPTOR = _INT64,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.Int64)
  ))
_sym_db.RegisterMessage(Int64)

Int64Array = _reflection.GeneratedProtocolMessageType('Int64Array', (_message.Message,), dict(
  DESCRIPTOR = _INT64ARRAY,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.Int64Array)
  ))
_sym_db.RegisterMessage(Int64Array)

Uint32 = _reflection.GeneratedProtocolMessageType('Uint32', (_message.Message,), dict(
  DESCRIPTOR = _UINT32,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.Uint32)
  ))
_sym_db.RegisterMessage(Uint32)

Uint32Array = _reflection.GeneratedProtocolMessageType('Uint32Array', (_message.Message,), dict(
  DESCRIPTOR = _UINT32ARRAY,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.Uint32Array)
  ))
_sym_db.RegisterMessage(Uint32Array)

Uint64 = _reflection.GeneratedProtocolMessageType('Uint64', (_message.Message,), dict(
  DESCRIPTOR = _UINT64,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.Uint64)
  ))
_sym_db.RegisterMessage(Uint64)

Uint64Array = _reflection.GeneratedProtocolMessageType('Uint64Array', (_message.Message,), dict(
  DESCRIPTOR = _UINT64ARRAY,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.Uint64Array)
  ))
_sym_db.RegisterMessage(Uint64Array)

Float32 = _reflection.GeneratedProtocolMessageType('Float32', (_message.Message,), dict(
  DESCRIPTOR = _FLOAT32,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.Float32)
  ))
_sym_db.RegisterMessage(Float32)

Float32Array = _reflection.GeneratedProtocolMessageType('Float32Array', (_message.Message,), dict(
  DESCRIPTOR = _FLOAT32ARRAY,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.Float32Array)
  ))
_sym_db.RegisterMessage(Float32Array)

Float64 = _reflection.GeneratedProtocolMessageType('Float64', (_message.Message,), dict(
  DESCRIPTOR = _FLOAT64,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.Float64)
  ))
_sym_db.RegisterMessage(Float64)

Float64Array = _reflection.GeneratedProtocolMessageType('Float64Array', (_message.Message,), dict(
  DESCRIPTOR = _FLOAT64ARRAY,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.Float64Array)
  ))
_sym_db.RegisterMessage(Float64Array)

Bool = _reflection.GeneratedProtocolMessageType('Bool', (_message.Message,), dict(
  DESCRIPTOR = _BOOL,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.Bool)
  ))
_sym_db.RegisterMessage(Bool)

BoolArray = _reflection.GeneratedProtocolMessageType('BoolArray', (_message.Message,), dict(
  DESCRIPTOR = _BOOLARRAY,
  __module__ = 'ntypes_pb2'
  # @@protoc_insertion_point(class_scope:ntypes.BoolArray)
  ))
_sym_db.RegisterMessage(BoolArray)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\006ntypes'))
# @@protoc_insertion_point(module_scope)
