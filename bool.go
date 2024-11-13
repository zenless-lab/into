package into

import (
	"errors"
	"reflect"
	"strconv"
)

// TryIntoBool converts a value of type T into a boolean value.
//
// It accepts any value that implements the convertable interface which allows conversion to bool.
// The generic type T must satisfy the convertable interface constraint.
//
// Parameters:
//   - value: The value to convert to bool. Must be of type T that implements convertable.
//
// Returns:
//   - bool: The converted boolean value
//   - error: An error if the conversion fails, nil if successful
//
// Example:
//
//	val := "true"
//	b, err := TryIntoBool(val)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// b == true
//
//	val2 := 1
//	b2, err := TryIntoBool(val2)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// b2 == true
func TryIntoBool[T convertable](value T) (bool, error) {
	result, err := toBool(value)
	return result.(bool), err
}

// toBool converts a value of any type to a boolean value. It is used internally by TryIntoBool.
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

// BoolToBool takes a boolean value and returns it directly with no error.
// It serves as a simple pass-through function for boolean values.
//
// Parameters:
//   - value: A boolean input value to be passed through
//
// Returns:
//   - bool: The same boolean value that was passed in
//   - error: Always nil since this operation cannot fail
//
// Example:
//
//	result, err := BoolToBool(true)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Result: %v\n", result) // Output: Result: true
func BoolToBool(value bool) (bool, error) {
	return value, nil
}

// Float32ToBool converts a float32 value to a boolean value.
// It returns true if the value is not equal to 0, false otherwise.
//
// Parameters:
//   - value: float32 number to be converted
//
// Returns:
//   - bool: true if value is not 0, false if value is 0
//   - error: always returns nil (reserved for future validation)
//
// Example:
//
//	result, err := Float32ToBool(1.0)  // returns true, nil
//	result, err := Float32ToBool(0.0)  // returns false, nil
//	result, err := Float32ToBool(-1.0) // returns true, nil
func Float32ToBool(value float32) (bool, error) {
	return value != 0, nil
}

// Float64ToBool converts a float64 value to a boolean value.
// It returns true if the value is not equal to 0, false otherwise.
//
// Parameters:
//   - value: float64 input value to be converted
//
// Returns:
//   - bool: true if value != 0, false if value == 0
//   - error: currently always returns nil, reserved for future validation
//
// Example:
//
//	result, err := Float64ToBool(1.0)    // returns true, nil
//	result, err := Float64ToBool(0.0)    // returns false, nil
//	result, err := Float64ToBool(-1.5)   // returns true, nil
func Float64ToBool(value float64) (bool, error) {
	return value != 0, nil
}

// IntToBool converts an integer value to a boolean.
// It returns true if the value is non-zero, false if the value is zero.
//
// Parameters:
//   - value: The integer value to convert
//
// Returns:
//   - bool: The boolean result (true if non-zero, false if zero)
//   - error: Always returns nil for this function
//
// Example:
//
//	b, _ := IntToBool(1)  // returns true
//	b, _ := IntToBool(0)  // returns false
//	b, _ := IntToBool(-1) // returns true
func IntToBool(value int) (bool, error) {
	return value != 0, nil
}

// Int8ToBool converts an int8 value to a bool.
//
// Parameters:
//   - value: The int8 value to convert (0 or non-0)
//
// Returns:
//   - bool: true if value is non-zero, false if value is zero
//   - error: always nil for this function
//
// Example:
//
//	b, err := Int8ToBool(1)  // returns true, nil
//	b, err := Int8ToBool(0)  // returns false, nil
//	b, err := Int8ToBool(-1) // returns true, nil
func Int8ToBool(value int8) (bool, error) {
	return value != 0, nil
}

// Int16ToBool converts an int16 value to a boolean.
// It returns true if value is not 0, false otherwise.
//
// Parameters:
//   - value: int16 input to convert to boolean
//
// Returns:
//   - bool: true if value != 0, false if value == 0
//   - error: always nil for this function 
//
// Example:
//
//	b, err := Int16ToBool(1)  // returns true, nil
//	b, err := Int16ToBool(0)  // returns false, nil
//	b, err := Int16ToBool(-1) // returns true, nil
func Int16ToBool(value int16) (bool, error) {
	return value != 0, nil
}

