package sugar

import (
	"testing"
)

func TestPtr_BasicTypes(t *testing.T) {
	// Test Ptr with basic types
	t.Run("int", func(t *testing.T) {
		value := 42
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != value {
			t.Errorf("Expected pointed value %d, got %d", value, *result)
		}

		// Verify it's a different memory location than the original
		value = 100
		if *result == value {
			t.Error("Pointer should point to a copy, not the original variable")
		}
	})

	t.Run("string", func(t *testing.T) {
		value := "test string"
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != value {
			t.Errorf("Expected pointed value %q, got %q", value, *result)
		}

		// Verify it's a different memory location than the original
		value = "modified"
		if *result == value {
			t.Error("Pointer should point to a copy, not the original variable")
		}
	})

	t.Run("bool", func(t *testing.T) {
		value := true
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != value {
			t.Errorf("Expected pointed value %t, got %t", value, *result)
		}
	})

	t.Run("float64", func(t *testing.T) {
		value := 3.14159
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != value {
			t.Errorf("Expected pointed value %f, got %f", value, *result)
		}
	})

	t.Run("byte", func(t *testing.T) {
		value := byte(255)
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != value {
			t.Errorf("Expected pointed value %d, got %d", value, *result)
		}
	})

	t.Run("rune", func(t *testing.T) {
		value := rune('A')
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != value {
			t.Errorf("Expected pointed value %c, got %c", value, *result)
		}
	})
}

func TestPtr_ZeroValues(t *testing.T) {
	// Test Ptr with zero values
	t.Run("int", func(t *testing.T) {
		value := 0
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != 0 {
			t.Errorf("Expected pointed value 0, got %d", *result)
		}
	})

	t.Run("string", func(t *testing.T) {
		value := ""
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != "" {
			t.Errorf("Expected pointed value empty string, got %q", *result)
		}
	})

	t.Run("bool", func(t *testing.T) {
		value := false
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != false {
			t.Errorf("Expected pointed value false, got %t", *result)
		}
	})

	t.Run("float64", func(t *testing.T) {
		value := 0.0
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != 0.0 {
			t.Errorf("Expected pointed value 0.0, got %f", *result)
		}
	})
}

func TestPtr_CompositeTypes(t *testing.T) {
	// Test Ptr with composite types
	t.Run("slice", func(t *testing.T) {
		value := []int{1, 2, 3}
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}

		// Compare slice contents
		if len(*result) != len(value) {
			t.Errorf("Expected slice length %d, got %d", len(value), len(*result))
		}
		for i, v := range value {
			if (*result)[i] != v {
				t.Errorf("Expected element %d to be %d, got %d", i, v, (*result)[i])
			}
		}
	})

	t.Run("map", func(t *testing.T) {
		value := map[string]int{"a": 1, "b": 2}
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}

		// Compare map contents
		if len(*result) != len(value) {
			t.Errorf("Expected map length %d, got %d", len(value), len(*result))
		}
		for k, v := range value {
			if (*result)[k] != v {
				t.Errorf("Expected map[%s] to be %d, got %d", k, v, (*result)[k])
			}
		}
	})

	t.Run("struct", func(t *testing.T) {
		type TestStruct struct {
			Name string
			Age  int
		}
		value := TestStruct{Name: "John", Age: 30}
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != value {
			t.Errorf("Expected pointed value %+v, got %+v", value, *result)
		}
	})

	t.Run("array", func(t *testing.T) {
		value := [3]int{1, 2, 3}
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != value {
			t.Errorf("Expected pointed value %v, got %v", value, *result)
		}
	})

	t.Run("nil_slice", func(t *testing.T) {
		var value []int = nil
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != nil {
			t.Errorf("Expected pointed value nil, got %v", *result)
		}
	})

	t.Run("nil_map", func(t *testing.T) {
		var value map[string]int = nil
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != nil {
			t.Errorf("Expected pointed value nil, got %v", *result)
		}
	})
}

func TestPtr_Pointers(t *testing.T) {
	// Test Ptr with pointer types
	t.Run("pointer_to_int", func(t *testing.T) {
		num := 42
		value := &num
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != value {
			t.Errorf("Expected pointed value %p, got %p", value, *result)
		}
		if **result != *value {
			t.Errorf("Expected double-dePtrerenced value %d, got %d", *value, **result)
		}
	})

	t.Run("nil_pointer", func(t *testing.T) {
		var value *int = nil
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != nil {
			t.Errorf("Expected pointed value nil, got %v", *result)
		}
	})
}

