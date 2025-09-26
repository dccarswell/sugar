# Sugar

A collection of generic utility functions that add syntactic sugar to common Go patterns. This library provides type-safe, zero-dependency utilities for error handling, pointer management, and value initialization.

## Installation

```bash
go get github.com/dccarswell/sugar
```

## Overview

The Sugar library provides five core utilities that simplify common Go programming patterns:

- **`Must[T]`** - Convert (value, error) pairs to value-or-panic
- **`Try[T]`** - Convert panics to (value, error) pairs  
- **`Ptr[T]`** - Create pointers to any value (especially literals)
- **`Zero[T]`** - Get zero values for any type
- **`Handle[T]`** - Customizable error handling with panic conversion

All functions are generic and work with any Go type, providing type safety and consistency across your codebase.

## Functions

### `Must[T any](v T, err error) T`

Converts Go's explicit error handling into panic-based error handling. Takes a (value, error) pair and either returns the value (if error is nil) or panics with the error.

**Use Cases:**
- Initialization code where errors should be fatal
- Converting functions that return (T, error) into functions that return T or panic
- Simplifying error handling in prototypes and CLI tools
- Test setup where failure should crash the test

**Examples:**
```go
// File operations
data := Must(os.ReadFile("config.json"))

// Parsing operations  
port := Must(strconv.Atoi(os.Getenv("PORT")))

// Network operations
conn := Must(net.Dial("tcp", "localhost:8080"))

// Global initialization
var db = Must(sql.Open("postgres", connectionString))

// Chaining operations
result := Must(processData(Must(loadConfig(Must(os.ReadFile("app.conf"))))))
```

**When to use:** Programming errors, initialization, testing, CLI tools  
**When to avoid:** Expected runtime conditions, library code, long-running services

---

### `Try[T any](f func() T) (T, error)`

Executes a function and converts any panics into regular Go errors. Provides a safe way to call potentially panicking code by transforming panic-based error handling into Go's standard (value, error) pattern.

**Use Cases:**
- Wrapping third-party libraries that may panic
- Safe execution of operations known to potentially panic
- Converting legacy panic-based code to error-based patterns
- Building robust systems that shouldn't crash on individual failures

**Examples:**
```go
// Safe array access
arr := []int{1, 2, 3}
result, err := Try(func() int {
    return arr[10] // Would panic with "index out of range"
})
if err != nil {
    log.Printf("Array access failed: %v", err)
    // result is 0 (zero value for int)
}

// Safe type assertion
var val interface{} = "hello"
result, err := Try(func() int {
    return val.(int) // Would panic with "interface conversion"
})

// Safe nil pointer dereference
var ptr *string = nil
result, err := Try(func() string {
    return *ptr // Would panic with "nil pointer dereference"  
})

// Wrapping third-party code
data, err := Try(func() []byte {
    return someThirdPartyLibrary.ProcessData(input)
})
```

**Performance:** ~4ns overhead for normal execution, ~200ns for panic recovery

---

### `Ptr[T any](v T) *T`

Creates a pointer to any value. Particularly useful for getting pointers to literal values, which Go doesn't allow directly with the `&` operator.

**Use Cases:**
- Optional struct fields that use pointers to indicate presence/absence
- APIs that require pointer parameters
- Working with JSON marshaling where nil pointers represent absent fields
- Creating test data with specific pointer values

**Examples:**
```go
// Getting pointers to literals
intPtr := Ptr(42)        // *int pointing to 42
stringPtr := Ptr("hello") // *string pointing to "hello"
boolPtr := Ptr(true)     // *bool pointing to true

// Optional fields in structs
type Config struct {
    Host    string
    Port    int
    Timeout *int   // Optional field
    Debug   *bool  // Optional field
}
config := Config{
    Host:    "localhost", 
    Port:    8080,
    Timeout: Ptr(30),    // Clean and concise
    Debug:   Ptr(true),  // No intermediate variables needed
}

// Function parameters requiring pointers
func updateValue(ptr *string) {
    *ptr = "updated"
}
updateValue(Ptr("initial")) // Pass pointer to literal

// Collections of pointers
numbers := []*int{Ptr(1), Ptr(2), Ptr(3)}
settings := map[string]*bool{
    "enabled": Ptr(true),
    "verbose": Ptr(false),
}
```

**Note:** Returns a pointer to a *copy* of the input value, not the original variable.

---

### `Zero[T any]() T`

Returns the zero value for any type T. Provides a clean way to get zero values in generic code where `var v T` syntax isn't convenient.

**Use Cases:**
- Generic functions that need zero values
- Clearing variables to their zero state
- Default return values in error conditions
- Initialization in generic code

