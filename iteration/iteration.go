package iteration

func Repeat(s string, howMany int) string {
	var repeated string

	for i := 0; i < howMany; i++ {
		repeated += s
	}

	return repeated
}
