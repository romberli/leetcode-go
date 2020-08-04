package leetcode

func relativePoint(relativeNum int, numRows int) (int, int) {
	if relativeNum < numRows {
		return relativeNum, 0
	}

	r := numRows - 1
	c := 0

	rightPartLen := relativeNum - numRows + 1
	for i := 0; i < rightPartLen; i++ {
		r--
		c++
	}

	return r, c
}

// this solution works but cost too much deadweight loss, especially for memory for resultSlice
// func convert(s string, numRows int) string {
// 	var (
// 		result string
// 	)
//
// 	length := len(s)
//
// 	if length <= numRows {
// 		return s
// 	}
//
// 	resultSlice := make([][]string, numRows)
//
// 	numCharsPerUnit := numRows
// 	if numRows > 1 {
// 		numCharsPerUnit = numRows * 2 - 2
// 	}
// 	columnsPerUnit := numCharsPerUnit - numRows + 1
//
// 	numUnitsComplete := length / numCharsPerUnit
// 	lastUnitCharsLength := length % numCharsPerUnit
// 	isLargerThanNumRows := lastUnitCharsLength > numRows
//
// 	numColumns := columnsPerUnit * numUnitsComplete
// 	if lastUnitCharsLength > 0 {
// 		numColumns += 1
// 	}
// 	if isLargerThanNumRows {
// 		numColumns += lastUnitCharsLength - numRows
// 	}
//
// 	for i := range resultSlice {
// 		resultSlice[i] = make([]string, numColumns)
// 	}
//
// 	for i, c := range s {
// 		p := i / numCharsPerUnit
// 		q := i % numCharsPerUnit
//
// 		row, relativeColumn := relativePoint(q, numRows)
// 		column := p * columnsPerUnit + relativeColumn
// 		resultSlice[row][column] = string(c)
// 	}
//
// 	for i := 0; i < numRows; i++ {
// 		for j := 0; j < numColumns; j++ {
// 			c := resultSlice[i][j]
// 			if c != "" {
// 				result += c
// 			}
// 		}
// 	}
//
// 	return result
// }

func convert(s string, numRows int) string {
	var (
		result []byte
		n      int
	)

	length := len(s)
	numColumns := make([]int, numRows)

	if length <= numRows || numRows == 1 {
		return s
	}

	numCharsPerUnit := numRows
	if numRows > 1 {
		numCharsPerUnit = numRows*2 - 2
	}

	numUnitsComplete := length / numCharsPerUnit
	lastUnitCharsLength := length % numCharsPerUnit
	lastUnitRightPartLength := lastUnitCharsLength - numRows

	lastUnitRightPartMinRow := 0
	if lastUnitRightPartLength > 0 {
		lastUnitRightPartMinRow = numRows - lastUnitRightPartLength - 1
	}

	for i := 0; i < numRows; i++ {
		numColumns[i] = numUnitsComplete
		if i > 0 && i < numRows-1 {
			numColumns[i] += numUnitsComplete
		}

		if lastUnitCharsLength > 0 {
			if lastUnitCharsLength <= numRows {
				if i < lastUnitCharsLength {
					numColumns[i] += 1
				}
			} else {
				numColumns[i] += 1
				if i >= lastUnitRightPartMinRow && i < numRows-1 {
					numColumns[i] += 1
				}
			}
		}
	}

	for i := 0; i < numRows; i++ {
		columns := numColumns[i]

		for j := 0; j < columns; j++ {
			p := j
			if i > 0 && i < numRows-1 {
				p = j / 2
			}

			q := j % 2
			if q == 0 {
				n = numCharsPerUnit*p + i
			} else {
				if i > 0 && i < numRows-1 {
					n = numCharsPerUnit*p + numCharsPerUnit - i
				} else {
					n = numCharsPerUnit*p + i
				}
			}

			result = append(result, s[n])
		}
	}

	return string(result)
}

// func convert(s string, numRows int) string {
// 	var (
// 		result []byte
// 	)
//
// 	length := len(s)
// 	numColumns := make([]int, numRows)
//
// 	if length <= numRows || numRows == 1 {
// 		return s
// 	}
//
// 	numCharsPerUnit := numRows
// 	if numRows > 1 {
// 		numCharsPerUnit = numRows * 2 - 2
// 	}
//
// 	numUnitsComplete := length / numCharsPerUnit
// 	lastUnitCharsLength := length % numCharsPerUnit
// 	lastUnitRightPartLength := lastUnitCharsLength - numRows
//
// 	lastUnitRightPartMinRow := 0
// 	if lastUnitRightPartLength > 0 {
// 		lastUnitRightPartMinRow = numRows - lastUnitRightPartLength - 1
// 	}
//
// 	for i := 0; i < numRows; i++ {
// 		numColumns[i] = numUnitsComplete
// 		if i > 0 && i < numRows - 1 {
// 			numColumns[i] += numUnitsComplete
// 		}
//
// 		if lastUnitCharsLength > 0 {
// 			if lastUnitCharsLength <= numRows {
// 				if i < lastUnitCharsLength {
// 					numColumns[i] += 1
// 				}
// 			} else {
// 				numColumns[i] += 1
// 				if i >= lastUnitRightPartMinRow && i < numRows - 1 {
// 					numColumns[i] += 1
// 				}
// 			}
// 		}
// 	}
//
// 	row := 0
// 	flag := 1
// 	opposite := 0
// 	pos := 0
// 	result = []byte(s)
// 	for _, c := range s {
// 		result[pos] = byte(c)
//
// 		pos = 0
// 		row += flag
// 		for j := 0; j < row; j++ {
// 			pos += numColumns[j]
// 		}
//
// 		if row == 0 || row == numRows - 1 {
// 			pos += opposite / 2
// 			if row == 0 {
// 				pos += 1
// 			}
//
// 			flag = -flag
// 			opposite++
// 		} else {
// 			pos += opposite
// 		}
// 	}
//
// 	return string(result)
// }
