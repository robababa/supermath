package supermath

import "strconv"

// A Superint is a representation of a mathematical integer with an arbitrary
// size, i.e. an unlimited number of digits
type Superint struct {
	digits   []int
	negative bool
}

func (si Superint) String() string {
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

func SimplifySuperint(si Superint) Superint {
	var answer Superint
	answer.negative = si.negative
	carryover := 0
	for _, digit := range si.digits {
		answer.digits = append(answer.digits, (digit+carryover)%10)
		carryover = (digit + carryover) / 10
	}
	for ; carryover > 0; carryover /= 10 {
		answer.digits = append(answer.digits, carryover%10)
	}
	return answer
}

// for now, this only adds positive Superints
func AddSuperints(s1 Superint, s2 Superint) Superint {
	var answer Superint
	sizeDiff := len(s1.digits) - len(s2.digits)
	if sizeDiff > 0 {
		for i := 1; i <= sizeDiff; i++ {
			s2.digits = append(s2.digits, 0)
		}
	} else if sizeDiff < 0 {
		for i := -1; i >= sizeDiff; i-- {
			s1.digits = append(s1.digits, 0)
		}
	}
	for index := range s1.digits {
		answer.digits = append(answer.digits, s1.digits[index]+s2.digits[index])
	}
	return SimplifySuperint(answer)
}

func zeroSuperint() Superint {
	return Superint{digits: []int{0}, negative: false}
}

func copySuperint(s Superint) Superint {
	return Superint{digits: append([]int{}, s.digits...), negative: s.negative}
}

func SuperintMultiple(s Superint, n int) Superint {
	if n == 0 {
		return zeroSuperint()
	} else if n == 1 {
		return s
	}
	sCopy := copySuperint(s)
	answer := zeroSuperint()
	for i := 1; i <= n; i++ {
		answer = AddSuperints(answer, sCopy)
	}
	return answer
}

func DigitSum(s Superint) int {
	answer := 0
	for _, num := range s.digits {
		answer += num
	}
	return answer
}
