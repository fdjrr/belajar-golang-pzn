package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit Test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Herlambang")

	if result != "Hello Herlambang" {
		t.Error("Result must be 'Hello Herlambang'")
	}
}

func TestHelloWorldFadjrir(t *testing.T) {
	result := HelloWorld("Fadjrir")

	if result != "Hello Fadjrir" {
		t.Fatal("Result must be 'Hello Fadjrir'")
	}
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Herlambang")

	require.Equal(t, "Hello Herlambang", result, "Result must be 'Hello Herlambang'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Herlambang")

	assert.Equal(t, "Hello Herlambang", result, "Result must be 'Hello Herlambang'")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't run on Mac")
	}

	result := HelloWorld("Herlambang")

	assert.Equal(t, "Hello Herlambang", result, "Result must be 'Hello Herlambang'")
}

func TestSubTest(t *testing.T) {
	t.Run("Herlambang", func(t *testing.T) {
		result := HelloWorld("Herlambang")

		assert.Equal(t, "Hello Herlambang", result, "Result must be 'Hello Herlambang'")
	})

	t.Run("Fadjrir", func(t *testing.T) {
		result := HelloWorld("Fadjrir")

		assert.Equal(t, "Hello Fadjrir", result, "Result must be 'Hello Fadjrir'")
	})
}

func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Herlambang",
			request:  "Herlambang",
			expected: "Hello Herlambang",
		},
		{
			name:     "Fadjrir",
			request:  "Fadjrir",
			expected: "Hello Fadjrir",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)

			assert.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloHerlambang(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Herlambang")
	}
}

func BenchmarkHelloFadjrir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Fadjrir")
	}
}

func BenchmarkHelloSub(b *testing.B) {
	b.Run("Herlambang", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Herlambang")
		}
	})

	b.Run("Fadjrir", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fadjrir")
		}
	})
}

func BenchmarkHelloTable(b *testing.B) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Herlambang",
			request:  "Herlambang",
			expected: "Hello Herlambang",
		},
		{
			name:     "Fadjrir",
			request:  "Fadjrir",
			expected: "Hello Fadjrir",
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(test.request)
			}
		})
	}
}