// Int32ToBool converts an int32 value to a boolean.
// It returns true if value is non-zero, false if value is zero.
//
// Parameters:
//   - value: The int32 value to convert
//
// Returns:
//   - bool: The boolean result (true if value != 0, false if value == 0)
//   - error: Always returns nil as this conversion cannot fail
//
// Example:
//
//	result, err := Int32ToBool(1) // returns true, nil
//	result, err := Int32ToBool(0) // returns false, nil
//	result, err := Int32ToBool(-1) // returns true, nil
func Int32ToBool(value int32) (bool, error) {
	return value != 0, nil
}

// Int64ToBool converts an int64 value to a boolean value.
// It returns true if the value is non-zero, false if the value is zero.
//
// Parameters:
//   - value: The int64 value to convert
//
// Returns:
//   - bool: true if value is non-zero, false if value is zero
//   - error: always returns nil (for interface consistency)
//
// Example:
//
//	b, err := Int64ToBool(1)  // returns true, nil
//	b, err := Int64ToBool(0)  // returns false, nil
//	b, err := Int64ToBool(-1) // returns true, nil
func Int64ToBool(value int64) (bool, error) {
	return value != 0, nil
}

// StringToBool converts a string to a boolean value.
//
// The function uses strconv.ParseBool internally and accepts "1", "t", "T", "true", "TRUE", "True"
// for true and "0", "f", "F", "false", "FALSE", "False" for false.
//
// Parameters:
//   - value: string to be converted to boolean
//
// Returns:
//   - bool: the parsed boolean value
//   - error: error if the string cannot be converted to a valid boolean
//
// Example:
//
//	b, err := StringToBool("true")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Value: %v\n", b) // Output: Value: true
//
//	b, err = StringToBool("invalid")
//	if err != nil {
//	    fmt.Printf("Error: %v\n", err) // Output: Error: invalid syntax
//	}
func StringToBool(value string) (bool, error) {
	return strconv.ParseBool(value)
}

// UintToBool converts a uint value to a boolean.
// It returns true if value is non-zero, false if value is zero.
//
// Parameters:
//   - value: uint number to convert
//
// Returns:
//   - bool: true if value is non-zero, false if value is zero
//   - error: always returns nil as this conversion cannot fail
//
// Example:
//
//	b, err := UintToBool(1) // returns true, nil
//	b, err := UintToBool(0) // returns false, nil
func UintToBool(value uint) (bool, error) {
	return value != 0, nil
}

// Uint8ToBool converts an uint8 value to a boolean.
// It returns true if the input value is non-zero, false otherwise.
//
// Parameters:
//   - value: uint8 value to convert
//
// Returns:
//   - bool: true if value is non-zero, false if value is 0
//   - error: always returns nil as this conversion cannot fail
//
// Example:
//
//	b, err := Uint8ToBool(1)  // returns true, nil
//	b, err := Uint8ToBool(0)  // returns false, nil
//	b, err := Uint8ToBool(42) // returns true, nil
func Uint8ToBool(value uint8) (bool, error) {
	return value != 0, nil
}

// Uint16ToBool converts a uint16 value to a boolean.
//
// Parameters:
//   - value: The uint16 value to convert to boolean
//
// Returns:
//   - bool: true if value is not 0, false if value is 0
//   - error: always returns nil for this function
//
// Example:
//
//	b, err := Uint16ToBool(1)    // returns true, nil
//	b, err := Uint16ToBool(0)    // returns false, nil
//	b, err := Uint16ToBool(1000) // returns true, nil
func Uint16ToBool(value uint16) (bool, error) {
	return value != 0, nil
}

// Uint32ToBool converts a uint32 value to a boolean.
// It returns true if the input value is non-zero, false if the value is zero.
//
// Parameters:
//   - value: The uint32 value to convert
//
// Returns:
//   - bool: The boolean result (true if value != 0, false if value == 0)
//   - error: Always returns nil as this conversion cannot fail
//
// Example:
//
//	b, err := Uint32ToBool(1) // returns true, nil
//	b, err := Uint32ToBool(0) // returns false, nil
func Uint32ToBool(value uint32) (bool, error) {
	return value != 0, nil
}

// Uint64ToBool converts an uint64 value to a boolean.
// The function returns true if value is not zero, false otherwise.
//
// Parameters:
//   - value: uint64 number to convert
//
// Returns:
//   - bool: true if value is not 0, false if value is 0
//   - error: always returns nil for this function
//
// Example:
//
//	b, err := Uint64ToBool(1) // returns true, nil
//	b, err := Uint64ToBool(0) // returns false, nil
func Uint64ToBool(value uint64) (bool, error) {
	return value != 0, nil
}
