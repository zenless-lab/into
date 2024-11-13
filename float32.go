package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// TryIntoFloat32 attempts to convert a value of type T to float32.
//
// Parameters:
//   - value: The input value to convert, must satisfy the convertable constraint
//
// Returns:
//   - float32: The converted float32 value
//   - error: Error if conversion fails
//
// float32 range: ±3.4E38 (approximately 7 decimal digits precision)
//
// Example:
//
//	// Convert int to float32
//	f, err := TryIntoFloat32(42)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("%f\n", f) // prints: 42.000000
//
//	// Convert string to float32
//	f, err = TryIntoFloat32("3.14")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("%f\n", f) // prints: 3.140000
func TryIntoFloat32[T convertable](value T) (float32, error) {
	resoult, err := toFloat32(value)
	return resoult.(float32), err
}

// toFloat32 converts a value of any type to a float32 value. It is used internally by TryIntoFloat32.
func toFloat32(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToFloat32(value.(float64))
	case reflect.Float32:
		return value.(float32), nil
	case reflect.Int:
		return IntToFloat32(value.(int))
	case reflect.Int8:
		return Int8ToFloat32(value.(int8))
	case reflect.Int16:
		return Int16ToFloat32(value.(int16))
	case reflect.Int32:
		return Int32ToFloat32(value.(int32))
	case reflect.Int64:
		return Int64ToFloat32(value.(int64))
	case reflect.Uint:
		return UintToFloat32(value.(uint))
	case reflect.Uint8:
		return Uint8ToFloat32(value.(uint8))
	case reflect.Uint16:
		return Uint16ToFloat32(value.(uint16))
	case reflect.Uint32:
		return Uint32ToFloat32(value.(uint32))
	case reflect.Uint64:
		return Uint64ToFloat32(value.(uint64))
	case reflect.String:
		return StringToFloat32(value.(string))
	case reflect.Bool:
		return BoolToFloat32(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

// BoolToFloat32 converts a boolean value to its float32 representation.
//
// Parameters:
//   - value: A boolean value to be converted to float32.
//
// Returns:
//   - float32: Returns 1.0 if the input is true, 0.0 if the input is false
//   - error: Always returns nil as this conversion cannot fail
//
// Example:
//
//	f, err := BoolToFloat32(true)  // f = 1.0, err = nil
//	f, err := BoolToFloat32(false) // f = 0.0, err = nil
func BoolToFloat32(value bool) (float32, error) {
	if value {
		return 1.0, nil
	}
	return 0.0, nil
}

// Float32ToFloat32 converts a float32 value to another float32 value.
//
// Parameters:
//   - value: The float32 input value to be converted.
//
// Returns:
//   - float32: The converted float32 value, which is identical to the input.
//   - error: Always returns nil as this conversion cannot fail.
//
// The function accepts any valid float32 value within the range of approximately
// ±3.4E38 (IEEE-754 single precision floating point).
//
// Example:
//
//	input := float32(3.14)
//	result, err := Float32ToFloat32(input)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Converted value: %f\n", result) // Output: Converted value: 3.140000
func Float32ToFloat32(value float32) (float32, error) {
	return value, nil
}

// Float64ToFloat32 converts a float64 value to a float32 value.
// If the input value is outside the range that can be represented by a float32,
// it returns an error.
//
// Parameters:
//   - value: the float64 value to be converted. The range of float64 is approximately ±1.7E308.
//
// Returns:
//   - float32: the converted float32 value. The range of float32 is approximately ±3.4E38.
//   - error: an error if the input value is out of the float32 range.
//
// Example:
//
//	result, err := Float64ToFloat32(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123.456
//
// Range:
//
//	The function ensures that the input value is within the range of float32,
//	which is approximately -3.4e38 to 3.4e38.
func Float64ToFloat32(value float64) (float32, error) {
	if value > math.MaxFloat32 || value < -math.MaxFloat32 {
		return 0, errors.New("value out of range for float32")
	}
	return float32(value), nil
}

// IntToFloat32 converts an integer value to a float32 value.
// If the integer value is outside the range that can be represented by a float32,
// it returns an error.
//
// Parameters:
//   - value: The integer value to be converted. The range of int is from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807.
//
// Returns:
//   - float32: The converted float32 value. The range of float32 is approximately ±3.4E38.
//   - error: An error if the value is out of range for float32.
//
// Example usage:
//
//	f, err := IntToFloat32(123)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(f) // Output: 123
func IntToFloat32(value int) (float32, error) {
	if float64(value) > math.MaxFloat32 || float64(value) < -math.MaxFloat32 {
		return 0, errors.New("value out of range for float32")
	}
	return float32(value), nil
}

// Int8ToFloat32 converts an int8 value to a float32 value.
//
// Parameters:
//   - value: The int8 value to be converted. The range of int8 is from -128 to 127.
//
// Returns:
//   - float32: The converted float32 value. The range of float32 is approximately ±3.4E38.
//   - error: An error value, which is always nil in this implementation.
//
// Example:
//
//	result, err := Int8ToFloat32(42)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(result) // Output: 42.0
//
// Note:
//
//	This function is designed to convert int8 values within the range of -128 to 127 to float32.
func Int8ToFloat32(value int8) (float32, error) {
	return float32(value), nil
}

// Int16ToFloat32 converts an int16 value to a float32 value.
//
// Parameters:
//   - value: The int16 value to be converted. The range of int16 is from -32768 to 32767.
//
// Returns:
//   - float32: The converted float32 value. The range of float32 is approximately ±3.4E38.
//   - error: Always returns nil as there is no error condition in this conversion.
//
// Example:
//
//	result, err := Int16ToFloat32(12345)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(result) // Output: 12345
func Int16ToFloat32(value int16) (float32, error) {
	return float32(value), nil
}

// Int32ToFloat32 converts an int32 value to a float32 value.
//
// Parameters:
//   - value: The int32 value to be converted. The range of int32 is from -2,147,483,648 to 2,147,483,647.
//
// Returns:
//   - float32: The converted float32 value. The range of float32 is approximately from -3.4e38 to 3.4e38.
//   - error: Always returns nil as this function does not produce an error.
//
// Example:
//
//	result, err := Int32ToFloat32(12345)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(result) // Output: 12345
func Int32ToFloat32(value int32) (float32, error) {
	return float32(value), nil
}

// Int64ToFloat32 converts an int64 value to a float32 value.
//
// Parameters:
// - value: The int64 value to be converted. The range of int64 is from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807.
//
// Returns:
// - float32: The converted float32 value. The range of float32 is approximately from -3.4e38 to 3.4e38.
// - error: An error if the value is out of the range for float32.
//
// The function checks if the int64 value is within the range of float32.
// If the value is out of range, it returns an error.
//
// Example:
//
//	result, err := Int64ToFloat32(123456789)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//	    fmt.Println("Result:", result)
//	}
func Int64ToFloat32(value int64) (float32, error) {
	if float64(value) > math.MaxFloat32 || float64(value) < -math.MaxFloat32 {
		return 0, errors.New("value out of range for float32")
	}
	return float32(value), nil
}

// StringToFloat32 converts a string representation of a floating-point number
// to a float32. It takes a string `value` as input and returns the converted
// float32 value and an error if the conversion fails.
//
// Parameters:
//   - value: A string representing the floating-point number to be converted.
//
// Returns:
//   - float32: The converted floating-point number. The range of float32 is approximately ±3.4E38.
//   - error: An error if the conversion fails, otherwise nil.
//
// Example:
//
//	f, err := StringToFloat32("3.14")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(f) // Output: 3.14
func StringToFloat32(value string) (float32, error) {
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

// UintToFloat32 converts an unsigned integer to a float32.
//
// Parameters:
//   - value: The unsigned integer to be converted. The value should be within the range of a uint. The range of uint is from 0 to 18446744073709551615.
//
// Returns:
//   - float32: The converted float32 value. The range of float32 is approximately ±3.4E38.
//   - error: Always returns nil as there is no error condition in this conversion.
//
// Example:
//
//	result, err := UintToFloat32(42)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(result) // Output: 42.0
func UintToFloat32(value uint) (float32, error) {
	return float32(value), nil
}

// Uint8ToFloat32 converts a uint8 value to a float32 value.
//
// Parameters:
//   - value: The uint8 value to be converted. The range of uint8 is 0 to 255.
//
// Returns:
//   - float32: The converted float32 value. The range of float32 is approximately ±3.4E38.
//   - error: Always returns nil as there is no error condition.
//
// Example:
//
//	result, err := Uint8ToFloat32(42)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(result) // Output: 42.0
func Uint8ToFloat32(value uint8) (float32, error) {
	return float32(value), nil
}

// Uint16ToFloat32 converts a uint16 value to a float32 value.
//
// Parameters:
//   - value: The uint16 value to be converted. The range of uint16 is from 0 to 65535.
//
// Returns:
//   - float32: The converted float32 value. The range of float32 is approximately ±3.4E38.
//   - error: Always returns nil as there is no error condition in this conversion.
//
// Example:
//
//	result, err := Uint16ToFloat32(12345)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(result) // Output: 12345
func Uint16ToFloat32(value uint16) (float32, error) {
	return float32(value), nil
}

// Uint32ToFloat32 converts a uint32 value to a float32 value.
// If the uint32 value exceeds the maximum limit of float32, it returns an error.
//
// Parameters:
//   - value: The uint32 value to be converted. The range of uint32 is from 0 to 4294967295.
//
// Returns:
//   - float32: The converted float32 value. The range of float32 is approximately ±3.4E38.
//   - error: An error if the value exceeds the float32 maximum limit.
//
// Example:
//
//	result, err := Uint32ToFloat32(1234567890)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 1.23456794e+09
func Uint32ToFloat32(value uint32) (float32, error) {
	if float64(value) > math.MaxFloat32 {
		return 0, errors.New("value exceeds float32 max limit")
	}
	return float32(value), nil
}

// Uint64ToFloat32 converts a uint64 value to a float32 value.
//
// Parameters:
//   - value: The uint64 value to be converted. The range of uint64 is from 0 to 18446744073709551615.
//
// Returns:
//   - float32: The converted float32 value if the conversion is successful. The range of float32 is approximately ±3.4E38.
//   - error: An error if the value exceeds the maximum limit of float32.
//
// Example:
//
//	result, err := Uint64ToFloat32(123456789)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 1.2345679e+08
func Uint64ToFloat32(value uint64) (float32, error) {
	if float64(value) > math.MaxFloat32 {
		return 0, errors.New("value exceeds float32 max limit")
	}
	return float32(value), nil
}
