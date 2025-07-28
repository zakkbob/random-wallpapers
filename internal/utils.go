package internal

func Clamp(n int, l int, u int) int {
	if n < l {
		return l
	}
	if n > u {
		return u
	}
	return n
}

