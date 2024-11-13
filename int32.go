package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

func TryIntoInt32[T convertable](value T) (int32, error) {
	result, err := toInt32(value)
	return result.(int32), err
}

func toInt32(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToInt32(value.(float64))
	case reflect.Float32:
		return Float32ToInt32(value.(float32))
	case reflect.Int:
		return IntToInt32(value.(int))
	case reflect.Int8:
		return Int8ToInt32(value.(int8))
	case reflect.Int16:
		return Int16ToInt32(value.(int16))
	case reflect.Int32:
		return value.(int32), nil
	case reflect.Int64:
		return Int64ToInt32(value.(int64))
	case reflect.Uint:
		return UintToInt32(value.(uint))
	case reflect.Uint8:
		return Uint8ToInt32(value.(uint8))
	case reflect.Uint16:
		return Uint16ToInt32(value.(uint16))
	case reflect.Uint32:
		return Uint32ToInt32(value.(uint32))
	case reflect.Uint64:
		return Uint64ToInt32(value.(uint64))
	case reflect.String:
		return StringToInt32(value.(string))
	case reflect.Bool:
		return BoolToInt32(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func BoolToInt32(value bool) (int32, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}
func Float32ToInt32(value float32) (int32, error) {
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, errors.New("value out of range for int32")
	}
	return int32(value), nil
}

func Float64ToInt32(value float64) (int32, error) {
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, errors.New("value out of range for int32")
	}
	return int32(value), nil
}

func IntToInt32(value int) (int32, error) {
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, errors.New("value out of range for int32")
	}
	return int32(value), nil
}

func Int8ToInt32(value int8) (int32, error) {
	return int32(value), nil
}

func Int16ToInt32(value int16) (int32, error) {
	return int32(value), nil
}

func Int32ToInt32(value int32) (int32, error) {
	return value, nil
}

func Int64ToInt32(value int64) (int32, error) {
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, errors.New("value out of range for int32")
	}
	return int32(value), nil
}

func StringToInt32(value string) (int32, error) {
	i, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

func UintToInt32(value uint) (int32, error) {
	if value > math.MaxInt32 {
		return 0, errors.New("value exceeds int32 max limit")
	}
	return int32(value), nil
}

func Uint8ToInt32(value uint8) (int32, error) {
	return int32(value), nil
}
func Uint16ToInt32(value uint16) (int32, error) {
	return int32(value), nil
}
func Uint32ToInt32(value uint32) (int32, error) {
	if value > math.MaxInt32 {
		return 0, errors.New("value exceeds int32 max limit")
	}
	return int32(value), nil
}
func Uint64ToInt32(value uint64) (int32, error) {
	if value > math.MaxInt32 {
		return 0, errors.New("value exceeds int32 max limit")
	}
	return int32(value), nil
}
