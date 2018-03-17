package sum

// Sum takes an array of numbers and returns sum
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumAll takes varying number of slices and return new slice containing the sums of each slice
func SumAll(slices ...[]int) (sums []int) {
	for _, numbers := range slices {
		sums = append(sums, Sum(numbers))
	}
	return
}

// SumAllTails takes varying number of slices and return new slice containing the sum of tails of each slice
func SumAllTails(slices ...[]int) (sums []int) {
	for _, numbers := range slices {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
