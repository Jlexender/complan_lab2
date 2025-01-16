package mapping

func GenerateRectangleSet(start, end complex128, step float64) []complex128 {
	var result []complex128
	for i := real(start); i < real(end); i += step {
		for j := imag(start); j < imag(end); j += step {
			result = append(result, complex(i, j))
		}
	}
	return result
}