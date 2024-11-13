package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

func TryIntoInt16[T convertable](value T) (int16, error) {
	result, err := toInt16(value)
	return result.(int16), err
}

func toInt16(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToInt16(value.(float64))
	case reflect.Float32:
		return Float32ToInt16(value.(float32))
	case reflect.Int:
		return IntToInt16(value.(int))
	case reflect.Int8:
		return Int8ToInt16(value.(int8))
	case reflect.Int16:
		return value.(int16), nil
	case reflect.Int32:
		return Int32ToInt16(value.(int32))
	case reflect.Int64:
		return Int64ToInt16(value.(int64))
	case reflect.Uint:
		return UintToInt16(value.(uint))
	case reflect.Uint8:
		return Uint8ToInt16(value.(uint8))
	case reflect.Uint16:
		return Uint16ToInt16(value.(uint16))
	case reflect.Uint32:
		return Uint32ToInt16(value.(uint32))
	case reflect.Uint64:
		return Uint64ToInt16(value.(uint64))
	case reflect.String:
		return StringToInt16(value.(string))
	case reflect.Bool:
		return BoolToInt16(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}
func BoolToInt16(value bool) (int16, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func Float32ToInt16(value float32) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}

func Float64ToInt16(value float64) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}

func IntToInt16(value int) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}

func Int8ToInt16(value int8) (int16, error) {
	return int16(value), nil
}
func Int16ToInt16(value int16) (int16, error) {
	return value, nil
}

func Int32ToInt16(value int32) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}
func Int64ToInt16(value int64) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}

func StringToInt16(value string) (int16, error) {
	i, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(i), nil
}

func UintToInt16(value uint) (int16, error) {
	if value > math.MaxInt16 {
		return 0, errors.New("value exceeds int16 max limit")
	}
	return int16(value), nil
}

func Uint8ToInt16(value uint8) (int16, error) {
	return int16(value), nil
}
func Uint16ToInt16(value uint16) (int16, error) {
	if value > math.MaxInt16 {
		return 0, errors.New("value exceeds int16 max limit")
	}
	return int16(value), nil
}

func Uint32ToInt16(value uint32) (int16, error) {
	if value > math.MaxInt16 {
		return 0, errors.New("value exceeds int16 max limit")
	}
	return int16(value), nil
}

func Uint64ToInt16(value uint64) (int16, error) {
	if value > math.MaxInt16 {
		return 0, errors.New("value exceeds int16 max limit")
	}
	return int16(value), nil
}