func TestPtr_Interface(t *testing.T) {
	// Test Ptr with interface types
	t.Run("interface{}", func(t *testing.T) {
		var value interface{} = "test string"
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != value {
			t.Errorf("Expected pointed value %v, got %v", value, *result)
		}
	})

	t.Run("error_interface", func(t *testing.T) {
		var value error = nil
		result := Ptr(value)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != nil {
			t.Errorf("Expected pointed value nil, got %v", *result)
		}
	})
}

func TestPtr_Literals(t *testing.T) {
	// Test Ptr with literals (common use case)
	t.Run("int_literal", func(t *testing.T) {
		result := Ptr(42)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != 42 {
			t.Errorf("Expected pointed value 42, got %d", *result)
		}
	})

	t.Run("string_literal", func(t *testing.T) {
		result := Ptr("hello")

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != "hello" {
			t.Errorf("Expected pointed value 'hello', got %q", *result)
		}
	})

	t.Run("bool_literal", func(t *testing.T) {
		result := Ptr(true)

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != true {
			t.Errorf("Expected pointed value true, got %t", *result)
		}
	})

	t.Run("struct_literal", func(t *testing.T) {
		type Point struct {
			X, Y int
		}
		result := Ptr(Point{X: 10, Y: 20})

		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		expected := Point{X: 10, Y: 20}
		if *result != expected {
			t.Errorf("Expected pointed value %+v, got %+v", expected, *result)
		}
	})
}

func TestPtr_Modification(t *testing.T) {
	// Test that modifying the pointed value works correctly
	t.Run("modify_through_pointer", func(t *testing.T) {
		result := Ptr(42)

		*result = 100

		if *result != 100 {
			t.Errorf("Expected modified value 100, got %d", *result)
		}
	})

	t.Run("modify_slice_through_pointer", func(t *testing.T) {
		result := Ptr([]int{1, 2, 3})

		(*result)[0] = 999

		if (*result)[0] != 999 {
			t.Errorf("Expected modified slice element 999, got %d", (*result)[0])
		}
	})

	t.Run("modify_struct_through_pointer", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		result := Ptr(Person{Name: "Alice", Age: 25})

		result.Name = "Bob"
		result.Age = 30

		if result.Name != "Bob" || result.Age != 30 {
			t.Errorf("Expected modified struct {Name: Bob, Age: 30}, got %+v", *result)
		}
	})
}

func TestPtr_UsagePatterns(t *testing.T) {
	// Test common usage patterns
	t.Run("optional_field", func(t *testing.T) {
		// Common pattern: creating pointers for optional struct fields
		type Config struct {
			Name    string
			Timeout *int
		}

		config := Config{
			Name:    "test",
			Timeout: Ptr(30), // Instead of &30 which doesn't work
		}

		if config.Timeout == nil {
			t.Error("Expected non-nil Timeout")
		}
		if *config.Timeout != 30 {
			t.Errorf("Expected Timeout 30, got %d", *config.Timeout)
		}
	})

	t.Run("function_parameter", func(t *testing.T) {
		// Function that takes a pointer parameter
		updateValue := func(ptr *string) {
			*ptr = "updated"
		}

		// Use Ptr to pass a literal
		result := Ptr("original")
		updateValue(result)

		if *result != "updated" {
			t.Errorf("Expected 'updated', got %q", *result)
		}
	})

	t.Run("comparison_with_manual_reference", func(t *testing.T) {
		// Compare Ptr() with manual variable + &
		value1 := 42
		ptr1 := &value1
		ptr2 := Ptr(42)

		// Both should point to the same value
		if *ptr1 != *ptr2 {
			t.Errorf("Expected same values: %d vs %d", *ptr1, *ptr2)
		}

		// But they should be different pointers
		if ptr1 == ptr2 {
			t.Error("Pointers should be different")
		}
	})
}

// Benchmark tests
func BenchmarkPtr_Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Ptr(42)
	}
}

func BenchmarkPtr_String(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Ptr("test string")
	}
}

func BenchmarkPtr_Struct(b *testing.B) {
	type TestStruct struct {
		Name string
		Age  int
	}
	s := TestStruct{Name: "John", Age: 30}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Ptr(s)
	}
}

func BenchmarkPtr_Slice(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Ptr(slice)
	}
}

func BenchmarkPtr_Map(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Ptr(m)
	}
}
