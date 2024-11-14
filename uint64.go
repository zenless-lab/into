package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// TryIntoUint64 attempts to convert a value of type T to a uint64 value.
//
// TryIntoUint64 attempts to convert a value of type T to a uint64 value.
// It supports various numeric and string types as input.
//
// Parameters:
//   - value: The value to be converted.
//
// Returns:
//   - uint64: The converted uint64 value if successful.
//   - error: An error if the conversion fails or the value is out of range for uint64.
func TryIntoUint64[T convertable](value T) (uint64, error) {
	result, err := toUint64(value)
	return result.(uint64), err
}

// toUint64 converts a value of any supported type to a uint64 value.
//
// toUint64 converts a value of any supported type to a uint64 value.
// It supports various numeric and string types as input.
//
// Parameters:
//   - value: The value to be converted.
//
// Returns:
//   - any: The converted uint64 value if successful.
//   - error: An error if the conversion fails or the value is out of range for uint64.
func toUint64(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint64(value.(float64))
	case reflect.Float32:
		return Float32ToUint64(value.(float32))
	case reflect.Int:
		return IntToUint64(value.(int))
	case reflect.Int8:
		return Int8ToUint64(value.(int8))
	case reflect.Int16:
		return Int16ToUint64(value.(int16))
	case reflect.Int32:
		return Int32ToUint64(value.(int32))
	case reflect.Int64:
		return Int64ToUint64(value.(int64))
	case reflect.Uint:
		return UintToUint64(value.(uint))
	case reflect.Uint8:
		return Uint8ToUint64(value.(uint8))
	case reflect.Uint16:
		return Uint16ToUint64(value.(uint16))
	case reflect.Uint32:
		return Uint32ToUint64(value.(uint32))
	case reflect.Uint64:
		return value.(uint64), nil
	case reflect.String:
		return StringToUint64(value.(string))
	case reflect.Bool:
		return BoolToUint64(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

// BoolToUint64 converts a boolean value to a uint64 value.
//
// BoolToUint64 converts a boolean value to a uint64 value.
// True is converted to 1, and false is converted to 0.
//
// Parameters:
//   - value: The boolean value to be converted.
//
// Returns:
//   - uint64: The converted uint64 value.
//   - error: nil.
func BoolToUint64(value bool) (uint64, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

// Float32ToUint64 converts a float32 value to a uint64 value.
//
// Float32ToUint64 converts a float32 value to a uint64 value.
// If the input value is negative or greater than math.MaxUint64, it returns an error.
//
// Parameters:
//   - value: The float32 value to be converted. The range of float32 is approximately ±3.4E38.
//     The function ensures that the input value is within the range of uint64,
//     which is 0 to 18,446,744,073,709,551,615.
//
// Returns:
//   - uint64: The converted uint64 value. The range of uint64 is 0 to 18,446,744,073,709,551,615.
//   - error: An error if the input value is out of range for uint64.
func Float32ToUint64(value float32) (uint64, error) {
	if value < 0 || value > math.MaxUint64 {
		return 0, errors.New("value out of range for uint64")
	}
	return uint64(value), nil
}

// Float64ToUint64 converts a float64 value to a uint64 value.
//
// Float64ToUint64 converts a float64 value to a uint64 value.
// If the input value is negative or greater than math.MaxUint64, it returns an error.
//
// Parameters:
//   - value: The float64 value to be converted. The range of float64 is approximately ±1.7E308.
//     The function ensures that the input value is within the range of uint64,
//     which is 0 to 18,446,744,073,709,551,615.
//
// Returns:
//   - uint64: The converted uint64 value. The range of uint64 is 0 to 18,446,744,073,709,551,615.
//   - error: An error if the input value is out of range for uint64.
func Float64ToUint64(value float64) (uint64, error) {
	if value < 0 || value > math.MaxUint64 {
		return 0, errors.New("value out of range for uint64")
	}
	return uint64(value), nil
}

// IntToUint64 converts an int value to a uint64 value.
//
// IntToUint64 converts an int value to a uint64 value.
// If the input value is negative, it returns an error.
//
// Parameters:
//   - value: The int value to be converted. The range of int is approximately ±2.1E9.
//     The function ensures that the input value is within the range of uint64,
//     which is 0 to 18,446,744,073,709,551,615.
//
// Returns:
//   - uint64: The converted uint64 value. The range of uint64 is 0 to 18,446,744,073,709,551,615.
//   - error: An error if the input value is out of range for uint64.
func IntToUint64(value int) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}

// Int8ToUint64 converts an int8 value to a uint64 value.
//
// Int8ToUint64 converts an int8 value to a uint64 value.
// If the input value is negative, it returns an error.
//
// Parameters:
//   - value: The int8 value to be converted. The range of int8 is -128 to 127.
//     The function ensures that the input value is within the range of uint64,
//     which is 0 to 18,446,744,073,709,551,615.
//
// Returns:
//   - uint64: The converted uint64 value. The range of uint64 is 0 to 18,446,744,073,709,551,615.
//   - error: An error if the input value is out of range for uint64.
func Int8ToUint64(value int8) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}

// Int16ToUint64 converts an int16 value to a uint64 value.
//
// Int16ToUint64 converts an int16 value to a uint64 value.
// If the input value is negative, it returns an error.
//
// Parameters:
//   - value: The int16 value to be converted. The range of int16 is -32768 to 32767.
//     The function ensures that the input value is within the range of uint64,
//     which is 0 to 18,446,744,073,709,551,615.
//
// Returns:
//   - uint64: The converted uint64 value. The range of uint64 is 0 to 18,446,744,073,709,551,615.
//   - error: An error if the input value is out of range for uint64.
func Int16ToUint64(value int16) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}

