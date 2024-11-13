package into

import (
	"errors"
	"math"
	"strconv"
	"time"
)

func Int32ToFloat64(value int32) (float64, error) {
	return float64(value), nil
}

func Int32ToFloat32(value int32) (float32, error) {
	return float32(value), nil
}

func Int32ToInt(value int32) (int, error) {
	return int(value), nil
}

func Int32ToInt8(value int32) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

func Int32ToInt16(value int32) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}

func Int32ToInt32(value int32) (int32, error) {
	return value, nil
}

func Int32ToInt64(value int32) (int64, error) {
	return int64(value), nil
}

func Int32ToUint(value int32) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}

func Int32ToUint8(value int32) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}

func Int32ToUint16(value int32) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}

func Int32ToUint32(value int32) (uint32, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint32")
	}
	return uint32(value), nil
}

func Int32ToUint64(value int32) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}

func Int32ToString(value int32) (string, error) {
	return strconv.Itoa(int(value)), nil
}

func Int32ToBool(value int32) (bool, error) {
	return value != 0, nil
}

func Int32ToTime(value int32) (time.Time, error) {
	timestamp := int64(value)
	return time.Unix(timestamp, 0), nil
}
