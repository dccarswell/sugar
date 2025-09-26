package sugar

import (
	"errors"
	"fmt"
	"testing"
)

func TestTry_NoError(t *testing.T) {
	// Test that when the function executes successfully, the value is returned and no error
	t.Run("int", func(t *testing.T) {
		result, err := Try(func() int {
			return 42
		})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 42 {
			t.Errorf("Expected result 42, got %d", result)
		}
	})

	t.Run("string", func(t *testing.T) {
		result, err := Try(func() string {
			return "hello world"
		})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != "hello world" {
			t.Errorf("Expected result 'hello world', got %q", result)
		}
	})

	t.Run("bool", func(t *testing.T) {
		result, err := Try(func() bool {
			return true
		})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != true {
			t.Errorf("Expected result true, got %t", result)
		}
	})

	t.Run("slice", func(t *testing.T) {
		expected := []int{1, 2, 3}
		result, err := Try(func() []int {
			return expected
		})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(result) != len(expected) {
			t.Errorf("Expected slice length %d, got %d", len(expected), len(result))
		}
		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected element %d to be %d, got %d", i, v, result[i])
			}
		}
	})

	t.Run("struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		expected := Person{Name: "Alice", Age: 30}
		result, err := Try(func() Person {
			return expected
		})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})
}

func TestTry_WithPanic(t *testing.T) {
	// Test that when the function panics, an error is returned and zero value is returned
	t.Run("panic_with_string", func(t *testing.T) {
		result, err := Try(func() int {
			panic("something went wrong")
		})

		if err == nil {
			t.Error("Expected error, got nil")
		}
		if result != 0 {
			t.Errorf("Expected zero value (0), got %d", result)
		}

		expectedErrorMsg := "panic: something went wrong"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Expected error message %q, got %q", expectedErrorMsg, err.Error())
		}
	})

	t.Run("panic_with_error", func(t *testing.T) {
		panicErr := errors.New("original error")
		result, err := Try(func() string {
			panic(panicErr)
		})

		if err == nil {
			t.Error("Expected error, got nil")
		}
		if result != "" {
			t.Errorf("Expected zero value (empty string), got %q", result)
		}

		expectedErrorMsg := fmt.Sprintf("panic: %v", panicErr)
		if err.Error() != expectedErrorMsg {
			t.Errorf("Expected error message %q, got %q", expectedErrorMsg, err.Error())
		}
	})

	t.Run("panic_with_int", func(t *testing.T) {
		result, err := Try(func() bool {
			panic(123)
		})

		if err == nil {
			t.Error("Expected error, got nil")
		}
		if result != false {
			t.Errorf("Expected zero value (false), got %t", result)
		}

		expectedErrorMsg := "panic: 123"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Expected error message %q, got %q", expectedErrorMsg, err.Error())
		}
	})

	t.Run("panic_with_nil", func(t *testing.T) {
		result, err := Try(func() []int {
			panic(nil)
		})

		if err == nil {
			t.Error("Expected error, got nil")
		}
		if result != nil {
			t.Errorf("Expected zero value (nil slice), got %v", result)
		}

		// Go's panic(nil) has special behavior
		expectedErrorMsg := "panic: panic called with nil argument"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Expected error message %q, got %q", expectedErrorMsg, err.Error())
		}
	})
}

