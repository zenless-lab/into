package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// TryIntoInt attempts to convert the given value to an int.
//
// TryIntoInt attempts to convert the given value to an int.
// It supports various data types, including float64, float32, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, string, and bool.
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - int: the converted int value.
//   - error: an error if the conversion fails.
//
// Example:
//
//	result, err := TryIntoInt(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func TryIntoInt[T convertable](value T) (int, error) {
	result, err := toInt(value)
	return result.(int), err
}

// toInt converts the given value to an int.
//
// toInt converts the given value to an int.
// It supports various data types, including float64, float32, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, string, and bool.
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - any: the converted int value.
//   - error: an error if the conversion fails.
//
// Example:
//
//	result, err := toInt(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func toInt(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToInt(value.(float64))
	case reflect.Float32:
		return Float32ToInt(value.(float32))
	case reflect.Int:
		return value.(int), nil
	case reflect.Int8:
		return Int8ToInt(value.(int8))
	case reflect.Int16:
		return Int16ToInt(value.(int16))
	case reflect.Int32:
		return Int32ToInt(value.(int32))
	case reflect.Int64:
		return Int64ToInt(value.(int64))
	case reflect.Uint:
		return UintToInt(value.(uint))
	case reflect.Uint8:
		return Uint8ToInt(value.(uint8))
	case reflect.Uint16:
		return Uint16ToInt(value.(uint16))
	case reflect.Uint32:
		return Uint32ToInt(value.(uint32))
	case reflect.Uint64:
		return Uint64ToInt(value.(uint64))
	case reflect.String:
		return StringToInt(value.(string))
	case reflect.Bool:
		return BoolToInt(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

// BoolToInt converts a bool value to an int.
//
// BoolToInt converts a bool value to an int.
// It returns 1 if the value is true, and 0 if the value is false.
//
// Parameters:
//   - value: the bool value to be converted.
//
// Returns:
//   - int: the converted int value.
//   - error: nil.
//
// Example:
//
//	result, err := BoolToInt(true)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 1
func BoolToInt(value bool) (int, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

// Float32ToInt converts a float32 value to an int.
//
// Float32ToInt converts a float32 value to an int.
// If the input value is outside the range that can be represented by an int,
// it returns an error.
//
// Parameters:
//   - value: the float32 value to be converted. The range of float32 is approximately ±3.4E38.
//     The function ensures that the input value is within the range of int,
//     which is approximately -2.147e9 to 2.147e9.
//
// Returns:
//   - int: the converted int value. The range of int is approximately -2.147e9 to 2.147e9.
//   - error: an error if the input value is out of the int range.
//
// Example:
//
//	result, err := Float32ToInt(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Float32ToInt(value float32) (int, error) {
	if value > math.MaxInt || value < math.MinInt {
		return 0, errors.New("value out of range for int")
	}
	return int(value), nil
}

// Float64ToInt converts a float64 value to an int.
//
// Float64ToInt converts a float64 value to an int.
// If the input value is outside the range that can be represented by an int,
// it returns an error.
//
// Parameters:
//   - value: the float64 value to be converted. The range of float64 is approximately ±1.7E308.
//     The function ensures that the input value is within the range of int,
//     which is approximately -2.147e9 to 2.147e9.
//
// Returns:
//   - int: the converted int value. The range of int is approximately -2.147e9 to 2.147e9.
//   - error: an error if the input value is out of the int range.
//
// Example:
//
//	result, err := Float64ToInt(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Float64ToInt(value float64) (int, error) {
	if value > math.MaxInt || value < math.MinInt {
		return 0, errors.New("value out of range for int")
	}
	return int(value), nil
}

// IntToInt converts an int value to an int.
//
// IntToInt converts an int value to an int.
// It simply returns the input value.
//
// Parameters:
//   - value: the int value to be converted. The range of int is approximately -2.147e9 to 2.147e9.
//
// Returns:
//   - int: the converted int value. The range of int is approximately -2.147e9 to 2.147e9.
//   - error: nil.
//
// Example:
//
//	result, err := IntToInt(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func IntToInt(value int) (int, error) {
	return value, nil
}

// Int8ToInt converts an int8 value to an int.
//
// Int8ToInt converts an int8 value to an int.
// It simply returns the input value.
//
// Parameters:
//   - value: the int8 value to be converted. The range of int8 is -128 to 127.
//
// Returns:
//   - int: the converted int value. The range of int is approximately -2.147e9 to 2.147e9.
//   - error: nil.
//
// Example:
//
//	result, err := Int8ToInt(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int8ToInt(value int8) (int, error) {
	return int(value), nil
}

// Int16ToInt converts an int16 value to an int.
//
// Int16ToInt converts an int16 value to an int.
// It simply returns the input value.
//
// Parameters:
//   - value: the int16 value to be converted. The range of int16 is -32768 to 32767.
//
// Returns:
//   - int: the converted int value. The range of int is approximately -2.147e9 to 2.147e9.
//   - error: nil.
//
// Example:
//
//	result, err := Int16ToInt(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int16ToInt(value int16) (int, error) {
	return int(value), nil
}

