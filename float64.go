package into

import (
	"errors"
	"reflect"
	"strconv"
)

// TryIntoFloat64 attempts to convert a value of generic type T to float64.
//
// The generic type T must satisfy the convertable constraint, which means it must be
// one of the following types:
// - Integer types (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64)
// - Float types (float32)
// - String type that represents a valid numeric value
//
// Parameters:
//   - value: The input value to be converted to float64
//
// Returns:
//   - float64: The converted float64 value
//   - error: Error if conversion fails, nil if successful
//
// Examples:
//
//	i := 42
//	f, err := TryIntoFloat64(i) // f = 42.0, err = nil
//
//	s := "3.14"
//	f, err := TryIntoFloat64(s) // f = 3.14, err = nil
//
//	s := "invalid"
//	f, err := TryIntoFloat64(s) // f = 0, err = error
//
// Note: Conversion may lose precision for very large integers or lead to rounding
// errors for certain decimal string representations.
func TryIntoFloat64[T convertable](value T) (float64, error) {
	result, err := toFloat64(value)
	return result.(float64), err
}

// toFloat64 converts a value of any type to a float64 value. It is used internally by TryIntoFloat64.
func toFloat64(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return value.(float64), nil
	case reflect.Float32:
		return Float32ToFloat64(value.(float32))
	case reflect.Int:
		return IntToFloat64(value.(int))
	case reflect.Int8:
		return Int8ToFloat64(value.(int8))
	case reflect.Int16:
		return Int16ToFloat64(value.(int16))
	case reflect.Int32:
		return Int32ToFloat64(value.(int32))
	case reflect.Int64:
		return Int64ToFloat64(value.(int64))
	case reflect.Uint:
		return UintToFloat64(value.(uint))
	case reflect.Uint8:
		return Uint8ToFloat64(value.(uint8))
	case reflect.Uint16:
		return Uint16ToFloat64(value.(uint16))
	case reflect.Uint32:
		return Uint32ToFloat64(value.(uint32))
	case reflect.Uint64:
		return Uint64ToFloat64(value.(uint64))
	case reflect.String:
		return StringToFloat64(value.(string))
	case reflect.Bool:
		return BoolToFloat64(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

// BoolToFloat64 converts a boolean value to a float64.
// If the input value is true, it returns 1.0. If the input value is false, it returns 0.0.
//
// Parameters:
//   - value: A boolean value to be converted.
//
// Returns:
//   - float64: The converted float64 value (1.0 for true, 0.0 for false).
//   - error: Always returns nil as there is no error condition.
//
// Example:
//
//	result, err := BoolToFloat64(true)
//	// result will be 1.0, err will be nil
func BoolToFloat64(value bool) (float64, error) {
	if value {
		return 1.0, nil
	}
	return 0.0, nil
}

// Float32ToFloat64 converts a float32 value to a float64 value.
//
// Parameters:
//   - value: The float32 value to be converted. The range of float32 is approximately 1.18e-38 to 3.4e+38.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately 2.23e-308 to 1.8e+308.
//   - error: Always returns nil as this function does not produce an error.
//
// Example:
//
//	result, err := Float32ToFloat64(3.14)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(result) // Output: 3.14
func Float32ToFloat64(value float32) (float64, error) {
	return float64(value), nil
}

// Float64ToFloat64 takes a float64 value and returns it unchanged.
//
// Parameters:
//   - value: The input float64 value. The range of float64 is approximately 2.23e-308 to 1.8e+308.
//
// Returns:
//   - float64: The same float64 value that was passed in. The range of float64 is approximately 2.23e-308 to 1.8e+308.
//   - error: Always returns nil as there is no error condition.
//
// Example:
//
//	result, err := Float64ToFloat64(3.14)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(result) // Output: 3.14
func Float64ToFloat64(value float64) (float64, error) {
	return value, nil
}

// IntToFloat64 converts an integer value to a float64 value.
//
// Parameters:
//   - value: an integer that needs to be converted to float64. which is int32 in 32-bit systems and int64 in 64-bit systems.
//
// Returns:
//   - float64: the converted float64 value. The range of float64 is approximately 2.23e-308 to 1.8e+308.
//   - error: always returns nil as there is no error condition.
//
// Example:
//
//	result, err := IntToFloat64(42)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(result) // Output: 42.0
func IntToFloat64(value int) (float64, error) {
	return float64(value), nil
}

// Int8ToFloat64 converts an int8 value to a float64 value.
//
// Parameters:
//   - value: an int8 value to be converted. The range of int8 is from -128 to 127.
//
// Returns:
//   - float64: the converted float64 value. The range of float64 is approximately 2.23e-308 to 1.8e+308.
//   - error: always returns nil as there is no error condition in this conversion.
//
// Example:
//
//	result, err := Int8ToFloat64(42)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(result) // Output: 42.0
func Int8ToFloat64(value int8) (float64, error) {
	return float64(value), nil
}

// Int16ToFloat64 converts an int16 value to a float64 value.
//
// Parameters:
//   - value: The int16 value to be converted. The range of int16 is from -32768 to 32767.
//
// Returns:
//   - float64: The converted float64 value.
//   - error: Always returns nil as this function does not produce an error.
//
// Example:
//
//	result, err := Int16ToFloat64(12345)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(result) // Output: 12345.0
func Int16ToFloat64(value int16) (float64, error) {
	return float64(value), nil
}

// Int32ToFloat64 converts an int32 value to a float64 value.
//
// Parameters:
//   - value: The int32 value to be converted. The range of int32 is from -2147483648 to 2147483647.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately from 2.23e-308 to 1.8e+308.
//   - error: Always returns nil as this conversion cannot fail.
//
// Example:
//
//	result, err := Int32ToFloat64(12345)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(result) // Output: 12345
func Int32ToFloat64(value int32) (float64, error) {
	return float64(value), nil
}

// Int64ToFloat64 converts an int64 value to a float64 value.
//
// Parameters:
//   - value: The int64 value to be converted. The range of int64 is from -9223372036854775808 to 9223372036854775807.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately from -1.8e308 to 1.8e308.
//   - error: Always returns nil as this function does not produce an error.
//
// Example:
//
//	result, err := Int64ToFloat64(123456789)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(result) // Output: 123456789
func Int64ToFloat64(value int64) (float64, error) {
	return float64(value), nil
}

// StringToFloat64 converts a string representation of a number to float64.
//
// The input string can be:
// - A decimal number: "3.14159"
// - An integer: "42"
// - Scientific notation: "1.234e-5"
//
// Parameters:
//   - value: string to be converted to float64.
//
// Returns:
//   - float64: the converted floating-point number. The range of float64 is approximately from 2.23e-308 to 1.8e+308.
//   - error: if the string cannot be parsed as a valid number
//
// Example:
//
//	f, err := StringToFloat64("3.14")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// f == 3.14
func StringToFloat64(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

// UintToFloat64 converts an unsigned integer to a float64.
//
// The function performs a direct type conversion from uint to float64.
//
// Parameters:
//   - value: uint number to convert. The range of uint is from 0 to 18446744073709551615.
//
// Returns:
//   - float64: converted floating-point number. The range of float64 is approximately from 2.23e-308 to 1.8e+308.
//   - error: currently always returns nil (reserved for future validation)
//
// Example:
//
//	val, err := UintToFloat64(42)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Converted value: %f\n", val) // Output: 42.000000
func UintToFloat64(value uint) (float64, error) {
	return float64(value), nil
}

// Uint8ToFloat64 converts an uint8 (0 to 255) to a float64 (±2.23E-308 to ±1.80E+308).
//
// Parameters:
//   - value: An unsigned 8-bit integer. The range of uint8 is from 0 to 255.
//
// Returns:
//   - float64: The converted floating-point number. The range of float64 is approximately from 2.23e-308 to 1.8e+308.
//   - error: Always returns nil as this conversion cannot fail
//
// Example:
//
//	value := uint8(128)
//	result, err := Uint8ToFloat64(value)
//	if err != nil {
//	    // Handle error
//	}
//	fmt.Printf("Converted value: %f\n", result) // Output: 128.000000
func Uint8ToFloat64(value uint8) (float64, error) {
	return float64(value), nil
}

// Uint16ToFloat64 converts a uint16 value to float64.
//
// Parameters:
//   - value: uint16 input. The range of uint16 is from 0 to 65535.
//
// Returns:
//   - float64: converted value. The range of float64 is approximately from 2.23e-308 to 1.8e+308.
//   - error: always returns nil as this conversion cannot fail
//
// Example:
//
//	val, err := Uint16ToFloat64(12345)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Printf("Converted value: %f\n", val) // prints: 12345.000000
func Uint16ToFloat64(value uint16) (float64, error) {
	return float64(value), nil
}

// Uint32ToFloat64 converts an unsigned 32-bit integer to a 64-bit floating-point number.
//
// Parameters:
//   - value: uint32 input value. The range of uint32 is from 0 to 4294967295.
//
// Returns:
//   - float64: converted floating-point value. The range of float64 is approximately from 2.23e-308 to 1.8e+308.
//   - error: always returns nil as this conversion cannot fail
//
// Example:
//
//	val := uint32(12345)
//	f, err := Uint32ToFloat64(val)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Converted %d to %f\n", val, f)
//	// Output: Converted 12345 to 12345.000000
func Uint32ToFloat64(value uint32) (float64, error) {
	return float64(value), nil
}

// Uint64ToFloat64 converts an unsigned 64-bit integer to a float64.
//
// Parameters:
//   - value: uint64 number to be converted. The range of uint64 is from 0 to 18446744073709551615.
//
// Returns:
//   - float64: converted floating point number. The range of float64 is approximately from 2.23e-308 to 1.8e+308.
//   - error: currently always returns nil, reserved for future validation
//
// Note that float64 has limited precision and can exactly represent integers
// only up to 2^53 (9007199254740992). Values larger than this may lose precision
// during conversion.
//
// Example:
//
//	val := uint64(123456789)
//	f, err := Uint64ToFloat64(val)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Converted value: %f\n", f)
func Uint64ToFloat64(value uint64) (float64, error) {
	return float64(value), nil
}
