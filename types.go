package into

// Bool represents a bool value.
//
// Include:
//   - bool: the bool value.
type Bool interface {
	~bool
}

// Byte represents a byte value.
//
// Include:
//   - byte: the byte value.
type Byte interface {
	~byte
}

// Complex represents a complex64 or complex128 value.
//
// Include:
//   - complex64 or complex128: the complex64 or complex128 value.
type Complex interface {
	~complex64 | ~complex128
}

// Float represents a float32 or float64 value.
//
// Include:
//   - float32 or float64: the float32 or float64 value.
type Float interface {
	~float64 | ~float32
}

// Int represents an integer value.
//
// Include:
//   - int, int8, int16, int32, or int64: the integer value.
type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Rune represents a rune value.
//
// Include:
//   - rune: the rune value.
type Rune interface {
	~rune
}

// String represents a string value.
//
// Include:
//   - string: the string value.
type String interface {
	~string
}

// Uint represents an unsigned integer value.
//
// Includes:
//   - uint, uint8, uint16, uint32, or uint64: the unsigned integer value.
type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// BuildIn represents a built-in type.
//
// Include:
//   - bool, byte, complex64, complex128, float32, float64, int, int8, int16, int32, int64, rune, string, uint, uint8, uint16, uint32, or uint64: the built-in type value.
type BuildIn interface {
	Bool | Byte | Complex | Float | Int | Rune | String | Uint
}

// convertable represents a type that can be converted to another type.
type convertable interface {
	Bool | Float | Int | String | Uint
}
