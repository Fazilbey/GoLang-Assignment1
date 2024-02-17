package main

import "fmt"

func SortSlice(slice []int) {
	n := len(slice)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

}

func IncrementOdd(slice []int) {
	n := len(slice)
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			slice[i] += 1
		}
	}
}

func PrintSlice(slice []int) {
	fmt.Print(slice)
}

func ReverseSlice(input []int) []int {
	length := len(input)
	reversed := make([]int, length)

	for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
		reversed[i] = input[j]
	}

	return reversed
}

func appendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(input []int) {
		dst(input)

		for _, fn := range src {
			fn(input)
		}
	}
}

func main() {
	numbers := []int{12, 214, 43, 11, 0, -1}
	//SortSlice(numbers)
	//fmt.Println(numbers)
	//
	//IncrementOdd(numbers)
	//fmt.Println(numbers)
	//
	//PrintSlice(numbers)
	//
	//fmt.Print(ReverseSlice(numbers))

	combinedFunc := appendFunc(
		func(slice []int) {},
		SortSlice,
		IncrementOdd,
		PrintSlice,
	)

	combinedFunc(numbers)
}
