# ntypes [![GoDoc](https://godoc.org/github.com/piotrkowalczuk/ntypes?status.svg)](http://godoc.org/github.com/piotrkowalczuk/ntypes)&nbsp;[![Build Status](https://travis-ci.org/piotrkowalczuk/ntypes.svg?branch=master)](https://travis-ci.org/piotrkowalczuk/ntypes)&nbsp;[![codecov.io](https://codecov.io/github/piotrkowalczuk/ntypes/coverage.svg?branch=master)](https://codecov.io/github/piotrkowalczuk/ntypes?branch=master)&nbsp;[![Code Climate](https://codeclimate.com/github/piotrkowalczuk/ntypes/badges/gpa.svg)](https://codeclimate.com/github/piotrkowalczuk/ntypes)&nbsp;[![Go Report Card](https://goreportcard.com/badge/github.com/piotrkowalczuk/ntypes)](https://goreportcard.com/report/github.com/piotrkowalczuk/ntypes)
Package provides set of types that helps to build complex protobuf messages that contains optional properties.

## Types

* [ntypes.String](https://godoc.org/github.com/piotrkowalczuk/ntypes#String)
* [ntypes.Bool](https://godoc.org/github.com/piotrkowalczuk/ntypes#Bool)
* [ntypes.Int](https://godoc.org/github.com/piotrkowalczuk/ntypes#Int)
* [ntypes.Int32](https://godoc.org/github.com/piotrkowalczuk/ntypes#Int32)
* [ntypes.Int64](https://godoc.org/github.com/piotrkowalczuk/ntypes#Int64)
* [ntypes.Uint32](https://godoc.org/github.com/piotrkowalczuk/ntypes#Uint32)
* [ntypes.Float32](https://godoc.org/github.com/piotrkowalczuk/ntypes#Float32)
* [ntypes.Float64](https://godoc.org/github.com/piotrkowalczuk/ntypes#Float64)


## Interfaces

Each type implements set of interfaces:

* [driver.Valuer](https://golang.org/pkg/database/sql/driver/#Valuer)
* [proto.Message](https://godoc.org/github.com/golang/protobuf/proto#Message)
* [sql.Scanner](https://golang.org/pkg/database/sql/#Scanner)
* [json.Marshaler](https://golang.org/pkg/encoding/json/#Marshaler)
* [json.Unmarshaler](https://golang.org/pkg/encoding/json/#Unmarshaler)


## Helpers

* [ntypes.True](https://godoc.org/github.com/piotrkowalczuk/ntypes#True)
* [ntypes.False](https://godoc.org/github.com/piotrkowalczuk/ntypes#False)

Besides that each type implementss `func <Type>Or(<builtin>) <builtin> {}` method. It returns given argument if receiver is not valid or is nil. For example:
[ntypes.Bool.BoolOr](https://godoc.org/github.com/piotrkowalczuk/ntypes#Bool.BoolOr).
