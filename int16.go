package into

import (
	"errors"
	"math"
	"strconv"
	"time"
)

func Int16ToFloat64(value int16) (float64, error) {
	return float64(value), nil
}

func Int16ToFloat32(value int16) (float32, error) {
	return float32(value), nil
}

func Int16ToInt(value int16) (int, error) {
	return int(value), nil
}

func Int16ToInt8(value int16) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

func Int16ToInt16(value int16) (int16, error) {
	return value, nil
}

func Int16ToInt32(value int16) (int32, error) {
	return int32(value), nil
}

func Int16ToInt64(value int16) (int64, error) {
	return int64(value), nil
}

func Int16ToUint(value int16) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}

func Int16ToUint8(value int16) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}

func Int16ToUint16(value int16) (uint16, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint16")
	}
	return uint16(value), nil
}

func Int16ToUint32(value int16) (uint32, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint32")
	}
	return uint32(value), nil
}

func Int16ToUint64(value int16) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}

func Int16ToString(value int16) (string, error) {
	return strconv.Itoa(int(value)), nil
}

func Int16ToBool(value int16) (bool, error) {
	return value != 0, nil
}

func Int16ToTime(value int16) (time.Time, error) {
	timestamp := int64(value)
	return time.Unix(timestamp, 0), nil
}
