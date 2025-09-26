// Package sugar provides utility functions that add syntactic sugar to common Go patterns.
// This file contains error handling utilities for converting error values into panic-based control flow.
package sugar

// Handler is a generic function type that processes an error and returns either nil
// (to indicate the error was handled successfully) or a new error (to indicate
// the error should cause a panic).
//
// Type parameter T represents the type of value being processed, though the handler
// function itself doesn't directly work with values of type T. The type parameter
// is used to create type-safe handler functions for specific value types.
//
// Example usage:
//
//	// Handler that logs errors and converts them to nil (swallows errors)
//	logHandler := func(err error) error {
//		log.Printf("Error occurred: %v", err)
//		return nil // Error handled, don't panic
//	}
//
//	// Handler that transforms specific errors
//	transformHandler := func(err error) error {
//		if strings.Contains(err.Error(), "temporary") {
//			return nil // Ignore temporary errors
//		}
//		return fmt.Errorf("critical error: %w", err) // Transform and re-raise
//	}
type Handler[T any] func(error) error

// Handle creates a function that processes (value, error) pairs using the provided
// error handler. This is useful for converting Go's explicit error handling into
// a more exception-like control flow where errors either get handled gracefully
// or cause panics.
//
// The returned function takes a value of type T and an error. If the error is nil,
// the value is returned unchanged. If the error is non-nil, it's passed to the
// handler function:
//   - If the handler returns nil, the original value is returned (error was handled)
//   - If the handler returns a non-nil error, that error is panicked with
//
// Type parameter T can be any type, making this function work with any kind of
// (value, error) pair that Go functions commonly return.
//
// Parameters:
//   - h: A Handler function that decides how to process errors
//
// Returns:
//   - A function that takes (T, error) and returns T, with error handling logic applied
//
// Example usage:
//
//	// Create a handler that logs and ignores network timeout errors
//	networkHandler := Handle[[]byte](func(err error) error {
//		var netErr net.Error
//		if errors.As(err, &netErr) && netErr.Timeout() {
//			log.Printf("Network timeout, retrying: %v", err)
//			return nil // Don't panic on timeouts
//		}
//		return err // Panic on other network errors
//	})
//
//	// Use with network operations
//	data := networkHandler(httpClient.Get(url))
//	// If Get() returns a timeout error, it gets logged and data receives zero value
//	// If Get() returns other errors, they cause a panic
//	// If Get() succeeds, data contains the response
//
// Common patterns:
//
//	// Always panic on any error (equivalent to Must pattern)
//	strictHandler := Handle[string](func(err error) error { return err })
//
//	// Always ignore errors (swallow all errors)
//	ignoreHandler := Handle[int](func(err error) error { return nil })
//
//	// Log and re-raise errors
//	logHandler := Handle[Data](func(err error) error {
//		log.Printf("Operation failed: %v", err)
//		return err
//	})
//
//	// Transform errors before panicking
//	contextHandler := Handle[Result](func(err error) error {
//		return fmt.Errorf("processing failed in module X: %w", err)
//	})
//
// The Handle function is particularly useful in scenarios where:
//   - You want to add logging or monitoring to error paths
//   - You need to transform or wrap errors with additional context
//   - You want to selectively ignore certain types of errors
//   - You're building error handling pipelines or middleware
//   - You want to convert between error handling styles in different parts of your application
func Handle[T any](h Handler[T]) func(T, error) T {
	return func(v T, err error) T {
		if err != nil {
			err = h(err)
			if err != nil {
				panic(err)
			}
		}
		return v
	}
}
