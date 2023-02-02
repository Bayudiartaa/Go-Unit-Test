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
			name:    "Ama",
			request: "Ama",
		},
		{
			name:    "Jani",
			request: "Jani",
		},
		{
			name:    "Ama Jani",
			request: "Ama Jani",
		},
		{
			name:    "Budi",
			request: "Budi Nugraha",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Hello(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Ama", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Ama")
		}
	})
	b.Run("Ma", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Ma")
		}
	})
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Ama")
	}
}

func TestTableHello(t *testing.T) {
	tests := []struct {
		name string
		request string
		excepted string
	}{
		{
			name : "Ama",
			request : "Ama",
			excepted: "Hello Ama",
	    },
	    {
		   name : "Wawan",
		   request : "Wawan",
		   excepted: "Hello Wawan",
	    },
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			result := Hello(test.request)
			require.Equal(t, result, test.excepted)
		})
	}
}
func TestSubTest(t *testing.T) {
	t.Run("Ama", func(t *testing.T) {
		result := Hello("Ama")
	require.Equal(t, "Hello Ama", result, "Result must be 'Hello Ama'")
	})

	t.Run("wawan", func(t *testing.T) {
		result := Hello("wawan")
	require.Equal(t, "Hello wawan", result, "Result must be 'Hello wawan'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")
    m.Run()
	fmt.Println("After Unit Test")
}

func TestHello(t *testing.T) {
   result := Hello("bayu")	
   if result != "Hello bayu" {
		//error
		t.Error("result should be" + result)
   }

   fmt.Println("testHello Done")
}

func TestHelloAssertion(t *testing.T){
	result := Hello("bayu")
	assert.Equal(t, "Hello bayu", result, "Result must be 'Hello bayu'")
}

func TestHelloRequire(t *testing.T){
	result := Hello("bayu")
	require.Equal(t, "Hello bayu", result, "Result must be 'Hello bayu'")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}

	result := Hello("windows")
	require.Equal(t, "Hello windows", result, "Result must be 'Hello windows'")
}
