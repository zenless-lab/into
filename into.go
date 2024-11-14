package into

import (
	"errors"
	"reflect"
)

// Into converts a value of type U to a value of type T.
//
// Into converts a value of type U to a value of type T. If the conversion fails, it panics.
//
// Parameters:
//   - value: the value to be converted. It must be a convertable type.
//
// Returns:
//   - T: the converted value of type T.
//
// Example:
//
//	var a float64 = 123.456
//	b := Into[int](a)
//	fmt.Println(b) // Output: 123
func Into[T convertable, U convertable](value U) T {
	result, err := TryInto[T, U](value)
	if err != nil {
		panic(err)
	}
	return result
}

// TryInto attempts to convert a value of type U to a value of type T.
//
// TryInto attempts to convert a value of type U to a value of type T. If the conversion fails, it returns an error.
//
// Parameters:
//   - value: the value to be converted. It must be a convertable type.
//
// Returns:
//   - T: the converted value of type T.
//   - error: an error if the conversion fails.
//
// Example:
//
//	var a float64 = 123.456
//	b, err := TryInto[int](a)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(b) // Output: 123
func TryInto[T convertable, U convertable](value U) (result T, err error) {
	resultType := reflect.TypeOf(result)

	var r any
	switch resultType.Kind() {
	case reflect.Float64:
		r, err = toFloat64(value)
	case reflect.Float32:
		r, err = toFloat32(value)
	case reflect.Int:
		r, err = toInt(value)
	case reflect.Int8:
		r, err = toInt8(value)
	case reflect.Int16:
		r, err = toInt16(value)
	case reflect.Int32:
		r, err = toInt32(value)
	case reflect.Int64:
		r, err = toInt64(value)
	case reflect.Uint:
		r, err = toUint(value)
	case reflect.Uint8:
		r, err = toUint8(value)
	case reflect.Uint16:
		r, err = toUint16(value)
	case reflect.Uint32:
		r, err = toUint32(value)
	case reflect.Uint64:
		r, err = toUint64(value)
	case reflect.String:
		r, err = toString(value)
	case reflect.Bool:
		r, err = toBool(value)
	default:
		err = errors.New("unsupported type")
		return
	}

	result = r.(T)
	return
}
