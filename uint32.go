package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// TryIntoUint32 attempts to convert a value of any type to uint32.
//
// TryIntoUint32 attempts to convert a value of any type to uint32.
// It supports converting values of types bool, float32, float64, int, int8,
// int16, int32, int64, uint, uint8, uint16, uint32, uint64, and string.
// If the input value is outside the range that can be represented by a uint32,
// it returns an error.
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: an error if the input value is out of the uint32 range or the type is unsupported.
//
// Example:
//
//	result, err := TryIntoUint32(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func TryIntoUint32[T convertable](value T) (uint32, error) {
	result, err := toUint32(value)
	return result.(uint32), err
}

func toUint32(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint32(value.(float64))
	case reflect.Float32:
		return Float32ToUint32(value.(float32))
	case reflect.Int:
		return IntToUint32(value.(int))
	case reflect.Int8:
		return Int8ToUint32(value.(int8))
	case reflect.Int16:
		return Int16ToUint32(value.(int16))
	case reflect.Int32:
		return Int32ToUint32(value.(int32))
	case reflect.Int64:
		return Int64ToUint32(value.(int64))
	case reflect.Uint:
		return UintToUint32(value.(uint))
	case reflect.Uint8:
		return Uint8ToUint32(value.(uint8))
	case reflect.Uint16:
		return Uint16ToUint32(value.(uint16))
	case reflect.Uint32:
		return value.(uint32), nil
	case reflect.Uint64:
		return Uint64ToUint32(value.(uint64))
	case reflect.String:
		return StringToUint32(value.(string))
	case reflect.Bool:
		return BoolToUint32(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

// BoolToUint32 converts a bool value to a uint32 value.
//
// BoolToUint32 converts a bool value to a uint32 value.
// True is converted to 1, and false is converted to 0.
//
// Parameters:
//   - value: the bool value to be converted.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: no error is returned.
//
// Example:
//
//	result, err := BoolToUint32(true)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 1
func BoolToUint32(value bool) (uint32, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

// Float32ToUint32 converts a float32 value to a uint32 value.
//
// Float32ToUint32 converts a float32 value to a uint32 value.
// If the input value is outside the range that can be represented by a uint32,
// it returns an error.
//
// Parameters:
//   - value: the float32 value to be converted. The range of float32 is approximately ±3.4E38.
//     The function ensures that the input value is within the range of uint32,
//     which is approximately 0 to 4.2E9.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: an error if the input value is out of the uint32 range.
//
// Example:
//
//	result, err := Float32ToUint32(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Float32ToUint32(value float32) (uint32, error) {
	if value < 0 || value > math.MaxUint32 {
		return 0, errors.New("value out of range for uint32")
	}
	return uint32(value), nil
}

// Float64ToUint32 converts a float64 value to a uint32 value.
//
// Float64ToUint32 converts a float64 value to a uint32 value.
// If the input value is outside the range that can be represented by a uint32,
// it returns an error.
//
// Parameters:
//   - value: the float64 value to be converted. The range of float64 is approximately ±1.7E308.
//     The function ensures that the input value is within the range of uint32,
//     which is approximately 0 to 4.2E9.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: an error if the input value is out of the uint32 range.
//
// Example:
//
//	result, err := Float64ToUint32(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Float64ToUint32(value float64) (uint32, error) {
	if value < 0 || value > math.MaxUint32 {
		return 0, errors.New("value out of range for uint32")
	}
	return uint32(value), nil
}

// IntToUint32 converts an int value to a uint32 value.
//
// IntToUint32 converts an int value to a uint32 value.
// If the input value is outside the range that can be represented by a uint32,
// it returns an error.
//
// Parameters:
//   - value: the int value to be converted. The range of int is approximately -9.2E18 to 9.2E18.
//     The function ensures that the input value is within the range of uint32,
//     which is approximately 0 to 4.2E9.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: an error if the input value is out of the uint32 range.
//
// Example:
//
//	result, err := IntToUint32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func IntToUint32(value int) (uint32, error) {
	if value < 0 || value > math.MaxUint32 {
		return 0, errors.New("value out of range for uint32")
	}
	return uint32(value), nil
}

// Int8ToUint32 converts an int8 value to a uint32 value.
//
// Int8ToUint32 converts an int8 value to a uint32 value.
// If the input value is negative, it returns an error.
//
// Parameters:
//   - value: the int8 value to be converted. The range of int8 is approximately -128 to 127.
//     The function ensures that the input value is non-negative,
//     which is within the range of uint32, which is approximately 0 to 4.2E9.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: an error if the input value is negative.
//
// Example:
//
//	result, err := Int8ToUint32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int8ToUint32(value int8) (uint32, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint32")
	}
	return uint32(value), nil
}

// Int16ToUint32 converts an int16 value to a uint32 value.
//
// Int16ToUint32 converts an int16 value to a uint32 value.
// If the input value is negative, it returns an error.
//
// Parameters:
//   - value: the int16 value to be converted. The range of int16 is approximately -32768 to 32767.
//     The function ensures that the input value is non-negative,
//     which is within the range of uint32, which is approximately 0 to 4.2E9.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: an error if the input value is negative.
//
// Example:
//
//	result, err := Int16ToUint32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int16ToUint32(value int16) (uint32, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint32")
	}
	return uint32(value), nil
}

