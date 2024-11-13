package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

func TryIntoUint8[T convertable](value T) (uint8, error) {
	result, err := toUint8(value)
	return result.(uint8), err
}

func toUint8(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint8(value.(float64))
	case reflect.Float32:
		return Float32ToUint8(value.(float32))
	case reflect.Int:
		return IntToUint8(value.(int))
	case reflect.Int8:
		return Int8ToUint8(value.(int8))
	case reflect.Int16:
		return Int16ToUint8(value.(int16))
	case reflect.Int32:
		return Int32ToUint8(value.(int32))
	case reflect.Int64:
		return Int64ToUint8(value.(int64))
	case reflect.Uint:
		return UintToUint8(value.(uint))
	case reflect.Uint8:
		return value.(uint8), nil
	case reflect.Uint16:
		return Uint16ToUint8(value.(uint16))
	case reflect.Uint32:
		return Uint32ToUint8(value.(uint32))
	case reflect.Uint64:
		return Uint64ToUint8(value.(uint64))
	case reflect.String:
		return StringToUint8(value.(string))
	case reflect.Bool:
		return BoolToUint8(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func BoolToUint8(value bool) (uint8, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func Float32ToUint8(value float32) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}
func Float64ToUint8(value float64) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}
func IntToUint8(value int) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}

func Int8ToUint8(value int8) (uint8, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint8")
	}
	return uint8(value), nil
}

func Int16ToUint8(value int16) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}
func Int32ToUint8(value int32) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}
func Int64ToUint8(value int64) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}
func StringToUint8(value string) (uint8, error) {
	i, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(i), nil
}
func UintToUint8(value uint) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, errors.New("value exceeds uint8 max limit")
	}
	return uint8(value), nil
}
func Uint8ToUint8(value uint8) (uint8, error) {
	return value, nil
}
func Uint16ToUint8(value uint16) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, errors.New("value exceeds uint8 max limit")
	}
	return uint8(value), nil
}

func Uint32ToUint8(value uint32) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, errors.New("value exceeds uint8 max limit")
	}
	return uint8(value), nil
}
func Uint64ToUint8(value uint64) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, errors.New("value exceeds uint8 max limit")
	}
	return uint8(value), nil
}
