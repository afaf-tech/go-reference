package go_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, FindMin[int](100, 200))
	assert.Equal(t, int64(100), FindMin[int64](int64(100), int64(200)))
}

func TestExperimentalMap(t *testing.T) {
	first := map[string]string{
		"Name": "Fikri",
	}
	second := map[string]string{
		"Name": "Fikri",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlices(t *testing.T) {
	first := []string{"Fikri"}
	second := []string{"Fikri"}

	assert.True(t, slices.Equal(first, second))
}
