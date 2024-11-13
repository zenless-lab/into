package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

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

func BoolToInt64(value bool) (int64, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func Float32ToInt64(value float32) (int64, error) {
	if value > math.MaxInt64 || value < math.MinInt64 {
		return 0, errors.New("value out of range for int64")
	}
	return int64(value), nil
}

func Float64ToInt64(value float64) (int64, error) {
	if value > math.MaxInt64 || value < math.MinInt64 {
		return 0, errors.New("value out of range for int64")
	}
	return int64(value), nil
}

func IntToInt64(value int) (int64, error) {
	return int64(value), nil
}

func Int8ToInt64(value int8) (int64, error) {
	return int64(value), nil
}

func Int16ToInt64(value int16) (int64, error) {
	return int64(value), nil
}

func Int32ToInt64(value int32) (int64, error) {
	return int64(value), nil
}

func Int64ToInt64(value int64) (int64, error) {
	return value, nil
}

func StringToInt64(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

func UintToInt64(value uint) (int64, error) {
	if value > math.MaxInt64 {
		return 0, errors.New("value exceeds int64 max limit")
	}
	return int64(value), nil
}
func Uint8ToInt64(value uint8) (int64, error) {
	return int64(value), nil
}

func Uint16ToInt64(value uint16) (int64, error) {
	return int64(value), nil
}
func Uint32ToInt64(value uint32) (int64, error) {
	return int64(value), nil
}

func Uint64ToInt64(value uint64) (int64, error) {
	if value > math.MaxInt64 {
		return 0, errors.New("value exceeds int64 max limit")
	}
	return int64(value), nil
}
