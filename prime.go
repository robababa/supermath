package supermath

func primeSlice(maxint int) []int {
	var answer []int
	// loop through prime candidates
CandidateLoop:
	for i := 2; i <= maxint; i++ {
		// see if any existing primes divide it
		for _, prime := range answer {
			if i%prime == 0 {
				continue CandidateLoop
			}
		}
		answer = append(answer, i)
	}
	return answer
}

func primeMap(maxint int) map[int]bool {
	answer := make(map[int]bool)
	for _, prime := range primeSlice(maxint) {
		answer[prime] = true
	}
	return answer
}
