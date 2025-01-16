package transform

func Transform(set []complex128, f func(complex128) complex128) []complex128 {
	result := make([]complex128, len(set))
	for i, v := range set {
		result[i] = f(v)
	}
	return result
}

func TransformInPlace(set []complex128, f func(complex128) complex128) {
	for i, v := range set {
		set[i] = f(v)
	}
}