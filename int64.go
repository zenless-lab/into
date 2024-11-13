package into

import (
	"errors"
	"math"
	"strconv"
	"time"
)

func Int64ToFloat64(value int64) (float64, error) {
	return float64(value), nil
}

func Int64ToFloat32(value int64) (float32, error) {
	if float64(value) > math.MaxFloat32 || float64(value) < -math.MaxFloat32 {
		return 0, errors.New("value out of range for float32")
	}
	return float32(value), nil
}

func Int64ToInt(value int64) (int, error) {
	if value > math.MaxInt || value < math.MinInt {
		return 0, errors.New("value out of range for int")
	}
	return int(value), nil
}

func Int64ToInt8(value int64) (int8, error) {
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}
	return int8(value), nil
}

func Int64ToInt16(value int64) (int16, error) {
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}
	return int16(value), nil
}

func Int64ToInt32(value int64) (int32, error) {
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, errors.New("value out of range for int32")
	}
	return int32(value), nil
}

func Int64ToInt64(value int64) (int64, error) {
	return value, nil
}

func Int64ToUint(value int64) (uint, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint")
	}
	return uint(value), nil
}

func Int64ToUint8(value int64) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}
	return uint8(value), nil
}

func Int64ToUint16(value int64) (uint16, error) {
	if value < 0 || value > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}
	return uint16(value), nil
}

func Int64ToUint32(value int64) (uint32, error) {
	if value < 0 || value > math.MaxUint32 {
		return 0, errors.New("value out of range for uint32")
	}
	return uint32(value), nil
}

func Int64ToUint64(value int64) (uint64, error) {
	if value < 0 {
		return 0, errors.New("negative value cannot be converted to uint64")
	}
	return uint64(value), nil
}

func Int64ToString(value int64) (string, error) {
	iValue, err := Int64ToInt(value)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(iValue), nil
}

func Int64ToBool(value int64) (bool, error) {
	return value != 0, nil
}

func Int64ToTime(value int64) (time.Time, error) {
	timestamp := value
	return time.Unix(timestamp, 0), nil
}
