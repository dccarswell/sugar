package sugar

// Ptr is a generic utility function that creates a pointer to any value.
// It takes a value of any type T and returns a pointer to that value.
// This is particularly useful when you need a pointer to a literal value
// or when working with APIs that require pointer parameters.
//
// The function creates a copy of the input value and returns a pointer to
// that copy, not to the original variable. This means modifying the value
// through the returned pointer won't affect the original variable that was
// passed in (if any).
//
// Type parameter T can be any type, making this function work with all Go
// types including basic types, structs, slices, maps, and interfaces.
//
// Parameters:
//   - v: The value to create a pointer to
//
// Returns:
//   - A pointer to a copy of the input value
//
// Example usage:
//
//	// Getting pointers to literals (common use case)
//	intPtr := Ptr(42)              // *int pointing to 42
//	stringPtr := Ptr("hello")      // *string pointing to "hello"
//	boolPtr := Ptr(true)          // *bool pointing to true
//
//	// Using with struct literals
//	type Person struct {
//	    Name string
//	    Age  int
//	}
//	personPtr := Ptr(Person{Name: "Alice", Age: 30})
//
//	// Useful for optional fields in structs
//	type Config struct {
//	    Host     string
//	    Port     int
//	    Timeout  *int    // Optional field
//	    Debug    *bool   // Optional field
//	}
//	config := Config{
//	    Host:    "localhost",
//	    Port:    8080,
//	    Timeout: Ptr(30),        // Instead of creating intermediate variable
//	    Debug:   Ptr(true),      // Clean and concise
//	}
//
//	// Function parameters that require pointers
//	func updateValue(ptr *string) {
//	    *ptr = "updated"
//	}
//	updateValue(Ptr("initial"))  // Pass pointer to literal
//
//	// Working with APIs that expect pointers
//	json.Marshal(map[string]interface{}{
//	    "name":    Ptr("John"),
//	    "age":     Ptr(25),
//	    "active":  Ptr(true),
//	})
//
// Common patterns:
//
//	// Instead of this verbose pattern:
//	temp := 42
//	intPtr := &temp
//
//	// Use this concise pattern:
//	intPtr := Ptr(42)
//
//	// Conditional pointer creation
//	var namePtr *string
//	if name != "" {
//	    namePtr = Ptr(name)
//	}
//
//	// Slice of pointers to literals
//	numbers := []*int{Ptr(1), Ptr(2), Ptr(3)}
//
//	// Map with pointer values
//	settings := map[string]*bool{
//	    "enabled":  Ptr(true),
//	    "verbose":  Ptr(false),
//	    "debug":    Ptr(true),
//	}
//
// The Ptr function is especially valuable when:
//   - Working with optional struct fields that use pointers to indicate presence/absence
//   - Interfacing with C libraries or APIs that require pointers
//   - Creating test data where you need pointers to specific values
//   - Building configuration objects with optional parameters
//   - Working with JSON marshaling where nil pointers represent absent fields
//   - Converting between value and pointer semantics in generic code
//
// Note: The returned pointer points to a copy of the input value, so:
//
//	original := 42
//	ptr := Ptr(original)
//	*ptr = 100
//	// original is still 42, *ptr is 100
//
// This behavior is usually what you want when creating pointers to literals
// or when you need independent pointer values.
func Ptr[T any](v T) *T {
	return &v
}
