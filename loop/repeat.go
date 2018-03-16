package repeat

// Repeat takes a character and returns a string in which the character is repeated count times
func Repeat(character string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += character
	}
	return
}
