package into

import (
	"errors"
	"reflect"
	"strconv"
)

func TryIntoFloat64[T convertable](value T) (float64, error) {
	result, err := toFloat64(value)
	return result.(float64), err
}

func toFloat64(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return value.(float64), nil
	case reflect.Float32:
		return Float32ToFloat64(value.(float32))
	case reflect.Int:
		return IntToFloat64(value.(int))
	case reflect.Int8:
		return Int8ToFloat64(value.(int8))
	case reflect.Int16:
		return Int16ToFloat64(value.(int16))
	case reflect.Int32:
		return Int32ToFloat64(value.(int32))
	case reflect.Int64:
		return Int64ToFloat64(value.(int64))
	case reflect.Uint:
		return UintToFloat64(value.(uint))
	case reflect.Uint8:
		return Uint8ToFloat64(value.(uint8))
	case reflect.Uint16:
		return Uint16ToFloat64(value.(uint16))
	case reflect.Uint32:
		return Uint32ToFloat64(value.(uint32))
	case reflect.Uint64:
		return Uint64ToFloat64(value.(uint64))
	case reflect.String:
		return StringToFloat64(value.(string))
	case reflect.Bool:
		return BoolToFloat64(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func BoolToFloat64(value bool) (float64, error) {
	if value {
		return 1.0, nil
	}
	return 0.0, nil
}
func Float32ToFloat64(value float32) (float64, error) {
	return float64(value), nil
}
func Float64ToFloat64(value float64) (float64, error) {
	return value, nil
}

func IntToFloat64(value int) (float64, error) {
	return float64(value), nil
}

func Int8ToFloat64(value int8) (float64, error) {
	return float64(value), nil
}
func Int16ToFloat64(value int16) (float64, error) {
	return float64(value), nil
}
func Int32ToFloat64(value int32) (float64, error) {
	return float64(value), nil
}
func Int64ToFloat64(value int64) (float64, error) {
	return float64(value), nil
}
func StringToFloat64(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

func UintToFloat64(value uint) (float64, error) {
	return float64(value), nil
}
func Uint8ToFloat64(value uint8) (float64, error) {
	return float64(value), nil
}

func Uint16ToFloat64(value uint16) (float64, error) {
	return float64(value), nil
}

func Uint32ToFloat64(value uint32) (float64, error) {
	return float64(value), nil
}

func Uint64ToFloat64(value uint64) (float64, error) {
	return float64(value), nil
}
