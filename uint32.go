package into

import (
	"errors"
	"math"
	"strconv"
	"time"
)

func Uint32ToFloat64(value uint32) (float64, error) {
	return float64(value), nil
}

func Uint32ToFloat32(value uint32) (float32, error) {
	if float64(value) > math.MaxFloat32 {
		return 0, errors.New("value exceeds float32 max limit")
	}
	return float32(value), nil
}

func Uint32ToInt(value uint32) (int, error) {
	if uint64(value) > math.MaxInt {
		return 0, errors.New("value exceeds int max limit")
	}
	return int(value), nil
}

func Uint32ToInt8(value uint32) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}

func Uint32ToInt16(value uint32) (int16, error) {
	if value > math.MaxInt16 {
		return 0, errors.New("value exceeds int16 max limit")
	}
	return int16(value), nil
}

func Uint32ToInt32(value uint32) (int32, error) {
	if value > math.MaxInt32 {
		return 0, errors.New("value exceeds int32 max limit")
	}
	return int32(value), nil
}

func Uint32ToInt64(value uint32) (int64, error) {
	return int64(value), nil
}

func Uint32ToUint(value uint32) (uint, error) {
	return uint(value), nil
}

func Uint32ToUint8(value uint32) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, errors.New("value exceeds uint8 max limit")
	}
	return uint8(value), nil
}

func Uint32ToUint16(value uint32) (uint16, error) {
	if value > math.MaxUint16 {
		return 0, errors.New("value exceeds uint16 max limit")
	}
	return uint16(value), nil
}

func Uint32ToUint32(value uint32) (uint32, error) {
	return value, nil
}

func Uint32ToUint64(value uint32) (uint64, error) {
	return uint64(value), nil
}

func Uint32ToString(value uint32) (string, error) {
	return strconv.FormatUint(uint64(value), 10), nil
}

func Uint32ToBool(value uint32) (bool, error) {
	return value != 0, nil
}

func Uint32ToTime(value uint32) (time.Time, error) {
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}
