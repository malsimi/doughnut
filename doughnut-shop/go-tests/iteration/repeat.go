package iteration

func Repeat(phrase string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += phrase
	}

	return repeated
}
