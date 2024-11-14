package into

import (
	"errors"
	"reflect"
	"strconv"
)

// TryIntoBool attempts to convert a value of any type to a boolean value.
//
// TryIntoBool attempts to convert a value of any type to a boolean value.
// It uses the `reflect` package to determine the type of the input value and
// calls the appropriate conversion function.
//
// Parameters:
//   - value: the value to be converted. It can be any type that can be
//		converted to a boolean, including int, float, string, and bool.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: an error if the conversion fails.
func TryIntoBool[T convertable](value T) (bool, error) {
	result, err := toBool(value)
	return result.(bool), err
}

// toBool converts a value of any type to a boolean value.
//
// toBool converts a value of any type to a boolean value.
// It uses the `reflect` package to determine the type of the input value and
// calls the appropriate conversion function.
//
// Parameters:
//   - value: the value to be converted. It can be any type that can be
//		converted to a boolean, including int, float, string, and bool.
//
// Returns:
//   - any: the converted boolean value.
//   - error: an error if the conversion fails.
func toBool(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToBool(value.(float64))
	case reflect.Float32:
		return Float32ToBool(value.(float32))
	case reflect.Int:
		return IntToBool(value.(int))
	case reflect.Int8:
		return Int8ToBool(value.(int8))
	case reflect.Int16:
		return Int16ToBool(value.(int16))
	case reflect.Int32:
		return Int32ToBool(value.(int32))
	case reflect.Int64:
		return Int64ToBool(value.(int64))
	case reflect.Uint:
		return UintToBool(value.(uint))
	case reflect.Uint8:
		return Uint8ToBool(value.(uint8))
	case reflect.Uint16:
		return Uint16ToBool(value.(uint16))
	case reflect.Uint32:
		return Uint32ToBool(value.(uint32))
	case reflect.Uint64:
		return Uint64ToBool(value.(uint64))
	case reflect.String:
		return StringToBool(value.(string))
	case reflect.Bool:
		return value.(bool), nil
	default:
		return false, errors.New("unsupported type")
	}
}

// BoolToBool converts a boolean value to a boolean value.
//
// BoolToBool converts a boolean value to a boolean value.
// It simply returns the input value without any modification.
//
// Parameters:
//   - value: the boolean value to be converted.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: nil.
func BoolToBool(value bool) (bool, error) {
	return value, nil
}

// Float32ToBool converts a float32 value to a boolean value.
//
// Float32ToBool converts a float32 value to a boolean value.
// It returns true if the input value is not equal to 0, and false otherwise.
//
// Parameters:
//   - value: the float32 value to be converted.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: nil.
func Float32ToBool(value float32) (bool, error) {
	return value != 0, nil
}

// Float64ToBool converts a float64 value to a boolean value.
//
// Float64ToBool converts a float64 value to a boolean value.
// It returns true if the input value is not equal to 0, and false otherwise.
//
// Parameters:
//   - value: the float64 value to be converted.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: nil.
func Float64ToBool(value float64) (bool, error) {
	return value != 0, nil
}

// IntToBool converts an int value to a boolean value.
//
// IntToBool converts an int value to a boolean value.
// It returns true if the input value is not equal to 0, and false otherwise.
//
// Parameters:
//   - value: the int value to be converted.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: nil.
func IntToBool(value int) (bool, error) {
	return value != 0, nil
}

// Int8ToBool converts an int8 value to a boolean value.
//
// Int8ToBool converts an int8 value to a boolean value.
// It returns true if the input value is not equal to 0, and false otherwise.
//
// Parameters:
//   - value: the int8 value to be converted.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: nil.
func Int8ToBool(value int8) (bool, error) {
	return value != 0, nil
}

// Int16ToBool converts an int16 value to a boolean value.
//
// Int16ToBool converts an int16 value to a boolean value.
// It returns true if the input value is not equal to 0, and false otherwise.
//
// Parameters:
//   - value: the int16 value to be converted.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: nil.
func Int16ToBool(value int16) (bool, error) {
	return value != 0, nil
}

// Int32ToBool converts an int32 value to a boolean value.
//
// Int32ToBool converts an int32 value to a boolean value.
// It returns true if the input value is not equal to 0, and false otherwise.
//
// Parameters:
//   - value: the int32 value to be converted.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: nil.
func Int32ToBool(value int32) (bool, error) {
	return value != 0, nil
}

// Int64ToBool converts an int64 value to a boolean value.
//
// Int64ToBool converts an int64 value to a boolean value.
// It returns true if the input value is not equal to 0, and false otherwise.
//
// Parameters:
//   - value: the int64 value to be converted.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: nil.
func Int64ToBool(value int64) (bool, error) {
	return value != 0, nil
}

// StringToBool converts a string value to a boolean value.
//
// StringToBool converts a string value to a boolean value.
// It uses the `strconv.ParseBool` function to parse the input string.
//
// Parameters:
//   - value: the string value to be converted.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: an error if the conversion fails.
func StringToBool(value string) (bool, error) {
	return strconv.ParseBool(value)
}

// UintToBool converts a uint value to a boolean value.
//
// UintToBool converts a uint value to a boolean value.
// It returns true if the input value is not equal to 0, and false otherwise.
//
// Parameters:
//   - value: the uint value to be converted.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: nil.
func UintToBool(value uint) (bool, error) {
	return value != 0, nil
}

// Uint8ToBool converts a uint8 value to a boolean value.
//
// Uint8ToBool converts a uint8 value to a boolean value.
// It returns true if the input value is not equal to 0, and false otherwise.
//
// Parameters:
//   - value: the uint8 value to be converted.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: nil.
func Uint8ToBool(value uint8) (bool, error) {
	return value != 0, nil
}

// Uint16ToBool converts a uint16 value to a boolean value.
//
// Uint16ToBool converts a uint16 value to a boolean value.
// It returns true if the input value is not equal to 0, and false otherwise.
//
// Parameters:
//   - value: the uint16 value to be converted.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: nil.
func Uint16ToBool(value uint16) (bool, error) {
	return value != 0, nil
}

// Uint32ToBool converts a uint32 value to a boolean value.
//
// Uint32ToBool converts a uint32 value to a boolean value.
// It returns true if the input value is not equal to 0, and false otherwise.
//
// Parameters:
//   - value: the uint32 value to be converted.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: nil.
func Uint32ToBool(value uint32) (bool, error) {
	return value != 0, nil
}

// Uint64ToBool converts a uint64 value to a boolean value.
//
// Uint64ToBool converts a uint64 value to a boolean value.
// It returns true if the input value is not equal to 0, and false otherwise.
//
// Parameters:
//   - value: the uint64 value to be converted.
//
// Returns:
//   - bool: the converted boolean value.
//   - error: nil.
func Uint64ToBool(value uint64) (bool, error) {
	return value != 0, nil
}