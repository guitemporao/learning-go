package iteration

const repeatedCount = 5

func Repeat(character string) string {
	// new var
	var repeated string

	for i := 0; i < repeatedCount ; i++ {
		// repeated = repeated + character
		repeated += character
	}

	return repeated
}
