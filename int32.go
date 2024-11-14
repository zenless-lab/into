package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// TryIntoInt32 attempts to convert a value of type T to int32.
//
// TryIntoInt32 attempts to convert a value of type T to int32.
// The supported types for conversion are:
//   - float64
//   - float32
//   - int
//   - int8
//   - int16
//   - int32
//   - int64
//   - uint
//   - uint8
//   - uint16
//   - uint32
//   - uint64
//   - string
//   - bool
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: an error if the conversion fails or the input value is out of range for int32.
//
// Example:
//
//	result, err := TryIntoInt32(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func TryIntoInt32[T convertable](value T) (int32, error) {
	result, err := toInt32(value)
	return result.(int32), err
}

func toInt32(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToInt32(value.(float64))
	case reflect.Float32:
		return Float32ToInt32(value.(float32))
	case reflect.Int:
		return IntToInt32(value.(int))
	case reflect.Int8:
		return Int8ToInt32(value.(int8))
	case reflect.Int16:
		return Int16ToInt32(value.(int16))
	case reflect.Int32:
		return value.(int32), nil
	case reflect.Int64:
		return Int64ToInt32(value.(int64))
	case reflect.Uint:
		return UintToInt32(value.(uint))
	case reflect.Uint8:
		return Uint8ToInt32(value.(uint8))
	case reflect.Uint16:
		return Uint16ToInt32(value.(uint16))
	case reflect.Uint32:
		return Uint32ToInt32(value.(uint32))
	case reflect.Uint64:
		return Uint64ToInt32(value.(uint64))
	case reflect.String:
		return StringToInt32(value.(string))
	case reflect.Bool:
		return BoolToInt32(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

// BoolToInt32 converts a bool value to int32.
//
// BoolToInt32 converts a bool value to int32.
// If value is true, it returns 1. Otherwise, it returns 0.
//
// Parameters:
//   - value: the bool value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: nil.
//
// Example:
//
//	result, err := BoolToInt32(true)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 1
func BoolToInt32(value bool) (int32, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

// Float32ToInt32 converts a float32 value to int32.
//
// Float32ToInt32 converts a float32 value to int32.
// The range of float32 is approximately ±3.4E38.
// The function ensures that the input value is within the range of int32,
// which is approximately -2147483648 to 2147483647.
//
// Parameters:
//   - value: the float32 value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: an error if the input value is out of the int32 range.
//
// Example:
//
//	result, err := Float32ToInt32(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Float32ToInt32(value float32) (int32, error) {
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, errors.New("value out of range for int32")
	}
	return int32(value), nil
}

// Float64ToInt32 converts a float64 value to int32.
//
// Float64ToInt32 converts a float64 value to int32.
// The range of float64 is approximately ±1.7E308.
// The function ensures that the input value is within the range of int32,
// which is approximately -2147483648 to 2147483647.
//
// Parameters:
//   - value: the float64 value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: an error if the input value is out of the int32 range.
//
// Example:
//
//	result, err := Float64ToInt32(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Float64ToInt32(value float64) (int32, error) {
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, errors.New("value out of range for int32")
	}
	return int32(value), nil
}

// IntToInt32 converts an int value to int32.
//
// IntToInt32 converts an int value to int32.
// The range of int is machine-dependent, typically -2147483648 to 2147483647.
// The function ensures that the input value is within the range of int32,
// which is approximately -2147483648 to 2147483647.
//
// Parameters:
//   - value: the int value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: an error if the input value is out of the int32 range.
//
// Example:
//
//	result, err := IntToInt32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func IntToInt32(value int) (int32, error) {
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, errors.New("value out of range for int32")
	}
	return int32(value), nil
}

// Int8ToInt32 converts an int8 value to int32.
//
// Int8ToInt32 converts an int8 value to int32.
// The range of int8 is -128 to 127.
// The function ensures that the input value is within the range of int32,
// which is approximately -2147483648 to 2147483647.
//
// Parameters:
//   - value: the int8 value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: nil.
//
// Example:
//
//	result, err := Int8ToInt32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int8ToInt32(value int8) (int32, error) {
	return int32(value), nil
}

// Int16ToInt32 converts an int16 value to int32.
//
// Int16ToInt32 converts an int16 value to int32.
// The range of int16 is -32768 to 32767.
// The function ensures that the input value is within the range of int32,
// which is approximately -2147483648 to 2147483647.
//
// Parameters:
//   - value: the int16 value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: nil.
//
// Example:
//
//	result, err := Int16ToInt32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int16ToInt32(value int16) (int32, error) {
	return int32(value), nil
}

