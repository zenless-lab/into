package into

import (
	"errors"
	"math"
	"strconv"
	"time"
)

func UintToFloat64(value uint) (float64, error) {
	return float64(value), nil
}

func UintToFloat32(value uint) (float32, error) {
	return float32(value), nil
}

func UintToInt(value uint) (int, error) {
	if value > math.MaxInt {
		return 0, errors.New("value exceeds int max limit")
	}
	return int(value), nil
}

func UintToInt8(value uint) (int8, error) {
	if value > math.MaxInt8 {
		return 0, errors.New("value exceeds int8 max limit")
	}
	return int8(value), nil
}

func UintToInt16(value uint) (int16, error) {
	if value > math.MaxInt16 {
		return 0, errors.New("value exceeds int16 max limit")
	}
	return int16(value), nil
}

func UintToInt32(value uint) (int32, error) {
	if value > math.MaxInt32 {
		return 0, errors.New("value exceeds int32 max limit")
	}
	return int32(value), nil
}

func UintToInt64(value uint) (int64, error) {
	if value > math.MaxInt64 {
		return 0, errors.New("value exceeds int64 max limit")
	}
	return int64(value), nil
}

func UintToUint(value uint) (uint, error) {
	return value, nil
}

func UintToUint8(value uint) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, errors.New("value exceeds uint8 max limit")
	}
	return uint8(value), nil
}

func UintToUint16(value uint) (uint16, error) {
	if value > math.MaxUint16 {
		return 0, errors.New("value exceeds uint16 max limit")
	}
	return uint16(value), nil
}

func UintToUint32(value uint) (uint32, error) {
	if value > math.MaxUint32 {
		return 0, errors.New("value exceeds uint32 max limit")
	}
	return uint32(value), nil
}

func UintToUint64(value uint) (uint64, error) {
	return uint64(value), nil
}

func UintToString(value uint) (string, error) {
	return strconv.FormatUint(uint64(value), 10), nil
}

func UintToBool(value uint) (bool, error) {
	return value != 0, nil
}

func UintToTime(value uint) (time.Time, error) {
	if value > uint(math.MaxInt64) {
		return time.Time{}, errors.New("value exceeds int64 max limit")
	}

	timestamp := int64(value)
	return time.Unix(timestamp, 0), nil
}
