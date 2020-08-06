package leetcode

func reverse(x int) int {
	var (
		result           int
		negative         bool
		isTenFigures     bool
		overFlowPossible bool
	)

	if x < 0 {
		negative = true
		x = -x
	}

	if x >= 1000000000 {
		isTenFigures = true
	}

	for i := 0; x > 0; i++ {
		v := x % 10
		if isTenFigures {
			switch i {
			case 0:
				if v > 2 {
					return 0
				}

				if v == 2 {
					overFlowPossible = true
				}
			case 1:
				if overFlowPossible && v > 1 {
					return 0
				}

				if v < 1 {
					overFlowPossible = false
				}
			case 2:
				if overFlowPossible && v > 4 {
					return 0
				}

				if v < 4 {
					overFlowPossible = false
				}
			case 3:
				if overFlowPossible && v > 7 {
					return 0
				}

				if v < 7 {
					overFlowPossible = false
				}
			case 4:
				if overFlowPossible && v > 4 {
					return 0
				}

				if v < 4 {
					overFlowPossible = false
				}
			case 5:
				if overFlowPossible && v > 8 {
					return 0
				}

				if v < 8 {
					overFlowPossible = false
				}
			case 6:
				if overFlowPossible && v > 3 {
					return 0
				}

				if v < 3 {
					overFlowPossible = false
				}
			case 7:
				if overFlowPossible && v > 6 {
					return 0
				}

				if v < 6 {
					overFlowPossible = false
				}
			case 8:
				if overFlowPossible && v > 4 {
					return 0
				}

				if v < 4 {
					overFlowPossible = false
				}
			case 9:
				if overFlowPossible && (!negative && v > 7 || v > 8) {
					return 0
				}
			}
		}

		result = result*10 + v
		x /= 10
	}

	if negative {
		return -result
	}

	return result
}
