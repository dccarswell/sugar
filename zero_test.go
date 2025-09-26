package sugar

import (
	"testing"
	"unsafe"
)

func TestZero_BasicTypes(t *testing.T) {
	// Test Zero with basic types
	t.Run("int", func(t *testing.T) {
		result := Zero[int]()
		expected := 0
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("int8", func(t *testing.T) {
		result := Zero[int8]()
		expected := int8(0)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("int16", func(t *testing.T) {
		result := Zero[int16]()
		expected := int16(0)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("int32", func(t *testing.T) {
		result := Zero[int32]()
		expected := int32(0)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("int64", func(t *testing.T) {
		result := Zero[int64]()
		expected := int64(0)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("uint", func(t *testing.T) {
		result := Zero[uint]()
		expected := uint(0)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("uint8", func(t *testing.T) {
		result := Zero[uint8]()
		expected := uint8(0)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("uint16", func(t *testing.T) {
		result := Zero[uint16]()
		expected := uint16(0)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("uint32", func(t *testing.T) {
		result := Zero[uint32]()
		expected := uint32(0)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("uint64", func(t *testing.T) {
		result := Zero[uint64]()
		expected := uint64(0)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("uintptr", func(t *testing.T) {
		result := Zero[uintptr]()
		expected := uintptr(0)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("float32", func(t *testing.T) {
		result := Zero[float32]()
		expected := float32(0.0)
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("float64", func(t *testing.T) {
		result := Zero[float64]()
		expected := 0.0
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("complex64", func(t *testing.T) {
		result := Zero[complex64]()
		expected := complex64(0)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("complex128", func(t *testing.T) {
		result := Zero[complex128]()
		expected := complex128(0)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("string", func(t *testing.T) {
		result := Zero[string]()
		expected := ""
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run("bool", func(t *testing.T) {
		result := Zero[bool]()
		expected := false
		if result != expected {
			t.Errorf("Expected %t, got %t", expected, result)
		}
	})

	t.Run("byte", func(t *testing.T) {
		result := Zero[byte]()
		expected := byte(0)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("rune", func(t *testing.T) {
		result := Zero[rune]()
		expected := rune(0)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

func TestZero_CompositeTypes(t *testing.T) {
	// Test Zero with composite types
	t.Run("slice", func(t *testing.T) {
		result := Zero[[]int]()
		if result != nil {
			t.Errorf("Expected nil slice, got %v", result)
		}
	})

	t.Run("map", func(t *testing.T) {
		result := Zero[map[string]int]()
		if result != nil {
			t.Errorf("Expected nil map, got %v", result)
		}
	})

	t.Run("channel", func(t *testing.T) {
		result := Zero[chan int]()
		if result != nil {
			t.Errorf("Expected nil channel, got %v", result)
		}
	})

	t.Run("function", func(t *testing.T) {
		result := Zero[func() int]()
		if result != nil {
			t.Errorf("Expected nil function, got non-nil function")
		}
	})

	t.Run("pointer", func(t *testing.T) {
		result := Zero[*int]()
		if result != nil {
			t.Errorf("Expected nil pointer, got %v", result)
		}
	})

	t.Run("interface", func(t *testing.T) {
		result := Zero[interface{}]()
		if result != nil {
			t.Errorf("Expected nil interface{}, got %v", result)
		}
	})

	t.Run("error_interface", func(t *testing.T) {
		result := Zero[error]()
		if result != nil {
			t.Errorf("Expected nil error, got %v", result)
		}
	})

	t.Run("array", func(t *testing.T) {
		result := Zero[[3]int]()
		expected := [3]int{0, 0, 0}
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("struct", func(t *testing.T) {
		type TestStruct struct {
			Name string
			Age  int
		}
		result := Zero[TestStruct]()
		expected := TestStruct{Name: "", Age: 0}
		if result != expected {
			t.Errorf("Expected %+v, got %+v", expected, result)
		}
	})
}

func TestZero_PointerTypes(t *testing.T) {
	// Test Zero with various pointer types
	t.Run("pointer_to_int", func(t *testing.T) {
		result := Zero[*int]()
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	t.Run("pointer_to_string", func(t *testing.T) {
		result := Zero[*string]()
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	t.Run("pointer_to_struct", func(t *testing.T) {
		type Person struct {
			Name string
		}
		result := Zero[*Person]()
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	t.Run("double_pointer", func(t *testing.T) {
		result := Zero[**int]()
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	t.Run("unsafe_pointer", func(t *testing.T) {
		result := Zero[unsafe.Pointer]()
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})
}

func TestZero_CustomTypes(t *testing.T) {
	// Test Zero with custom defined types
	t.Run("type_alias", func(t *testing.T) {
		type MyInt int
		result := Zero[MyInt]()
		expected := MyInt(0)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("custom_struct", func(t *testing.T) {
		type Config struct {
			Host     string
			Port     int
			Enabled  bool
			Timeout  float64
			Settings map[string]interface{}
		}
		result := Zero[Config]()

		// Check each field individually since struct contains uncomparable map
		if result.Host != "" {
			t.Errorf("Expected Host to be empty string, got %q", result.Host)
		}
		if result.Port != 0 {
			t.Errorf("Expected Port to be 0, got %d", result.Port)
		}
		if result.Enabled != false {
			t.Errorf("Expected Enabled to be false, got %t", result.Enabled)
		}
		if result.Timeout != 0.0 {
			t.Errorf("Expected Timeout to be 0.0, got %f", result.Timeout)
		}
		if result.Settings != nil {
			t.Errorf("Expected Settings to be nil, got %v", result.Settings)
		}
	})

	t.Run("nested_struct", func(t *testing.T) {
		type Address struct {
			Street string
			City   string
		}
		type Person struct {
			Name    string
			Age     int
			Address Address
		}
		result := Zero[Person]()
		expected := Person{
			Name:    "",
			Age:     0,
			Address: Address{Street: "", City: ""},
		}
		if result != expected {
			t.Errorf("Expected %+v, got %+v", expected, result)
		}
	})

	t.Run("struct_with_pointer_fields", func(t *testing.T) {
		type OptionalFields struct {
			Required string
			Optional *string
			Number   *int
		}
		result := Zero[OptionalFields]()
		expected := OptionalFields{
			Required: "",
			Optional: nil,
			Number:   nil,
		}
		if result != expected {
			t.Errorf("Expected %+v, got %+v", expected, result)
		}
	})
}

func TestZero_ChannelTypes(t *testing.T) {
	// Test Zero with different channel types
	t.Run("unbuffered_channel", func(t *testing.T) {
		result := Zero[chan int]()
		if result != nil {
			t.Errorf("Expected nil channel, got %v", result)
		}
	})

	t.Run("receive_only_channel", func(t *testing.T) {
		result := Zero[<-chan string]()
		if result != nil {
			t.Errorf("Expected nil receive-only channel, got %v", result)
		}
	})

	t.Run("send_only_channel", func(t *testing.T) {
		result := Zero[chan<- bool]()
		if result != nil {
			t.Errorf("Expected nil send-only channel, got %v", result)
		}
	})
}

func TestZero_FunctionTypes(t *testing.T) {
	// Test Zero with function types
	t.Run("simple_function", func(t *testing.T) {
		result := Zero[func()]()
		if result != nil {
			t.Errorf("Expected nil function, got non-nil function")
		}
	})

	t.Run("function_with_params", func(t *testing.T) {
		result := Zero[func(int, string) bool]()
		if result != nil {
			t.Errorf("Expected nil function, got non-nil function")
		}
	})

	t.Run("function_with_return", func(t *testing.T) {
		result := Zero[func() (int, error)]()
		if result != nil {
			t.Errorf("Expected nil function, got non-nil function")
		}
	})
}

func TestZero_InterfaceTypes(t *testing.T) {
	// Test Zero with interface types
	t.Run("empty_interface", func(t *testing.T) {
		result := Zero[interface{}]()
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	t.Run("custom_interface", func(t *testing.T) {
		type Stringer interface {
			String() string
		}
		result := Zero[Stringer]()
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	t.Run("io_writer_interface", func(t *testing.T) {
		// Use a real interface type instead of comparable
		type Writer interface {
			Write([]byte) (int, error)
		}
		result := Zero[Writer]()
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})
}

func TestZero_GenericTypes(t *testing.T) {
	// Test Zero with generic-related types
	t.Run("any_type", func(t *testing.T) {
		result := Zero[any]()
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})
}

func TestZero_ConsistencyWithDeclaration(t *testing.T) {
	// Test that Zero[T]() produces the same result as var v T; v
	t.Run("int_consistency", func(t *testing.T) {
		var declared int
		zero := Zero[int]()
		if declared != zero {
			t.Errorf("Zero[int]() (%d) != declared var (%d)", zero, declared)
		}
	})

	t.Run("string_consistency", func(t *testing.T) {
		var declared string
		zero := Zero[string]()
		if declared != zero {
			t.Errorf("Zero[string]() (%q) != declared var (%q)", zero, declared)
		}
	})

	t.Run("bool_consistency", func(t *testing.T) {
		var declared bool
		zero := Zero[bool]()
		if declared != zero {
			t.Errorf("Zero[bool]() (%t) != declared var (%t)", zero, declared)
		}
	})

	t.Run("slice_consistency", func(t *testing.T) {
		var declared []int
		zero := Zero[[]int]()
		if len(declared) != len(zero) || (declared == nil) != (zero == nil) {
			t.Errorf("Zero[[]int]() (%v) != declared var (%v)", zero, declared)
		}
	})

	t.Run("struct_consistency", func(t *testing.T) {
		type TestStruct struct {
			Name string
			Age  int
		}
		var declared TestStruct
		zero := Zero[TestStruct]()
		if declared != zero {
			t.Errorf("Zero[TestStruct]() (%+v) != declared var (%+v)", zero, declared)
		}
	})
}

func TestZero_UsagePatterns(t *testing.T) {
	// Test common usage patterns
	t.Run("initialization_alternative", func(t *testing.T) {
		// Instead of var x int, use Zero[int]()
		result := Zero[int]()
		if result != 0 {
			t.Errorf("Expected 0, got %d", result)
		}
	})

	t.Run("clearing_variables", func(t *testing.T) {
		// Simulating clearing a variable to its zero value
		type State struct {
			Counter int
			Message string
		}

		state := State{Counter: 42, Message: "active"}
		state = Zero[State]() // Clear to zero value

		expected := State{Counter: 0, Message: ""}
		if state != expected {
			t.Errorf("Expected %+v, got %+v", expected, state)
		}
	})

	t.Run("default_return_value", func(t *testing.T) {
		// Function that returns zero value on error
		getValue := func(valid bool) int {
			if !valid {
				return Zero[int]()
			}
			return 42
		}

		result := getValue(false)
		if result != 0 {
			t.Errorf("Expected 0, got %d", result)
		}

		result = getValue(true)
		if result != 42 {
			t.Errorf("Expected 42, got %d", result)
		}
	})
}

// Benchmark tests
func BenchmarkZero_Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Zero[int]()
	}
}

func BenchmarkZero_String(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Zero[string]()
	}
}

func BenchmarkZero_Slice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Zero[[]int]()
	}
}

func BenchmarkZero_Struct(b *testing.B) {
	type TestStruct struct {
		Name    string
		Age     int
		Enabled bool
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Zero[TestStruct]()
	}
}

func BenchmarkZero_vs_Declaration(b *testing.B) {
	b.Run("with_zero", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Zero[int]()
		}
	})

	b.Run("with_declaration", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var v int
			_ = v
		}
	})
}
