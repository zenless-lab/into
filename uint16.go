package into

import (
	"errors"
	"math"
	"strconv"
	"time"
)

func Uint16ToFloat64(value uint16) (float64, error) {
	return float64(value), nil
}

func Uint16ToFloat32(value uint16) (float32, error) {
	return float32(value), nil
}

func Uint16ToInt(value uint16) (int, error) {
	return int(value), nil
}

func Uint16ToInt8(value uint16) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}

func Uint16ToInt16(value uint16) (int16, error) {
	if value > math.MaxInt16 {
		return 0, errors.New("value exceeds int16 max limit")
	}
	return int16(value), nil
}

func Uint16ToInt32(value uint16) (int32, error) {
	return int32(value), nil
}

func Uint16ToInt64(value uint16) (int64, error) {
	return int64(value), nil
}

func Uint16ToUint(value uint16) (uint, error) {
	return uint(value), nil
}

func Uint16ToUint8(value uint16) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, errors.New("value exceeds uint8 max limit")
	}
	return uint8(value), nil
}

func Uint16ToUint16(value uint16) (uint16, error) {
	return value, nil
}

func Uint16ToUint32(value uint16) (uint32, error) {
	return uint32(value), nil
}

func Uint16ToUint64(value uint16) (uint64, error) {
	return uint64(value), nil
}

func Uint16ToString(value uint16) (string, error) {
	return strconv.FormatUint(uint64(value), 10), nil
}

func Uint16ToBool(value uint16) (bool, error) {
	return value != 0, nil
}

func Uint16ToTime(value uint16) (time.Time, error) {
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}
