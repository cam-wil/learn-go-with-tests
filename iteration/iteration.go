package iteration

func Repeat(character string, iters int) string {
	var repeated string
	if iters <= 0 {
		iters = 5
	}

	for i := 0; i < iters; i++ {
		repeated += character
	}
	return repeated
}
