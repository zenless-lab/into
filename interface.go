package into

type ITryInto[T any] interface {
	TryInto() (T, error)
}

type IInto[T any] interface {
	Into() T
}

// TODO: Implement TryFrom and From Converters
type ITryFrom[T any] interface {
	TryFrom() (T, error)
}

type IFrom[T any] interface {
	From() T
}
