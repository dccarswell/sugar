package sugar

// Must is a generic utility function that converts Go's explicit error handling
// into panic-based error handling. It takes a (value, error) pair and either
// returns the value (if error is nil) or panics with the error (if error is non-nil).
//
// This function is particularly useful for:
//   - Simplifying error handling in scenarios where errors should be fatal
//   - Converting functions that return (T, error) into functions that return T or panic
//   - Initializing variables where failure should crash the program
//   - Prototyping and testing code where error handling can be deferred
//   - Working with operations that "must succeed" in the current context
//
// Type parameter T can be any type, making this function work with any Go function
// that follows the common (value, error) return pattern.
//
// Parameters:
//   - v: The value to return if no error occurred
//   - err: The error to check; if non-nil, causes a panic
//
// Returns:
//   - The value v if err is nil
//
// Panics:
//   - If err is non-nil, panics with the error as the panic value
//
// Example usage:
//
//	// Converting file operations
//	data := Must(os.ReadFile("config.json"))
//	// Equivalent to:
//	// data, err := os.ReadFile("config.json")
//	// if err != nil {
//	//     panic(err)
//	// }
//
//	// Converting parsing operations
//	port := Must(strconv.Atoi(os.Getenv("PORT")))
//
//	// Converting network operations
//	conn := Must(net.Dial("tcp", "localhost:8080"))
//
//	// Chaining operations
//	result := Must(processData(Must(loadConfig(Must(os.ReadFile("app.conf"))))))
//
// Common patterns:
//
//	// Initialization that must succeed
//	var globalDB = Must(sql.Open("postgres", connectionString))
//
//	// Configuration loading
//	func loadConfig() Config {
//	    return Must(json.Unmarshal(Must(os.ReadFile("config.json")), &Config{}))
//	}
//
//	// Testing scenarios
//	func TestSomething(t *testing.T) {
//	    data := Must(generateTestData()) // Test setup must succeed
//	    // ... test logic
//	}
//
// Warning: Must should be used judiciously. It's appropriate when:
//   - The error represents a programming error or misconfiguration
//   - Failure should terminate the program (like during initialization)
//   - You're in a context where panics are acceptable (tests, CLI tools, etc.)
//   - The operation is expected to always succeed in normal circumstances
//
// Avoid Must when:
//   - The error represents expected runtime conditions (user input, network issues)
//   - You're writing library code that should not crash the caller's program
//   - Recovery from the error is possible and desired
//   - The function is part of a long-running service that should handle errors gracefully
//
// The Must function is inspired by similar utilities in other languages and
// Go libraries, providing a concise way to handle the common case where an
// error should be treated as fatal.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
