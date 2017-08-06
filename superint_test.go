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

func ExampleStringToSuperint() {
	fmt.Println(
		String(StringToSuperint("-3")),
		String(StringToSuperint("0")),
		String(StringToSuperint("452")),
		String(StringToSuperint("-1337")),
	)
	//Output: -3 0 452 -1337
}

func ExampleIntToSuperint() {
	fmt.Println(
		String(IntToSuperint(-3)),
		String(IntToSuperint(0)),
		String(IntToSuperint(452)),
		String(IntToSuperint(-1337)),
	)
	//Output: -3 0 452 -1337
}
