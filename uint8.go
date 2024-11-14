package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// TryIntoUint8 attempts to convert the given value to uint8.
//
// TryIntoUint8 attempts to convert the given value to uint8.
// It returns the converted value and an error if the conversion fails.
//
// Parameters:
//   - value: The value to be converted.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: An error if the conversion fails.
func TryIntoUint8[T convertable](value T) (uint8, error) {
	result, err := toUint8(value)
	return result.(uint8), err
}

// toUint8 converts a value of any type to a uint8.
//
// toUint8 converts a value of any type to a uint8.
// It supports conversion from various numeric types, strings and booleans.
//
// Parameters:
//   - value: The value to be converted.
//
// Returns:
//   - any: The converted uint8 value.
//   - error: An error if the conversion fails.
func toUint8(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint8(value.(float64))
	case reflect.Float32:
		return Float32ToUint8(value.(float32))
	case reflect.Int:
		return IntToUint8(value.(int))
	case reflect.Int8:
		return Int8ToUint8(value.(int8))
	case reflect.Int16:
		return Int16ToUint8(value.(int16))
	case reflect.Int32:
		return Int32ToUint8(value.(int32))
	case reflect.Int64:
		return Int64ToUint8(value.(int64))
	case reflect.Uint:
		return UintToUint8(value.(uint))
	case reflect.Uint8:
		return value.(uint8), nil
	case reflect.Uint16:
		return Uint16ToUint8(value.(uint16))
	case reflect.Uint32:
		return Uint32ToUint8(value.(uint32))
	case reflect.Uint64:
		return Uint64ToUint8(value.(uint64))
	case reflect.String:
		return StringToUint8(value.(string))
	case reflect.Bool:
		return BoolToUint8(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

// BoolToUint8 converts a bool to uint8.
//
// BoolToUint8 converts a bool to uint8.
// True is converted to 1 and false is converted to 0.
//
// Parameters:
//   - value: The bool value to be converted.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: nil.
func BoolToUint8(value bool) (uint8, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

// Float32ToUint8 converts a float32 to uint8.
//
// Float32ToUint8 converts a float32 to uint8.
// The function ensures that the input value is within the range of uint8.
//
// Parameters:
//   - value: The float32 value to be converted. The range of float32 is approximately ±3.4E38.
//     The function ensures that the input value is within the range of uint8,
//     which is 0 to 255.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: An error if the input value is out of the uint8 range.
func Float32ToUint8(value float32) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}

// Float64ToUint8 converts a float64 to uint8.
//
// Float64ToUint8 converts a float64 to uint8.
// The function ensures that the input value is within the range of uint8.
//
// Parameters:
//   - value: The float64 value to be converted. The range of float64 is approximately ±1.7E308.
//     The function ensures that the input value is within the range of uint8,
//     which is 0 to 255.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: An error if the input value is out of the uint8 range.
func Float64ToUint8(value float64) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}

// IntToUint8 converts a int to uint8.
//
// IntToUint8 converts a int to uint8.
// The function ensures that the input value is within the range of uint8.
//
// Parameters:
//   - value: The int value to be converted. The range of int is -2147483648 to 2147483647.
//     The function ensures that the input value is within the range of uint8,
//     which is 0 to 255.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: An error if the input value is out of the uint8 range.
func IntToUint8(value int) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}

// Int8ToUint8 converts a int8 to uint8.
//
// Int8ToUint8 converts a int8 to uint8.
// The function ensures that the input value is within the range of uint8.
//
// Parameters:
//   - value: The int8 value to be converted. The range of int8 is -128 to 127.
//     The function ensures that the input value is within the range of uint8,
//     which is 0 to 255.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: An error if the input value is out of the uint8 range.
func Int8ToUint8(value int8) (uint8, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint8")
	}
	return uint8(value), nil
}

// Int16ToUint8 converts a int16 to uint8.
//
// Int16ToUint8 converts a int16 to uint8.
// The function ensures that the input value is within the range of uint8.
//
// Parameters:
//   - value: The int16 value to be converted. The range of int16 is -32768 to 32767.
//     The function ensures that the input value is within the range of uint8,
//     which is 0 to 255.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: An error if the input value is out of the uint8 range.
func Int16ToUint8(value int16) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}

