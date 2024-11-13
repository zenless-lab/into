package into

import (
	"strconv"
)

func BoolToFloat64(value bool) (float64, error) {
	if value {
		return 1.0, nil
	}
	return 0.0, nil
}

func BoolToFloat32(value bool) (float32, error) {
	if value {
		return 1.0, nil
	}
	return 0.0, nil
}

func BoolToInt(value bool) (int, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func BoolToInt8(value bool) (int8, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func BoolToInt16(value bool) (int16, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func BoolToInt32(value bool) (int32, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func BoolToInt64(value bool) (int64, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func BoolToUint(value bool) (uint, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func BoolToUint8(value bool) (uint8, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func BoolToUint16(value bool) (uint16, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func BoolToUint32(value bool) (uint32, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func BoolToUint64(value bool) (uint64, error) {
	if value {
		return 1, nil
	}
	return 0, nil
}

func BoolToString(value bool) (string, error) {
	return strconv.FormatBool(value), nil
}

func BoolToBool(value bool) (bool, error) {
	return value, nil
}
