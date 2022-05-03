package go_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

// type parameter inheritance
func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePrecident interface {
	GetName() string
	GetVicePrecidentName() string
}

type MyVicePrecident struct {
	Name string
}

func (m *MyVicePrecident) GetName() string {
	return m.Name
}

func (m *MyVicePrecident) GetVicePrecidentName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, GetName[Manager](&MyManager{Name: "Fikri"}), "Fikri")
	assert.Equal(t, GetName[VicePrecident](&MyVicePrecident{Name: "Fikri"}), "Fikri")
	// assert.Equal(t, GetName[string]("Fikri"), "Fikri") // not allowed due to not inherit from Employee or has no method GetMember

}
