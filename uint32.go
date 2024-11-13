package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

func TryIntoUint32[T convertable](value T) (uint32, error) {
	result, err := toUint32(value)
	return result.(uint32), err
}

func toUint32(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint32(value.(float64))
	case reflect.Float32:
		return Float32ToUint32(value.(float32))
	case reflect.Int:
		return IntToUint32(value.(int))
	case reflect.Int8:
		return Int8ToUint32(value.(int8))
	case reflect.Int16:
		return Int16ToUint32(value.(int16))
	case reflect.Int32:
		return Int32ToUint32(value.(int32))
	case reflect.Int64:
		return Int64ToUint32(value.(int64))
	case reflect.Uint:
		return UintToUint32(value.(uint))
	case reflect.Uint8:
		return Uint8ToUint32(value.(uint8))
	case reflect.Uint16:
		return Uint16ToUint32(value.(uint16))
	case reflect.Uint32:
		return value.(uint32), nil
	case reflect.Uint64:
		return Uint64ToUint32(value.(uint64))
	case reflect.String:
		return StringToUint32(value.(string))
	case reflect.Bool:
		return BoolToUint32(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func BoolToUint32(value bool) (uint32, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}
func Float32ToUint32(value float32) (uint32, error) {
	if value < 0 || value > math.MaxUint32 {
		return 0, errors.New("value out of range for uint32")
	}
	return uint32(value), nil
}

func Float64ToUint32(value float64) (uint32, error) {
	if value < 0 || value > math.MaxUint32 {
		return 0, errors.New("value out of range for uint32")
	}
	return uint32(value), nil
}

func IntToUint32(value int) (uint32, error) {
	if value < 0 || value > math.MaxUint32 {
		return 0, errors.New("value out of range for uint32")
	}
	return uint32(value), nil
}
func Int8ToUint32(value int8) (uint32, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint32")
	}
	return uint32(value), nil
}
func Int16ToUint32(value int16) (uint32, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint32")
	}
	return uint32(value), nil
}
func Int32ToUint32(value int32) (uint32, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint32")
	}
	return uint32(value), nil
}
func Int64ToUint32(value int64) (uint32, error) {
	if value < 0 || value > math.MaxUint32 {
		return 0, errors.New("value out of range for uint32")
	}
	return uint32(value), nil
}
func StringToUint32(value string) (uint32, error) {
	i, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(i), nil
}
func UintToUint32(value uint) (uint32, error) {
	if value > math.MaxUint32 {
		return 0, errors.New("value exceeds uint32 max limit")
	}
	return uint32(value), nil
}
func Uint8ToUint32(value uint8) (uint32, error) {
	return uint32(value), nil
}
func Uint16ToUint32(value uint16) (uint32, error) {
	return uint32(value), nil
}
func Uint32ToUint32(value uint32) (uint32, error) {
	return value, nil
}
func Uint64ToUint32(value uint64) (uint32, error) {
	if value > math.MaxUint32 {
		return 0, errors.New("value exceeds uint32 max limit")
	}
	return uint32(value), nil
}
