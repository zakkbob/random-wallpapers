package internal

type Number interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Clamp[T Number](n T, l T, u T) T {
	if n < l {
		return l
	}
	if n > u {
		return u
	}
	return n
}
