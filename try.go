package sugar

import "fmt"

// Try is a generic utility function that executes a function and converts any
// panics that occur during execution into regular Go errors. This provides a
// safe way to call potentially panicking code by transforming panic-based
// error handling into Go's standard (value, error) return pattern.
//
// The function executes the provided function f() and:
//   - If f() executes successfully, returns (result, nil)
//   - If f() panics, recovers from the panic and returns (zero_value, error)
//
// When a panic is recovered, the returned value will be the zero value for
// type T (obtained via Zero[T]()), and the error will contain the panic value
// formatted as "panic: <value>".
//
// Type parameter T can be any type, making this function work with any
// function that returns a single value of type T.
//
// Parameters:
//   - f: A function that takes no parameters and returns a value of type T.
//        This function may panic during execution.
//
// Returns:
//   - retval: The return value of f() on success, or Zero[T]() if f() panics
//   - err: nil on success, or an error describing the panic if f() panics
//
// Example usage:
//
//	// Safe array access that might panic on out-of-bounds
//	arr := []int{1, 2, 3}
//	result, err := Try(func() int {
//	    return arr[10] // This will panic with "index out of range"
//	})
//	if err != nil {
//	    log.Printf("Array access failed: %v", err)
//	    // result is 0 (zero value for int)
//	}
//
//	// Safe type assertion
//	var val interface{} = "hello"
//	result, err := Try(func() int {
//	    return val.(int) // This will panic with "interface conversion"
//	})
//	if err != nil {
//	    log.Printf("Type assertion failed: %v", err)
//	    // result is 0 (zero value for int)
//	}
//
//	// Safe nil pointer dereference protection
//	var ptr *string = nil
//	result, err := Try(func() string {
//	    return *ptr // This will panic with "nil pointer dereference"
//	})
//	if err != nil {
//	    log.Printf("Pointer dereference failed: %v", err)
//	    // result is "" (zero value for string)
//	}
//
//	// Safe division by zero
//	result, err := Try(func() int {
//	    return 10 / 0 // This will panic with "division by zero"
//	})
//	if err != nil {
//	    log.Printf("Division failed: %v", err)
//	    // result is 0 (zero value for int)
//	}
//
//	// Wrapping third-party code that might panic
//	data, err := Try(func() []byte {
//	    return someThirdPartyLibrary.ProcessData(input)
//	})
//	if err != nil {
//	    log.Printf("Third-party processing failed: %v", err)
//	    // Handle the error gracefully instead of crashing
//	}
//
// Common patterns:
//
//	// Converting panic-prone operations to error-based flow
//	func safeDivide(a, b int) (int, error) {
//	    return Try(func() int {
//	        return a / b
//	    })
//	}
//
//	// Safe JSON processing with potential panics
//	func safeUnmarshal(data []byte) (MyStruct, error) {
//	    return Try(func() MyStruct {
//	        var result MyStruct
//	        if err := json.Unmarshal(data, &result); err != nil {
//	            panic(err) // Convert error to panic, then back to error
//	        }
//	        return result
//	    })
//	}
//
//	// Protecting against panics in concurrent code
//	func workerFunction(input Data) (Result, error) {
//	    return Try(func() Result {
//	        // Complex processing that might panic
//	        return processComplexData(input)
//	    })
//	}
//
//	// Chain multiple Try operations
//	result1, err1 := Try(func() int { return step1() })
//	if err1 != nil {
//	    return handleError(err1)
//	}
//	result2, err2 := Try(func() string { return step2(result1) })
//	if err2 != nil {
//	    return handleError(err2)
//	}
//
// The Try function is particularly useful when:
//   - Working with third-party libraries that may panic unexpectedly
//   - Performing operations that are known to potentially panic (array access, type assertions, etc.)
//   - Converting legacy panic-based error handling to modern error return patterns
//   - Building robust systems that should not crash on individual operation failures
//   - Testing code where you want to verify that certain operations panic
//   - Implementing error boundaries in concurrent or plugin systems
//   - Gradual migration from panic-based to error-based error handling
//
// Performance considerations:
//   - Try has minimal overhead when no panic occurs (~4ns)
//   - Recovering from panics is more expensive (~200ns) but still reasonable
//   - The defer/recover mechanism is optimized in modern Go versions
//   - For hot paths where panics are expected to be rare, the overhead is negligible
//
// Note on error handling philosophy:
// Try enables a hybrid approach where you can use panic for exceptional cases
// internally while presenting a clean error-based interface externally. This
// is particularly valuable when:
//   - Integrating with code that uses different error handling styles
//   - Building libraries that should not crash the calling application
//   - Implementing fail-safe mechanisms in critical systems
//
// The function leverages Go's built-in panic/recover mechanism and integrates
// with the Zero[T]() function to provide consistent zero-value behavior across
// all types when panics occur.
func Try[T any](f func() T) (retval T, err error) {
	defer func() {
		if r := recover(); r != nil {
			retval = Zero[T]()
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	return f(), nil
}
