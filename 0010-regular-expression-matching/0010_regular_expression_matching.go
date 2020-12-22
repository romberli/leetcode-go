package leetcode

func getNextCh(s string, pos int) uint8 {
	length := len(s)
	if pos+1 < length {
		return s[pos+1]
	}

	return 0
}

func mayNull(s string) bool {
	lenS := len(s)
	for i := 0; i < lenS; i++ {
		if s[i] == '*' {
			continue
		}

		if s[i] == '.' {
			if (i+1 < lenS && s[i+1] != '*') || i == lenS-1 {
				return false
			}

			i++
			continue
		}

		if s[i] != '.' && (i+1 < lenS && s[i+1] != '*') || (i == lenS-1 && s[i] != '*') {
			return false
		}
	}

	return true
}

func match(s, p string, result *bool) {
	var (
		prevP, currentP, nextP uint8
		i, j                   int
	)

	if s == p {
		*result = true
		return
	}

	if s == "" && p == "" || p == ".*" {
		*result = true
		return
	}

	if s != "" && p == "" {
		return
	}

	// 0 1 2 3
	// a b c d
	lenS := len(s)
	lenP := len(p)
	for {
		if i == lenS {
			if j == lenP {
				*result = true
				return
			}

			if mayNull(p[j:]) {
				*result = true
				return
			}

			return
		}

		if j >= lenP {
			return
		}

		currentP = p[j]
		nextP = getNextCh(p, j)
		switch currentP {
		case '.':
			if nextP == '*' {
				prevP = currentP
				j++
				continue
			}

			i++
			j++
			continue
		case '*':
			if prevP == '*' || prevP == 0 {
				// previous char is not valid
				*result = false
				return
			}

			if prevP == '.' {
				if j == lenP-1 {
					// .* is the end of the string p
					*result = true
					return
				}

				match(s[i:], p[j+1:], result)
				if *result == true {
					return
				}

				i++
				continue
			}

			// previous char is a normal char
			if prevP != s[i] {
				j++
				continue
			}

			match(s[i:], p[j+1:], result)
			if *result == true {
				return
			}

			if j == lenP-1 && i == lenS-1 {
				j++
			}

			i++
			continue
		default:
			if nextP == '*' {
				prevP = currentP
				j++
				continue
			}

			if currentP != s[i] {
				return
			}

			i++
			j++
			continue
		}
	}
}

func isMatch(s string, p string) bool {
	var result bool

	if s == "" && p == "" || p == ".*" {
		return true
	}

	if s != "" && p == "" {
		return false
	}

	match(s, p, &result)
	return result
}
