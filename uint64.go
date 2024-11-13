package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

func TryIntoUint64[T convertable](value T) (uint64, error) {
	result, err := toUint64(value)
	return result.(uint64), err
}

func toUint64(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint64(value.(float64))
	case reflect.Float32:
		return Float32ToUint64(value.(float32))
	case reflect.Int:
		return IntToUint64(value.(int))
	case reflect.Int8:
		return Int8ToUint64(value.(int8))
	case reflect.Int16:
		return Int16ToUint64(value.(int16))
	case reflect.Int32:
		return Int32ToUint64(value.(int32))
	case reflect.Int64:
		return Int64ToUint64(value.(int64))
	case reflect.Uint:
		return UintToUint64(value.(uint))
	case reflect.Uint8:
		return Uint8ToUint64(value.(uint8))
	case reflect.Uint16:
		return Uint16ToUint64(value.(uint16))
	case reflect.Uint32:
		return Uint32ToUint64(value.(uint32))
	case reflect.Uint64:
		return value.(uint64), nil
	case reflect.String:
		return StringToUint64(value.(string))
	case reflect.Bool:
		return BoolToUint64(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func BoolToUint64(value bool) (uint64, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func Float32ToUint64(value float32) (uint64, error) {
	if value < 0 || value > math.MaxUint64 {
		return 0, errors.New("value out of range for uint64")
	}
	return uint64(value), nil
}

func Float64ToUint64(value float64) (uint64, error) {
	if value < 0 || value > math.MaxUint64 {
		return 0, errors.New("value out of range for uint64")
	}
	return uint64(value), nil
}
func IntToUint64(value int) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}

func Int8ToUint64(value int8) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}

func Int16ToUint64(value int16) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}

func Int32ToUint64(value int32) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}
func Int64ToUint64(value int64) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}
func StringToUint64(value string) (uint64, error) {
	return strconv.ParseUint(value, 10, 64)
}
func UintToUint64(value uint) (uint64, error) {
	return uint64(value), nil
}
func Uint8ToUint64(value uint8) (uint64, error) {
	return uint64(value), nil
}
func Uint16ToUint64(value uint16) (uint64, error) {
	return uint64(value), nil
}
func Uint32ToUint64(value uint32) (uint64, error) {
	return uint64(value), nil
}
func Uint64ToUint64(value uint64) (uint64, error) {
	return value, nil
}
