package supermath

import "fmt"

func ExampleString() {
	fmt.Println(
		String(Superint{[]int{3}, true}),
		String(Superint{[]int{0}, false}),
		String(Superint{[]int{2, 5, 4}, false}),
		String(Superint{[]int{7, 3, 3, 1}, true}),
	)
	//Output: -3 0 452 -1337
}
