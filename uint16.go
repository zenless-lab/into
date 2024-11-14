package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// TryIntoUint16 attempts to convert the given value to uint16.
//
// TryIntoUint16 attempts to convert the given value to uint16.
// If the conversion is successful, it returns the converted value and nil error.
// Otherwise, it returns 0 and an error.
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - uint16: the converted uint16 value.
//   - error: an error if the conversion fails.
func TryIntoUint16[T convertable](value T) (uint16, error) {
	result, err := toUint16(value)
	return result.(uint16), err
}

// toUint16 converts the given value to uint16 based on its type.
//
// toUint16 converts the given value to uint16 based on its type.
// It supports various types including float64, float32, int, int8, int16, int32, int64,
// uint, uint8, uint16, uint32, uint64, string, and bool.
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - any: the converted uint16 value.
//   - error: an error if the conversion fails.
func toUint16(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint16(value.(float64))
	case reflect.Float32:
		return Float32ToUint16(value.(float32))
	case reflect.Int:
		return IntToUint16(value.(int))
	case reflect.Int8:
		return Int8ToUint16(value.(int8))
	case reflect.Int16:
		return Int16ToUint16(value.(int16))
	case reflect.Int32:
		return Int32ToUint16(value.(int32))
	case reflect.Int64:
		return Int64ToUint16(value.(int64))
	case reflect.Uint:
		return UintToUint16(value.(uint))
	case reflect.Uint8:
		return Uint8ToUint16(value.(uint8))
	case reflect.Uint16:
		return value.(uint16), nil
	case reflect.Uint32:
		return Uint32ToUint16(value.(uint32))
	case reflect.Uint64:
		return Uint64ToUint16(value.(uint64))
	case reflect.String:
		return StringToUint16(value.(string))
	case reflect.Bool:
		return BoolToUint16(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

// BoolToUint16 converts a bool value to a uint16 value.
//
// BoolToUint16 converts a bool value to a uint16 value.
// True is converted to 1 and false is converted to 0.
//
// Parameters:
//   - value: the bool value to be converted.
//
// Returns:
//   - uint16: the converted uint16 value.
//   - error: nil.
func BoolToUint16(value bool) (uint16, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

// Float32ToUint16 converts a float32 value to a uint16 value.
//
// Float32ToUint16 converts a float32 value to a uint16 value.
// It checks if the value is within the range of uint16 (0 to 65535).
// If the value is out of range, it returns an error.
//
// Parameters:
//   - value: the float32 value to be converted. The range of float32 is approximately ±3.4E38.
//     The function ensures that the input value is within the range of uint16,
//     which is 0 to 65535.
//
// Returns:
//   - uint16: the converted uint16 value. The range of uint16 is 0 to 65535.
//   - error: an error if the input value is out of the uint16 range.
func Float32ToUint16(value float32) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}

// Float64ToUint16 converts a float64 value to a uint16 value.
//
// Float64ToUint16 converts a float64 value to a uint16 value.
// It checks if the value is within the range of uint16 (0 to 65535).
// If the value is out of range, it returns an error.
//
// Parameters:
//   - value: the float64 value to be converted. The range of float64 is approximately ±1.7E308.
//     The function ensures that the input value is within the range of uint16,
//     which is 0 to 65535.
//
// Returns:
//   - uint16: the converted uint16 value. The range of uint16 is 0 to 65535.
//   - error: an error if the input value is out of the uint16 range.
func Float64ToUint16(value float64) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}

// IntToUint16 converts an int value to a uint16 value.
//
// IntToUint16 converts an int value to a uint16 value.
// It checks if the value is within the range of uint16 (0 to 65535).
// If the value is out of range, it returns an error.
//
// Parameters:
//   - value: the int value to be converted. The range of int is -2,147,483,648 to 2,147,483,647.
//     The function ensures that the input value is within the range of uint16,
//     which is 0 to 65535.
//
// Returns:
//   - uint16: the converted uint16 value. The range of uint16 is 0 to 65535.
//   - error: an error if the input value is out of the uint16 range.
func IntToUint16(value int) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}

// Int8ToUint16 converts an int8 value to a uint16 value.
//
// Int8ToUint16 converts an int8 value to a uint16 value.
// It checks if the value is within the range of uint16 (0 to 65535).
// If the value is out of range, it returns an error.
//
// Parameters:
//   - value: the int8 value to be converted. The range of int8 is -128 to 127.
//     The function ensures that the input value is within the range of uint16,
//     which is 0 to 65535.
//
// Returns:
//   - uint16: the converted uint16 value. The range of uint16 is 0 to 65535.
//   - error: an error if the input value is out of the uint16 range.
func Int8ToUint16(value int8) (uint16, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint16")
	}
	return uint16(value), nil
}

// Int16ToUint16 converts an int16 value to a uint16 value.
//
// Int16ToUint16 converts an int16 value to a uint16 value.
// It checks if the value is within the range of uint16 (0 to 65535).
// If the value is out of range, it returns an error.
//
// Parameters:
//   - value: the int16 value to be converted. The range of int16 is -32,768 to 32,767.
//     The function ensures that the input value is within the range of uint16,
//     which is 0 to 65535.
//
// Returns:
//   - uint16: the converted uint16 value. The range of uint16 is 0 to 65535.
//   - error: an error if the input value is out of the uint16 range.
func Int16ToUint16(value int16) (uint16, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint16")
	}
	return uint16(value), nil
}

