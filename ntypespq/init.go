package ntypespq

import (
	"database/sql/driver"

	"github.com/lib/pq"
	"github.com/piotrkowalczuk/ntypes"
)

func init() {
	ntypes.StringArrayScan = func(data interface{}) ([]string, error) {
		var arr pq.StringArray
		if err := arr.Scan(data); err != nil {
			return nil, err
		}
		return arr, nil
	}
	ntypes.StringArrayValue = func(arr []string) (driver.Value, error) {
		return pq.StringArray(arr).Value()
	}
	ntypes.Int64ArrayScan = func(data interface{}) ([]int64, error) {
		var arr pq.Int64Array
		if err := arr.Scan(data); err != nil {
			return nil, err
		}
		return arr, nil
	}
	ntypes.Int64ArrayValue = func(arr []int64) (driver.Value, error) {
		return pq.Int64Array(arr).Value()
	}
	ntypes.Float64ArrayScan = func(data interface{}) ([]float64, error) {
		var arr pq.Float64Array
		if err := arr.Scan(data); err != nil {
			return nil, err
		}
		return arr, nil
	}
	ntypes.Float64ArrayValue = func(arr []float64) (driver.Value, error) {
		return pq.Float64Array(arr).Value()
	}
	ntypes.BoolArrayScan = func(data interface{}) ([]bool, error) {
		var arr pq.BoolArray
		if err := arr.Scan(data); err != nil {
			return nil, err
		}
		return arr, nil
	}
	ntypes.BoolArrayValue = func(arr []bool) (driver.Value, error) {
		return pq.BoolArray(arr).Value()
	}
}
