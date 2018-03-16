package sum

// Sum takes an array of numbers and returns sum
func Sum(numbers [5]int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
