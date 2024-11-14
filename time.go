package into

import (
	"errors"
	"math"
	"reflect"
	"time"
)

// TryIntoTime converts the given value to a time.Time value.
//
// TryIntoTime attempts to convert the given value to a time.Time value.
// It supports the following types:
//   - int
//   - int8
//   - int16
//   - int32
//   - int64
//   - uint
//   - uint8
//   - uint16
//   - uint32
//   - string
//   - time.Time
//
// If the given value is not one of the supported types, it returns an error.
// If the given value is a string, it attempts to parse it using the following formats:
//   - "2006-01-02T15:04:05.999999-0700MST"
//   - "2006-01-02T15:04:05.999999-0700"
//   - "2006-01-02T15:04:05.999999"
//   - "2006-01-02T15:04:05"
//   - "2006-01-02"
//   - "2006-01-02 15:04:05"
//   - "15:04:05"
//   - "Jan 2, 2006"
//   - "2006-01-02 15:04"
//   - "2006-01-02 15:04:05.999999-0700"
//   - "2006-01-02 15:04:05.999999Z07:00"
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - time.Time: the converted time.Time value.
//   - error: an error if the conversion fails.
//
// Example:
//
//	result, err := TryIntoTime(1678867200)
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(result) // Output: 2023-03-15T00:00:00Z
func TryIntoTime[T String | Int | Uint | time.Time](value T) (time.Time, error) {
	result, err := toTime(value)
	return result.(time.Time), err
}

// toTime converts the given value to a time.Time value.
//
// toTime attempts to convert the given value to a time.Time value.
// It supports the following types:
//   - int
//   - int8
//   - int16
//   - int32
//   - int64
//   - uint
//   - uint8
//   - uint16
//   - uint32
//   - string
//
// If the given value is not one of the supported types, it returns an error.
// If the given value is a string, it attempts to parse it using the following formats:
//   - "2006-01-02T15:04:05.999999-0700MST"
//   - "2006-01-02T15:04:05.999999-0700"
//   - "2006-01-02T15:04:05.999999"
//   - "2006-01-02T15:04:05"
//   - "2006-01-02"
//   - "2006-01-02 15:04:05"
//   - "15:04:05"
//   - "Jan 2, 2006"
//   - "2006-01-02 15:04"
//   - "2006-01-02 15:04:05.999999-0700"
//   - "2006-01-02 15:04:05.999999Z07:00"
//
// Parameters:
//   - value: the value to be converted.
//
// Returns:
//   - any: the converted value.
//   - error: an error if the conversion fails.
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

// IntToTime converts the given int value to a time.Time value.
//
// IntToTime converts the given int value to a time.Time value, representing
// the number of seconds since the Unix epoch.
//
// Parameters:
//   - value: the int value to be converted.
//
// Returns:
//   - time.Time: the converted time.Time value.
//   - error: nil.
func IntToTime(value int) (time.Time, error) {
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}

// Int8ToTime converts the given int8 value to a time.Time value.
//
// Int8ToTime converts the given int8 value to a time.Time value, representing
// the number of seconds since the Unix epoch.
//
// Parameters:
//   - value: the int8 value to be converted.
//
// Returns:
//   - time.Time: the converted time.Time value.
//   - error: nil.
func Int8ToTime(value int8) (time.Time, error) {
	timestamp := int64(value)
	return time.Unix(timestamp, 0), nil
}

// Int16ToTime converts the given int16 value to a time.Time value.
//
// Int16ToTime converts the given int16 value to a time.Time value, representing
// the number of seconds since the Unix epoch.
//
// Parameters:
//   - value: the int16 value to be converted.
//
// Returns:
//   - time.Time: the converted time.Time value.
//   - error: nil.
func Int16ToTime(value int16) (time.Time, error) {
	timestamp := int64(value)
	return time.Unix(timestamp, 0), nil
}

// Int32ToTime converts the given int32 value to a time.Time value.
//
// Int32ToTime converts the given int32 value to a time.Time value, representing
// the number of seconds since the Unix epoch.
//
// Parameters:
//   - value: the int32 value to be converted.
//
// Returns:
//   - time.Time: the converted time.Time value.
//   - error: nil.
func Int32ToTime(value int32) (time.Time, error) {
	timestamp := int64(value)
	return time.Unix(timestamp, 0), nil
}

