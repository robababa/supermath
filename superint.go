package supermath

import "strconv"

// A Superint is a representation of a mathematical integer with an arbitrary
// size, i.e. an unlimited number of digits
type Superint struct {
	digits   []int
	negative bool
}

func String(si Superint) string {
	var answer string
	for _, n := range si.digits {
		answer = string(strconv.Itoa(n)) + answer
	}
	if si.negative {
		answer = "-" + answer
	}
	return answer
}

// Take a string and return its corresponding Superint
func StringToSuperint(s string) Superint {
	var si Superint
	//TODO add validation later
	if s[0] == '-' {
		si.negative = true
		s = s[1:]
	}
	// "12345" would be stored as [5, 4, 3, 2, 1]
	for i := len(s) - 1; i >= 0; i-- {
		nextDigit, _ := strconv.Atoi(string(s[i]))
		si.digits = append(si.digits, nextDigit)
	}
	return si
}

// Take an int and return its corresponding Superint
func IntToSuperint(n int) Superint {
	return StringToSuperint(strconv.Itoa(n))
}
