package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

func TryIntoUint16[T convertable](value T) (uint16, error) {
	result, err := toUint16(value)
	return result.(uint16), err
}

func toUint16(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint16(value.(float64))
	case reflect.Float32:
		return Float32ToUint16(value.(float32))
	case reflect.Int:
		return IntToUint16(value.(int))
	case reflect.Int8:
		return Int8ToUint16(value.(int8))
	case reflect.Int16:
		return Int16ToUint16(value.(int16))
	case reflect.Int32:
		return Int32ToUint16(value.(int32))
	case reflect.Int64:
		return Int64ToUint16(value.(int64))
	case reflect.Uint:
		return UintToUint16(value.(uint))
	case reflect.Uint8:
		return Uint8ToUint16(value.(uint8))
	case reflect.Uint16:
		return value.(uint16), nil
	case reflect.Uint32:
		return Uint32ToUint16(value.(uint32))
	case reflect.Uint64:
		return Uint64ToUint16(value.(uint64))
	case reflect.String:
		return StringToUint16(value.(string))
	case reflect.Bool:
		return BoolToUint16(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func BoolToUint16(value bool) (uint16, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func Float32ToUint16(value float32) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}

func Float64ToUint16(value float64) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}

func IntToUint16(value int) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}
func Int8ToUint16(value int8) (uint16, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint16")
	}
	return uint16(value), nil
}
func Int16ToUint16(value int16) (uint16, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint16")
	}
	return uint16(value), nil
}

func Int32ToUint16(value int32) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}
func Int64ToUint16(value int64) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}
func StringToUint16(value string) (uint16, error) {
	i, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(i), nil
}
func UintToUint16(value uint) (uint16, error) {
	if value > math.MaxUint16 {
		return 0, errors.New("value exceeds uint16 max limit")
	}
	return uint16(value), nil
}
func Uint8ToUint16(value uint8) (uint16, error) {
	return uint16(value), nil
}
func Uint16ToUint16(value uint16) (uint16, error) {
	return value, nil
}
func Uint32ToUint16(value uint32) (uint16, error) {
	if value > math.MaxUint16 {
		return 0, errors.New("value exceeds uint16 max limit")
	}
	return uint16(value), nil
}
func Uint64ToUint16(value uint64) (uint16, error) {
	if value > math.MaxUint16 {
		return 0, errors.New("value exceeds uint16 max limit")
	}
	return uint16(value), nil
}
