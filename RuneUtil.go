package utility

func IndexRunes(s []rune, subRunes []rune) int {
	ret, _ := subIndexRunes(s, subRunes, 0)
	return ret
}

func subIndexRunes(s []rune, subRunes []rune, nest int) (ret int, maxNest int) {
	ret = -1
	maxNest = nest

	for i := 0; i < len(s); {
		if i+len(subRunes) < len(s) {
			if s[i] == subRunes[0] {
				ret = i

				if len(subRunes) > 1 {
					tempRet, tempMaxNest := subIndexRunes(s[i+1:], subRunes[1:], nest+1)
					if tempRet != 0 {
						ret = -1
						maxNest = tempMaxNest
					}
				}

				if nest > 0 {
					break
				} else {
					if ret >= 0 {
						break
					} else {
						i += maxNest
					}
				}
			} else {
				if nest > 0 {
					break
				} else {
					i++
				}
			}
		} else {
			break
		}
	}

	return ret, maxNest
}

func LastIndexRunes(s []rune, subRunes []rune) int {
	ret := lastIndexRunes(s, subRunes)

	if ret+1-len(subRunes) >= 0 {
		ret = ret + 1 - len(subRunes)
	}

	return ret
}

func lastIndexRunes(s []rune, subRunes []rune) int {
	ret := -1

	for i := len(s) - 1; i >= (len(subRunes) - 1); i-- {
		for j := len(subRunes) - 1; j >= 0; j-- {
			if s[i-(len(subRunes)-1-j)] == subRunes[j] {
				ret = i - (len(subRunes) - 1 - j)

				if len(subRunes) > 1 {
					tempRet := lastIndexRunes(s[0:i-(len(subRunes)-1-j)], subRunes[0:j])
					if tempRet != (i - (len(subRunes) - 1 - j) - 1) {
						ret = -1
					}
				}
			}
			break
		}

		if ret >= 0 {
			break
		}
	}

	return ret
}

func HasPrefix(s, prefix []rune) bool {
	return IndexRunes(s, prefix) == 0
}

func HasSuffix(s, suffix []rune) bool {
	return LastIndexRunes(s, suffix) == (len(s) - len(suffix))
}
