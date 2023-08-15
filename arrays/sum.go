package arrays

// func Sum(numbers [5]int) int {
// 	return 0
// }

// func Sum(numbers [5]int) int {
// 	sum := 0

// 	for i := 0; i < 5; i++ {
// 		sum += numbers[i]
// 	}

// 	return sum
// }

// func Sum(numbers [5]int) int {
// 	sum := 0

// 	for _, number := range numbers {
// 		sum += number
// 	}

// 	return sum
// }

// range lets you iterate over an array.
// On each iteration, range returns two values - the index and the value.
// We are choosing to ignore the index value by using _

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// func SumAll(numbersToSum ...[]int) []int {
// 	// sum := 0
// 	lengthOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengthOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}

// 	return sums
// }

// func SumAll(numbersToSum ...[]int) []int {
// 	return nil
// }

func SumAllTails(numbersToSum ...[]int) []int {
	// sum := 0
	var sums []int
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers)
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
