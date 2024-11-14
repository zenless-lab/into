package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// TryIntoInt64 converts a value of any supported type to an int64.
//
// TryIntoInt64 attempts to convert a value of any supported type to an int64.
// The supported types are:
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
//   - int64: the converted int64 value.
//   - error: an error if the conversion failed or the input value is out of range.
//
// Example:
//
//	result, err := TryIntoInt64(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func TryIntoInt64[T convertable](value T) (int64, error) {
	result, err := toInt64(value)
	return result.(int64), err
}

func toInt64(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToInt64(value.(float64))
	case reflect.Float32:
		return Float32ToInt64(value.(float32))
	case reflect.Int:
		return IntToInt64(value.(int))
	case reflect.Int8:
		return Int8ToInt64(value.(int8))
	case reflect.Int16:
		return Int16ToInt64(value.(int16))
	case reflect.Int32:
		return Int32ToInt64(value.(int32))
	case reflect.Int64:
		return value.(int64), nil
	case reflect.Uint:
		return UintToInt64(value.(uint))
	case reflect.Uint8:
		return Uint8ToInt64(value.(uint8))
	case reflect.Uint16:
		return Uint16ToInt64(value.(uint16))
	case reflect.Uint32:
		return Uint32ToInt64(value.(uint32))
	case reflect.Uint64:
		return Uint64ToInt64(value.(uint64))
	case reflect.String:
		return StringToInt64(value.(string))
	case reflect.Bool:
		return BoolToInt64(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

// BoolToInt64 converts a bool value to an int64.
//
// BoolToInt64 converts a bool value to an int64.
// True will be converted to 1, and false will be converted to 0.
//
// Parameters:
//   - value: the bool value to be converted.
//
// Returns:
//   - int64: the converted int64 value.
//   - error: nil
//
// Example:
//
//	result, err := BoolToInt64(true)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 1
func BoolToInt64(value bool) (int64, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

// Float32ToInt64 converts a float32 value to an int64.
//
// Float32ToInt64 converts a float32 value to an int64.
// If the input value is outside the range that can be represented by an int64,
// it returns an error.
//
// Parameters:
//   - value: the float32 value to be converted. The range of float32 is approximately ±3.4E38.
//     The function ensures that the input value is within the range of int64,
//     which is approximately ±9.22E18.
//
// Returns:
//   - int64: the converted int64 value. The range of int64 is approximately ±9.22E18.
//   - error: an error if the input value is out of the int64 range.
//
// Example:
//
//	result, err := Float32ToInt64(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Float32ToInt64(value float32) (int64, error) {
	if value > math.MaxInt64 || value < math.MinInt64 {
		return 0, errors.New("value out of range for int64")
	}
	return int64(value), nil
}

// Float64ToInt64 converts a float64 value to an int64.
//
// Float64ToInt64 converts a float64 value to an int64.
// If the input value is outside the range that can be represented by an int64,
// it returns an error.
//
// Parameters:
//   - value: the float64 value to be converted. The range of float64 is approximately ±1.7E308.
//     The function ensures that the input value is within the range of int64,
//     which is approximately ±9.22E18.
//
// Returns:
//   - int64: the converted int64 value. The range of int64 is approximately ±9.22E18.
//   - error: an error if the input value is out of the int64 range.
//
// Example:
//
//	result, err := Float64ToInt64(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Float64ToInt64(value float64) (int64, error) {
	if value > math.MaxInt64 || value < math.MinInt64 {
		return 0, errors.New("value out of range for int64")
	}
	return int64(value), nil
}

// IntToInt64 converts an int value to an int64.
//
// IntToInt64 converts an int value to an int64.
// The range of int is approximately ±2.1E9, while the range of int64 is approximately ±9.22E18.
// As int value is always in range of int64, no check is performed before conversion.
//
// Parameters:
//   - value: the int value to be converted.
//
// Returns:
//   - int64: the converted int64 value.
//   - error: nil
//
// Example:
//
//	result, err := IntToInt64(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func IntToInt64(value int) (int64, error) {
	return int64(value), nil
}

// Int8ToInt64 converts an int8 value to an int64.
//
// Int8ToInt64 converts an int8 value to an int64.
// The range of int8 is approximately ±1.2E2, while the range of int64 is approximately ±9.22E18.
// As int8 value is always in range of int64, no check is performed before conversion.
//
// Parameters:
//   - value: the int8 value to be converted.
//
// Returns:
//   - int64: the converted int64 value.
//   - error: nil
//
// Example:
//
//	result, err := Int8ToInt64(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int8ToInt64(value int8) (int64, error) {
	return int64(value), nil
}

// Int16ToInt64 converts an int16 value to an int64.
//
// Int16ToInt64 converts an int16 value to an int64.
// The range of int16 is approximately ±3.2E4, while the range of int64 is approximately ±9.22E18.
// As int16 value is always in range of int64, no check is performed before conversion.
//
// Parameters:
//   - value: the int16 value to be converted.
//
// Returns:
//   - int64: the converted int64 value.
//   - error: nil
//
// Example:
//
//	result, err := Int16ToInt64(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int16ToInt64(value int16) (int64, error) {
	return int64(value), nil
}

