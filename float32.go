package into

import (
	"errors"
	"math"
	"strconv"
	"time"
)

func Float32ToFloat64(value float32) (float64, error) {
	return float64(value), nil
}

func Float32ToFloat32(value float32) (float32, error) {
	return value, nil
}

func Float32ToInt(value float32) (int, error) {
	if value > math.MaxInt || value < math.MinInt {
		return 0, errors.New("value out of range for int")
	}
	return int(value), nil
}

func Float32ToInt8(value float32) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

func Float32ToInt16(value float32) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}

func Float32ToInt32(value float32) (int32, error) {
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, errors.New("value out of range for int32")
	}
	return int32(value), nil
}

func Float32ToInt64(value float32) (int64, error) {
	if value > math.MaxInt64 || value < math.MinInt64 {
		return 0, errors.New("value out of range for int64")
	}
	return int64(value), nil
}

func Float32ToUint(value float32) (uint, error) {
	if value < 0 || value > math.MaxUint {
		return 0, errors.New("value out of range for uint")
	}
	return uint(value), nil
}

func Float32ToUint8(value float32) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}

func Float32ToUint16(value float32) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}

func Float32ToUint32(value float32) (uint32, error) {
	if value < 0 || value > math.MaxUint32 {
		return 0, errors.New("value out of range for uint32")
	}
	return uint32(value), nil
}

func Float32ToUint64(value float32) (uint64, error) {
	if value < 0 || value > math.MaxUint64 {
		return 0, errors.New("value out of range for uint64")
	}
	return uint64(value), nil
}

func Float32ToString(value float32) (string, error) {
	return strconv.FormatFloat(float64(value), 'f', -1, 32), nil
}

func Float32ToBool(value float32) (bool, error) {
	return value != 0, nil
}

func Float32ToTime(value float32) (time.Time, error) {
	return time.Time{}, errors.New("float32 cannot be converted to time.Time")
}
