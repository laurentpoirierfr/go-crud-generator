package helpers_test

import (
	"testing"

	"github.com/laurentpoirierfr/crud-generator-go/pkg/helpers"
	"gotest.tools/v3/assert"
)

// Exemple de structures
type StructA struct {
	Name  string
	Age   int
	Email string
}

type StructB struct {
	Name  string
	Age   int
	Phone string
}

func TestCopyStruct(t *testing.T) {
	a := &StructA{Name: "Alice", Age: 30, Email: "alice@example.com"}
	b := &StructB{Name: "Bob", Phone: "123-456-7890"}

	t.Logf("Avant copie: %+v\n", b)

	err := helpers.CopyStruct(a, b)

	assert.Equal(t, a.Age, b.Age, "they should be equal")
	assert.Equal(t, a.Name, b.Name, "they should be equal")

	if err != nil {
		t.Error("Erreur:", err)
	}

	t.Logf("Après copie: %+v\n", b)

}
