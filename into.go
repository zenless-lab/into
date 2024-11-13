package into

import (
	"errors"
	"reflect"
	"time"
)

type Float interface {
	~float64 | ~float32
}

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type String interface {
	~string
}

type Bool interface {
	~bool
}

type buildin interface {
	Float | Int | Uint | String | Bool
}

func Into[U buildin, T buildin](value T) U {
	result, err := TryInto[U, T](value)
	if err != nil {
		panic(err)
	}
	return result
}

func TryInto[U buildin, T buildin](value T) (result U, err error) {
	resultType := reflect.TypeOf(result)

	var r any
	switch resultType.Kind() {
	case reflect.Float64:
		r, err = toFloat64(value)
	case reflect.Float32:
		r, err = toFloat32(value)
	case reflect.Int:
		r, err = toInt(value)
	case reflect.Int8:
		r, err = toInt8(value)
	case reflect.Int16:
		r, err = toInt16(value)
	case reflect.Int32:
		r, err = toInt32(value)
	case reflect.Int64:
		r, err = toInt64(value)
	case reflect.Uint:
		r, err = toUint(value)
	case reflect.Uint8:
		r, err = toUint8(value)
	case reflect.Uint16:
		r, err = toUint16(value)
	case reflect.Uint32:
		r, err = toUint32(value)
	case reflect.Uint64:
		r, err = toUint64(value)
	case reflect.String:
		r, err = toString(value)
	case reflect.Bool:
		r, err = toBool(value)
	case reflect.Struct:
		if resultType.Name() == "Time" {
			r, err = toTime(value)
		} else {
			err = errors.New("unsupported type")
		}
	default:
		intoResult, isInto := checkAndCallInto[U](value)
		if isInto {
			result = intoResult
			return
		}

		tryIntoResult, tryIntoErr, isTryInto := checkAndCallTryInto[U](value)
		if isTryInto {
			if tryIntoErr == nil {
				result = tryIntoResult
				return
			}
			err = tryIntoErr
			return
		}
		err = errors.New("unsupported type")
		return
	}

	result = r.(U)
	return
}

func checkAndCallInto[T any](value any) (result T, isInto bool) {
	valueType := reflect.TypeOf(value)
	if valueType.Kind() != reflect.Struct {
		return
	}

	intoType := reflect.TypeOf((*IInto[T])(nil)).Elem()
	if valueType.Implements(intoType) {
		intoValue := value.(IInto[T])
		result = intoValue.Into()
		isInto = true
	}
	return
}

func checkAndCallTryInto[T any](value any) (result T, err error, isTryInto bool) {
	valueType := reflect.TypeOf(value)
	if valueType.Kind() != reflect.Struct {
		return
	}

	tryIntoType := reflect.TypeOf((*ITryInto[T])(nil)).Elem()
	if valueType.Implements(tryIntoType) {
		tryIntoValue := value.(ITryInto[T])
		result, err = tryIntoValue.TryInto()
		isTryInto = true
	}
	return
}

func TryIntoFloat64[T buildin](value T) (float64, error) {
	result, err := toFloat64(value)
	return result.(float64), err
}

