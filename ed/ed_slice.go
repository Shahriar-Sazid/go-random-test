package ed

var memoArray [100][100]float32

func progressiveED(s, t string, progressSoFar, steps int) float32 {
	var distance float32
	runeS, runeT := []rune(s), []rune(t)
	for i := progressSoFar; i < progressSoFar+steps; i++ {
		distance = edInternal(runeS[:min(len(runeS), i+1)], runeT[:min(len(runeT), i+1)])
	}
	return distance
}

func edInternal(s, t []rune) float32 {
	if len(s) >= len(t) {
		for i := 1; i < len(t); i++ {
			if s[len(s)-1] == t[i-1] {
				memoArray[i][len(s)] = memoArray[i-1][len(s)-1]
			} else {
				memoArray[i][len(s)] = 1 + min(
					memoArray[i-1][len(s)-1],
					memoArray[i-1][len(s)],
					memoArray[i][len(s)-1])
			}
		}
	}

	if len(t) >= len(s) {
		for j := 1; j < len(s); j++ {
			if t[len(t)-1] == s[j-1] {
				memoArray[len(t)][j] = memoArray[len(t)-1][j-1]
			} else {
				memoArray[len(t)][j] = 1 + min(
					memoArray[len(t)-1][j-1],
					memoArray[len(t)-1][j],
					memoArray[len(t)][j-1])
			}
		}
	}

	if s[len(s)-1] == t[len(t)-1] {
		memoArray[len(t)][len(s)] = memoArray[len(t)-1][len(s)-1]
	} else {
		memoArray[len(t)][len(s)] = 1 + min(
			memoArray[len(t)-1][len(s)-1],
			memoArray[len(t)-1][len(s)],
			memoArray[len(t)][len(s)-1])
	}

	return memoArray[len(t)][len(s)]
}
