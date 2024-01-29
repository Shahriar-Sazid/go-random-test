package ed

var memoArray [100][100]float32

func progressiveED(s, t string) float32 {
	maxLen := max(len(s), len(t))
	var distance float32
	for i := 0; i < maxLen; i++ {
		distance = edInternal(s[:min(len(s), i+1)], t[:min(len(t), i+1)])
	}
	return distance
}

func edInternal(s, t string) float32 {
	for i := 1; i < len(t); i++ {
		memoArray[i][len(s)] = min(
			memoArray[i-1][len(s)-1],
			memoArray[i-1][len(s)],
			memoArray[i][len(s)-1])
		if s[len(s)-1] != t[i-1] {
			memoArray[i][len(s)] = memoArray[i][len(s)] + 1
		}
	}

	for j := 1; j < len(s); j++ {
		memoArray[len(t)][j] = min(
			memoArray[len(t)-1][j-1],
			memoArray[len(t)-1][j],
			memoArray[len(t)][j-1])
		if t[len(t)-1] != s[j-1] {
			memoArray[len(t)][j] = memoArray[len(t)][j] + 1
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
