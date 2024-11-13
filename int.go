package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

func TryIntoInt[T convertable](value T) (int, error) {
	result, err := toInt(value)
	return result.(int), err
}

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

func BoolToInt(value bool) (int, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func Float32ToInt(value float32) (int, error) {
	if value > math.MaxInt || value < math.MinInt {
		return 0, errors.New("value out of range for int")
	}
	return int(value), nil
}

func Float64ToInt(value float64) (int, error) {
	if value > math.MaxInt || value < math.MinInt {
		return 0, errors.New("value out of range for int")
	}
	return int(value), nil
}

func IntToInt(value int) (int, error) {
	return value, nil
}

func Int8ToInt(value int8) (int, error) {
	return int(value), nil
}

func Int16ToInt(value int16) (int, error) {
	return int(value), nil
}

func Int32ToInt(value int32) (int, error) {
	return int(value), nil
}

func Int64ToInt(value int64) (int, error) {
	if value > math.MaxInt || value < math.MinInt {
		return 0, errors.New("value out of range for int")
	}
	return int(value), nil
}

func StringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

func UintToInt(value uint) (int, error) {
	if value > math.MaxInt {
		return 0, errors.New("value exceeds int max limit")
	}
	return int(value), nil
}

func Uint8ToInt(value uint8) (int, error) {
	return int(value), nil
}

func Uint16ToInt(value uint16) (int, error) {
	return int(value), nil
}

func Uint32ToInt(value uint32) (int, error) {
	if uint64(value) > math.MaxInt {
		return 0, errors.New("value exceeds int max limit")
	}
	return int(value), nil
}

func Uint64ToInt(value uint64) (int, error) {
	if value > math.MaxInt {
		return 0, errors.New("value exceeds int max limit")
	}
	return int(value), nil
}
