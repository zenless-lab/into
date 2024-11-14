package into

import (
	"errors"
	"reflect"
	"strconv"
)

// TryIntoFloat64 attempts to convert the given value to a float64.
//
// TryIntoFloat64 attempts to convert the given value to a float64.
// If the input value is not a float64, it attempts to convert it using the appropriate conversion function.
// If the input value cannot be converted to a float64, it returns an error.
//
// Parameters:
//   - value: The value to be converted. The range of float64 is approximately ±1.7E308.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: An error if the input value cannot be converted to a float64.
//
// Example:
//
//	result, err := TryIntoFloat64(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123.456
func TryIntoFloat64[T convertable](value T) (float64, error) {
	result, err := toFloat64(value)
	return result.(float64), err
}

// toFloat64 converts the given value to a float64.
//
// toFloat64 converts the given value to a float64.
// If the input value is not a float64, it attempts to convert it using the appropriate conversion function.
// If the input value cannot be converted to a float64, it returns an error.
//
// Parameters:
//   - value: The value to be converted.
//
// Returns:
//   - any: The converted float64 value.
//   - error: An error if the input value cannot be converted to a float64.
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

// BoolToFloat64 converts a bool value to a float64 value.
//
// BoolToFloat64 converts a bool value to a float64 value.
// True is converted to 1.0, and false is converted to 0.0.
//
// Parameters:
//   - value: The bool value to be converted.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: No error is returned.
func BoolToFloat64(value bool) (float64, error) {
	if value {
		return 1.0, nil
	}
	return 0.0, nil
}

// Float32ToFloat64 converts a float32 value to a float64 value.
//
// Float32ToFloat64 converts a float32 value to a float64 value.
// The range of float32 is approximately ±3.4E38, while the range of float64 is approximately ±1.7E308.
// Therefore, all values within the range of float32 can be accurately represented as a float64.
//
// Parameters:
//   - value: The float32 value to be converted. The range of float32 is approximately ±3.4E38.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: No error is returned.
func Float32ToFloat64(value float32) (float64, error) {
	return float64(value), nil
}

// Float64ToFloat64 converts a float64 value to a float64 value.
//
// Float64ToFloat64 converts a float64 value to a float64 value.
// This function is a no-op, as the input and output types are the same.
//
// Parameters:
//   - value: The float64 value to be converted. The range of float64 is approximately ±1.7E308.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: No error is returned.
func Float64ToFloat64(value float64) (float64, error) {
	return value, nil
}

// IntToFloat64 converts an int value to a float64 value.
//
// IntToFloat64 converts an int value to a float64 value.
// The range of int is -2^31 to 2^31 - 1, while the range of float64 is approximately ±1.7E308.
// Therefore, all values within the range of int can be accurately represented as a float64.
//
// Parameters:
//   - value: The int value to be converted. The range of int is -2^31 to 2^31 - 1.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: No error is returned.
func IntToFloat64(value int) (float64, error) {
	return float64(value), nil
}

// Int8ToFloat64 converts an int8 value to a float64 value.
//
// Int8ToFloat64 converts an int8 value to a float64 value.
// The range of int8 is -128 to 127, while the range of float64 is approximately ±1.7E308.
// Therefore, all values within the range of int8 can be accurately represented as a float64.
//
// Parameters:
//   - value: The int8 value to be converted. The range of int8 is -128 to 127.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: No error is returned.
func Int8ToFloat64(value int8) (float64, error) {
	return float64(value), nil
}

// Int16ToFloat64 converts an int16 value to a float64 value.
//
// Int16ToFloat64 converts an int16 value to a float64 value.
// The range of int16 is -32768 to 32767, while the range of float64 is approximately ±1.7E308.
// Therefore, all values within the range of int16 can be accurately represented as a float64.
//
// Parameters:
//   - value: The int16 value to be converted. The range of int16 is -32768 to 32767.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: No error is returned.
func Int16ToFloat64(value int16) (float64, error) {
	return float64(value), nil
}

