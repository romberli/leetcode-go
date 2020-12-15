package leetcode

const (
	MaxInt = 2147483647
	MinInt = -2147483648
)

func myAtoi(str string) int {
	var result, v int

	startFlag := true
	firstFlag := true
	ZeroFlag := true
	negative := false
	overFlow := false
	overFlowPossibleHigh := false
	overFlowPossibleLow := false
	n := 0

loop:
	for i := 0; i < len(str); i++ {
		c := str[i]
		s := string(c)

		if s == " " && startFlag {
			continue
		}

		startFlag = false

		if firstFlag {
			if s == "+" {
				firstFlag = false
				continue
			}

			if s == "-" {
				negative = true
				firstFlag = false
				continue
			}

			firstFlag = false
		}

		switch s {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			if ZeroFlag && s == "0" {
				continue
			}

			ZeroFlag = false
			v = int(c - '0')

			switch n {
			case 0:
				if v > 2 {
					overFlowPossibleHigh = true
				} else if v == 2 {
					overFlowPossibleLow = true
				} else {
					overFlowPossibleLow = false
				}
			case 1:
				if !overFlowPossibleHigh && overFlowPossibleLow && v < 1 {
					overFlowPossibleLow = false
				}
				if !overFlowPossibleHigh && overFlowPossibleLow && v > 1 {
					overFlowPossibleHigh = true
				}
			case 2:
				if !overFlowPossibleHigh && overFlowPossibleLow && v < 4 {
					overFlowPossibleLow = false
				}
				if !overFlowPossibleHigh && overFlowPossibleLow && v > 4 {
					overFlowPossibleHigh = true
				}
			case 3:
				if !overFlowPossibleHigh && overFlowPossibleLow && v < 7 {
					overFlowPossibleLow = false
				}
				if !overFlowPossibleHigh && overFlowPossibleLow && v > 7 {
					overFlowPossibleHigh = true
				}
			case 4:
				if !overFlowPossibleHigh && overFlowPossibleLow && v < 4 {
					overFlowPossibleLow = false
				}
				if !overFlowPossibleHigh && overFlowPossibleLow && v > 4 {
					overFlowPossibleHigh = true
				}
			case 5:
				if !overFlowPossibleHigh && overFlowPossibleLow && v < 8 {
					overFlowPossibleLow = false
				}
				if !overFlowPossibleHigh && overFlowPossibleLow && v > 8 {
					overFlowPossibleHigh = true
				}
			case 6:
				if !overFlowPossibleHigh && overFlowPossibleLow && v < 3 {
					overFlowPossibleLow = false
				}
				if !overFlowPossibleHigh && overFlowPossibleLow && v > 3 {
					overFlowPossibleHigh = true
				}
			case 7:
				if !overFlowPossibleHigh && overFlowPossibleLow && v < 6 {
					overFlowPossibleLow = false
				}
				if !overFlowPossibleHigh && overFlowPossibleLow && v > 6 {
					overFlowPossibleHigh = true
				}
			case 8:
				if !overFlowPossibleHigh && overFlowPossibleLow && v < 4 {
					overFlowPossibleLow = false
				}
				if !overFlowPossibleHigh && overFlowPossibleLow && v > 4 {
					overFlowPossibleHigh = true
				}
			case 9:
				if overFlowPossibleHigh || (overFlowPossibleLow && (!negative && v > 7 || v > 8)) {
					overFlow = true
					break loop
				}
			default:
				overFlow = true
				break loop
			}

			result = result*10 + v
			n += 1
		default:
			break loop
		}
	}

	if overFlow {
		if negative {
			return MinInt
		}

		return MaxInt
	}

	if negative {
		result *= -1
	}

	return result
}
