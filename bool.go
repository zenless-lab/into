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

func BoolToBool(value bool) (bool, error) {
	return value, nil
}

func Float32ToBool(value float32) (bool, error) {
	return value != 0, nil
}
func Float64ToBool(value float64) (bool, error) {
	return value != 0, nil
}
func IntToBool(value int) (bool, error) {
	return value != 0, nil
}
func Int8ToBool(value int8) (bool, error) {
	return value != 0, nil
}

func Int16ToBool(value int16) (bool, error) {
	return value != 0, nil
}
func Int32ToBool(value int32) (bool, error) {
	return value != 0, nil
}

func Int64ToBool(value int64) (bool, error) {
	return value != 0, nil
}
func StringToBool(value string) (bool, error) {
	return strconv.ParseBool(value)
}
func UintToBool(value uint) (bool, error) {
	return value != 0, nil
}
func Uint8ToBool(value uint8) (bool, error) {
	return value != 0, nil
}

func Uint16ToBool(value uint16) (bool, error) {
	return value != 0, nil
}
func Uint32ToBool(value uint32) (bool, error) {
	return value != 0, nil
}

func Uint64ToBool(value uint64) (bool, error) {
	return value != 0, nil
}
