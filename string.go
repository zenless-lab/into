package into

import (
	"errors"
	"reflect"
	"strconv"
)

func TryIntoString[T convertable | ~[]byte | ~[]rune](value T) (string, error) {
	result, err := toString(value)
	return result.(string), err
}

func toString(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToString(value.(float64))
	case reflect.Float32:
		return Float32ToString(value.(float32))
	case reflect.Int:
		return IntToString(value.(int))
	case reflect.Int8:
		return Int8ToString(value.(int8))
	case reflect.Int16:
		return Int16ToString(value.(int16))
	case reflect.Int32:
		return Int32ToString(value.(int32))
	case reflect.Int64:
		return Int64ToString(value.(int64))
	case reflect.Uint:
		return UintToString(value.(uint))
	case reflect.Uint8:
		return Uint8ToString(value.(uint8))
	case reflect.Uint16:
		return Uint16ToString(value.(uint16))
	case reflect.Uint32:
		return Uint32ToString(value.(uint32))
	case reflect.Uint64:
		return Uint64ToString(value.(uint64))
	case reflect.String:
		return value.(string), nil
	case reflect.Bool:
		return BoolToString(value.(bool))
	default:
		return "", errors.New("unsupported type")
	}
}

func BoolToString(value bool) (string, error) {
	return strconv.FormatBool(value), nil
}

func BytesToString(value []byte) (string, error) {
	return string(value), nil
}

func ErrorToString(err error) (string, error) {
	return err.Error(), nil
}

func Float32ToString(value float32) (string, error) {
	return strconv.FormatFloat(float64(value), 'f', -1, 32), nil
}

func Float64ToString(value float64) (string, error) {
	return strconv.FormatFloat(value, 'f', -1, 64), nil
}
func IntToString(value int) (string, error) {
	return strconv.Itoa(value), nil
}
func Int8ToString(value int8) (string, error) {
	return strconv.Itoa(int(value)), nil
}
func Int16ToString(value int16) (string, error) {
	return strconv.Itoa(int(value)), nil
}

func Int32ToString(value int32) (string, error) {
	return strconv.Itoa(int(value)), nil
}

func Int64ToString(value int64) (string, error) {
	return strconv.FormatInt(value, 10), nil
}

func RuneToString(value rune) (string, error) {
	return string(value), nil
}
func StringToString(value string) (string, error) {
	return value, nil
}
func UintToString(value uint) (string, error) {
	return strconv.FormatUint(uint64(value), 10), nil
}

func Uint8ToString(value uint8) (string, error) {
	return strconv.FormatUint(uint64(value), 10), nil
}
func Uint16ToString(value uint16) (string, error) {
	return strconv.FormatUint(uint64(value), 10), nil
}
func Uint32ToString(value uint32) (string, error) {
	return strconv.FormatUint(uint64(value), 10), nil
}
func Uint64ToString(value uint64) (string, error) {
	return strconv.FormatUint(value, 10), nil
}
