package into

import (
	"errors"
	"math"
	"reflect"
	"time"
)

func TryIntoTime[T String | Int | Uint | time.Time](value T) (time.Time, error) {
	result, err := toTime(value)
	return result.(time.Time), err
}

func toTime(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Int:
		return IntToTime(value.(int))
	case reflect.Int8:
		return Int8ToTime(value.(int8))
	case reflect.Int16:
		return Int16ToTime(value.(int16))
	case reflect.Int32:
		return Int32ToTime(value.(int32))
	case reflect.Int64:
		return Int64ToTime(value.(int64))
	case reflect.Uint:
		return UintToTime(value.(uint))
	case reflect.Uint8:
		return Uint8ToTime(value.(uint8))
	case reflect.Uint16:
		return Uint16ToTime(value.(uint16))
	case reflect.Uint32:
		return Uint32ToTime(value.(uint32))
	case reflect.String:
		return StringToTime(value.(string))
	default:
		return time.Time{}, errors.New("unsupported type")
	}
}

func IntToTime(value int) (time.Time, error) {
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}

func Int8ToTime(value int8) (time.Time, error) {
	timestamp := int64(value)
	return time.Unix(timestamp, 0), nil
}
func Int16ToTime(value int16) (time.Time, error) {
	timestamp := int64(value)
	return time.Unix(timestamp, 0), nil
}

func Int32ToTime(value int32) (time.Time, error) {
	timestamp := int64(value)
	return time.Unix(timestamp, 0), nil
}

func Int64ToTime(value int64) (time.Time, error) {
	timestamp := value
	return time.Unix(timestamp, 0), nil
}

func StringToTime(value string) (time.Time, error) {
	return parseDateWith(value, time.UTC, timeFormats)
}
func StringToDuration(value string) (time.Duration, error) {
	return time.ParseDuration(value)
}

func UintToTime(value uint) (time.Time, error) {
	if value > uint(math.MaxInt64) {
		return time.Time{}, errors.New("value exceeds int64 max limit")
	}

	timestamp := int64(value)
	return time.Unix(timestamp, 0), nil
}

func Uint8ToTime(value uint8) (time.Time, error) {
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}

func Uint16ToTime(value uint16) (time.Time, error) {
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}

func Uint32ToTime(value uint32) (time.Time, error) {
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}

func Uint64ToTime(value uint64) (time.Time, error) {
	if value > uint64(math.MaxInt64) {
		return time.Time{}, errors.New("value exceeds int64 max limit")
	}
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}
