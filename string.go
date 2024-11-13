package into

import (
	"strconv"
	"time"
)

func StringToFloat64(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

func StringToFloat32(value string) (float32, error) {
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

func StringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

func StringToInt8(value string) (int8, error) {
	i, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(i), nil
}

func StringToInt16(value string) (int16, error) {
	i, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(i), nil
}

func StringToInt32(value string) (int32, error) {
	i, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

func StringToInt64(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

func StringToUint(value string) (uint, error) {
	i, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}

func StringToUint8(value string) (uint8, error) {
	i, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(i), nil
}

func StringToUint16(value string) (uint16, error) {
	i, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(i), nil
}

func StringToUint32(value string) (uint32, error) {
	i, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(i), nil
}

func StringToUint64(value string) (uint64, error) {
	return strconv.ParseUint(value, 10, 64)
}

func StringToString(value string) (string, error) {
	return value, nil
}

func StringToBool(value string) (bool, error) {
	return strconv.ParseBool(value)
}

func StringToTime(value string) (time.Time, error) {
	return parseDateWith(value, time.UTC, timeFormats)
}

func StringToDuration(value string) (time.Duration, error) {
	return time.ParseDuration(value)
}
