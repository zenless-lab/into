package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// TryIntoInt16 attempts to convert a value of type T to an int16.
//
// TryIntoInt16 attempts to convert a value of type T to an int16.
// It supports conversion from various types, including float64, float32,
// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
// string, and bool.
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - int16: the converted int16 value.
//   - error: an error if the conversion fails or if the value is out of range
//     for int16.
func TryIntoInt16[T convertable](value T) (int16, error) {
	result, err := toInt16(value)
	return result.(int16), err
}

// toInt16 is a helper function for TryIntoInt16 that handles conversion
// based on the type of the input value.
//
// toInt16 is a helper function for TryIntoInt16 that handles conversion
// based on the type of the input value. It supports conversion from various
// types, including float64, float32, int, int8, int16, int32, int64, uint,
// uint8, uint16, uint32, uint64, string, and bool.
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - any: the converted value.
//   - error: an error if the conversion fails or if the value is out of range
//     for int16.
func toInt16(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToInt16(value.(float64))
	case reflect.Float32:
		return Float32ToInt16(value.(float32))
	case reflect.Int:
		return IntToInt16(value.(int))
	case reflect.Int8:
		return Int8ToInt16(value.(int8))
	case reflect.Int16:
		return value.(int16), nil
	case reflect.Int32:
		return Int32ToInt16(value.(int32))
	case reflect.Int64:
		return Int64ToInt16(value.(int64))
	case reflect.Uint:
		return UintToInt16(value.(uint))
	case reflect.Uint8:
		return Uint8ToInt16(value.(uint8))
	case reflect.Uint16:
		return Uint16ToInt16(value.(uint16))
	case reflect.Uint32:
		return Uint32ToInt16(value.(uint32))
	case reflect.Uint64:
		return Uint64ToInt16(value.(uint64))
	case reflect.String:
		return StringToInt16(value.(string))
	case reflect.Bool:
		return BoolToInt16(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

// BoolToInt16 converts a boolean value to an int16.
//
// BoolToInt16 converts a boolean value to an int16. True is converted to 1,
// and false is converted to 0.
//
// Parameters:
//   - value: the boolean value to be converted.
//
// Returns:
//   - int16: the converted int16 value.
//   - error: nil.
func BoolToInt16(value bool) (int16, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

// Float32ToInt16 converts a float32 value to an int16.
//
// Float32ToInt16 converts a float32 value to an int16. If the value is
// outside the range of int16, it returns an error.
//
// Parameters:
//   - value: the float32 value to be converted. The range of float32 is
//     approximately ±3.4E38. The function ensures that the input value is
//     within the range of int16, which is approximately -32,768 to 32,767.
//
// Returns:
//   - int16: the converted int16 value. The range of int16 is
//     approximately -32,768 to 32,767.
//   - error: an error if the input value is out of the int16 range.
func Float32ToInt16(value float32) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}

// Float64ToInt16 converts a float64 value to an int16.
//
// Float64ToInt16 converts a float64 value to an int16. If the value is
// outside the range of int16, it returns an error.
//
// Parameters:
//   - value: the float64 value to be converted. The range of float64 is
//     approximately ±1.7E308. The function ensures that the input value is
//     within the range of int16, which is approximately -32,768 to 32,767.
//
// Returns:
//   - int16: the converted int16 value. The range of int16 is
//     approximately -32,768 to 32,767.
//   - error: an error if the input value is out of the int16 range.
func Float64ToInt16(value float64) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}

// IntToInt16 converts an int value to an int16.
//
// IntToInt16 converts an int value to an int16. If the value is outside the
// range of int16, it returns an error.
//
// Parameters:
//   - value: the int value to be converted. The range of int is
//     approximately -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807.
//     The function ensures that the input value is within the range of int16,
//     which is approximately -32,768 to 32,767.
//
// Returns:
//   - int16: the converted int16 value. The range of int16 is
//     approximately -32,768 to 32,767.
//   - error: an error if the input value is out of the int16 range.
func IntToInt16(value int) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}

// Int8ToInt16 converts an int8 value to an int16.
//
// Int8ToInt16 converts an int8 value to an int16. The range of int8 is
// -128 to 127, which is a subset of the range of int16 (-32,768 to 32,767).
// Therefore, no range check is necessary.
//
// Parameters:
//   - value: the int8 value to be converted. The range of int8 is
//     approximately -128 to 127.
//
// Returns:
//   - int16: the converted int16 value. The range of int16 is
//     approximately -32,768 to 32,767.
//   - error: nil.
func Int8ToInt16(value int8) (int16, error) {
	return int16(value), nil
}

// Int16ToInt16 converts an int16 value to an int16.
//
// Int16ToInt16 converts an int16 value to an int16. This function simply
// returns the input value without any conversion or range check.
//
// Parameters:
//   - value: the int16 value to be converted. The range of int16 is
//     approximately -32,768 to 32,767.
//
// Returns:
//   - int16: the converted int16 value. The range of int16 is
//     approximately -32,768 to 32,767.
//   - error: nil.
func Int16ToInt16(value int16) (int16, error) {
	return value, nil
}

