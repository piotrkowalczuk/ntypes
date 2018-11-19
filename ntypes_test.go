package ntypes_test

import (
	"encoding/json"
	"testing"

	"github.com/piotrkowalczuk/ntypes"
)

// TODO: implement
// This test should prove that omitempty works as expected.
// Currently there is no way to omit struct.
// LINK: https://github.com/golang/go/issues/11939
func TestMarshalJSON(t *testing.T) {
	x := struct {
		BoolVal ntypes.Bool  `json:",omitempty"`
		BoolPtr *ntypes.Bool `json:",omitempty"`
	}{
		BoolVal: ntypes.Bool{},
		BoolPtr: &ntypes.Bool{},
	}

	buf, err := json.Marshal(&x)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(buf))

}