func toFloat64(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return value.(float64), nil
	case reflect.Float32:
		return Float32ToFloat64(value.(float32))
	case reflect.Int:
		return IntToFloat64(value.(int))
	case reflect.Int8:
		return Int8ToFloat64(value.(int8))
	case reflect.Int16:
		return Int16ToFloat64(value.(int16))
	case reflect.Int32:
		return Int32ToFloat64(value.(int32))
	case reflect.Int64:
		return Int64ToFloat64(value.(int64))
	case reflect.Uint:
		return UintToFloat64(value.(uint))
	case reflect.Uint8:
		return Uint8ToFloat64(value.(uint8))
	case reflect.Uint16:
		return Uint16ToFloat64(value.(uint16))
	case reflect.Uint32:
		return Uint32ToFloat64(value.(uint32))
	case reflect.Uint64:
		return Uint64ToFloat64(value.(uint64))
	case reflect.String:
		return StringToFloat64(value.(string))
	case reflect.Bool:
		return BoolToFloat64(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func TryIntoFloat32[T buildin](value T) (float32, error) {
	resoult, err := toFloat32(value)
	return resoult.(float32), err
}

func toFloat32(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToFloat32(value.(float64))
	case reflect.Float32:
		return value.(float32), nil
	case reflect.Int:
		return IntToFloat32(value.(int))
	case reflect.Int8:
		return Int8ToFloat32(value.(int8))
	case reflect.Int16:
		return Int16ToFloat32(value.(int16))
	case reflect.Int32:
		return Int32ToFloat32(value.(int32))
	case reflect.Int64:
		return Int64ToFloat32(value.(int64))
	case reflect.Uint:
		return UintToFloat32(value.(uint))
	case reflect.Uint8:
		return Uint8ToFloat32(value.(uint8))
	case reflect.Uint16:
		return Uint16ToFloat32(value.(uint16))
	case reflect.Uint32:
		return Uint32ToFloat32(value.(uint32))
	case reflect.Uint64:
		return Uint64ToFloat32(value.(uint64))
	case reflect.String:
		return StringToFloat32(value.(string))
	case reflect.Bool:
		return BoolToFloat32(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func TryIntoInt[T buildin](value T) (int, error) {
	result, err := toInt(value)
	return result.(int), err
}

func toInt(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToInt(value.(float64))
	case reflect.Float32:
		return Float32ToInt(value.(float32))
	case reflect.Int:
		return value.(int), nil
	case reflect.Int8:
		return Int8ToInt(value.(int8))
	case reflect.Int16:
		return Int16ToInt(value.(int16))
	case reflect.Int32:
		return Int32ToInt(value.(int32))
	case reflect.Int64:
		return Int64ToInt(value.(int64))
	case reflect.Uint:
		return UintToInt(value.(uint))
	case reflect.Uint8:
		return Uint8ToInt(value.(uint8))
	case reflect.Uint16:
		return Uint16ToInt(value.(uint16))
	case reflect.Uint32:
		return Uint32ToInt(value.(uint32))
	case reflect.Uint64:
		return Uint64ToInt(value.(uint64))
	case reflect.String:
		return StringToInt(value.(string))
	case reflect.Bool:
		return BoolToInt(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func TryIntoInt8[T buildin](value T) (int8, error) {
	result, err := toInt8(value)
	return result.(int8), err
}

func toInt8(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToInt8(value.(float64))
	case reflect.Float32:
		return Float32ToInt8(value.(float32))
	case reflect.Int:
		return IntToInt8(value.(int))
	case reflect.Int8:
		return value.(int8), nil
	case reflect.Int16:
		return Int16ToInt8(value.(int16))
	case reflect.Int32:
		return Int32ToInt8(value.(int32))
	case reflect.Int64:
		return Int64ToInt8(value.(int64))
	case reflect.Uint:
		return UintToInt8(value.(uint))
	case reflect.Uint8:
		return Uint8ToInt8(value.(uint8))
	case reflect.Uint16:
		return Uint16ToInt8(value.(uint16))
	case reflect.Uint32:
		return Uint32ToInt8(value.(uint32))
	case reflect.Uint64:
		return Uint64ToInt8(value.(uint64))
	case reflect.String:
		return StringToInt8(value.(string))
	case reflect.Bool:
		return BoolToInt8(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func TryIntoInt16[T buildin](value T) (int16, error) {
	result, err := toInt16(value)
	return result.(int16), err
}

func toInt16(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToInt16(value.(float64))
	case reflect.Float32:
		return Float32ToInt16(value.(float32))
	case reflect.Int:
		return IntToInt16(value.(int))
	case reflect.Int8:
		return Int8ToInt16(value.(int8))
	case reflect.Int16:
		return value.(int16), nil
	case reflect.Int32:
		return Int32ToInt16(value.(int32))
	case reflect.Int64:
		return Int64ToInt16(value.(int64))
	case reflect.Uint:
		return UintToInt16(value.(uint))
	case reflect.Uint8:
		return Uint8ToInt16(value.(uint8))
	case reflect.Uint16:
		return Uint16ToInt16(value.(uint16))
	case reflect.Uint32:
		return Uint32ToInt16(value.(uint32))
	case reflect.Uint64:
		return Uint64ToInt16(value.(uint64))
	case reflect.String:
		return StringToInt16(value.(string))
	case reflect.Bool:
		return BoolToInt16(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func TryIntoInt32[T buildin](value T) (int32, error) {
	result, err := toInt32(value)
	return result.(int32), err
}

func toInt32(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToInt32(value.(float64))
	case reflect.Float32:
		return Float32ToInt32(value.(float32))
	case reflect.Int:
		return IntToInt32(value.(int))
	case reflect.Int8:
		return Int8ToInt32(value.(int8))
	case reflect.Int16:
		return Int16ToInt32(value.(int16))
	case reflect.Int32:
		return value.(int32), nil
	case reflect.Int64:
		return Int64ToInt32(value.(int64))
	case reflect.Uint:
		return UintToInt32(value.(uint))
	case reflect.Uint8:
		return Uint8ToInt32(value.(uint8))
	case reflect.Uint16:
		return Uint16ToInt32(value.(uint16))
	case reflect.Uint32:
		return Uint32ToInt32(value.(uint32))
	case reflect.Uint64:
		return Uint64ToInt32(value.(uint64))
	case reflect.String:
		return StringToInt32(value.(string))
	case reflect.Bool:
		return BoolToInt32(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func TryIntoInt64[T buildin](value T) (int64, error) {
	result, err := toInt64(value)
	return result.(int64), err
}

func toInt64(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToInt64(value.(float64))
	case reflect.Float32:
		return Float32ToInt64(value.(float32))
	case reflect.Int:
		return IntToInt64(value.(int))
	case reflect.Int8:
		return Int8ToInt64(value.(int8))
	case reflect.Int16:
		return Int16ToInt64(value.(int16))
	case reflect.Int32:
		return Int32ToInt64(value.(int32))
	case reflect.Int64:
		return value.(int64), nil
	case reflect.Uint:
		return UintToInt64(value.(uint))
	case reflect.Uint8:
		return Uint8ToInt64(value.(uint8))
	case reflect.Uint16:
		return Uint16ToInt64(value.(uint16))
	case reflect.Uint32:
		return Uint32ToInt64(value.(uint32))
	case reflect.Uint64:
		return Uint64ToInt64(value.(uint64))
	case reflect.String:
		return StringToInt64(value.(string))
	case reflect.Bool:
		return BoolToInt64(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func TryIntoUint[T buildin](value T) (uint, error) {
	result, err := toUint(value)
	return result.(uint), err
}

func toUint(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint(value.(float64))
	case reflect.Float32:
		return Float32ToUint(value.(float32))
	case reflect.Int:
		return IntToUint(value.(int))
	case reflect.Int8:
		return Int8ToUint(value.(int8))
	case reflect.Int16:
		return Int16ToUint(value.(int16))
	case reflect.Int32:
		return Int32ToUint(value.(int32))
	case reflect.Int64:
		return Int64ToUint(value.(int64))
	case reflect.Uint:
		return value.(uint), nil
	case reflect.Uint8:
		return Uint8ToUint(value.(uint8))
	case reflect.Uint16:
		return Uint16ToUint(value.(uint16))
	case reflect.Uint32:
		return Uint32ToUint(value.(uint32))
	case reflect.Uint64:
		return Uint64ToUint(value.(uint64))
	case reflect.String:
		return StringToUint(value.(string))
	case reflect.Bool:
		return BoolToUint(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func TryIntoUint8[T buildin](value T) (uint8, error) {
	result, err := toUint8(value)
	return result.(uint8), err
}

func toUint8(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint8(value.(float64))
	case reflect.Float32:
		return Float32ToUint8(value.(float32))
	case reflect.Int:
		return IntToUint8(value.(int))
	case reflect.Int8:
		return Int8ToUint8(value.(int8))
	case reflect.Int16:
		return Int16ToUint8(value.(int16))
	case reflect.Int32:
		return Int32ToUint8(value.(int32))
	case reflect.Int64:
		return Int64ToUint8(value.(int64))
	case reflect.Uint:
		return UintToUint8(value.(uint))
	case reflect.Uint8:
		return value.(uint8), nil
	case reflect.Uint16:
		return Uint16ToUint8(value.(uint16))
	case reflect.Uint32:
		return Uint32ToUint8(value.(uint32))
	case reflect.Uint64:
		return Uint64ToUint8(value.(uint64))
	case reflect.String:
		return StringToUint8(value.(string))
	case reflect.Bool:
		return BoolToUint8(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func TryIntoUint16[T buildin](value T) (uint16, error) {
	result, err := toUint16(value)
	return result.(uint16), err
}

func toUint16(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint16(value.(float64))
	case reflect.Float32:
		return Float32ToUint16(value.(float32))
	case reflect.Int:
		return IntToUint16(value.(int))
	case reflect.Int8:
		return Int8ToUint16(value.(int8))
	case reflect.Int16:
		return Int16ToUint16(value.(int16))
	case reflect.Int32:
		return Int32ToUint16(value.(int32))
	case reflect.Int64:
		return Int64ToUint16(value.(int64))
	case reflect.Uint:
		return UintToUint16(value.(uint))
	case reflect.Uint8:
		return Uint8ToUint16(value.(uint8))
	case reflect.Uint16:
		return value.(uint16), nil
	case reflect.Uint32:
		return Uint32ToUint16(value.(uint32))
	case reflect.Uint64:
		return Uint64ToUint16(value.(uint64))
	case reflect.String:
		return StringToUint16(value.(string))
	case reflect.Bool:
		return BoolToUint16(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func TryIntoUint32[T buildin](value T) (uint32, error) {
	result, err := toUint32(value)
	return result.(uint32), err
}

func toUint32(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint32(value.(float64))
	case reflect.Float32:
		return Float32ToUint32(value.(float32))
	case reflect.Int:
		return IntToUint32(value.(int))
	case reflect.Int8:
		return Int8ToUint32(value.(int8))
	case reflect.Int16:
		return Int16ToUint32(value.(int16))
	case reflect.Int32:
		return Int32ToUint32(value.(int32))
	case reflect.Int64:
		return Int64ToUint32(value.(int64))
	case reflect.Uint:
		return UintToUint32(value.(uint))
	case reflect.Uint8:
		return Uint8ToUint32(value.(uint8))
	case reflect.Uint16:
		return Uint16ToUint32(value.(uint16))
	case reflect.Uint32:
		return value.(uint32), nil
	case reflect.Uint64:
		return Uint64ToUint32(value.(uint64))
	case reflect.String:
		return StringToUint32(value.(string))
	case reflect.Bool:
		return BoolToUint32(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func TryIntoUint64[T buildin](value T) (uint64, error) {
	result, err := toUint64(value)
	return result.(uint64), err
}

func toUint64(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToUint64(value.(float64))
	case reflect.Float32:
		return Float32ToUint64(value.(float32))
	case reflect.Int:
		return IntToUint64(value.(int))
	case reflect.Int8:
		return Int8ToUint64(value.(int8))
	case reflect.Int16:
		return Int16ToUint64(value.(int16))
	case reflect.Int32:
		return Int32ToUint64(value.(int32))
	case reflect.Int64:
		return Int64ToUint64(value.(int64))
	case reflect.Uint:
		return UintToUint64(value.(uint))
	case reflect.Uint8:
		return Uint8ToUint64(value.(uint8))
	case reflect.Uint16:
		return Uint16ToUint64(value.(uint16))
	case reflect.Uint32:
		return Uint32ToUint64(value.(uint32))
	case reflect.Uint64:
		return value.(uint64), nil
	case reflect.String:
		return StringToUint64(value.(string))
	case reflect.Bool:
		return BoolToUint64(value.(bool))
	default:
		return 0, errors.New("unsupported type")
	}
}

func TryIntoString[T buildin](value T) (string, error) {
	result, err := toString(value)
	return result.(string), err
}

func toString(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToString(value.(float64))
	case reflect.Float32:
		return Float32ToString(value.(float32))
	case reflect.Int:
		return IntToString(value.(int))
	case reflect.Int8:
		return Int8ToString(value.(int8))
	case reflect.Int16:
		return Int16ToString(value.(int16))
	case reflect.Int32:
		return Int32ToString(value.(int32))
	case reflect.Int64:
		return Int64ToString(value.(int64))
	case reflect.Uint:
		return UintToString(value.(uint))
	case reflect.Uint8:
		return Uint8ToString(value.(uint8))
	case reflect.Uint16:
		return Uint16ToString(value.(uint16))
	case reflect.Uint32:
		return Uint32ToString(value.(uint32))
	case reflect.Uint64:
		return Uint64ToString(value.(uint64))
	case reflect.String:
		return value.(string), nil
	case reflect.Bool:
		return BoolToString(value.(bool))
	default:
		return "", errors.New("unsupported type")
	}
}

func TryIntoBool[T buildin](value T) (bool, error) {
	result, err := toBool(value)
	return result.(bool), err
}

func toBool(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Float64:
		return Float64ToBool(value.(float64))
	case reflect.Float32:
		return Float32ToBool(value.(float32))
	case reflect.Int:
		return IntToBool(value.(int))
	case reflect.Int8:
		return Int8ToBool(value.(int8))
	case reflect.Int16:
		return Int16ToBool(value.(int16))
	case reflect.Int32:
		return Int32ToBool(value.(int32))
	case reflect.Int64:
		return Int64ToBool(value.(int64))
	case reflect.Uint:
		return UintToBool(value.(uint))
	case reflect.Uint8:
		return Uint8ToBool(value.(uint8))
	case reflect.Uint16:
		return Uint16ToBool(value.(uint16))
	case reflect.Uint32:
		return Uint32ToBool(value.(uint32))
	case reflect.Uint64:
		return Uint64ToBool(value.(uint64))
	case reflect.String:
		return StringToBool(value.(string))
	case reflect.Bool:
		return value.(bool), nil
	default:
		return false, errors.New("unsupported type")
	}
}

func TryIntoTime[T ~string](value T) (time.Time, error) {
	result, err := toTime(value)
	return result.(time.Time), err
}

func toTime(value any) (any, error) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.String:
		return StringToTime(value.(string))
	default:
		return time.Time{}, errors.New("unsupported type")
	}
}
