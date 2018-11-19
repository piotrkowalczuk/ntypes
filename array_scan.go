package ntypes

import (
	"database/sql/driver"
	"errors"
)

var (
	// StringArrayScan is used by StringArray Scan method.
	StringArrayScan = func(interface{}) ([]string, error) {
		return nil, errors.New("ntypes: StringArrayScan not set")
	}
	// StringArrayValue is used by StringArray Value method.
	StringArrayValue = func([]string) (driver.Value, error) {
		return nil, errors.New("ntypes: StringArrayValue not set")
	}
	// DoubleArrayScan is used by DoubleArray Scan method.
	DoubleArrayScan = func(interface{}) ([]float64, error) {
		return nil, errors.New("ntypes: DoubleArrayScan not set")
	}
	// DoubleArrayValue is used by DoubleArray Value method.
	DoubleArrayValue = func([]float64) (driver.Value, error) {
		return nil, errors.New("ntypes: DoubleArrayValue not set")
	}
	// Int64ArrayScan is used by Int64Array Scan method.
	Int64ArrayScan = func(interface{}) ([]int64, error) {
		return nil, errors.New("ntypes: DoubleArrayValue not set")
	}
	// Int64ArrayValue is used by Int64Array Value method.
	Int64ArrayValue = func([]int64) (driver.Value, error) {
		return nil, errors.New("ntypes: Int64ArrayValue not set")
	}
	// BoolArrayScan is used by BoolArray Scan method.
	BoolArrayScan = func(interface{}) ([]bool, error) {
		return nil, errors.New("ntypes: BoolArrayScan not set")
	}
	// BoolArrayValue is used by BoolArray Value method.
	BoolArrayValue = func([]bool) (driver.Value, error) {
		return nil, errors.New("ntypes: BoolArrayValue not set")
	}
	// BytesArrayScan is used by BytesArray Scan method.
	BytesArrayScan = func(interface{}) ([][]byte, error) {
		return nil, errors.New("ntypes: BytesArrayScan not set")
	}
	// ByteArrayValue is used by BytesArray Value method.
	BytesArrayValue = func([][]byte) (driver.Value, error) {
		return nil, errors.New("ntypes: BytesArrayValue not set")
	}
)
