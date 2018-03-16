package repeat

func repeat(character string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += character
	}
	return
}