// Int32ToInt64 converts an int32 value to an int64.
//
// Int32ToInt64 converts an int32 value to an int64.
// The range of int32 is approximately ±2.1E9, while the range of int64 is approximately ±9.22E18.
// As int32 value is always in range of int64, no check is performed before conversion.
//
// Parameters:
//   - value: the int32 value to be converted.
//
// Returns:
//   - int64: the converted int64 value.
//   - error: nil
//
// Example:
//
//	result, err := Int32ToInt64(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int32ToInt64(value int32) (int64, error) {
	return int64(value), nil
}

// Int64ToInt64 converts an int64 value to an int64.
//
// Int64ToInt64 converts an int64 value to an int64.
// No conversion is performed.
//
// Parameters:
//   - value: the int64 value to be converted.
//
// Returns:
//   - int64: the converted int64 value.
//   - error: nil
//
// Example:
//
//	result, err := Int64ToInt64(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int64ToInt64(value int64) (int64, error) {
	return value, nil
}

// StringToInt64 converts a string value to an int64.
//
// StringToInt64 converts a string value to an int64.
// The function uses strconv.ParseInt to perform the conversion.
//
// Parameters:
//   - value: the string value to be converted.
//
// Returns:
//   - int64: the converted int64 value.
//   - error: an error if the conversion failed.
//
// Example:
//
//	result, err := StringToInt64("123")
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func StringToInt64(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

// UintToInt64 converts a uint value to an int64.
//
// UintToInt64 converts a uint value to an int64.
// If the input value is outside the range that can be represented by an int64,
// it returns an error.
//
// Parameters:
//   - value: the uint value to be converted. The range of uint is approximately 0 to 4.2E9.
//     The function ensures that the input value is within the range of int64,
//     which is approximately ±9.22E18.
//
// Returns:
//   - int64: the converted int64 value. The range of int64 is approximately ±9.22E18.
//   - error: an error if the input value is out of the int64 range.
//
// Example:
//
//	result, err := UintToInt64(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func UintToInt64(value uint) (int64, error) {
	if value > math.MaxInt64 {
		return 0, errors.New("value exceeds int64 max limit")
	}
	return int64(value), nil
}

// Uint8ToInt64 converts a uint8 value to an int64.
//
// Uint8ToInt64 converts a uint8 value to an int64.
// The range of uint8 is approximately 0 to 2.5E2, while the range of int64 is approximately ±9.22E18.
// As uint8 value is always in range of int64, no check is performed before conversion.
//
// Parameters:
//   - value: the uint8 value to be converted.
//
// Returns:
//   - int64: the converted int64 value.
//   - error: nil
//
// Example:
//
//	result, err := Uint8ToInt64(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint8ToInt64(value uint8) (int64, error) {
	return int64(value), nil
}

// Uint16ToInt64 converts a uint16 value to an int64.
//
// Uint16ToInt64 converts a uint16 value to an int64.
// The range of uint16 is approximately 0 to 6.5E4, while the range of int64 is approximately ±9.22E18.
// As uint16 value is always in range of int64, no check is performed before conversion.
//
// Parameters:
//   - value: the uint16 value to be converted.
//
// Returns:
//   - int64: the converted int64 value.
//   - error: nil
//
// Example:
//
//	result, err := Uint16ToInt64(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint16ToInt64(value uint16) (int64, error) {
	return int64(value), nil
}

// Uint32ToInt64 converts a uint32 value to an int64.
//
// Uint32ToInt64 converts a uint32 value to an int64.
// The range of uint32 is approximately 0 to 4.2E9, while the range of int64 is approximately ±9.22E18.
// As uint32 value is always in range of int64, no check is performed before conversion.
//
// Parameters:
//   - value: the uint32 value to be converted.
//
// Returns:
//   - int64: the converted int64 value.
//   - error: nil
//
// Example:
//
//	result, err := Uint32ToInt64(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint32ToInt64(value uint32) (int64, error) {
	return int64(value), nil
}

// Uint64ToInt64 converts a uint64 value to an int64.
//
// Uint64ToInt64 converts a uint64 value to an int64.
// If the input value is outside the range that can be represented by an int64,
// it returns an error.
//
// Parameters:
//   - value: the uint64 value to be converted. The range of uint64 is approximately 0 to 1.8E19.
//     The function ensures that the input value is within the range of int64,
//     which is approximately ±9.22E18.
//
// Returns:
//   - int64: the converted int64 value. The range of int64 is approximately ±9.22E18.
//   - error: an error if the input value is out of the int64 range.
//
// Example:
//
//	result, err := Uint64ToInt64(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint64ToInt64(value uint64) (int64, error) {
	if value > math.MaxInt64 {
		return 0, errors.New("value exceeds int64 max limit")
	}
	return int64(value), nil
}