// Int32ToUint32 converts an int32 value to a uint32 value.
//
// Int32ToUint32 converts an int32 value to a uint32 value.
// If the input value is negative, it returns an error.
//
// Parameters:
//   - value: the int32 value to be converted. The range of int32 is approximately -2147483648 to 2147483647.
//     The function ensures that the input value is non-negative,
//     which is within the range of uint32, which is approximately 0 to 4.2E9.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: an error if the input value is negative.
//
// Example:
//
//	result, err := Int32ToUint32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int32ToUint32(value int32) (uint32, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint32")
	}
	return uint32(value), nil
}

// Int64ToUint32 converts an int64 value to a uint32 value.
//
// Int64ToUint32 converts an int64 value to a uint32 value.
// If the input value is outside the range that can be represented by a uint32,
// it returns an error.
//
// Parameters:
//   - value: the int64 value to be converted. The range of int64 is approximately -9.2E18 to 9.2E18.
//     The function ensures that the input value is within the range of uint32,
//     which is approximately 0 to 4.2E9.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: an error if the input value is out of the uint32 range.
//
// Example:
//
//	result, err := Int64ToUint32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int64ToUint32(value int64) (uint32, error) {
	if value < 0 || value > math.MaxUint32 {
		return 0, errors.New("value out of range for uint32")
	}
	return uint32(value), nil
}

// StringToUint32 converts a string value to a uint32 value.
//
// StringToUint32 converts a string value to a uint32 value.
// If the input value is not a valid unsigned integer, it returns an error.
//
// Parameters:
//   - value: the string value to be converted.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: an error if the input value is not a valid unsigned integer.
//
// Example:
//
//	result, err := StringToUint32("123")
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func StringToUint32(value string) (uint32, error) {
	i, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(i), nil
}

// UintToUint32 converts a uint value to a uint32 value.
//
// UintToUint32 converts a uint value to a uint32 value.
// If the input value is outside the range that can be represented by a uint32,
// it returns an error.
//
// Parameters:
//   - value: the uint value to be converted. The range of uint is approximately 0 to 1.8E19.
//     The function ensures that the input value is within the range of uint32,
//     which is approximately 0 to 4.2E9.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: an error if the input value is out of the uint32 range.
//
// Example:
//
//	result, err := UintToUint32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func UintToUint32(value uint) (uint32, error) {
	if value > math.MaxUint32 {
		return 0, errors.New("value exceeds uint32 max limit")
	}
	return uint32(value), nil
}

// Uint8ToUint32 converts a uint8 value to a uint32 value.
//
// Uint8ToUint32 converts a uint8 value to a uint32 value.
// The input value is always within the range of uint32.
//
// Parameters:
//   - value: the uint8 value to be converted. The range of uint8 is approximately 0 to 255.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: no error is returned.
//
// Example:
//
//	result, err := Uint8ToUint32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint8ToUint32(value uint8) (uint32, error) {
	return uint32(value), nil
}

// Uint16ToUint32 converts a uint16 value to a uint32 value.
//
// Uint16ToUint32 converts a uint16 value to a uint32 value.
// The input value is always within the range of uint32.
//
// Parameters:
//   - value: the uint16 value to be converted. The range of uint16 is approximately 0 to 65535.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: no error is returned.
//
// Example:
//
//	result, err := Uint16ToUint32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint16ToUint32(value uint16) (uint32, error) {
	return uint32(value), nil
}

// Uint32ToUint32 converts a uint32 value to a uint32 value.
//
// Uint32ToUint32 converts a uint32 value to a uint32 value.
// The input value is always within the range of uint32.
//
// Parameters:
//   - value: the uint32 value to be converted. The range of uint32 is approximately 0 to 4.2E9.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: no error is returned.
//
// Example:
//
//	result, err := Uint32ToUint32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint32ToUint32(value uint32) (uint32, error) {
	return value, nil
}

// Uint64ToUint32 converts a uint64 value to a uint32 value.
//
// Uint64ToUint32 converts a uint64 value to a uint32 value.
// If the input value is outside the range that can be represented by a uint32,
// it returns an error.
//
// Parameters:
//   - value: the uint64 value to be converted. The range of uint64 is approximately 0 to 1.8E19.
//     The function ensures that the input value is within the range of uint32,
//     which is approximately 0 to 4.2E9.
//
// Returns:
//   - uint32: the converted uint32 value. The range of uint32 is approximately 0 to 4.2E9.
//   - error: an error if the input value is out of the uint32 range.
//
// Example:
//
//	result, err := Uint64ToUint32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint64ToUint32(value uint64) (uint32, error) {
	if value > math.MaxUint32 {
		return 0, errors.New("value exceeds uint32 max limit")
	}
	return uint32(value), nil
}
