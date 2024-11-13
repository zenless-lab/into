package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

func TryIntoUint[T convertable](value T) (uint, error) {
	result, err := toUint(value)
	return result.(uint), err
}

func toUint(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint(value.(float64))
	case reflect.Float32:
		return Float32ToUint(value.(float32))
	case reflect.Int:
		return IntToUint(value.(int))
	case reflect.Int8:
		return Int8ToUint(value.(int8))
	case reflect.Int16:
		return Int16ToUint(value.(int16))
	case reflect.Int32:
		return Int32ToUint(value.(int32))
	case reflect.Int64:
		return Int64ToUint(value.(int64))
	case reflect.Uint:
		return value.(uint), nil
	case reflect.Uint8:
		return Uint8ToUint(value.(uint8))
	case reflect.Uint16:
		return Uint16ToUint(value.(uint16))
	case reflect.Uint32:
		return Uint32ToUint(value.(uint32))
	case reflect.Uint64:
		return Uint64ToUint(value.(uint64))
	case reflect.String:
		return StringToUint(value.(string))
	case reflect.Bool:
		return BoolToUint(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func BoolToUint(value bool) (uint, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}
func Float32ToUint(value float32) (uint, error) {
	if value < 0 || value > math.MaxUint {
		return 0, errors.New("value out of range for uint")
	}
	return uint(value), nil
}
func Float64ToUint(value float64) (uint, error) {
	if value < 0 || value > math.MaxUint {
		return 0, errors.New("value out of range for uint")
	}
	return uint(value), nil
}

func IntToUint(value int) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}
func Int8ToUint(value int8) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}

func Int16ToUint(value int16) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}

func Int32ToUint(value int32) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}

func Int64ToUint(value int64) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}

func StringToUint(value string) (uint, error) {
	i, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}
func UintToUint(value uint) (uint, error) {
	return value, nil
}
func Uint8ToUint(value uint8) (uint, error) {
	return uint(value), nil
}
func Uint16ToUint(value uint16) (uint, error) {
	return uint(value), nil
}
func Uint32ToUint(value uint32) (uint, error) {
	return uint(value), nil
}
func Uint64ToUint(value uint64) (uint, error) {
	if value > math.MaxUint {
		return 0, errors.New("value exceeds uint max limit")
	}
	return uint(value), nil
}