// Int32ToInt converts an int32 value to an int.
//
// Int32ToInt converts an int32 value to an int.
// It simply returns the input value.
//
// Parameters:
//   - value: the int32 value to be converted. The range of int32 is -2147483648 to 2147483647.
//
// Returns:
//   - int: the converted int value. The range of int is approximately -2.147e9 to 2.147e9.
//   - error: nil.
//
// Example:
//
//	result, err := Int32ToInt(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int32ToInt(value int32) (int, error) {
	return int(value), nil
}

// Int64ToInt converts an int64 value to an int.
//
// Int64ToInt converts an int64 value to an int.
// If the input value is outside the range that can be represented by an int,
// it returns an error.
//
// Parameters:
//   - value: the int64 value to be converted. The range of int64 is approximately ±9.2E18.
//     The function ensures that the input value is within the range of int,
//     which is approximately -2.147e9 to 2.147e9.
//
// Returns:
//   - int: the converted int value. The range of int is approximately -2.147e9 to 2.147e9.
//   - error: an error if the input value is out of the int range.
//
// Example:
//
//	result, err := Int64ToInt(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int64ToInt(value int64) (int, error) {
	if value > math.MaxInt || value < math.MinInt {
		return 0, errors.New("value out of range for int")
	}
	return int(value), nil
}

// StringToInt converts a string value to an int.
//
// StringToInt converts a string value to an int.
// If the input value is not a valid integer, it returns an error.
//
// Parameters:
//   - value: the string value to be converted.
//
// Returns:
//   - int: the converted int value.
//   - error: an error if the input value is not a valid integer.
//
// Example:
//
//	result, err := StringToInt("123")
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func StringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

// UintToInt converts a uint value to an int.
//
// UintToInt converts a uint value to an int.
// If the input value is outside the range that can be represented by an int,
// it returns an error.
//
// Parameters:
//   - value: the uint value to be converted. The range of uint is 0 to 1.844e19.
//     The function ensures that the input value is within the range of int,
//     which is approximately -2.147e9 to 2.147e9.
//
// Returns:
//   - int: the converted int value. The range of int is approximately -2.147e9 to 2.147e9.
//   - error: an error if the input value is out of the int range.
//
// Example:
//
//	result, err := UintToInt(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func UintToInt(value uint) (int, error) {
	if value > math.MaxInt {
		return 0, errors.New("value exceeds int max limit")
	}
	return int(value), nil
}

// Uint8ToInt converts a uint8 value to an int.
//
// Uint8ToInt converts a uint8 value to an int.
// It simply returns the input value.
//
// Parameters:
//   - value: the uint8 value to be converted. The range of uint8 is 0 to 255.
//
// Returns:
//   - int: the converted int value. The range of int is approximately -2.147e9 to 2.147e9.
//   - error: nil.
//
// Example:
//
//	result, err := Uint8ToInt(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint8ToInt(value uint8) (int, error) {
	return int(value), nil
}

// Uint16ToInt converts a uint16 value to an int.
//
// Uint16ToInt converts a uint16 value to an int.
// It simply returns the input value.
//
// Parameters:
//   - value: the uint16 value to be converted. The range of uint16 is 0 to 65535.
//
// Returns:
//   - int: the converted int value. The range of int is approximately -2.147e9 to 2.147e9.
//   - error: nil.
//
// Example:
//
//	result, err := Uint16ToInt(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint16ToInt(value uint16) (int, error) {
	return int(value), nil
}

// Uint32ToInt converts a uint32 value to an int.
//
// Uint32ToInt converts a uint32 value to an int.
// If the input value is outside the range that can be represented by an int,
// it returns an error.
//
// Parameters:
//   - value: the uint32 value to be converted. The range of uint32 is 0 to 4.295e9.
//     The function ensures that the input value is within the range of int,
//     which is approximately -2.147e9 to 2.147e9.
//
// Returns:
//   - int: the converted int value. The range of int is approximately -2.147e9 to 2.147e9.
//   - error: an error if the input value is out of the int range.
//
// Example:
//
//	result, err := Uint32ToInt(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint32ToInt(value uint32) (int, error) {
	if uint64(value) > math.MaxInt {
		return 0, errors.New("value exceeds int max limit")
	}
	return int(value), nil
}

// Uint64ToInt converts a uint64 value to an int.
//
// Uint64ToInt converts a uint64 value to an int.
// If the input value is outside the range that can be represented by an int,
// it returns an error.
//
// Parameters:
//   - value: the uint64 value to be converted. The range of uint64 is 0 to 1.844e19.
//     The function ensures that the input value is within the range of int,
//     which is approximately -2.147e9 to 2.147e9.
//
// Returns:
//   - int: the converted int value. The range of int is approximately -2.147e9 to 2.147e9.
//   - error: an error if the input value is out of the int range.
//
// Example:
//
//	result, err := Uint64ToInt(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint64ToInt(value uint64) (int, error) {
	if value > math.MaxInt {
		return 0, errors.New("value exceeds int max limit")
	}
	return int(value), nil
}
