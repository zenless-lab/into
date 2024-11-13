package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

// TryIntoInt attempts to convert a value of type T into an integer.
//
// The generic type T must satisfy the convertable interface constraint.
// This function can convert values from:
//   - string representations of integers
//   - floating point numbers (truncates decimal portion)
//   - boolean values (true=1, false=0)
//   - integer types (int8, int16, int32, int64, uint8, etc.)
//
// Parameters:
//   - value: The input value to be converted to int
//
// Returns:
//   - int: The converted integer value
//   - error: An error if conversion fails or value is out of range for int type
//
// Example:
//   i, err := TryIntoInt(123.45)    // returns 123, nil
//   i, err := TryIntoInt("42")      // returns 42, nil
//   i, err := TryIntoInt(true)      // returns 1, nil
//   i, err := TryIntoInt("invalid") // returns 0, error
func TryIntoInt[T convertable](value T) (int, error) {
	result, err := toInt(value)
	return result.(int), err
}

// toInt converts a value of any type to an integer. It is used internally by TryIntoInt.
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

// BoolToInt converts a boolean value to its integer representation.
//
// Parameters:
//   - value: A boolean value to be converted
//
// Returns:
//   - int: Returns 1 if true, 0 if false
//   - error: Always returns nil for this simple conversion
//
// Example:
//
//	i, err := BoolToInt(true)  // returns 1, nil
//	i, err := BoolToInt(false) // returns 0, nil
func BoolToInt(value bool) (int, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

// Float32ToInt converts a float32 value to int.
//
// Parameters:
//   - value: The float32 number to be converted. Must be within the range of [math.MinInt, math.MaxInt].
//
// Returns:
//   - int: The converted integer value
//   - error: Returns nil if conversion is successful, or an error if value is out of range
//
// Range:
// 	Source Range: float32 (-3.4028235e+38 to 3.4028235e+38)
// 	Target Range: int (architecture dependent, typically -9223372036854775808 to 9223372036854775807 on 64-bit systems)
//
// Example:
//
//	i, err := Float32ToInt(123.45)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Converted value: %d\n", i) // Output: 123
func Float32ToInt(value float32) (int, error) {
	if value > math.MaxInt || value < math.MinInt {
		return 0, errors.New("value out of range for int")
	}
	return int(value), nil
}

// Float64ToInt converts a float64 value to an int.
//
// Parameters:
//   - value: The float64 value to be converted. Must be within the range [math.MinInt, math.MaxInt].
//
// Returns:
//   - int: The converted integer value
//   - error: Returns nil if conversion is successful, or an error if value is out of valid int range
//
// Range:
//   Source (float64): -1.8e+308 to +1.8e+308
//   Target (int): -9223372036854775808 to 9223372036854775807 (on 64-bit systems)
//
// Example:
//   i, err := Float64ToInt(123.45)
//   if err != nil {
//       log.Fatal(err)
//   }
//   fmt.Printf("Converted value: %d\n", i) // Output: Converted value: 123
func Float64ToInt(value float64) (int, error) {
	if value > math.MaxInt || value < math.MinInt {
		return 0, errors.New("value out of range for int")
	}
	return int(value), nil
}

// IntToInt converts an integer value to another integer value.
//
// Parameters:
//   - value: input integer value (range: platform-dependent, typically -2^31 to 2^31-1 on 32-bit systems,
//     -2^63 to 2^63-1 on 64-bit systems)
//
// Returns:
//   - int: same value as input (range same as input)
//   - error: always nil for this simple conversion
//
// Example:
//
//	input := 42
//	result, err := IntToInt(input)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Converted value: %d\n", result) // Output: Converted value: 42
func IntToInt(value int) (int, error) {
	return value, nil
}

// Int8ToInt converts an int8 value to int.
//
// Parameters:
//   - value: The int8 value to convert (-128 to 127)
//
// Returns:
//   - int: The converted int value (platform dependent size, minimum -2^31 to 2^31-1)
//   - error: Always returns nil for this function as conversion is safe
//
// Example:
//
//	n, err := Int8ToInt(100)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Converted value: %d\n", n)
//
// Note that int8 range (-128 to 127) is always safely convertible to int
// as int is guaranteed to be at least 32 bits on all platforms.
func Int8ToInt(value int8) (int, error) {
	return int(value), nil
}

// Int16ToInt converts int16 to int.
//
// Parameters:
//   - value: int16 value to be converted (-32768 to 32767)
//
// Returns:
//   - int: converted integer value (platform dependent size, minimum -2^31 to 2^31-1 on 32-bit systems)
//   - error: nil as this conversion cannot fail due to int having larger range than int16
//
// Example:
//   i16 := int16(12345)
//   i, err := Int16ToInt(i16)
//   if err != nil {
//       // handle error
//   }
//   fmt.Printf("Converted %d to %d\n", i16, i) // Output: Converted 12345 to 12345
func Int16ToInt(value int16) (int, error) {
	return int(value), nil
}

// Int32ToInt converts an int32 value to int.
//
// Parameters:
//   - value: An int32 number to be converted, range from -2147483648 to 2147483647
//
// Returns:
//   - int: The converted int value. On 32-bit systems, range is same as int32 (-2147483648 to 2147483647).
//     On 64-bit systems, range is -9223372036854775808 to 9223372036854775807
//   - error: Always returns nil as this conversion cannot fail
//
// Example:
//
//	val := int32(123)
//	result, err := Int32ToInt(val)
//	if err != nil {
//	    // Handle error
//	}
//	fmt.Printf("Converted value: %d\n", result) // Output: Converted value: 123
func Int32ToInt(value int32) (int, error) {
	return int(value), nil
}

// Int64ToInt safely converts an int64 to int with range validation.
//
// This function checks if the input int64 value can be safely converted to int
// without losing data by verifying it falls within int's valid range
// (math.MinInt to math.MaxInt).
//
// Parameters:
//   - value: The int64 value to be converted
//
// Returns:
//   - int: The converted integer value if conversion is successful
//   - error: An error if value exceeds int's range, nil otherwise
//
// Example:
//   val, err := Int64ToInt(123)
//   if err != nil {
//       log.Fatal(err)
//   }
//   fmt.Printf("Converted value: %d\n", val)
//
// Ranges:
//   - Input (int64): -9223372036854775808 to 9223372036854775807
//   - Output (int): -9223372036854775808 to 9223372036854775807 (64-bit)
//                   -2147483648 to 2147483647 (32-bit)
func Int64ToInt(value int64) (int, error) {
	if value > math.MaxInt || value < math.MinInt {
		return 0, errors.New("value out of range for int")
	}
	return int(value), nil
}

// StringToInt converts a string representation of an integer to an int value.
//
// Parameters:
//   - value: A string containing a base-10 integer number to be converted.
//     Valid input range: "-9223372036854775808" to "9223372036854775807" for 64-bit systems,
//     or "-2147483648" to "2147483647" for 32-bit systems.
//
// Returns:
//   - int: The converted integer value
//   - error: Returns nil if conversion is successful, or strconv.ErrRange if value overflows int,
//     or strconv.ErrSyntax if value is not a valid integer format
//
// Example:
//   i, err := StringToInt("123")    // Returns: 123, nil
//   i, err := StringToInt("-456")   // Returns: -456, nil
//   i, err := StringToInt("abc")    // Returns: 0, error
//   i, err := StringToInt("")       // Returns: 0, error
func StringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

// UintToInt converts an unsigned integer (uint) to a signed integer (int).
//
// The function checks if the input value exceeds the maximum value of int
// (math.MaxInt) to prevent overflow. On 32-bit systems, math.MaxInt is 2^31-1,
// and on 64-bit systems, it's 2^63-1.
//
// Parameters:
//   - value: An unsigned integer to be converted
//
// Returns:
//   - int: The converted signed integer value
//   - error: An error if the input exceeds math.MaxInt, nil otherwise
//
// Example:
//   val, err := UintToInt(42)
//   if err != nil {
//       log.Fatal(err)
//   }
//   fmt.Println(val) // Output: 42
//
// Ranges:
// 	Source range: 0 to uint max (2^32-1 on 32-bit systems, 2^64-1 on 64-bit systems)
// 	Target range: 0 to int max (2^31-1 on 32-bit systems, 2^63-1 on 64-bit systems)
func UintToInt(value uint) (int, error) {
	if value > math.MaxInt {
		return 0, errors.New("value exceeds int max limit")
	}
	return int(value), nil
}

// Uint8ToInt converts an uint8 (0 to 255) to int.
//
// Parameters:
//   - value: uint8 input value in range [0, 255]
//
// Returns:
//   - int: converted integer value
//   - error: currently always returns nil since uint8 range fits into int
//
// Example:
//
//	val, err := Uint8ToInt(255)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Converted value: %d\n", val)
//
// Note: As uint8 has range of [0, 255] and int has minimum range of [-2^15, 2^15-1],
// this conversion is always safe and will not overflow on any platform.
func Uint8ToInt(value uint8) (int, error) {
	return int(value), nil
}

// Uint16ToInt converts a uint16 value to int.
//
// Parameters:
//   - value: uint16 number to be converted (range: 0 to 65535)
//
// Returns:
//   - int: converted integer value
//   - error: nil if successful (no error scenario as uint16 range fits in int)
//
// Range:
//   - Source (uint16): 0 to 65535
//   - Target (int): minimum value of int (platform dependent, at least -2^31) to 65535
//
// Example:
//   val := uint16(12345)
//   result, err := Uint16ToInt(val)
//   if err != nil {
//       log.Fatal(err)
//   }
//   fmt.Printf("Converted value: %d\n", result) // Output: Converted value: 12345
func Uint16ToInt(value uint16) (int, error) {
	return int(value), nil
}

// Uint32ToInt converts a uint32 value to int type with range checking.
//
// Parameters:
//   - value: uint32 number to be converted (range: 0 to 4294967295)
//
// Returns:
//   - int: converted integer value
//   - error: returns error if value exceeds int max limit
//     (int max = 2147483647 on 32-bit systems, 9223372036854775807 on 64-bit systems)
//
// Example:
//
//	val := uint32(42)
//	result, err := Uint32ToInt(val)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Converted value: %d\n", result)
func Uint32ToInt(value uint32) (int, error) {
	if uint64(value) > math.MaxInt {
		return 0, errors.New("value exceeds int max limit")
	}
	return int(value), nil
}

// Uint64ToInt converts an unsigned 64-bit integer to a signed integer.
//
// The function checks if the input value exceeds the maximum value of int
// (math.MaxInt, which is 9223372036854775807 on 64-bit systems and 
// 2147483647 on 32-bit systems).
//
// Parameters:
//   - value: uint64 number to be converted (range: 0 to 18446744073709551615)
//
// Returns:
//   - int: converted value if within valid range
//   - error: returns error if value exceeds math.MaxInt
//
// Example:
//
//	num, err := Uint64ToInt(123)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(num) // Output: 123
//
//	// Error case
//	num, err = Uint64ToInt(math.MaxUint64)
//	if err != nil {
//	    fmt.Println(err) // Output: value exceeds int max limit
//	}
func Uint64ToInt(value uint64) (int, error) {
	if value > math.MaxInt {
		return 0, errors.New("value exceeds int max limit")
	}
	return int(value), nil
}
