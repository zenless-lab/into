package into

import (
	"errors"
	"math"
	"strconv"
	"time"
)

func IntToFloat64(value int) (float64, error) {
	return float64(value), nil
}

func IntToFloat32(value int) (float32, error) {
	if float64(value) > math.MaxFloat32 || float64(value) < -math.MaxFloat32 {
		return 0, errors.New("value out of range for float32")
	}
	return float32(value), nil
}

func IntToInt(value int) (int, error) {
	return value, nil
}

func IntToInt8(value int) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

func IntToInt16(value int) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}

func IntToInt32(value int) (int32, error) {
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, errors.New("value out of range for int32")
	}
	return int32(value), nil
}

func IntToInt64(value int) (int64, error) {
	return int64(value), nil
}

func IntToUint(value int) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}

func IntToUint8(value int) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}

func IntToUint16(value int) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}

func IntToUint32(value int) (uint32, error) {
	if value < 0 || value > math.MaxUint32 {
		return 0, errors.New("value out of range for uint32")
	}
	return uint32(value), nil
}

func IntToUint64(value int) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}

func IntToString(value int) (string, error) {
	return strconv.Itoa(value), nil
}

func IntToBool(value int) (bool, error) {
	return value != 0, nil
}

func IntToTime(value int) (time.Time, error) {
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}
