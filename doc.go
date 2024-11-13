/*
Package into provides a contemporary type conversion library for Go 1.18+, leveraging
generics to offer safe and flexible type conversions between basic Go types.

Key Features:

    • Generic conversion functions with compile-time type checking
    • Runtime boundary checks for numeric conversions
    • Consistent error handling
    • IDE-friendly direct conversion functions
    • Support for basic Go types (bool, numeric types, string)

Basic Usage:

    // Generic conversion with error handling
    result, err := into.TryInto[int32](someFloat64)
    if err != nil {
        // Handle error
    }

    // Direct conversion (faster, no reflection)
    val, err := into.Float64ToInt32(someFloat64)

File Structure:

Each file is named after its target type (e.g., int32.go contains conversions to int32)
and provides both generic and direct conversion functions.

Performance Note:

The generic functions use reflection and may have performance overhead. For
performance-critical code, use the direct conversion functions (e.g., Float64ToInt32).

For more information and examples, see: https://github.com/zenless-lab/into
*/
package into
