package supermath

import (
	"fmt"
	"sort"
)

func ExamplePrimeSlice() {
	fmt.Println(PrimeSlice(10))
	fmt.Println(PrimeSlice(20))
	//Output:
	//[2 3 5 7]
	//[2 3 5 7 11 13 17 19]
}

func sortKeys(m map[int]bool) []int {
	answer := []int{}
	for key := range m {
		answer = append(answer, key)
	}
	sort.Ints(answer)
	return answer
}

func ExamplePrimeMap() {
	fmt.Println(sortKeys(PrimeMap(10)))
	fmt.Println(sortKeys(PrimeMap(20)))
	//Output:
	//[2 3 5 7]
	//[2 3 5 7 11 13 17 19]
}

func ExamplePrimeFactors() {
	fmt.Println(PrimeFactors([]int{2,3,5,7}, 30))
	fmt.Println(PrimeFactors([]int{2,3,5,7}, 45))
	//Output:
	//[2 3 5]
	//[3 3 5]
}

