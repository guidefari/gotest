package iteration

func Repeat(s string) string {
	var repeated string
	repeatCount := 5

	for i := 0; i < repeatCount; i++ {
		repeated += s
	}

	return repeated
}
