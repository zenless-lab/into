package into

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

func TryIntoFloat32[T convertable](value T) (float32, error) {
	resoult, err := toFloat32(value)
	return resoult.(float32), err
}

func toFloat32(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToFloat32(value.(float64))
	case reflect.Float32:
		return value.(float32), nil
	case reflect.Int:
		return IntToFloat32(value.(int))
	case reflect.Int8:
		return Int8ToFloat32(value.(int8))
	case reflect.Int16:
		return Int16ToFloat32(value.(int16))
	case reflect.Int32:
		return Int32ToFloat32(value.(int32))
	case reflect.Int64:
		return Int64ToFloat32(value.(int64))
	case reflect.Uint:
		return UintToFloat32(value.(uint))
	case reflect.Uint8:
		return Uint8ToFloat32(value.(uint8))
	case reflect.Uint16:
		return Uint16ToFloat32(value.(uint16))
	case reflect.Uint32:
		return Uint32ToFloat32(value.(uint32))
	case reflect.Uint64:
		return Uint64ToFloat32(value.(uint64))
	case reflect.String:
		return StringToFloat32(value.(string))
	case reflect.Bool:
		return BoolToFloat32(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func BoolToFloat32(value bool) (float32, error) {
	if value {
		return 1.0, nil
	}
	return 0.0, nil
}

func Float32ToFloat32(value float32) (float32, error) {
	return value, nil
}

func Float64ToFloat32(value float64) (float32, error) {
	if value > math.MaxFloat32 || value < -math.MaxFloat32 {
		return 0, errors.New("value out of range for float32")
	}
	return float32(value), nil
}

func IntToFloat32(value int) (float32, error) {
	if float64(value) > math.MaxFloat32 || float64(value) < -math.MaxFloat32 {
		return 0, errors.New("value out of range for float32")
	}
	return float32(value), nil
}
func Int8ToFloat32(value int8) (float32, error) {
	return float32(value), nil
}

func Int16ToFloat32(value int16) (float32, error) {
	return float32(value), nil
}

func Int32ToFloat32(value int32) (float32, error) {
	return float32(value), nil
}

func Int64ToFloat32(value int64) (float32, error) {
	if float64(value) > math.MaxFloat32 || float64(value) < -math.MaxFloat32 {
		return 0, errors.New("value out of range for float32")
	}
	return float32(value), nil
}

func StringToFloat32(value string) (float32, error) {
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

func UintToFloat32(value uint) (float32, error) {
	return float32(value), nil
}

func Uint8ToFloat32(value uint8) (float32, error) {
	return float32(value), nil
}

func Uint16ToFloat32(value uint16) (float32, error) {
	return float32(value), nil
}
func Uint32ToFloat32(value uint32) (float32, error) {
	if float64(value) > math.MaxFloat32 {
		return 0, errors.New("value exceeds float32 max limit")
	}
	return float32(value), nil
}

func Uint64ToFloat32(value uint64) (float32, error) {
	if float64(value) > math.MaxFloat32 {
		return 0, errors.New("value exceeds float32 max limit")
	}
	return float32(value), nil
}
