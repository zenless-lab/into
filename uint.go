package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// TryIntoUint attempts to convert a value of any type to a uint.
//
// TryIntoUint attempts to convert a value of any type to a uint.
// It uses reflection to determine the type of the input value and calls the appropriate conversion function.
//
// Parameters:
//   - value: the value to be converted. It can be of any type that can be converted to a uint.
//
// Returns:
//   - uint: the converted uint value.
//   - error: an error if the input value cannot be converted to a uint.
//
// Example:
//
//	result, err := TryIntoUint(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func TryIntoUint[T convertable](value T) (uint, error) {
	result, err := toUint(value)
	return result.(uint), err
}

// toUint is a private function that performs the actual conversion to uint.
//
// toUint is a private function that performs the actual conversion to uint.
// It uses a switch statement to determine the type of the input value and calls the appropriate conversion function.
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - any: the converted uint value.
//   - error: an error if the input value cannot be converted to a uint.
func toUint(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint(value.(float64))
	case reflect.Float32:
		return Float32ToUint(value.(float32))
	case reflect.Int:
		return IntToUint(value.(int))
	case reflect.Int8:
		return Int8ToUint(value.(int8))
	case reflect.Int16:
		return Int16ToUint(value.(int16))
	case reflect.Int32:
		return Int32ToUint(value.(int32))
	case reflect.Int64:
		return Int64ToUint(value.(int64))
	case reflect.Uint:
		return value.(uint), nil
	case reflect.Uint8:
		return Uint8ToUint(value.(uint8))
	case reflect.Uint16:
		return Uint16ToUint(value.(uint16))
	case reflect.Uint32:
		return Uint32ToUint(value.(uint32))
	case reflect.Uint64:
		return Uint64ToUint(value.(uint64))
	case reflect.String:
		return StringToUint(value.(string))
	case reflect.Bool:
		return BoolToUint(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

// BoolToUint converts a bool value to a uint value.
//
// BoolToUint converts a bool value to a uint value.
// If the input value is true, it returns 1. Otherwise, it returns 0.
//
// Parameters:
//   - value: the bool value to be converted.
//
// Returns:
//   - uint: the converted uint value.
//   - error: nil.
func BoolToUint(value bool) (uint, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

// Float32ToUint converts a float32 value to a uint value.
//
// Float32ToUint converts a float32 value to a uint value.
// If the input value is outside the range of uint, it returns an error.
//
// Parameters:
//   - value: the float32 value to be converted. The range of float32 is approximately ±3.4E38.
//     The function ensures that the input value is within the range of uint,
//     which is approximately 0 to 1.8E19.
//
// Returns:
//   - uint: the converted uint value. The range of uint is approximately 0 to 1.8E19.
//   - error: an error if the input value is out of the uint range.
func Float32ToUint(value float32) (uint, error) {
	if value < 0 || value > math.MaxUint {
		return 0, errors.New("value out of range for uint")
	}
	return uint(value), nil
}

// Float64ToUint converts a float64 value to a uint value.
//
// Float64ToUint converts a float64 value to a uint value.
// If the input value is outside the range of uint, it returns an error.
//
// Parameters:
//   - value: the float64 value to be converted. The range of float64 is approximately ±1.7E308.
//     The function ensures that the input value is within the range of uint,
//     which is approximately 0 to 1.8E19.
//
// Returns:
//   - uint: the converted uint value. The range of uint is approximately 0 to 1.8E19.
//   - error: an error if the input value is out of the uint range.
func Float64ToUint(value float64) (uint, error) {
	if value < 0 || value > math.MaxUint {
		return 0, errors.New("value out of range for uint")
	}
	return uint(value), nil
}

// IntToUint converts an int value to a uint value.
//
// IntToUint converts an int value to a uint value.
// If the input value is negative, it returns an error.
//
// Parameters:
//   - value: the int value to be converted. The range of int is approximately -9E18 to 9E18.
//     The function ensures that the input value is within the range of uint,
//     which is approximately 0 to 1.8E19.
//
// Returns:
//   - uint: the converted uint value. The range of uint is approximately 0 to 1.8E19.
//   - error: an error if the input value is negative.
func IntToUint(value int) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}

// Int8ToUint converts an int8 value to a uint value.
//
// Int8ToUint converts an int8 value to a uint value.
// If the input value is negative, it returns an error.
//
// Parameters:
//   - value: the int8 value to be converted. The range of int8 is approximately -128 to 127.
//     The function ensures that the input value is within the range of uint,
//     which is approximately 0 to 1.8E19.
//
// Returns:
//   - uint: the converted uint value. The range of uint is approximately 0 to 1.8E19.
//   - error: an error if the input value is negative.
func Int8ToUint(value int8) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}

// Int16ToUint converts an int16 value to a uint value.
//
// Int16ToUint converts an int16 value to a uint value.
// If the input value is negative, it returns an error.
//
// Parameters:
//   - value: the int16 value to be converted. The range of int16 is approximately -32768 to 32767.
//     The function ensures that the input value is within the range of uint,
//     which is approximately 0 to 1.8E19.
//
// Returns:
//   - uint: the converted uint value. The range of uint is approximately 0 to 1.8E19.
//   - error: an error if the input value is negative.
func Int16ToUint(value int16) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}

