package supermath

import "math"

func PrimeSlice(maxint int) []int {
	var answer []int
	// loop through prime candidates
CandidateLoop:
	for i := 2; i <= maxint; i++ {
		// see if any existing primes divide it
		maxPossiblePrimeDivisor := int(math.Sqrt(float64(i)))
		for _, prime := range answer {
			if i%prime == 0 {
				// i is composite, go to the next i
				continue CandidateLoop
			} else if prime > maxPossiblePrimeDivisor {
				// we've already tested all primes less than or equal to sqrt(i), so i is prime
				answer = append(answer, i)
				continue CandidateLoop
			}
		}
		answer = append(answer, i)
	}
	return answer
}

func PrimeMap(maxint int) map[int]bool {
	answer := make(map[int]bool)
	for _, prime := range PrimeSlice(maxint) {
		answer[prime] = true
	}
	return answer
}
