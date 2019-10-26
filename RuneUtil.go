package utility

func IndexRunes(s []rune, subRunes []rune) int {
	ret := -1

	for i := range s {
		if i+len(subRunes) < len(s) {
			for j := range subRunes {
				if s[i+j] == subRunes[j] {
					ret = i + j

					if len(subRunes) > 1 {
						tempRet := IndexRunes(s[i+j+1:], subRunes[j+1:])
						if tempRet != 0 {
							ret = -1
						}
					}
				}
				break
			}

			if ret >= 0 {
				break
			}
		} else {
			break
		}
	}

	return ret
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