**Examples:**
```go
// Basic zero values
intZero := Zero[int]()        // 0
stringZero := Zero[string]()  // ""
boolZero := Zero[bool]()      // false
sliceZero := Zero[[]int]()    // nil

// Struct zero values
type Person struct {
    Name string
    Age  int
}
personZero := Zero[Person]() // {Name: "", Age: 0}

// Generic function usage
func getDefaultValue[T any]() T {
    return Zero[T]()
}

// Clearing variables
func resetState[T any](ptr *T) {
    *ptr = Zero[T]()
}

// Error handling with defaults
func safeDivide(a, b int) int {
    if b == 0 {
        return Zero[int]() // Return 0 instead of panicking
    }
    return a / b
}
```

---

### `Handle[T any](h Handler[T]) func(T, error) T`

Creates a customizable error handler that processes (value, error) pairs. Allows you to define custom logic for handling errors - either converting them to nil (handled) or transforming them before panicking.

**Type:** `Handler[T any] func(error) error`

**Use Cases:**
- Adding logging or monitoring to error paths
- Transforming or wrapping errors with additional context  
- Selectively ignoring certain types of errors
- Building error handling pipelines or middleware
- Converting between different error handling styles

**Examples:**
```go
// Handler that logs and ignores network timeouts
networkHandler := Handle[[]byte](func(err error) error {
    var netErr net.Error
    if errors.As(err, &netErr) && netErr.Timeout() {
        log.Printf("Network timeout, retrying: %v", err)
        return nil // Don't panic on timeouts
    }
    return err // Panic on other network errors
})

// Use with network operations  
data := networkHandler(httpClient.Get(url))

// Always panic on any error (equivalent to Must)
strictHandler := Handle[string](func(err error) error { 
    return err 
})

// Always ignore errors
ignoreHandler := Handle[int](func(err error) error { 
    return nil 
})

// Log and re-raise errors
logHandler := Handle[Data](func(err error) error {
    log.Printf("Operation failed: %v", err)
    return err
})

// Transform errors before panicking
contextHandler := Handle[Result](func(err error) error {
    return fmt.Errorf("processing failed in module X: %w", err)
})
```

## Performance

All functions are designed to be lightweight:

- **Must:** ~0.22ns (equivalent to direct calls)
- **Try:** ~4ns normal execution, ~200ns panic recovery
- **Ptr:** ~0.22ns (equivalent to taking address)
- **Zero:** ~0.22ns (equivalent to variable declaration)
- **Handle:** ~4ns (similar to Try when no error occurs)

## Error Handling Philosophy

This library supports multiple error handling approaches:

1. **Explicit errors** (traditional Go): `value, err := function()`
2. **Panic-based** (exceptional cases): `value := Must(function())`
3. **Safe execution** (unknown code): `value, err := Try(func() { ... })`
4. **Custom handling** (middleware): `value := handler(function())`

Choose the approach that best fits your specific use case and context.

## Examples

### Configuration Loading
```go
type Config struct {
    Host     string
    Port     int
    Timeout  *int
    Debug    *bool
    Database *DatabaseConfig
}

func LoadConfig(filename string) Config {
    // Must: Configuration loading should be fatal if it fails
    data := Must(os.ReadFile(filename))
    
    config := Config{
        Host: "localhost",
        Port: 8080,
        // Ptr: Clean optional field initialization
        Timeout: Ptr(30),
        Debug:   Ptr(false),
    }
    
    // Try: JSON parsing might panic, convert to error
    parsed, err := Try(func() Config {
        Must(json.Unmarshal(data, &config))
        return config
    })
    
    if err != nil {
        log.Printf("Config parsing failed, using defaults: %v", err)
        return config // Return default config
    }
    
    return parsed
}
```

### Safe Data Processing
```go
func ProcessUserData(input []byte) (Result, error) {
    // Handle: Custom error handling with logging
    processor := Handle[Result](func(err error) error {
        log.Printf("Data processing error: %v", err)
        
        // Ignore validation errors, but panic on system errors
        if strings.Contains(err.Error(), "validation") {
            return nil
        }
        return fmt.Errorf("system error: %w", err)
    })
    
    return Try(func() Result {
        // This might panic on malformed data
        return processor(parseAndValidate(input))
    })
}
```

### API Response Building
```go
type APIResponse struct {
    Data    *json.RawMessage `json:"data,omitempty"`
    Error   *string          `json:"error,omitempty"`
    Code    *int             `json:"code,omitempty"`
    Success *bool            `json:"success,omitempty"`
}

func BuildResponse(data interface{}, err error) APIResponse {
    if err != nil {
        return APIResponse{
            Error:   Ptr(err.Error()),
            Code:    Ptr(500),
            Success: Ptr(false),
        }
    }
    
    jsonData := Must(json.Marshal(data))
    return APIResponse{
        Data:    Ptr(json.RawMessage(jsonData)),
        Success: Ptr(true),
    }
}
```

## Contributing

Contributions are welcome! Please ensure all code:
- Includes comprehensive tests
- Follows Go conventions and best practices
- Includes clear documentation
- Maintains backward compatibility

## License

[License information here]