// Int32ToFloat64 converts an int32 value to a float64 value.
//
// Int32ToFloat64 converts an int32 value to a float64 value.
// The range of int32 is -2^31 to 2^31 - 1, while the range of float64 is approximately ±1.7E308.
// Therefore, all values within the range of int32 can be accurately represented as a float64.
//
// Parameters:
//   - value: The int32 value to be converted. The range of int32 is -2^31 to 2^31 - 1.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: No error is returned.
func Int32ToFloat64(value int32) (float64, error) {
	return float64(value), nil
}

// Int64ToFloat64 converts an int64 value to a float64 value.
//
// Int64ToFloat64 converts an int64 value to a float64 value.
// The range of int64 is -2^63 to 2^63 - 1, while the range of float64 is approximately ±1.7E308.
// Therefore, all values within the range of int64 can be accurately represented as a float64.
//
// Parameters:
//   - value: The int64 value to be converted. The range of int64 is -2^63 to 2^63 - 1.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: No error is returned.
func Int64ToFloat64(value int64) (float64, error) {
	return float64(value), nil
}

// StringToFloat64 converts a string value to a float64 value.
//
// StringToFloat64 converts a string value to a float64 value.
// If the input string cannot be parsed as a float64, it returns an error.
//
// Parameters:
//   - value: The string value to be converted.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: An error if the input string cannot be parsed as a float64.
func StringToFloat64(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

// UintToFloat64 converts a uint value to a float64 value.
//
// UintToFloat64 converts a uint value to a float64 value.
// The range of uint is 0 to 2^32 - 1, while the range of float64 is approximately ±1.7E308.
// Therefore, all values within the range of uint can be accurately represented as a float64.
//
// Parameters:
//   - value: The uint value to be converted. The range of uint is 0 to 2^32 - 1.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: No error is returned.
func UintToFloat64(value uint) (float64, error) {
	return float64(value), nil
}

// Uint8ToFloat64 converts a uint8 value to a float64 value.
//
// Uint8ToFloat64 converts a uint8 value to a float64 value.
// The range of uint8 is 0 to 255, while the range of float64 is approximately ±1.7E308.
// Therefore, all values within the range of uint8 can be accurately represented as a float64.
//
// Parameters:
//   - value: The uint8 value to be converted. The range of uint8 is 0 to 255.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: No error is returned.
func Uint8ToFloat64(value uint8) (float64, error) {
	return float64(value), nil
}

// Uint16ToFloat64 converts a uint16 value to a float64 value.
//
// Uint16ToFloat64 converts a uint16 value to a float64 value.
// The range of uint16 is 0 to 65535, while the range of float64 is approximately ±1.7E308.
// Therefore, all values within the range of uint16 can be accurately represented as a float64.
//
// Parameters:
//   - value: The uint16 value to be converted. The range of uint16 is 0 to 65535.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: No error is returned.
func Uint16ToFloat64(value uint16) (float64, error) {
	return float64(value), nil
}

// Uint32ToFloat64 converts a uint32 value to a float64 value.
//
// Uint32ToFloat64 converts a uint32 value to a float64 value.
// The range of uint32 is 0 to 2^32 - 1, while the range of float64 is approximately ±1.7E308.
// Therefore, all values within the range of uint32 can be accurately represented as a float64.
//
// Parameters:
//   - value: The uint32 value to be converted. The range of uint32 is 0 to 2^32 - 1.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: No error is returned.
func Uint32ToFloat64(value uint32) (float64, error) {
	return float64(value), nil
}

// Uint64ToFloat64 converts a uint64 value to a float64 value.
//
// Uint64ToFloat64 converts a uint64 value to a float64 value.
// The range of uint64 is 0 to 2^64 - 1, while the range of float64 is approximately ±1.7E308.
// Therefore, all values within the range of uint64 can be accurately represented as a float64.
//
// Parameters:
//   - value: The uint64 value to be converted. The range of uint64 is 0 to 2^64 - 1.
//
// Returns:
//   - float64: The converted float64 value. The range of float64 is approximately ±1.7E308.
//   - error: No error is returned.
func Uint64ToFloat64(value uint64) (float64, error) {
	return float64(value), nil
}
