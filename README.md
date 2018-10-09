# ntypes [![GoDoc](https://godoc.org/github.com/piotrkowalczuk/ntypes?status.svg)](http://godoc.org/github.com/piotrkowalczuk/ntypes)

[![CircleCI](https://circleci.com/gh/piotrkowalczuk/ntypes/tree/master.svg?style=svg)](https://circleci.com/gh/piotrkowalczuk/ntypes/tree/master)
[![codecov.io](https://codecov.io/github/piotrkowalczuk/ntypes/coverage.svg?branch=master)](https://codecov.io/github/piotrkowalczuk/ntypes?branch=master)
[![Code Climate](https://codeclimate.com/github/piotrkowalczuk/ntypes/badges/gpa.svg)](https://codeclimate.com/github/piotrkowalczuk/ntypes)
[![Go Report Card](https://goreportcard.com/badge/github.com/piotrkowalczuk/ntypes)](https://goreportcard.com/report/github.com/piotrkowalczuk/ntypes)
[![pypi](https://img.shields.io/pypi/v/protobuf-ntypes.svg)](https://pypi.python.org/pypi/protobuf-ntypes)
[![Download](https://img.shields.io/bintray/v/piotrkowalczuk/maven/ntypes.svg)](https://bintray.com/piotrkowalczuk/maven/ntypes/_latestVersion)


Package provides set of types that helps to build complex protobuf messages that contains optional properties. 
API can be considered as stable.

## Types

### Basic 

* [ntypes.Bytes](https://godoc.org/github.com/piotrkowalczuk/ntypes#Bytes)
* [ntypes.String](https://godoc.org/github.com/piotrkowalczuk/ntypes#String)
* [ntypes.Bool](https://godoc.org/github.com/piotrkowalczuk/ntypes#Bool)
* [ntypes.Int32](https://godoc.org/github.com/piotrkowalczuk/ntypes#Int32)
* [ntypes.Int64](https://godoc.org/github.com/piotrkowalczuk/ntypes#Int64)
* [ntypes.Uint32](https://godoc.org/github.com/piotrkowalczuk/ntypes#Uint32)
* [ntypes.Uint64](https://godoc.org/github.com/piotrkowalczuk/ntypes#Uint64)
* [ntypes.Float32](https://godoc.org/github.com/piotrkowalczuk/ntypes#Float32)
* [ntypes.Float64](https://godoc.org/github.com/piotrkowalczuk/ntypes#Float64)

### Arrays

* [ntypes.BytesArray](https://godoc.org/github.com/piotrkowalczuk/ntypes#BytesArray)
* [ntypes.StringArray](https://godoc.org/github.com/piotrkowalczuk/ntypes#StringArray)
* [ntypes.BoolArray](https://godoc.org/github.com/piotrkowalczuk/ntypes#BoolArray)
* [ntypes.Int32Array](https://godoc.org/github.com/piotrkowalczuk/ntypes#Int32Array)
* [ntypes.Int64Array](https://godoc.org/github.com/piotrkowalczuk/ntypes#Int64Array)
* [ntypes.Uint32Array](https://godoc.org/github.com/piotrkowalczuk/ntypes#Uint32Array)
* [ntypes.Uint64Array](https://godoc.org/github.com/piotrkowalczuk/ntypes#Uint64Array)
* [ntypes.Float32Array](https://godoc.org/github.com/piotrkowalczuk/ntypes#Float32Array)
* [ntypes.Float64Array](https://godoc.org/github.com/piotrkowalczuk/ntypes#Float64Array)

Arrays support is not an SQL standard. Extra import needs to be added to make it working with [postgres driver](https://github.com/lib/pq):
 
 ```
 import _ "github.com/piotrkowalczuk/ntypes/ntypespq"
 ```

## Interfaces

Some type implements set of interfaces:

* [driver.Valuer](https://golang.org/pkg/database/sql/driver/#Valuer) excluding Uint64 and some arrays
* [proto.Message](https://godoc.org/github.com/golang/protobuf/proto#Message)
* [sql.Scanner](https://golang.org/pkg/database/sql/#Scanner) excluding Uint64 and some arrays
* [json.Marshaler](https://golang.org/pkg/encoding/json/#Marshaler)
* [json.Unmarshaler](https://golang.org/pkg/encoding/json/#Unmarshaler)

## Helpers

* [ntypes.True](https://godoc.org/github.com/piotrkowalczuk/ntypes#True)
* [ntypes.False](https://godoc.org/github.com/piotrkowalczuk/ntypes#False)

Besides that each type implements `func <Type>Or(<builtin>) <builtin> {}` method. 
It returns given argument if receiver is not valid or is nil. 
For example:
[ntypes.Bool.BoolOr](https://godoc.org/github.com/piotrkowalczuk/ntypes#Bool.BoolOr).
