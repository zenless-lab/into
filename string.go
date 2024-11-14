package into

import (
	"errors"
	"reflect"
	"strconv"
)

// TryIntoString attempts to convert a value to a string.
//
// TryIntoString attempts to convert a value of type T to a string.
// Supported types include:
//   - bool
//   - int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
//   - float32, float64
//   - string
//   - []byte
//   - []rune
//
// Parameters:
//   - value: The value to be converted. The range of the input value is determined by the type T.
//
// Returns:
//   - string: The converted string value.
//   - error: An error if the conversion fails.
func TryIntoString[T convertable | ~[]byte | ~[]rune](value T) (string, error) {
	result, err := toString(value)
	return result.(string), err
}

func toString(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToString(value.(float64))
	case reflect.Float32:
		return Float32ToString(value.(float32))
	case reflect.Int:
		return IntToString(value.(int))
	case reflect.Int8:
		return Int8ToString(value.(int8))
	case reflect.Int16:
		return Int16ToString(value.(int16))
	case reflect.Int32:
		return Int32ToString(value.(int32))
	case reflect.Int64:
		return Int64ToString(value.(int64))
	case reflect.Uint:
		return UintToString(value.(uint))
	case reflect.Uint8:
		return Uint8ToString(value.(uint8))
	case reflect.Uint16:
		return Uint16ToString(value.(uint16))
	case reflect.Uint32:
		return Uint32ToString(value.(uint32))
	case reflect.Uint64:
		return Uint64ToString(value.(uint64))
	case reflect.String:
		return value.(string), nil
	case reflect.Bool:
		return BoolToString(value.(bool))
	default:
		return "", errors.New("unsupported type")
	}
}

// BoolToString converts a boolean value to a string.
//
// BoolToString converts a boolean value to a string.
//
// Parameters:
//   - value: The boolean value to be converted. The range of bool is true or false.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func BoolToString(value bool) (string, error) {
	return strconv.FormatBool(value), nil
}

// BytesToString converts a byte slice to a string.
//
// BytesToString converts a byte slice to a string.
//
// Parameters:
//   - value: The byte slice to be converted. The range of []byte is determined by the length of the slice.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func BytesToString(value []byte) (string, error) {
	return string(value), nil
}

// ErrorToString converts an error to a string.
//
// ErrorToString converts an error to a string.
//
// Parameters:
//   - err: The error to be converted. The range of error is determined by the error message.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func ErrorToString(err error) (string, error) {
	return err.Error(), nil
}

// Float32ToString converts a float32 value to a string.
//
// Float32ToString converts a float32 value to a string.
//
// Parameters:
//   - value: The float32 value to be converted. The range of float32 is approximately ±3.4E38.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func Float32ToString(value float32) (string, error) {
	return strconv.FormatFloat(float64(value), 'f', -1, 32), nil
}

// Float64ToString converts a float64 value to a string.
//
// Float64ToString converts a float64 value to a string.
//
// Parameters:
//   - value: The float64 value to be converted. The range of float64 is approximately ±1.7E308.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func Float64ToString(value float64) (string, error) {
	return strconv.FormatFloat(value, 'f', -1, 64), nil
}

// IntToString converts an int value to a string.
//
// IntToString converts an int value to a string.
//
// Parameters:
//   - value: The int value to be converted. The range of int is platform-dependent.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func IntToString(value int) (string, error) {
	return strconv.Itoa(value), nil
}

// Int8ToString converts an int8 value to a string.
//
// Int8ToString converts an int8 value to a string.
//
// Parameters:
//   - value: The int8 value to be converted. The range of int8 is -128 to 127.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func Int8ToString(value int8) (string, error) {
	return strconv.Itoa(int(value)), nil
}

// Int16ToString converts an int16 value to a string.
//
// Int16ToString converts an int16 value to a string.
//
// Parameters:
//   - value: The int16 value to be converted. The range of int16 is -32,768 to 32,767.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func Int16ToString(value int16) (string, error) {
	return strconv.Itoa(int(value)), nil
}

// Int32ToString converts an int32 value to a string.
//
// Int32ToString converts an int32 value to a string.
//
// Parameters:
//   - value: The int32 value to be converted. The range of int32 is -2,147,483,648 to 2,147,483,647.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func Int32ToString(value int32) (string, error) {
	return strconv.Itoa(int(value)), nil
}

// Int64ToString converts an int64 value to a string.
//
// Int64ToString converts an int64 value to a string.
//
// Parameters:
//   - value: The int64 value to be converted. The range of int64 is -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func Int64ToString(value int64) (string, error) {
	return strconv.FormatInt(value, 10), nil
}

// RuneToString converts a rune value to a string.
//
// RuneToString converts a rune value to a string.
//
// Parameters:
//   - value: The rune value to be converted. The range of rune is 0 to 1,114,111.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func RuneToString(value rune) (string, error) {
	return string(value), nil
}

// StringToString converts a string value to a string.
//
// StringToString converts a string value to a string.
//
// Parameters:
//   - value: The string value to be converted. The range of string is unlimited.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func StringToString(value string) (string, error) {
	return value, nil
}

// UintToString converts a uint value to a string.
//
// UintToString converts a uint value to a string.
//
// Parameters:
//   - value: The uint value to be converted. The range of uint is 0 to 4,294,967,295.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func UintToString(value uint) (string, error) {
	return strconv.FormatUint(uint64(value), 10), nil
}

// Uint8ToString converts a uint8 value to a string.
//
// Uint8ToString converts a uint8 value to a string.
//
// Parameters:
//   - value: The uint8 value to be converted. The range of uint8 is 0 to 255.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func Uint8ToString(value uint8) (string, error) {
	return strconv.FormatUint(uint64(value), 10), nil
}

// Uint16ToString converts a uint16 value to a string.
//
// Uint16ToString converts a uint16 value to a string.
//
// Parameters:
//   - value: The uint16 value to be converted. The range of uint16 is 0 to 65,535.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func Uint16ToString(value uint16) (string, error) {
	return strconv.FormatUint(uint64(value), 10), nil
}

// Uint32ToString converts a uint32 value to a string.
//
// Uint32ToString converts a uint32 value to a string.
//
// Parameters:
//   - value: The uint32 value to be converted. The range of uint32 is 0 to 4,294,967,295.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func Uint32ToString(value uint32) (string, error) {
	return strconv.FormatUint(uint64(value), 10), nil
}

// Uint64ToString converts a uint64 value to a string.
//
// Uint64ToString converts a uint64 value to a string.
//
// Parameters:
//   - value: The uint64 value to be converted. The range of uint64 is 0 to 18,446,744,073,709,551,615.
//
// Returns:
//   - string: The converted string value.
//   - error: nil.
func Uint64ToString(value uint64) (string, error) {
	return strconv.FormatUint(value, 10), nil
}
