package into

import (
	"errors"
	"math"
	"strconv"
	"time"
)

func Float64ToFloat64(value float64) (float64, error) {
	return value, nil
}

func Float64ToFloat32(value float64) (float32, error) {
	if value > math.MaxFloat32 || value < -math.MaxFloat32 {
		return 0, errors.New("value out of range for float32")
	}
	return float32(value), nil
}

func Float64ToInt(value float64) (int, error) {
	if value > math.MaxInt || value < math.MinInt {
		return 0, errors.New("value out of range for int")
	}
	return int(value), nil
}

func Float64ToInt8(value float64) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

func Float64ToInt16(value float64) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}

func Float64ToInt32(value float64) (int32, error) {
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, errors.New("value out of range for int32")
	}
	return int32(value), nil
}

func Float64ToInt64(value float64) (int64, error) {
	if value > math.MaxInt64 || value < math.MinInt64 {
		return 0, errors.New("value out of range for int64")
	}
	return int64(value), nil
}

func Float64ToUint(value float64) (uint, error) {
	if value < 0 || value > math.MaxUint {
		return 0, errors.New("value out of range for uint")
	}
	return uint(value), nil
}

func Float64ToUint8(value float64) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}

func Float64ToUint16(value float64) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}

func Float64ToUint32(value float64) (uint32, error) {
	if value < 0 || value > math.MaxUint32 {
		return 0, errors.New("value out of range for uint32")
	}
	return uint32(value), nil
}

func Float64ToUint64(value float64) (uint64, error) {
	if value < 0 || value > math.MaxUint64 {
		return 0, errors.New("value out of range for uint64")
	}
	return uint64(value), nil
}

func Float64ToString(value float64) (string, error) {
	return strconv.FormatFloat(value, 'f', -1, 64), nil
}

func Float64ToBool(value float64) (bool, error) {
	return value != 0, nil
}

func Float64ToTime(value float64) (time.Time, error) {
	return time.Time{}, errors.New("float64 cannot be converted to time.Time")
}
