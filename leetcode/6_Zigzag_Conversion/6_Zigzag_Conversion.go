package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	step := 2*numRows - 2
	size := len(s)
	convertArr := make([]byte, size)
	pos := 0

	// first row
	for i := 0; i < len(s); i += step {
		convertArr[pos] = s[i]
		pos++
	}

	// middle row
	for row := 1; row < numRows - 1; row++ {
		zStep := 2*numRows - 2*row - 2
		for i := row; i < size; i += step {
			convertArr[pos] = s[i]
			pos++

			if i + zStep < size {
				convertArr[pos] = s[i + zStep]
				pos++
			}
		}
	}

	// last row
	for i := numRows - 1; i < len(s); i += step {
		convertArr[pos] = s[i]
		pos++
	}

	return string(convertArr)
}
