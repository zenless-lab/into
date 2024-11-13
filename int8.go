package into

import (
	"errors"
	"strconv"
	"time"
)

func Int8ToFloat64(value int8) (float64, error) {
	return float64(value), nil
}

func Int8ToFloat32(value int8) (float32, error) {
	return float32(value), nil
}

func Int8ToInt(value int8) (int, error) {
	return int(value), nil
}

func Int8ToInt8(value int8) (int8, error) {
	return value, nil
}

func Int8ToInt16(value int8) (int16, error) {
	return int16(value), nil
}

func Int8ToInt32(value int8) (int32, error) {
	return int32(value), nil
}

func Int8ToInt64(value int8) (int64, error) {
	return int64(value), nil
}

func Int8ToUint(value int8) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}

func Int8ToUint8(value int8) (uint8, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint8")
	}
	return uint8(value), nil
}

func Int8ToUint16(value int8) (uint16, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint16")
	}
	return uint16(value), nil
}

func Int8ToUint32(value int8) (uint32, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint32")
	}
	return uint32(value), nil
}

func Int8ToUint64(value int8) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}

func Int8ToString(value int8) (string, error) {
	return strconv.Itoa(int(value)), nil
}

func Int8ToBool(value int8) (bool, error) {
	return value != 0, nil
}

func Int8ToTime(value int8) (time.Time, error) {
	timestamp := int64(value)
	return time.Unix(timestamp, 0), nil
}
