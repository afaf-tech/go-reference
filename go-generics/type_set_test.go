package go_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int
type Number interface {
	// ~ means type approximation(allowing another type with type ~ able to join )
	~int | int8 | int16 | int32 | int64 | float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min[int](100, 200))
	assert.Equal(t, int64(100), Min[int64](int64(100), int64(200)))
	assert.Equal(t, float64(100), Min[float64](float64(100), float64(200)))
	// assert.Equal(t, 100, Min[string](100, 200)) not includes in type Number type set
	assert.Equal(t, Age(100), Min[Age](Age(100), Age(200))) // use ~ (tilde) in type Number
}
func TestMinTypeInference(t *testing.T) {
	assert.Equal(t, 100, Min(100, 200))
	assert.Equal(t, int64(100), Min(int64(100), int64(200)))
	assert.Equal(t, float64(100), Min(float64(100), float64(200)))
	// assert.Equal(t, 100, Min[string](100, 200)) not includes in type Number type set
	assert.Equal(t, Age(100), Min(Age(100), Age(200))) // use ~ (tilde) in type Number
}
