package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

func TryIntoInt8[T convertable](value T) (int8, error) {
	result, err := toInt8(value)
	return result.(int8), err
}

func toInt8(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToInt8(value.(float64))
	case reflect.Float32:
		return Float32ToInt8(value.(float32))
	case reflect.Int:
		return IntToInt8(value.(int))
	case reflect.Int8:
		return value.(int8), nil
	case reflect.Int16:
		return Int16ToInt8(value.(int16))
	case reflect.Int32:
		return Int32ToInt8(value.(int32))
	case reflect.Int64:
		return Int64ToInt8(value.(int64))
	case reflect.Uint:
		return UintToInt8(value.(uint))
	case reflect.Uint8:
		return Uint8ToInt8(value.(uint8))
	case reflect.Uint16:
		return Uint16ToInt8(value.(uint16))
	case reflect.Uint32:
		return Uint32ToInt8(value.(uint32))
	case reflect.Uint64:
		return Uint64ToInt8(value.(uint64))
	case reflect.String:
		return StringToInt8(value.(string))
	case reflect.Bool:
		return BoolToInt8(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func BoolToInt8(value bool) (int8, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}
func Float32ToInt8(value float32) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}
func Float64ToInt8(value float64) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}
func IntToInt8(value int) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

func Int8ToInt8(value int8) (int8, error) {
	return value, nil
}

func Int16ToInt8(value int16) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}
func Int32ToInt8(value int32) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

func Int64ToInt8(value int64) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

func StringToInt8(value string) (int8, error) {
	i, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(i), nil
}
func UintToInt8(value uint) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}

func Uint8ToInt8(value uint8) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}
func Uint16ToInt8(value uint16) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}

func Uint32ToInt8(value uint32) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}

func Uint64ToInt8(value uint64) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}
