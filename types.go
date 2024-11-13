package into

type Bool interface {
	~bool
}

type Byte interface {
	~byte
}

type Complex interface {
	~complex64 | ~complex128
}

type Float interface {
	~float64 | ~float32
}

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Rune interface {
	~rune
}

type String interface {
	~string
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type BuildIn interface {
	Bool | Byte | Complex | Float | Int | Rune | String | Uint
}

type convertable interface {
	Bool | Float | Int | String | Uint
}
