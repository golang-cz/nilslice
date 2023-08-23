package nilslice

import (
	"encoding/json"
	"testing"
)

func TestInitializeStruct(t *testing.T) {
	type X struct {
		Keys     []string `json:"keys"`
		Values   []**int  `json:"values"`
		Children []*X     `json:"children"`
	}

	example := &X{
		Children: []*X{
			&X{
				Children: []*X{
					&X{},
				},
			},
		},
	}

	origJSON := `{"keys":null,"values":null,"children":[{"keys":null,"values":null,"children":[{"keys":null,"values":null,"children":null}]}]}`

	expectedJSON := `{"keys":[],"values":[],"children":[{"keys":[],"values":[],"children":[{"keys":[],"values":[],"children":[]}]}]}`

	b, _ := json.Marshal(example)
	if string(b) != origJSON {
		t.Fatalf("unexpected original JSON: %s", string(b))
	}

	b, _ = json.Marshal(Initialize(example))
	if string(b) != expectedJSON {
		t.Fatalf("unexpected JSON: %s", string(b))
	}
}