// Int32ToUint16 converts an int32 value to a uint16 value.
//
// Int32ToUint16 converts an int32 value to a uint16 value.
// It checks if the value is within the range of uint16 (0 to 65535).
// If the value is out of range, it returns an error.
//
// Parameters:
//   - value: the int32 value to be converted. The range of int32 is -2,147,483,648 to 2,147,483,647.
//     The function ensures that the input value is within the range of uint16,
//     which is 0 to 65535.
//
// Returns:
//   - uint16: the converted uint16 value. The range of uint16 is 0 to 65535.
//   - error: an error if the input value is out of the uint16 range.
func Int32ToUint16(value int32) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}

// Int64ToUint16 converts an int64 value to a uint16 value.
//
// Int64ToUint16 converts an int64 value to a uint16 value.
// It checks if the value is within the range of uint16 (0 to 65535).
// If the value is out of range, it returns an error.
//
// Parameters:
//   - value: the int64 value to be converted. The range of int64 is -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807.
//     The function ensures that the input value is within the range of uint16,
//     which is 0 to 65535.
//
// Returns:
//   - uint16: the converted uint16 value. The range of uint16 is 0 to 65535.
//   - error: an error if the input value is out of the uint16 range.
func Int64ToUint16(value int64) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}

// StringToUint16 converts a string value to a uint16 value.
//
// StringToUint16 converts a string value to a uint16 value.
// It parses the string as an unsigned integer with base 10.
// If the string is invalid, it returns an error.
//
// Parameters:
//   - value: the string value to be converted.
//
// Returns:
//   - uint16: the converted uint16 value. The range of uint16 is 0 to 65535.
//   - error: an error if the string is invalid.
func StringToUint16(value string) (uint16, error) {
	i, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(i), nil
}

// UintToUint16 converts a uint value to a uint16 value.
//
// UintToUint16 converts a uint value to a uint16 value.
// It checks if the value is within the range of uint16 (0 to 65535).
// If the value is out of range, it returns an error.
//
// Parameters:
//   - value: the uint value to be converted. The range of uint is 0 to 4,294,967,295.
//     The function ensures that the input value is within the range of uint16,
//     which is 0 to 65535.
//
// Returns:
//   - uint16: the converted uint16 value. The range of uint16 is 0 to 65535.
//   - error: an error if the input value is out of the uint16 range.
func UintToUint16(value uint) (uint16, error) {
	if value > math.MaxUint16 {
		return 0, errors.New("value exceeds uint16 max limit")
	}
	return uint16(value), nil
}

// Uint8ToUint16 converts a uint8 value to a uint16 value.
//
// Uint8ToUint16 converts a uint8 value to a uint16 value.
// It always succeeds without any range check.
//
// Parameters:
//   - value: the uint8 value to be converted. The range of uint8 is 0 to 255.
//
// Returns:
//   - uint16: the converted uint16 value. The range of uint16 is 0 to 65535.
//   - error: nil.
func Uint8ToUint16(value uint8) (uint16, error) {
	return uint16(value), nil
}

// Uint16ToUint16 converts a uint16 value to a uint16 value.
//
// Uint16ToUint16 converts a uint16 value to a uint16 value.
// It always succeeds without any range check.
//
// Parameters:
//   - value: the uint16 value to be converted. The range of uint16 is 0 to 65535.
//
// Returns:
//   - uint16: the converted uint16 value. The range of uint16 is 0 to 65535.
//   - error: nil.
func Uint16ToUint16(value uint16) (uint16, error) {
	return value, nil
}

// Uint32ToUint16 converts a uint32 value to a uint16 value.
//
// Uint32ToUint16 converts a uint32 value to a uint16 value.
// It checks if the value is within the range of uint16 (0 to 65535).
// If the value is out of range, it returns an error.
//
// Parameters:
//   - value: the uint32 value to be converted. The range of uint32 is 0 to 4,294,967,295.
//     The function ensures that the input value is within the range of uint16,
//     which is 0 to 65535.
//
// Returns:
//   - uint16: the converted uint16 value. The range of uint16 is 0 to 65535.
//   - error: an error if the input value is out of the uint16 range.
func Uint32ToUint16(value uint32) (uint16, error) {
	if value > math.MaxUint16 {
		return 0, errors.New("value exceeds uint16 max limit")
	}
	return uint16(value), nil
}

// Uint64ToUint16 converts a uint64 value to a uint16 value.
//
// Uint64ToUint16 converts a uint64 value to a uint16 value.
// It checks if the value is within the range of uint16 (0 to 65535).
// If the value is out of range, it returns an error.
//
// Parameters:
//   - value: the uint64 value to be converted. The range of uint64 is 0 to 18,446,744,073,709,551,615.
//     The function ensures that the input value is within the range of uint16,
//     which is 0 to 65535.
//
// Returns:
//   - uint16: the converted uint16 value. The range of uint16 is 0 to 65535.
//   - error: an error if the input value is out of the uint16 range.
func Uint64ToUint16(value uint64) (uint16, error) {
	if value > math.MaxUint16 {
		return 0, errors.New("value exceeds uint16 max limit")
	}
	return uint16(value), nil
}
