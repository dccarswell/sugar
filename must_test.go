package sugar

import (
	"errors"
	"testing"
)

func TestMust_NoError(t *testing.T) {
	// Test that when no error is present, the value is returned unchanged
	result := Must("test value", nil)
	expected := "test value"

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMust_WithError(t *testing.T) {
	// Test that when an error is present, Must panics with that error
	testErr := errors.New("test error")

	defer func() {
		if r := recover(); r != nil {
			if r != testErr {
				t.Errorf("Expected panic with %v, got %v", testErr, r)
			}
		} else {
			t.Error("Expected function to panic, but it didn't")
		}
	}()

	Must("test value", testErr)
}

func TestMust_DifferentTypes(t *testing.T) {
	// Test Must with different generic types
	t.Run("int", func(t *testing.T) {
		result := Must(123, nil)
		if result != 123 {
			t.Errorf("Expected 123, got %d", result)
		}
	})

	t.Run("bool", func(t *testing.T) {
		result := Must(true, nil)
		if result != true {
			t.Errorf("Expected true, got %t", result)
		}
	})

	t.Run("float64", func(t *testing.T) {
		result := Must(3.14, nil)
		if result != 3.14 {
			t.Errorf("Expected 3.14, got %f", result)
		}
	})

	t.Run("slice", func(t *testing.T) {
		expected := []int{1, 2, 3}
		result := Must(expected, nil)
		if len(result) != len(expected) {
			t.Errorf("Expected slice length %d, got %d", len(expected), len(result))
		}
		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected element %d to be %d, got %d", i, v, result[i])
			}
		}
	})

	t.Run("map", func(t *testing.T) {
		expected := map[string]int{"a": 1, "b": 2}
		result := Must(expected, nil)
		if len(result) != len(expected) {
			t.Errorf("Expected map length %d, got %d", len(expected), len(result))
		}
		for k, v := range expected {
			if result[k] != v {
				t.Errorf("Expected map[%s] to be %d, got %d", k, v, result[k])
			}
		}
	})

	t.Run("struct", func(t *testing.T) {
		type TestStruct struct {
			Name string
			Age  int
		}
		expected := TestStruct{Name: "John", Age: 30}
		result := Must(expected, nil)
		if result != expected {
			t.Errorf("Expected %+v, got %+v", expected, result)
		}
	})

	t.Run("pointer", func(t *testing.T) {
		value := 42
		expected := &value
		result := Must(expected, nil)
		if result != expected {
			t.Errorf("Expected %p, got %p", expected, result)
		}
		if *result != *expected {
			t.Errorf("Expected pointer value %d, got %d", *expected, *result)
		}
	})
}

func TestMust_ZeroValues(t *testing.T) {
	// Test Must with zero values of different types
	t.Run("int", func(t *testing.T) {
		result := Must(0, nil)
		if result != 0 {
			t.Errorf("Expected 0, got %d", result)
		}
	})

	t.Run("string", func(t *testing.T) {
		var zeroString string
		result := Must(zeroString, nil)
		if result != "" {
			t.Errorf("Expected empty string, got %q", result)
		}
	})

	t.Run("bool", func(t *testing.T) {
		result := Must(false, nil)
		if result != false {
			t.Errorf("Expected false, got %t", result)
		}
	})

	t.Run("slice", func(t *testing.T) {
		var nilSlice []int
		result := Must(nilSlice, nil)
		if result != nil {
			t.Errorf("Expected nil slice, got %v", result)
		}
	})

	t.Run("map", func(t *testing.T) {
		var nilMap map[string]int
		result := Must(nilMap, nil)
		if result != nil {
			t.Errorf("Expected nil map, got %v", result)
		}
	})

	t.Run("pointer", func(t *testing.T) {
		var nilPtr *int
		result := Must(nilPtr, nil)
		if result != nil {
			t.Errorf("Expected nil pointer, got %v", result)
		}
	})
}

func TestMust_DifferentErrorTypes(t *testing.T) {
	// Test Must with different types of errors
	t.Run("simple_error", func(t *testing.T) {
		testErr := errors.New("simple error")

		defer func() {
			if r := recover(); r != testErr {
				t.Errorf("Expected panic with %v, got %v", testErr, r)
			}
		}()

		Must(42, testErr)
	})

	t.Run("wrapped_error", func(t *testing.T) {
		baseErr := errors.New("base error")
		wrappedErr := errors.Join(baseErr, errors.New("additional context"))

		defer func() {
			if r := recover(); r != wrappedErr {
				t.Errorf("Expected panic with %v, got %v", wrappedErr, r)
			}
		}()

		Must("test", wrappedErr)
	})
}

func TestMust_UsagePatterns(t *testing.T) {
	// Test common usage patterns
	t.Run("function_return", func(t *testing.T) {
		// Simulate a function that returns (value, error)
		getValue := func() (string, error) {
			return "success", nil
		}

		result := Must(getValue())
		if result != "success" {
			t.Errorf("Expected 'success', got %q", result)
		}
	})

	t.Run("function_return_with_error", func(t *testing.T) {
		testErr := errors.New("function failed")

		// Simulate a function that returns (value, error)
		getValue := func() (string, error) {
			return "", testErr
		}

		defer func() {
			if r := recover(); r != testErr {
				t.Errorf("Expected panic with %v, got %v", testErr, r)
			}
		}()

		Must(getValue())
	})

	t.Run("chaining", func(t *testing.T) {
		// Test that Must can be used in a chain
		getValue1 := func() (int, error) { return 5, nil }
		getValue2 := func(x int) (int, error) { return x * 2, nil }

		result := Must(getValue2(Must(getValue1())))
		expected := 10

		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

func TestMust_Interface(t *testing.T) {
	// Test Must with interface types
	t.Run("interface{}", func(t *testing.T) {
		var expected interface{} = "test string"
		result := Must(expected, nil)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("error_interface", func(t *testing.T) {
		// Test with error interface (but no error, so it returns the value)
		var expected error = nil
		result := Must(expected, nil)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

// Benchmark tests
func BenchmarkMust_NoError(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Must(42, nil)
	}
}

func BenchmarkMust_NoError_String(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Must("test string", nil)
	}
}

func BenchmarkMust_NoError_Slice(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Must(slice, nil)
	}
}

func BenchmarkMust_NoError_Struct(b *testing.B) {
	type TestStruct struct {
		Name string
		Age  int
	}
	s := TestStruct{Name: "John", Age: 30}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Must(s, nil)
	}
}
