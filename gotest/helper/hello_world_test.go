package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "fikri",
			request: "fikri",
		},
		{
			name:    "al",
			request: "al",
		},
		{
			name:    "budi nugraha",
			request: "budi nugraha",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}

}

func BenchmarkSub(b *testing.B) {
	//sub benchmark
	b.Run("Fikri", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fikri")
		}
	})
	b.Run("Kurnia", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Kurnia")
		}
	})
}
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Akmal")
	}
}
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Fikri",
			request:  "Fikri",
			expected: "Hello Fikri",
		},
		{
			name:     "Akmal",
			request:  "Akmal",
			expected: "Hello Akmal",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
func TestSubTest(t *testing.T) {
	t.Run("Fikri", func(t *testing.T) {
		result := HelloWorld("Fikri")
		require.Equal(t, "Hello Fikri", result, "Result must be 'Hello Fikri'")
	})
	t.Run("Akmal", func(t *testing.T) {
		result := HelloWorld("Akmal")
		require.Equal(t, "Hello Akmal", result, "Result must be 'Hello Akmal'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Before test")

	m.Run()

	// after unit test
	fmt.Println("After test")
}

func TestSkipping(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("cannot run on linux")
	}

	result := HelloWorld("Farah")
	require.Equal(t, "Hello Farah", result)
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Fani")
	require.Equal(t, "Hello Fani", result, "Result must be 'Hello Fani'")
	fmt.Println("Test HelloWorld with Require") // will not be printed // call failnow
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Fani")
	assert.Equal(t, "Hello Fani", result, "Result must be 'Hello Fani'")
	fmt.Println("Test HelloWorld with Assert") // will be printed if failed -- call fail
}
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		// the next code is still executed
		//t.Fail()
		// t.Error give extra information about the error
		t.Error("Result must be Eko")
	}

	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldFikri(t *testing.T) {
	result := HelloWorld("Fikri")

	if result != "Hello Fikri" {
		// the next code will not be executed
		//t.FailNow()
		// t.Fatal give an extra information about the error
		t.Fatal("Result must be Fikri")
	}

	fmt.Println("TestHelloWorldFikri done")

}
