package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// TryIntoInt8 attempts to convert a value of any type to int8.
//
// TryIntoInt8 attempts to convert a value of any type to int8.
// It returns the converted value and an error if the conversion fails.
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - int8: the converted int8 value.
//   - error: an error if the conversion fails.
func TryIntoInt8[T convertable](value T) (int8, error) {
	result, err := toInt8(value)
	return result.(int8), err
}

// toInt8 is a helper function for TryIntoInt8.
//
// toInt8 is a helper function for TryIntoInt8.
// It attempts to convert a value of any type to int8.
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - any: the converted int8 value.
//   - error: an error if the conversion fails.
func toInt8(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToInt8(value.(float64))
	case reflect.Float32:
		return Float32ToInt8(value.(float32))
	case reflect.Int:
		return IntToInt8(value.(int))
	case reflect.Int8:
		return value.(int8), nil
	case reflect.Int16:
		return Int16ToInt8(value.(int16))
	case reflect.Int32:
		return Int32ToInt8(value.(int32))
	case reflect.Int64:
		return Int64ToInt8(value.(int64))
	case reflect.Uint:
		return UintToInt8(value.(uint))
	case reflect.Uint8:
		return Uint8ToInt8(value.(uint8))
	case reflect.Uint16:
		return Uint16ToInt8(value.(uint16))
	case reflect.Uint32:
		return Uint32ToInt8(value.(uint32))
	case reflect.Uint64:
		return Uint64ToInt8(value.(uint64))
	case reflect.String:
		return StringToInt8(value.(string))
	case reflect.Bool:
		return BoolToInt8(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

// BoolToInt8 converts a bool value to int8.
//
// BoolToInt8 converts a bool value to int8.
//
// Parameters:
//   - value: the bool value to be converted.
//
// Returns:
//   - int8: the converted int8 value.
//   - error: nil.
func BoolToInt8(value bool) (int8, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

// Float32ToInt8 converts a float32 value to int8.
//
// Float32ToInt8 converts a float32 value to int8.
// If the input value is outside the range of int8,
// it returns an error.
//
// Parameters:
//   - value: the float32 value to be converted. The range of float32 is approximately ±3.4E38.
//     The function ensures that the input value is within the range of int8,
//     which is approximately -128 to 127.
//
// Returns:
//   - int8: the converted int8 value. The range of int8 is approximately -128 to 127.
//   - error: an error if the input value is out of the int8 range.
func Float32ToInt8(value float32) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

// Float64ToInt8 converts a float64 value to int8.
//
// Float64ToInt8 converts a float64 value to int8.
// If the input value is outside the range of int8,
// it returns an error.
//
// Parameters:
//   - value: the float64 value to be converted. The range of float64 is approximately ±1.7E308.
//     The function ensures that the input value is within the range of int8,
//     which is approximately -128 to 127.
//
// Returns:
//   - int8: the converted int8 value. The range of int8 is approximately -128 to 127.
//   - error: an error if the input value is out of the int8 range.
func Float64ToInt8(value float64) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

// IntToInt8 converts an int value to int8.
//
// IntToInt8 converts an int value to int8.
// If the input value is outside the range of int8,
// it returns an error.
//
// Parameters:
//   - value: the int value to be converted. The range of int is platform dependent.
//     The function ensures that the input value is within the range of int8,
//     which is approximately -128 to 127.
//
// Returns:
//   - int8: the converted int8 value. The range of int8 is approximately -128 to 127.
//   - error: an error if the input value is out of the int8 range.
func IntToInt8(value int) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

// Int8ToInt8 converts an int8 value to int8.
//
// Int8ToInt8 converts an int8 value to int8.
// It simply returns the input value.
//
// Parameters:
//   - value: the int8 value to be converted. The range of int8 is approximately -128 to 127.
//
// Returns:
//   - int8: the converted int8 value. The range of int8 is approximately -128 to 127.
//   - error: nil.
func Int8ToInt8(value int8) (int8, error) {
	return value, nil
}

// Int16ToInt8 converts an int16 value to int8.
//
// Int16ToInt8 converts an int16 value to int8.
// If the input value is outside the range of int8,
// it returns an error.
//
// Parameters:
//   - value: the int16 value to be converted. The range of int16 is approximately -32768 to 32767.
//     The function ensures that the input value is within the range of int8,
//     which is approximately -128 to 127.
//
// Returns:
//   - int8: the converted int8 value. The range of int8 is approximately -128 to 127.
//   - error: an error if the input value is out of the int8 range.
func Int16ToInt8(value int16) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

// Int32ToInt8 converts an int32 value to int8.
//
// Int32ToInt8 converts an int32 value to int8.
// If the input value is outside the range of int8,
// it returns an error.
//
// Parameters:
//   - value: the int32 value to be converted. The range of int32 is approximately -2147483648 to 2147483647.
//     The function ensures that the input value is within the range of int8,
//     which is approximately -128 to 127.
//
// Returns:
//   - int8: the converted int8 value. The range of int8 is approximately -128 to 127.
//   - error: an error if the input value is out of the int8 range.
func Int32ToInt8(value int32) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

// Int64ToInt8 converts an int64 value to int8.
//
// Int64ToInt8 converts an int64 value to int8.
// If the input value is outside the range of int8,
// it returns an error.
//
// Parameters:
//   - value: the int64 value to be converted. The range of int64 is approximately -9223372036854775808 to 9223372036854775807.
//     The function ensures that the input value is within the range of int8,
//     which is approximately -128 to 127.
//
// Returns:
//   - int8: the converted int8 value. The range of int8 is approximately -128 to 127.
//   - error: an error if the input value is out of the int8 range.
func Int64ToInt8(value int64) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

// StringToInt8 converts a string value to int8.
//
// StringToInt8 converts a string value to int8.
// If the input value is not a valid integer,
// it returns an error.
//
// Parameters:
//   - value: the string value to be converted.
//
// Returns:
//   - int8: the converted int8 value.
//   - error: an error if the input value is not a valid integer.
func StringToInt8(value string) (int8, error) {
	i, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(i), nil
}

// UintToInt8 converts a uint value to int8.
//
// UintToInt8 converts a uint value to int8.
// If the input value is outside the range of int8,
// it returns an error.
//
// Parameters:
//   - value: the uint value to be converted. The range of uint is platform dependent.
//     The function ensures that the input value is within the range of int8,
//     which is approximately -128 to 127.
//
// Returns:
//   - int8: the converted int8 value. The range of int8 is approximately -128 to 127.
//   - error: an error if the input value is out of the int8 range.
func UintToInt8(value uint) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}

// Uint8ToInt8 converts a uint8 value to int8.
//
// Uint8ToInt8 converts a uint8 value to int8.
// If the input value is outside the range of int8,
// it returns an error.
//
// Parameters:
//   - value: the uint8 value to be converted. The range of uint8 is approximately 0 to 255.
//     The function ensures that the input value is within the range of int8,
//     which is approximately -128 to 127.
//
// Returns:
//   - int8: the converted int8 value. The range of int8 is approximately -128 to 127.
//   - error: an error if the input value is out of the int8 range.
func Uint8ToInt8(value uint8) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}

// Uint16ToInt8 converts a uint16 value to int8.
//
// Uint16ToInt8 converts a uint16 value to int8.
// If the input value is outside the range of int8,
// it returns an error.
//
// Parameters:
//   - value: the uint16 value to be converted. The range of uint16 is approximately 0 to 65535.
//     The function ensures that the input value is within the range of int8,
//     which is approximately -128 to 127.
//
// Returns:
//   - int8: the converted int8 value. The range of int8 is approximately -128 to 127.
//   - error: an error if the input value is out of the int8 range.
func Uint16ToInt8(value uint16) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}

// Uint32ToInt8 converts a uint32 value to int8.
//
// Uint32ToInt8 converts a uint32 value to int8.
// If the input value is outside the range of int8,
// it returns an error.
//
// Parameters:
//   - value: the uint32 value to be converted. The range of uint32 is approximately 0 to 4294967295.
//     The function ensures that the input value is within the range of int8,
//     which is approximately -128 to 127.
//
// Returns:
//   - int8: the converted int8 value. The range of int8 is approximately -128 to 127.
//   - error: an error if the input value is out of the int8 range.
func Uint32ToInt8(value uint32) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}

// Uint64ToInt8 converts a uint64 value to int8.
//
// Uint64ToInt8 converts a uint64 value to int8.
// If the input value is outside the range of int8,
// it returns an error.
//
// Parameters:
//   - value: the uint64 value to be converted. The range of uint64 is approximately 0 to 18446744073709551615.
//     The function ensures that the input value is within the range of int8,
//     which is approximately -128 to 127.
//
// Returns:
//   - int8: the converted int8 value. The range of int8 is approximately -128 to 127.
//   - error: an error if the input value is out of the int8 range.
func Uint64ToInt8(value uint64) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}
