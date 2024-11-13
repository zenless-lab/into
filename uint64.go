package into

import (
	"errors"
	"math"
	"strconv"
	"time"
)

func Uint64ToFloat64(value uint64) (float64, error) {
	return float64(value), nil
}

func Uint64ToFloat32(value uint64) (float32, error) {
	if float64(value) > math.MaxFloat32 {
		return 0, errors.New("value exceeds float32 max limit")
	}
	return float32(value), nil
}

func Uint64ToInt(value uint64) (int, error) {
	if value > math.MaxInt {
		return 0, errors.New("value exceeds int max limit")
	}
	return int(value), nil
}

func Uint64ToInt8(value uint64) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}

func Uint64ToInt16(value uint64) (int16, error) {
	if value > math.MaxInt16 {
		return 0, errors.New("value exceeds int16 max limit")
	}
	return int16(value), nil
}

func Uint64ToInt32(value uint64) (int32, error) {
	if value > math.MaxInt32 {
		return 0, errors.New("value exceeds int32 max limit")
	}
	return int32(value), nil
}

func Uint64ToInt64(value uint64) (int64, error) {
	if value > math.MaxInt64 {
		return 0, errors.New("value exceeds int64 max limit")
	}
	return int64(value), nil
}

func Uint64ToUint(value uint64) (uint, error) {
	if value > math.MaxUint {
		return 0, errors.New("value exceeds uint max limit")
	}
	return uint(value), nil
}

func Uint64ToUint8(value uint64) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, errors.New("value exceeds uint8 max limit")
	}
	return uint8(value), nil
}

func Uint64ToUint16(value uint64) (uint16, error) {
	if value > math.MaxUint16 {
		return 0, errors.New("value exceeds uint16 max limit")
	}
	return uint16(value), nil
}

func Uint64ToUint32(value uint64) (uint32, error) {
	if value > math.MaxUint32 {
		return 0, errors.New("value exceeds uint32 max limit")
	}
	return uint32(value), nil
}

func Uint64ToUint64(value uint64) (uint64, error) {
	return value, nil
}

func Uint64ToString(value uint64) (string, error) {
	return strconv.FormatUint(value, 10), nil
}

func Uint64ToBool(value uint64) (bool, error) {
	return value != 0, nil
}

func Uint64ToTime(value uint64) (time.Time, error) {
	if value > uint64(math.MaxInt64) {
		return time.Time{}, errors.New("value exceeds int64 max limit")
	}
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}