func TestTry_ZeroValues(t *testing.T) {
	// Test zero values are returned correctly on panic
	t.Run("int_zero", func(t *testing.T) {
		result, err := Try(func() int {
			panic("test")
		})

		if err == nil {
			t.Error("Expected error, got nil")
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("string_zero", func(t *testing.T) {
		result, err := Try(func() string {
			panic("test")
		})

		if err == nil {
			t.Error("Expected error, got nil")
		}
		if result != "" {
			t.Errorf("Expected zero value empty string, got %q", result)
		}
	})

	t.Run("bool_zero", func(t *testing.T) {
		result, err := Try(func() bool {
			panic("test")
		})

		if err == nil {
			t.Error("Expected error, got nil")
		}
		if result != false {
			t.Errorf("Expected zero value false, got %t", result)
		}
	})

	t.Run("slice_zero", func(t *testing.T) {
		result, err := Try(func() []int {
			panic("test")
		})

		if err == nil {
			t.Error("Expected error, got nil")
		}
		if result != nil {
			t.Errorf("Expected zero value nil, got %v", result)
		}
	})

	t.Run("map_zero", func(t *testing.T) {
		result, err := Try(func() map[string]int {
			panic("test")
		})

		if err == nil {
			t.Error("Expected error, got nil")
		}
		if result != nil {
			t.Errorf("Expected zero value nil, got %v", result)
		}
	})

	t.Run("pointer_zero", func(t *testing.T) {
		result, err := Try(func() *int {
			panic("test")
		})

		if err == nil {
			t.Error("Expected error, got nil")
		}
		if result != nil {
			t.Errorf("Expected zero value nil, got %v", result)
		}
	})

	t.Run("struct_zero", func(t *testing.T) {
		type TestStruct struct {
			Name string
			Age  int
		}
		result, err := Try(func() TestStruct {
			panic("test")
		})

		if err == nil {
			t.Error("Expected error, got nil")
		}
		expected := TestStruct{} // zero value
		if result != expected {
			t.Errorf("Expected zero value %+v, got %+v", expected, result)
		}
	})
}

func TestTry_ComplexPanicScenarios(t *testing.T) {
	// Test various panic scenarios
	t.Run("panic_in_nested_call", func(t *testing.T) {
		nestedFunc := func() int {
			panic("nested panic")
		}

		result, err := Try(func() int {
			return nestedFunc()
		})

		if err == nil {
			t.Error("Expected error, got nil")
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("panic_after_partial_execution", func(t *testing.T) {
		counter := 0
		result, err := Try(func() string {
			counter++
			if counter > 0 {
				panic("panic after increment")
			}
			return "success"
		})

		if err == nil {
			t.Error("Expected error, got nil")
		}
		if result != "" {
			t.Errorf("Expected zero value empty string, got %q", result)
		}
		// Verify that the counter was incremented before panic
		if counter != 1 {
			t.Errorf("Expected counter to be 1, got %d", counter)
		}
	})

	t.Run("multiple_try_calls", func(t *testing.T) {
		// First call succeeds
		result1, err1 := Try(func() int { return 10 })
		if err1 != nil || result1 != 10 {
			t.Errorf("First call failed: result=%d, err=%v", result1, err1)
		}

		// Second call panics
		result2, err2 := Try(func() int { panic("second call panic") })
		if err2 == nil || result2 != 0 {
			t.Errorf("Second call should have failed: result=%d, err=%v", result2, err2)
		}

		// Third call succeeds
		result3, err3 := Try(func() int { return 20 })
		if err3 != nil || result3 != 20 {
			t.Errorf("Third call failed: result=%d, err=%v", result3, err3)
		}
	})
}

func TestTry_UsagePatterns(t *testing.T) {
	// Test common usage patterns
	t.Run("division_by_zero", func(t *testing.T) {
		divide := func(a, b int) int {
			if b == 0 {
				panic("division by zero")
			}
			return a / b
		}

		// Normal division
		result, err := Try(func() int {
			return divide(10, 2)
		})
		if err != nil || result != 5 {
			t.Errorf("Normal division failed: result=%d, err=%v", result, err)
		}

		// Division by zero
		result, err = Try(func() int {
			return divide(10, 0)
		})
		if err == nil || result != 0 {
			t.Errorf("Division by zero should have failed: result=%d, err=%v", result, err)
		}
	})

	t.Run("array_bounds_check", func(t *testing.T) {
		arr := []int{1, 2, 3}

		// Valid access
		result, err := Try(func() int {
			return arr[1]
		})
		if err != nil || result != 2 {
			t.Errorf("Valid access failed: result=%d, err=%v", result, err)
		}

		// Out of bounds access
		result, err = Try(func() int {
			return arr[10] // This will panic
		})
		if err == nil || result != 0 {
			t.Errorf("Out of bounds access should have failed: result=%d, err=%v", result, err)
		}
	})

	t.Run("nil_pointer_dereference", func(t *testing.T) {
		var ptr *int = nil

		result, err := Try(func() int {
			return *ptr // This will panic
		})

		if err == nil {
			t.Error("Expected error for nil pointer dereference")
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("type_assertion_panic", func(t *testing.T) {
		var val interface{} = "string value"

		result, err := Try(func() int {
			return val.(int) // This will panic - wrong type
		})

		if err == nil {
			t.Error("Expected error for failed type assertion")
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})
}

func TestTry_FunctionTypes(t *testing.T) {
	// Test with different function signatures
	t.Run("no_params_function", func(t *testing.T) {
		getValue := func() string { return "test" }

		result, err := Try(getValue)
		if err != nil || result != "test" {
			t.Errorf("No params function failed: result=%q, err=%v", result, err)
		}
	})

	t.Run("closure_with_capture", func(t *testing.T) {
		x := 42
		result, err := Try(func() int {
			return x * 2
		})
		if err != nil || result != 84 {
			t.Errorf("Closure failed: result=%d, err=%v", result, err)
		}
	})

	t.Run("anonymous_function", func(t *testing.T) {
		result, err := Try(func() float64 {
			return 3.14159
		})
		if err != nil || result != 3.14159 {
			t.Errorf("Anonymous function failed: result=%f, err=%v", result, err)
		}
	})
}

// Benchmark tests
func BenchmarkTry_NoError(b *testing.B) {
	f := func() int { return 42 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Try(f)
	}
}

func BenchmarkTry_WithPanic(b *testing.B) {
	f := func() int { panic("benchmark panic") }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Try(f)
	}
}

func BenchmarkTry_ComplexOperation(b *testing.B) {
	f := func() []int {
		result := make([]int, 1000)
		for i := range result {
			result[i] = i * 2
		}
		return result
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Try(f)
	}
}

func BenchmarkTry_vs_DirectCall(b *testing.B) {
	f := func() int { return 42 }

	b.Run("with_try", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Try(f)
		}
	})

	b.Run("direct_call", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f()
		}
	})
}