// Int32ToUint64 converts an int32 value to a uint64 value.
//
// Int32ToUint64 converts an int32 value to a uint64 value.
// If the input value is negative, it returns an error.
//
// Parameters:
//   - value: The int32 value to be converted. The range of int32 is approximately ±2.1E9.
//     The function ensures that the input value is within the range of uint64,
//     which is 0 to 18,446,744,073,709,551,615.
//
// Returns:
//   - uint64: The converted uint64 value. The range of uint64 is 0 to 18,446,744,073,709,551,615.
//   - error: An error if the input value is out of range for uint64.
func Int32ToUint64(value int32) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}

// Int64ToUint64 converts an int64 value to a uint64 value.
//
// Int64ToUint64 converts an int64 value to a uint64 value.
// If the input value is negative, it returns an error.
//
// Parameters:
//   - value: The int64 value to be converted. The range of int64 is approximately ±9.2E18.
//     The function ensures that the input value is within the range of uint64,
//     which is 0 to 18,446,744,073,709,551,615.
//
// Returns:
//   - uint64: The converted uint64 value. The range of uint64 is 0 to 18,446,744,073,709,551,615.
//   - error: An error if the input value is out of range for uint64.
func Int64ToUint64(value int64) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}

// StringToUint64 converts a string value to a uint64 value.
//
// StringToUint64 converts a string value to a uint64 value.
// If the input value is not a valid unsigned integer, it returns an error.
//
// Parameters:
//   - value: The string value to be converted.
//
// Returns:
//   - uint64: The converted uint64 value.
//   - error: An error if the input value is not a valid unsigned integer.
func StringToUint64(value string) (uint64, error) {
	return strconv.ParseUint(value, 10, 64)
}

// UintToUint64 converts a uint value to a uint64 value.
//
// UintToUint64 converts a uint value to a uint64 value.
//
// Parameters:
//   - value: The uint value to be converted. The range of uint is 0 to 4294967295.
//     The function ensures that the input value is within the range of uint64,
//     which is 0 to 18,446,744,073,709,551,615.
//
// Returns:
//   - uint64: The converted uint64 value. The range of uint64 is 0 to 18,446,744,073,709,551,615.
//   - error: nil.
func UintToUint64(value uint) (uint64, error) {
	return uint64(value), nil
}

// Uint8ToUint64 converts a uint8 value to a uint64 value.
//
// Uint8ToUint64 converts a uint8 value to a uint64 value.
//
// Parameters:
//   - value: The uint8 value to be converted. The range of uint8 is 0 to 255.
//     The function ensures that the input value is within the range of uint64,
//     which is 0 to 18,446,744,073,709,551,615.
//
// Returns:
//   - uint64: The converted uint64 value. The range of uint64 is 0 to 18,446,744,073,709,551,615.
//   - error: nil.
func Uint8ToUint64(value uint8) (uint64, error) {
	return uint64(value), nil
}

// Uint16ToUint64 converts a uint16 value to a uint64 value.
//
// Uint16ToUint64 converts a uint16 value to a uint64 value.
//
// Parameters:
//   - value: The uint16 value to be converted. The range of uint16 is 0 to 65535.
//     The function ensures that the input value is within the range of uint64,
//     which is 0 to 18,446,744,073,709,551,615.
//
// Returns:
//   - uint64: The converted uint64 value. The range of uint64 is 0 to 18,446,744,073,709,551,615.
//   - error: nil.
func Uint16ToUint64(value uint16) (uint64, error) {
	return uint64(value), nil
}

// Uint32ToUint64 converts a uint32 value to a uint64 value.
//
// Uint32ToUint64 converts a uint32 value to a uint64 value.
//
// Parameters:
//   - value: The uint32 value to be converted. The range of uint32 is 0 to 4294967295.
//     The function ensures that the input value is within the range of uint64,
//     which is 0 to 18,446,744,073,709,551,615.
//
// Returns:
//   - uint64: The converted uint64 value. The range of uint64 is 0 to 18,446,744,073,709,551,615.
//   - error: nil.
func Uint32ToUint64(value uint32) (uint64, error) {
	return uint64(value), nil
}

// Uint64ToUint64 converts a uint64 value to a uint64 value.
//
// Uint64ToUint64 converts a uint64 value to a uint64 value.
//
// Parameters:
//   - value: The uint64 value to be converted. The range of uint64 is 0 to 18,446,744,073,709,551,615.
//
// Returns:
//   - uint64: The converted uint64 value.
//   - error: nil.
func Uint64ToUint64(value uint64) (uint64, error) {
	return value, nil
}