// Int64ToTime converts the given int64 value to a time.Time value.
//
// Int64ToTime converts the given int64 value to a time.Time value, representing
// the number of seconds since the Unix epoch.
//
// Parameters:
//   - value: the int64 value to be converted.
//
// Returns:
//   - time.Time: the converted time.Time value.
//   - error: nil.
func Int64ToTime(value int64) (time.Time, error) {
	timestamp := value
	return time.Unix(timestamp, 0), nil
}

// StringToTime converts the given string value to a time.Time value.
//
// StringToTime converts the given string value to a time.Time value, using
// the following formats:
//   - "2006-01-02T15:04:05.999999-0700MST"
//   - "2006-01-02T15:04:05.999999-0700"
//   - "2006-01-02T15:04:05.999999"
//   - "2006-01-02T15:04:05"
//   - "2006-01-02"
//   - "2006-01-02 15:04:05"
//   - "15:04:05"
//   - "Jan 2, 2006"
//   - "2006-01-02 15:04"
//   - "2006-01-02 15:04:05.999999-0700"
//   - "2006-01-02 15:04:05.999999Z07:00"
//
// Parameters:
//   - value: the string value to be converted.
//
// Returns:
//   - time.Time: the converted time.Time value.
//   - error: an error if the conversion fails.
func StringToTime(value string) (time.Time, error) {
	return parseDateWith(value, time.UTC, timeFormats)
}

// StringToDuration converts the given string value to a time.Duration value.
//
// StringToDuration converts the given string value to a time.Duration value,
// using the standard time duration format, e.g. "1h30m".
//
// Parameters:
//   - value: the string value to be converted.
//
// Returns:
//   - time.Duration: the converted time.Duration value.
//   - error: an error if the conversion fails.
func StringToDuration(value string) (time.Duration, error) {
	return time.ParseDuration(value)
}

// UintToTime converts the given uint value to a time.Time value.
//
// UintToTime converts the given uint value to a time.Time value, representing
// the number of seconds since the Unix epoch.
//
// Parameters:
//   - value: the uint value to be converted.
//
// Returns:
//   - time.Time: the converted time.Time value.
//   - error: an error if the value exceeds the int64 max limit.
func UintToTime(value uint) (time.Time, error) {
	if value > uint(math.MaxInt64) {
		return time.Time{}, errors.New("value exceeds int64 max limit")
	}

	timestamp := int64(value)
	return time.Unix(timestamp, 0), nil
}

// Uint8ToTime converts the given uint8 value to a time.Time value.
//
// Uint8ToTime converts the given uint8 value to a time.Time value, representing
// the number of seconds since the Unix epoch.
//
// Parameters:
//   - value: the uint8 value to be converted.
//
// Returns:
//   - time.Time: the converted time.Time value.
//   - error: nil.
func Uint8ToTime(value uint8) (time.Time, error) {
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}

// Uint16ToTime converts the given uint16 value to a time.Time value.
//
// Uint16ToTime converts the given uint16 value to a time.Time value, representing
// the number of seconds since the Unix epoch.
//
// Parameters:
//   - value: the uint16 value to be converted.
//
// Returns:
//   - time.Time: the converted time.Time value.
//   - error: nil.
func Uint16ToTime(value uint16) (time.Time, error) {
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}

// Uint32ToTime converts the given uint32 value to a time.Time value.
//
// Uint32ToTime converts the given uint32 value to a time.Time value, representing
// the number of seconds since the Unix epoch.
//
// Parameters:
//   - value: the uint32 value to be converted.
//
// Returns:
//   - time.Time: the converted time.Time value.
//   - error: nil.
func Uint32ToTime(value uint32) (time.Time, error) {
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}

// Uint64ToTime converts the given uint64 value to a time.Time value.
//
// Uint64ToTime converts the given uint64 value to a time.Time value, representing
// the number of seconds since the Unix epoch.
//
// Parameters:
//   - value: the uint64 value to be converted.
//
// Returns:
//   - time.Time: the converted time.Time value.
//   - error: an error if the value exceeds the int64 max limit.
func Uint64ToTime(value uint64) (time.Time, error) {
	if value > uint64(math.MaxInt64) {
		return time.Time{}, errors.New("value exceeds int64 max limit")
	}
	timestamp := int64(value)

	return time.Unix(timestamp, 0), nil
}
