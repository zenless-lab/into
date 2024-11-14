package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// TryIntoFloat32 tries to convert a value of type T to float32.
//
// TryIntoFloat32 attempts to convert a value of type T to float32. It uses
// reflection to determine the type of the input value and calls the
// appropriate conversion function.
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - float32: the converted float32 value.
//   - error: an error if the conversion fails or if the input value is
//     out of range for float32.
//
// Example:
//
//	result, err := TryIntoFloat32[int](123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func TryIntoFloat32[T convertable](value T) (float32, error) {
	resoult, err := toFloat32(value)
	return resoult.(float32), err
}

// toFloat32 converts a value to float32.
//
// toFloat32 converts a value to float32 using reflection to determine
// the type of the input value and calling the appropriate conversion function.
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - any: the converted value.
//   - error: an error if the conversion fails or if the input value is
//     out of range for float32.
func toFloat32(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToFloat32(value.(float64))
	case reflect.Float32:
		return value.(float32), nil
	case reflect.Int:
		return IntToFloat32(value.(int))
	case reflect.Int8:
		return Int8ToFloat32(value.(int8))
	case reflect.Int16:
		return Int16ToFloat32(value.(int16))
	case reflect.Int32:
		return Int32ToFloat32(value.(int32))
	case reflect.Int64:
		return Int64ToFloat32(value.(int64))
	case reflect.Uint:
		return UintToFloat32(value.(uint))
	case reflect.Uint8:
		return Uint8ToFloat32(value.(uint8))
	case reflect.Uint16:
		return Uint16ToFloat32(value.(uint16))
	case reflect.Uint32:
		return Uint32ToFloat32(value.(uint32))
	case reflect.Uint64:
		return Uint64ToFloat32(value.(uint64))
	case reflect.String:
		return StringToFloat32(value.(string))
	case reflect.Bool:
		return BoolToFloat32(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

// BoolToFloat32 converts a bool value to float32.
//
// BoolToFloat32 converts a bool value to float32. True is converted to
// 1.0 and false is converted to 0.0.
//
// Parameters:
//   - value: the bool value to be converted.
//
// Returns:
//   - float32: the converted float32 value.
//   - error: no error is returned.
//
// Example:
//
//	result, err := BoolToFloat32(true)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 1
func BoolToFloat32(value bool) (float32, error) {
	if value {
		return 1.0, nil
	}
	return 0.0, nil
}

// Float32ToFloat32 converts a float32 value to float32.
//
// Float32ToFloat32 converts a float32 value to float32. This function
// does not perform any conversion and simply returns the input value.
//
// Parameters:
//   - value: the float32 value to be converted.
//
// Returns:
//   - float32: the converted float32 value.
//   - error: no error is returned.
//
// Example:
//
//	result, err := Float32ToFloat32(1.23)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 1.23
func Float32ToFloat32(value float32) (float32, error) {
	return value, nil
}

// Float64ToFloat32 converts a float64 value to float32.
//
// Float64ToFloat32 converts a float64 value to float32. If the input value
// is outside the range that can be represented by a float32, it returns
// an error.
//
// Parameters:
//   - value: the float64 value to be converted. The range of float64 is
//     approximately ±1.7E308. The function ensures that the input value is
//     within the range of float32, which is approximately -3.4e38 to
//     3.4e38.
//
// Returns:
//   - float32: the converted float32 value. The range of float32 is
//     approximately ±3.4E38.
//   - error: an error if the input value is out of the float32 range.
//
// Example:
//
//	result, err := Float64ToFloat32(123.456)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123.456
func Float64ToFloat32(value float64) (float32, error) {
	if value > math.MaxFloat32 || value < -math.MaxFloat32 {
		return 0, errors.New("value out of range for float32")
	}
	return float32(value), nil
}

// IntToFloat32 converts an int value to float32.
//
// IntToFloat32 converts an int value to float32. If the input value is
// outside the range that can be represented by a float32, it returns an
// error.
//
// Parameters:
//   - value: the int value to be converted. The range of int is
//     approximately ±2.1E9. The function ensures that the input value is
//     within the range of float32, which is approximately -3.4e38 to
//     3.4e38.
//
// Returns:
//   - float32: the converted float32 value. The range of float32 is
//     approximately ±3.4E38.
//   - error: an error if the input value is out of the float32 range.
//
// Example:
//
//	result, err := IntToFloat32(123456789)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123456789
func IntToFloat32(value int) (float32, error) {
	if float64(value) > math.MaxFloat32 || float64(value) < -math.MaxFloat32 {
		return 0, errors.New("value out of range for float32")
	}
	return float32(value), nil
}

// Int8ToFloat32 converts an int8 value to float32.
//
// Int8ToFloat32 converts an int8 value to float32. This function does not
// perform any range checks because the range of int8 is within the range of
// float32.
//
// Parameters:
//   - value: the int8 value to be converted. The range of int8 is
//     -128 to 127. The range of float32 is approximately -3.4e38 to
//     3.4e38.
//
// Returns:
//   - float32: the converted float32 value.
//   - error: no error is returned.
//
// Example:
//
//	result, err := Int8ToFloat32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Int8ToFloat32(value int8) (float32, error) {
	return float32(value), nil
}

// Int16ToFloat32 converts an int16 value to float32.
//
// Int16ToFloat32 converts an int16 value to float32. This function does
// not perform any range checks because the range of int16 is within the range
// of float32.
//
// Parameters:
//   - value: the int16 value to be converted. The range of int16 is
//     -32768 to 32767. The range of float32 is approximately -3.4e38 to
//     3.4e38.
//
// Returns:
//   - float32: the converted float32 value.
//   - error: no error is returned.
//
// Example:
//
//	result, err := Int16ToFloat32(12345)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 12345
func Int16ToFloat32(value int16) (float32, error) {
	return float32(value), nil
}

// Int32ToFloat32 converts an int32 value to float32.
//
// Int32ToFloat32 converts an int32 value to float32. This function does
// not perform any range checks because the range of int32 is within the range
// of float32.
//
// Parameters:
//   - value: the int32 value to be converted. The range of int32 is
//     -2147483648 to 2147483647. The range of float32 is approximately
//     -3.4e38 to 3.4e38.
//
// Returns:
//   - float32: the converted float32 value.
//   - error: no error is returned.
//
// Example:
//
//	result, err := Int32ToFloat32(123456789)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123456789
func Int32ToFloat32(value int32) (float32, error) {
	return float32(value), nil
}

// Int64ToFloat32 converts an int64 value to float32.
//
// Int64ToFloat32 converts an int64 value to float32. If the input value is
// outside the range that can be represented by a float32, it returns an
// error.
//
// Parameters:
//   - value: the int64 value to be converted. The range of int64 is
//     approximately ±9.2E18. The function ensures that the input value is
//     within the range of float32, which is approximately -3.4e38 to
//     3.4e38.
//
// Returns:
//   - float32: the converted float32 value. The range of float32 is
//     approximately ±3.4E38.
//   - error: an error if the input value is out of the float32 range.
//
// Example:
//
//	result, err := Int64ToFloat32(1234567890123456789)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 1234567890123456789
func Int64ToFloat32(value int64) (float32, error) {
	if float64(value) > math.MaxFloat32 || float64(value) < -math.MaxFloat32 {
		return 0, errors.New("value out of range for float32")
	}
	return float32(value), nil
}

// StringToFloat32 converts a string value to float32.
//
// StringToFloat32 converts a string value to float32. If the input value is
// not a valid number, it returns an error.
//
// Parameters:
//   - value: the string value to be converted.
//
// Returns:
//   - float32: the converted float32 value.
//   - error: an error if the input value is not a valid number.
//
// Example:
//
//	result, err := StringToFloat32("123.456")
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123.456
func StringToFloat32(value string) (float32, error) {
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

// UintToFloat32 converts a uint value to float32.
//
// UintToFloat32 converts a uint value to float32. This function does not
// perform any range checks because the range of uint is within the range of
// float32.
//
// Parameters:
//   - value: the uint value to be converted. The range of uint is
//     0 to 4294967295. The range of float32 is approximately -3.4e38 to
//     3.4e38.
//
// Returns:
//   - float32: the converted float32 value.
//   - error: no error is returned.
//
// Example:
//
//	result, err := UintToFloat32(123456789)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123456789
func UintToFloat32(value uint) (float32, error) {
	return float32(value), nil
}

// Uint8ToFloat32 converts a uint8 value to float32.
//
// Uint8ToFloat32 converts a uint8 value to float32. This function does not
// perform any range checks because the range of uint8 is within the range of
// float32.
//
// Parameters:
//   - value: the uint8 value to be converted. The range of uint8 is
//     0 to 255. The range of float32 is approximately -3.4e38 to
//     3.4e38.
//
// Returns:
//   - float32: the converted float32 value.
//   - error: no error is returned.
//
// Example:
//
//	result, err := Uint8ToFloat32(123)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 123
func Uint8ToFloat32(value uint8) (float32, error) {
	return float32(value), nil
}

// Uint16ToFloat32 converts a uint16 value to float32.
//
// Uint16ToFloat32 converts a uint16 value to float32. This function does
// not perform any range checks because the range of uint16 is within the
// range of float32.
//
// Parameters:
//   - value: the uint16 value to be converted. The range of uint16 is
//     0 to 65535. The range of float32 is approximately -3.4e38 to
//     3.4e38.
//
// Returns:
//   - float32: the converted float32 value.
//   - error: no error is returned.
//
// Example:
//
//	result, err := Uint16ToFloat32(12345)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 12345
func Uint16ToFloat32(value uint16) (float32, error) {
	return float32(value), nil
}

// Uint32ToFloat32 converts a uint32 value to float32.
//
// Uint32ToFloat32 converts a uint32 value to float32. If the input value
// is outside the range that can be represented by a float32, it returns an
// error.
//
// Parameters:
//   - value: the uint32 value to be converted. The range of uint32 is
//     0 to 4294967295. The function ensures that the input value is within
//     the range of float32, which is approximately -3.4e38 to 3.4e38.
//
// Returns:
//   - float32: the converted float32 value. The range of float32 is
//     approximately ±3.4E38.
//   - error: an error if the input value exceeds the float32 maximum limit.
//
// Example:
//
//	result, err := Uint32ToFloat32(1234567890)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 1234567890
func Uint32ToFloat32(value uint32) (float32, error) {
	if float64(value) > math.MaxFloat32 {
		return 0, errors.New("value exceeds float32 max limit")
	}
	return float32(value), nil
}

// Uint64ToFloat32 converts a uint64 value to float32.
//
// Uint64ToFloat32 converts a uint64 value to float32. If the input value
// is outside the range that can be represented by a float32, it returns an
// error.
//
// Parameters:
//   - value: the uint64 value to be converted. The range of uint64 is
//     0 to 18446744073709551615. The function ensures that the input value is
//     within the range of float32, which is approximately -3.4e38 to 3.4e38.
//
// Returns:
//   - float32: the converted float32 value. The range of float32 is
//     approximately ±3.4E38.
//   - error: an error if the input value exceeds the float32 maximum limit.
//
// Example:
//
//	result, err := Uint64ToFloat32(1234567890123456789)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 1234567890123456789
func Uint64ToFloat32(value uint64) (float32, error) {
	if float64(value) > math.MaxFloat32 {
		return 0, errors.New("value exceeds float32 max limit")
	}
	return float32(value), nil
}