// Int32ToInt32 converts an int32 value to int32.
//
// Int32ToInt32 converts an int32 value to int32.
// The range of int32 is approximately -2147483648 to 2147483647.
// The function ensures that the input value is within the range of int32,
// which is approximately -2147483648 to 2147483647.
//
// Parameters:
//   - value: the int32 value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: nil.
//
// Example:
//
//	result, err := Int32ToInt32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int32ToInt32(value int32) (int32, error) {
	return value, nil
}

// Int64ToInt32 converts an int64 value to int32.
//
// Int64ToInt32 converts an int64 value to int32.
// The range of int64 is approximately ±9.2E18.
// The function ensures that the input value is within the range of int32,
// which is approximately -2147483648 to 2147483647.
//
// Parameters:
//   - value: the int64 value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: an error if the input value is out of the int32 range.
//
// Example:
//
//	result, err := Int64ToInt32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int64ToInt32(value int64) (int32, error) {
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, errors.New("value out of range for int32")
	}
	return int32(value), nil
}

// StringToInt32 converts a string value to int32.
//
// StringToInt32 converts a string value to int32.
// The input string should represent an integer in base 10.
//
// Parameters:
//   - value: the string value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: an error if the conversion fails.
//
// Example:
//
//	result, err := StringToInt32("123")
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func StringToInt32(value string) (int32, error) {
	i, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

// UintToInt32 converts a uint value to int32.
//
// UintToInt32 converts a uint value to int32.
// The range of uint is machine-dependent, typically 0 to 4294967295.
// The function ensures that the input value is within the range of int32,
// which is approximately -2147483648 to 2147483647.
//
// Parameters:
//   - value: the uint value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: an error if the input value is out of the int32 range.
//
// Example:
//
//	result, err := UintToInt32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func UintToInt32(value uint) (int32, error) {
	if value > math.MaxInt32 {
		return 0, errors.New("value exceeds int32 max limit")
	}
	return int32(value), nil
}

// Uint8ToInt32 converts a uint8 value to int32.
//
// Uint8ToInt32 converts a uint8 value to int32.
// The range of uint8 is 0 to 255.
// The function ensures that the input value is within the range of int32,
// which is approximately -2147483648 to 2147483647.
//
// Parameters:
//   - value: the uint8 value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: nil.
//
// Example:
//
//	result, err := Uint8ToInt32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint8ToInt32(value uint8) (int32, error) {
	return int32(value), nil
}

// Uint16ToInt32 converts a uint16 value to int32.
//
// Uint16ToInt32 converts a uint16 value to int32.
// The range of uint16 is 0 to 65535.
// The function ensures that the input value is within the range of int32,
// which is approximately -2147483648 to 2147483647.
//
// Parameters:
//   - value: the uint16 value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: nil.
//
// Example:
//
//	result, err := Uint16ToInt32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint16ToInt32(value uint16) (int32, error) {
	return int32(value), nil
}

// Uint32ToInt32 converts a uint32 value to int32.
//
// Uint32ToInt32 converts a uint32 value to int32.
// The range of uint32 is 0 to 4294967295.
// The function ensures that the input value is within the range of int32,
// which is approximately -2147483648 to 2147483647.
//
// Parameters:
//   - value: the uint32 value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: an error if the input value is out of the int32 range.
//
// Example:
//
//	result, err := Uint32ToInt32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint32ToInt32(value uint32) (int32, error) {
	if value > math.MaxInt32 {
		return 0, errors.New("value exceeds int32 max limit")
	}
	return int32(value), nil
}

// Uint64ToInt32 converts a uint64 value to int32.
//
// Uint64ToInt32 converts a uint64 value to int32.
// The range of uint64 is 0 to 18446744073709551615.
// The function ensures that the input value is within the range of int32,
// which is approximately -2147483648 to 2147483647.
//
// Parameters:
//   - value: the uint64 value to be converted.
//
// Returns:
//   - int32: the converted int32 value.
//   - error: an error if the input value is out of the int32 range.
//
// Example:
//
//	result, err := Uint64ToInt32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint64ToInt32(value uint64) (int32, error) {
	if value > math.MaxInt32 {
		return 0, errors.New("value exceeds int32 max limit")
	}
	return int32(value), nil
}
