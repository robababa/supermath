package supermath

import "fmt"

func ExampleString() {
	fmt.Println(
		Superint{[]int{3}, true},
		Superint{[]int{0}, false},
		Superint{[]int{2, 5, 4}, false},
		Superint{[]int{7, 3, 3, 1}, true},
	)
	//Output: -3 0 452 -1337
}

func ExampleStringToSuperint() {
	fmt.Println(
		StringToSuperint("-3"),
		StringToSuperint("0"),
		StringToSuperint("452"),
		StringToSuperint("-1337"),
	)
	//Output: -3 0 452 -1337
}

func ExampleIntToSuperint() {
	fmt.Println(
		IntToSuperint(-3),
		IntToSuperint(0),
		IntToSuperint(452),
		IntToSuperint(-1337),
	)
	//Output: -3 0 452 -1337
}

func ExampleSimplifySuperint() {
	fmt.Println(
		SimplifySuperint(Superint{[]int{1024}, false}),
		SimplifySuperint(Superint{[]int{1007, 33}, true}),
		SimplifySuperint(Superint{[]int{0}, false}),
		SimplifySuperint(Superint{[]int{3, 2, 1}, true}),
	)
	// Output: 1024 -1337 0 -123
}

//func ExampleAddSuperints() {
//	fmt.Println(
//
//	)
//}