// Int32ToUint converts an int32 value to a uint value.
//
// Int32ToUint converts an int32 value to a uint value.
// If the input value is negative, it returns an error.
//
// Parameters:
//   - value: the int32 value to be converted. The range of int32 is approximately -2E9 to 2E9.
//     The function ensures that the input value is within the range of uint,
//     which is approximately 0 to 1.8E19.
//
// Returns:
//   - uint: the converted uint value. The range of uint is approximately 0 to 1.8E19.
//   - error: an error if the input value is negative.
func Int32ToUint(value int32) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}

// Int64ToUint converts an int64 value to a uint value.
//
// Int64ToUint converts an int64 value to a uint value.
// If the input value is negative, it returns an error.
//
// Parameters:
//   - value: the int64 value to be converted. The range of int64 is approximately -9E18 to 9E18.
//     The function ensures that the input value is within the range of uint,
//     which is approximately 0 to 1.8E19.
//
// Returns:
//   - uint: the converted uint value. The range of uint is approximately 0 to 1.8E19.
//   - error: an error if the input value is negative.
func Int64ToUint(value int64) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}

// StringToUint converts a string value to a uint value.
//
// StringToUint converts a string value to a uint value.
// If the input value is not a valid number or if the value exceeds the maximum value of uint, it returns an error.
//
// Parameters:
//   - value: the string value to be converted.
//
// Returns:
//   - uint: the converted uint value. The range of uint is approximately 0 to 1.8E19.
//   - error: an error if the input value is not a valid number or if the value exceeds the maximum value of uint.
func StringToUint(value string) (uint, error) {
	i, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, err
	}
	if i > math.MaxUint {
		return 0, errors.New("value exceeds uint max limit")
	}
	return uint(i), nil
}

// UintToUint converts a uint value to a uint value.
//
// UintToUint converts a uint value to a uint value.
// It simply returns the input value.
//
// Parameters:
//   - value: the uint value to be converted. The range of uint is approximately 0 to 1.8E19.
//
// Returns:
//   - uint: the converted uint value. The range of uint is approximately 0 to 1.8E19.
//   - error: nil.
func UintToUint(value uint) (uint, error) {
	return value, nil
}

// Uint8ToUint converts a uint8 value to a uint value.
//
// Uint8ToUint converts a uint8 value to a uint value.
// It simply returns the input value.
//
// Parameters:
//   - value: the uint8 value to be converted. The range of uint8 is approximately 0 to 255.
//
// Returns:
//   - uint: the converted uint value. The range of uint is approximately 0 to 1.8E19.
//   - error: nil.
func Uint8ToUint(value uint8) (uint, error) {
	return uint(value), nil
}

// Uint16ToUint converts a uint16 value to a uint value.
//
// Uint16ToUint converts a uint16 value to a uint value.
// It simply returns the input value.
//
// Parameters:
//   - value: the uint16 value to be converted. The range of uint16 is approximately 0 to 65535.
//
// Returns:
//   - uint: the converted uint value. The range of uint is approximately 0 to 1.8E19.
//   - error: nil.
func Uint16ToUint(value uint16) (uint, error) {
	return uint(value), nil
}

// Uint32ToUint converts a uint32 value to a uint value.
//
// Uint32ToUint converts a uint32 value to a uint value.
// It simply returns the input value.
//
// Parameters:
//   - value: the uint32 value to be converted. The range of uint32 is approximately 0 to 4.3E9.
//
// Returns:
//   - uint: the converted uint value. The range of uint is approximately 0 to 1.8E19.
//   - error: nil.
func Uint32ToUint(value uint32) (uint, error) {
	return uint(value), nil
}

// Uint64ToUint converts a uint64 value to a uint value.
//
// Uint64ToUint converts a uint64 value to a uint value.
// If the input value exceeds the maximum value of uint, it returns an error.
//
// Parameters:
//   - value: the uint64 value to be converted. The range of uint64 is approximately 0 to 1.8E19.
//
// Returns:
//   - uint: the converted uint value. The range of uint is approximately 0 to 1.8E19.
//   - error: an error if the input value exceeds the maximum value of uint.
func Uint64ToUint(value uint64) (uint, error) {
	if value > math.MaxUint {
		return 0, errors.New("value exceeds uint max limit")
	}
	return uint(value), nil
}
