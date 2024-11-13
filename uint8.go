package into

import (
	"errors"
	"math"
	"strconv"
	"time"
)

func Uint8ToFloat64(value uint8) (float64, error) {
	return float64(value), nil
}

func Uint8ToFloat32(value uint8) (float32, error) {
	return float32(value), nil
}

func Uint8ToInt(value uint8) (int, error) {
	return int(value), nil
}

func Uint8ToInt8(value uint8) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}

func Uint8ToInt16(value uint8) (int16, error) {
	return int16(value), nil
}

func Uint8ToInt32(value uint8) (int32, error) {
	return int32(value), nil
}

func Uint8ToInt64(value uint8) (int64, error) {
	return int64(value), nil
}

func Uint8ToUint(value uint8) (uint, error) {
	return uint(value), nil
}

func Uint8ToUint8(value uint8) (uint8, error) {
	return value, nil
}

func Uint8ToUint16(value uint8) (uint16, error) {
	return uint16(value), nil
}

func Uint8ToUint32(value uint8) (uint32, error) {
	return uint32(value), nil
}

func Uint8ToUint64(value uint8) (uint64, error) {
	return uint64(value), nil
}

func Uint8ToString(value uint8) (string, error) {
	return strconv.FormatUint(uint64(value), 10), nil
}

func Uint8ToBool(value uint8) (bool, error) {
	return value != 0, nil
}

func Uint8ToTime(value uint8) (time.Time, error) {
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}