// Int32ToUint8 converts a int32 to uint8.
//
// Int32ToUint8 converts a int32 to uint8.
// The function ensures that the input value is within the range of uint8.
//
// Parameters:
//   - value: The int32 value to be converted. The range of int32 is -2147483648 to 2147483647.
//     The function ensures that the input value is within the range of uint8,
//     which is 0 to 255.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: An error if the input value is out of the uint8 range.
func Int32ToUint8(value int32) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}

// Int64ToUint8 converts a int64 to uint8.
//
// Int64ToUint8 converts a int64 to uint8.
// The function ensures that the input value is within the range of uint8.
//
// Parameters:
//   - value: The int64 value to be converted. The range of int64 is -9223372036854775808 to 9223372036854775807.
//     The function ensures that the input value is within the range of uint8,
//     which is 0 to 255.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: An error if the input value is out of the uint8 range.
func Int64ToUint8(value int64) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}

// StringToUint8 converts a string to uint8.
//
// StringToUint8 converts a string to uint8.
// The function ensures that the input string represents a valid unsigned integer.
//
// Parameters:
//   - value: The string value to be converted.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: An error if the input string is not a valid unsigned integer or if the value is out of the uint8 range.
func StringToUint8(value string) (uint8, error) {
	i, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(i), nil
}

// UintToUint8 converts a uint to uint8.
//
// UintToUint8 converts a uint to uint8.
// The function ensures that the input value is within the range of uint8.
//
// Parameters:
//   - value: The uint value to be converted. The range of uint is 0 to 4294967295.
//     The function ensures that the input value is within the range of uint8,
//     which is 0 to 255.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: An error if the input value is out of the uint8 range.
func UintToUint8(value uint) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, errors.New("value exceeds uint8 max limit")
	}
	return uint8(value), nil
}

// Uint8ToUint8 converts a uint8 to uint8.
//
// Uint8ToUint8 converts a uint8 to uint8.
// The function simply returns the input value.
//
// Parameters:
//   - value: The uint8 value to be converted.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: nil.
func Uint8ToUint8(value uint8) (uint8, error) {
	return value, nil
}

// Uint16ToUint8 converts a uint16 to uint8.
//
// Uint16ToUint8 converts a uint16 to uint8.
// The function ensures that the input value is within the range of uint8.
//
// Parameters:
//   - value: The uint16 value to be converted. The range of uint16 is 0 to 65535.
//     The function ensures that the input value is within the range of uint8,
//     which is 0 to 255.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: An error if the input value is out of the uint8 range.
func Uint16ToUint8(value uint16) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, errors.New("value exceeds uint8 max limit")
	}
	return uint8(value), nil
}

// Uint32ToUint8 converts a uint32 to uint8.
//
// Uint32ToUint8 converts a uint32 to uint8.
// The function ensures that the input value is within the range of uint8.
//
// Parameters:
//   - value: The uint32 value to be converted. The range of uint32 is 0 to 4294967295.
//     The function ensures that the input value is within the range of uint8,
//     which is 0 to 255.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: An error if the input value is out of the uint8 range.
func Uint32ToUint8(value uint32) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, errors.New("value exceeds uint8 max limit")
	}
	return uint8(value), nil
}

// Uint64ToUint8 converts a uint64 to uint8.
//
// Uint64ToUint8 converts a uint64 to uint8.
// The function ensures that the input value is within the range of uint8.
//
// Parameters:
//   - value: The uint64 value to be converted. The range of uint64 is 0 to 18446744073709551615.
//     The function ensures that the input value is within the range of uint8,
//     which is 0 to 255.
//
// Returns:
//   - uint8: The converted uint8 value.
//   - error: An error if the input value is out of the uint8 range.
func Uint64ToUint8(value uint64) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, errors.New("value exceeds uint8 max limit")
	}
	return uint8(value), nil
}