// Int32ToInt16 converts an int32 value to an int16.
//
// Int32ToInt16 converts an int32 value to an int16. If the value is outside
// the range of int16, it returns an error.
//
// Parameters:
//   - value: the int32 value to be converted. The range of int32 is
//     approximately -2,147,483,648 to 2,147,483,647. The function ensures
//     that the input value is within the range of int16, which is
//     approximately -32,768 to 32,767.
//
// Returns:
//   - int16: the converted int16 value. The range of int16 is
//     approximately -32,768 to 32,767.
//   - error: an error if the input value is out of the int16 range.
func Int32ToInt16(value int32) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}

// Int64ToInt16 converts an int64 value to an int16.
//
// Int64ToInt16 converts an int64 value to an int16. If the value is outside
// the range of int16, it returns an error.
//
// Parameters:
//   - value: the int64 value to be converted. The range of int64 is
//     approximately -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807.
//     The function ensures that the input value is within the range of int16,
//     which is approximately -32,768 to 32,767.
//
// Returns:
//   - int16: the converted int16 value. The range of int16 is
//     approximately -32,768 to 32,767.
//   - error: an error if the input value is out of the int16 range.
func Int64ToInt16(value int64) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}

// StringToInt16 converts a string value to an int16.
//
// StringToInt16 converts a string value to an int16. It parses the string
// using base 10. If the string cannot be parsed as an integer, it returns an
// error. If the parsed integer is outside the range of int16, it returns an
// error.
//
// Parameters:
//   - value: the string value to be converted.
//
// Returns:
//   - int16: the converted int16 value. The range of int16 is
//     approximately -32,768 to 32,767.
//   - error: an error if the string cannot be parsed as an integer or if the
//     parsed integer is out of the int16 range.
func StringToInt16(value string) (int16, error) {
	i, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(i), nil
}

// UintToInt16 converts a uint value to an int16.
//
// UintToInt16 converts a uint value to an int16. If the value is outside the
// range of int16, it returns an error.
//
// Parameters:
//   - value: the uint value to be converted. The range of uint is
//     approximately 0 to 18,446,744,073,709,551,615. The function ensures that
//     the input value is within the range of int16, which is
//     approximately -32,768 to 32,767.
//
// Returns:
//   - int16: the converted int16 value. The range of int16 is
//     approximately -32,768 to 32,767.
//   - error: an error if the input value is out of the int16 range.
func UintToInt16(value uint) (int16, error) {
	if value > math.MaxInt16 {
		return 0, errors.New("value exceeds int16 max limit")
	}
	return int16(value), nil
}

// Uint8ToInt16 converts a uint8 value to an int16.
//
// Uint8ToInt16 converts a uint8 value to an int16. The range of uint8 is
// 0 to 255, which is a subset of the range of int16 (-32,768 to 32,767).
// Therefore, no range check is necessary.
//
// Parameters:
//   - value: the uint8 value to be converted. The range of uint8 is
//     approximately 0 to 255.
//
// Returns:
//   - int16: the converted int16 value. The range of int16 is
//     approximately -32,768 to 32,767.
//   - error: nil.
func Uint8ToInt16(value uint8) (int16, error) {
	return int16(value), nil
}

// Uint16ToInt16 converts a uint16 value to an int16.
//
// Uint16ToInt16 converts a uint16 value to an int16. If the value is outside
// the range of int16, it returns an error.
//
// Parameters:
//   - value: the uint16 value to be converted. The range of uint16 is
//     approximately 0 to 65,535. The function ensures that the input value is
//     within the range of int16, which is approximately -32,768 to 32,767.
//
// Returns:
//   - int16: the converted int16 value. The range of int16 is
//     approximately -32,768 to 32,767.
//   - error: an error if the input value is out of the int16 range.
func Uint16ToInt16(value uint16) (int16, error) {
	if value > math.MaxInt16 {
		return 0, errors.New("value exceeds int16 max limit")
	}
	return int16(value), nil
}

// Uint32ToInt16 converts a uint32 value to an int16.
//
// Uint32ToInt16 converts a uint32 value to an int16. If the value is outside
// the range of int16, it returns an error.
//
// Parameters:
//   - value: the uint32 value to be converted. The range of uint32 is
//     approximately 0 to 4,294,967,295. The function ensures that the input
//     value is within the range of int16, which is approximately -32,768 to
//     32,767.
//
// Returns:
//   - int16: the converted int16 value. The range of int16 is
//     approximately -32,768 to 32,767.
//   - error: an error if the input value is out of the int16 range.
func Uint32ToInt16(value uint32) (int16, error) {
	if value > math.MaxInt16 {
		return 0, errors.New("value exceeds int16 max limit")
	}
	return int16(value), nil
}

// Uint64ToInt16 converts a uint64 value to an int16.
//
// Uint64ToInt16 converts a uint64 value to an int16. If the value is outside
// the range of int16, it returns an error.
//
// Parameters:
//   - value: the uint64 value to be converted. The range of uint64 is
//     approximately 0 to 18,446,744,073,709,551,615. The function ensures that
//     the input value is within the range of int16, which is
//     approximately -32,768 to 32,767.
//
// Returns:
//   - int16: the converted int16 value. The range of int16 is
//     approximately -32,768 to 32,767.
//   - error: an error if the input value is out of the int16 range.
func Uint64ToInt16(value uint64) (int16, error) {
	if value > math.MaxInt16 {
		return 0, errors.New("value exceeds int16 max limit")
	}
	return int16(value), nil
}
