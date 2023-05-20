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
		request string
		name    string
	}{
		{
			name:    "Riko",
			request: "Riko",
		},
		{
			name:    "Pernando",
			request: "Pernando",
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
	b.Run("Riko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Riko")
		}
	})

	b.Run("Pernando", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Pernando")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Riko")
	}
}

func BenchmarkHelloIqbal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Iqbal")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Riko",
			request:  "Riko",
			expected: "Hello Riko",
		},
		{
			name:     "Pernando",
			request:  "Pernando",
			expected: "Hello Pernando",
		},
		{
			name:     "Permadi",
			request:  "Permadi",
			expected: "Hello Permadi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result, "Result should be "+test.expected)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Riko", func(t *testing.T) {
		result := HelloWorld("Riko")
		assert.Equal(t, "Hello Riko", result, "Result should be Hello Riko")
	})

	t.Run("Pernando", func(t *testing.T) {
		result := HelloWorld("Pernando")
		assert.Equal(t, "Hello Pernando", result, "Result should be Hello Pernando")
	})
}

func TestMain(m *testing.M) {
	// before unit test
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
	// after unit test
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't run on MAC")
	}

	result := HelloWorld("Riko")
	assert.Equal(t, "Hello Riko", result, "Result should be Hello Riko")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Riko")
	assert.Equal(t, "Hello Riko", result, "Result should be Hello Riko")
	fmt.Println("Testing HelloWorldAssertion Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Riko")
	require.Equal(t, "Hello Riko", result, "Result should be Hello Riko")
	fmt.Println("Testing HelloWorldRequire Done")
}

func TestHelloWorldRiko(t *testing.T) {
	result := HelloWorld("Riko")

	if result != "Hello Riko" {
		t.Error("Result should be Hello Riko")
	}
}

func TestHelloWorldPernando(t *testing.T) {
	result := HelloWorld("Riko Pernando")

	if result != "Hello Riko Pernando" {
		t.Error("Result should be Hello Pernando")
	}
}
