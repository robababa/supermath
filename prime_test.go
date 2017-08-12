package supermath

import (
	"fmt"
	"sort"
)

func ExamplePrimeSlice() {
	fmt.Println(primeSlice(10))
	fmt.Println(primeSlice(20))
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
	fmt.Println(sortKeys(primeMap(10)))
	fmt.Println(sortKeys(primeMap(20)))
	//Output:
	//[2 3 5 7]
	//[2 3 5 7 11 13 17 19]
}

