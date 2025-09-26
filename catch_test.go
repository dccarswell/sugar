package sugar

import (
	"errors"
	"testing"
)

func TestHandle_NoError(t *testing.T) {
	// Test that when no error is present, the value is returned unchanged
	handler := Handle[string](func(err error) error {
		t.Fatal("Handler should not be called when there's no error")
		return err
	})

	result := handler("test value", nil)
	expected := "test value"

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestHandle_ErrorHandledSuccessfully(t *testing.T) {
	// Test that when an error is handled successfully (returns nil), the original value is returned
	originalErr := errors.New("original error")
	handlerCalled := false

	handler := Handle[int](func(err error) error {
		handlerCalled = true
		if err != originalErr {
			t.Errorf("Expected handler to receive %v, got %v", originalErr, err)
		}
		return nil // Handler successfully handles the error
	})

	result := handler(42, originalErr)
	expected := 42

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	if !handlerCalled {
		t.Error("Handler should have been called")
	}
}

func TestHandle_ErrorHandlerReturnsError(t *testing.T) {
	// Test that when the error handler returns an error, it panics
	originalErr := errors.New("original error")
	handlerErr := errors.New("handler error")
	handlerCalled := false

	handler := Handle[string](func(err error) error {
		handlerCalled = true
		if err != originalErr {
			t.Errorf("Expected handler to receive %v, got %v", originalErr, err)
		}
		return handlerErr // Handler returns a new error
	})

	defer func() {
		if r := recover(); r != nil {
			if r != handlerErr {
				t.Errorf("Expected panic with %v, got %v", handlerErr, r)
			}
		} else {
			t.Error("Expected function to panic, but it didn't")
		}
	}()

	handler("test", originalErr)

	if !handlerCalled {
		t.Error("Handler should have been called")
	}
}

func TestHandle_DifferentTypes(t *testing.T) {
	// Test with different generic types
	t.Run("int", func(t *testing.T) {
		handler := Handle[int](func(err error) error { return nil })
		result := handler(123, nil)
		if result != 123 {
			t.Errorf("Expected 123, got %d", result)
		}
	})

	t.Run("bool", func(t *testing.T) {
		handler := Handle[bool](func(err error) error { return nil })
		result := handler(true, nil)
		if result != true {
			t.Errorf("Expected true, got %t", result)
		}
	})

	t.Run("slice", func(t *testing.T) {
		handler := Handle[[]int](func(err error) error { return nil })
		expected := []int{1, 2, 3}
		result := handler(expected, nil)
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
		type TestStruct struct {
			Name string
			Age  int
		}
		handler := Handle[TestStruct](func(err error) error { return nil })
		expected := TestStruct{Name: "John", Age: 30}
		result := handler(expected, nil)
		if result != expected {
			t.Errorf("Expected %+v, got %+v", expected, result)
		}
	})
}

func TestHandle_ErrorTransformation(t *testing.T) {
	// Test that the handler can transform errors
	originalErr := errors.New("original error")
	transformedErr := errors.New("transformed error")

	handler := Handle[string](func(err error) error {
		if err.Error() == "original error" {
			return transformedErr
		}
		return err
	})

	defer func() {
		if r := recover(); r != nil {
			if r != transformedErr {
				t.Errorf("Expected panic with transformed error %v, got %v", transformedErr, r)
			}
		} else {
			t.Error("Expected function to panic with transformed error")
		}
	}()

	handler("test", originalErr)
}

func TestHandle_NilHandler(t *testing.T) {
	// Test behavior when handler is nil (should panic when called)
	var nilHandler Handler[string]

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when calling nil handler")
		}
	}()

	handleFunc := Handle[string](nilHandler)
	handleFunc("test", errors.New("test error"))
}

func TestHandle_ZeroValues(t *testing.T) {
	// Test with zero values
	handler := Handle[int](func(err error) error { return nil })

	result := handler(0, nil)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}

	var zeroString string
	stringHandler := Handle[string](func(err error) error { return nil })
	stringResult := stringHandler(zeroString, nil)
	if stringResult != "" {
		t.Errorf("Expected empty string, got %q", stringResult)
	}
}

func TestHandle_MultipleHandlers(t *testing.T) {
	// Test that multiple handlers can be created independently
	handler1 := Handle[int](func(err error) error { return nil })
	handler2 := Handle[string](func(err error) error { return errors.New("always fail") })

	// First handler should work normally
	result1 := handler1(100, nil)
	if result1 != 100 {
		t.Errorf("Expected 100, got %d", result1)
	}

	// Second handler should panic on error
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected second handler to panic")
		}
	}()

	handler2("test", errors.New("some error"))
}

// Benchmark tests
func BenchmarkHandle_NoError(b *testing.B) {
	handler := Handle[int](func(err error) error { return nil })

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler(42, nil)
	}
}

func BenchmarkHandle_WithErrorHandled(b *testing.B) {
	err := errors.New("test error")
	handler := Handle[int](func(e error) error { return nil })

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler(42, err)
	}
}
