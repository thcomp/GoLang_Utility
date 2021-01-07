package utility

import (
	"fmt"
	"strconv"
	"strings"
)

func SnakeToCamelCase(text string) string {
	var ret = text

	splitTextArray := strings.Split(text, `_`)
	if len(splitTextArray) > 0 {
		var builder StringBuilder

		for _, splitText := range splitTextArray {
			byteArray := strings.Split(splitText, ``)
			for i, char := range byteArray {
				if i == 0 {
					char = strings.ToUpper(char)
				}
				builder.Append(char)
			}
		}

		ret = builder.String()
	}

	return ret
}

func CamelToSnakeCase(text string) string {
	var ret = text

	byteArray := []byte(text)
	if len(byteArray) > 0 {
		var builder StringBuilder

		for i, _ := range byteArray {
			if byteArray[i] >= 65 && byteArray[i] <= 90 {
				if i > 0 {
					builder.Append(`_`)
				}
				byteArray[i] += 32
			}
			builder.AppendByte(byteArray[i])
		}

		ret = builder.String()
	}

	return ret
}

func ToLowerFirst(text string) string {
	var ret = text

	byteArray := []byte(text)
	if len(byteArray) > 0 {
		if byteArray[0] >= 65 && byteArray[0] <= 90 {
			byteArray[0] += 32
		}
		ret = string(byteArray)
	}

	return ret
}

func ToUpperFirst(text string) string {
	var ret = text

	byteArray := []byte(text)
	if len(byteArray) > 0 {
		if byteArray[0] >= 97 && byteArray[0] <= 122 {
			byteArray[0] -= 32
		}
		ret = string(byteArray)
	}

	return ret
}

const noneSpecificChar = 0
const startSepText = 1
const endSepText = 2
const startInvalidSpaceStarter = 4
const endInvalidSpaceStarter = 8

func Split(originalText string, sepTextArray []string, invalidSpaceStarters []rune, caseSensitive bool) [][]rune {
	var ret [][]rune
	var originalTextRunes []rune = []rune(originalText)
	var caseFixedOriginalTextRunes []rune
	var caseFixedSepTextArray []string
	var caseFixedInvalidSpaceStarters []rune
	if caseSensitive {
		caseFixedOriginalTextRunes = []rune(originalText)
		caseFixedSepTextArray = sepTextArray
		caseFixedInvalidSpaceStarters = invalidSpaceStarters
	} else {
		caseFixedOriginalTextRunes = []rune(strings.ToLower(originalText))
		for _, sepText := range sepTextArray {
			caseFixedSepTextArray = append(caseFixedSepTextArray, strings.ToLower(sepText))
		}
		caseFixedInvalidSpaceStarters = []rune(strings.ToLower(string(invalidSpaceStarters)))
	}
	var findPositionArray []int = make([]int, len(caseFixedOriginalTextRunes))
	var openedInvalidSpace []bool = make([]bool, len(invalidSpaceStarters))
	var openedInvalidSpaceCount = 0

	for i := 0; i < len(caseFixedOriginalTextRunes); {
		var increase = 1

		for j, caseFixedInvalidSpaceStarter := range caseFixedInvalidSpaceStarters {
			if caseFixedInvalidSpaceStarter == caseFixedOriginalTextRunes[i] {
				if openedInvalidSpace[j] {
					openedInvalidSpace[j] = false
					findPositionArray[i] = endInvalidSpaceStarter
					openedInvalidSpaceCount--
				} else {
					openedInvalidSpace[j] = true
					findPositionArray[i] = startInvalidSpaceStarter
					openedInvalidSpaceCount++
				}
				break
			}
		}

		if openedInvalidSpaceCount == 0 {
			for _, caseFixedSepText := range caseFixedSepTextArray {
				if strings.HasPrefix(string(caseFixedOriginalTextRunes[i:]), caseFixedSepText) {
					findPositionArray[i] |= startSepText

					caseFixedSepTextLen := len([]rune(caseFixedSepText))
					findPositionArray[i+caseFixedSepTextLen-1] |= endSepText
					increase = caseFixedSepTextLen
					break
				}
			}
		}

		i += increase
	}

	var startIndex = 0
	for i, specify := range findPositionArray {
		if (specify & startSepText) == startSepText {
			tempAppendRunes := originalTextRunes[startIndex:i]
			if len(tempAppendRunes) > 0 {
				ret = append(ret, tempAppendRunes)
			}
			startIndex = i
		}
		if (specify & endSepText) == endSepText {
			tempAppendRunes := originalTextRunes[startIndex : i+1]
			if len(tempAppendRunes) > 0 {
				ret = append(ret, tempAppendRunes)
			}
			startIndex = i + 1
		}
		if (specify & startInvalidSpaceStarter) == startInvalidSpaceStarter {
		}
		if (specify & endInvalidSpaceStarter) == endInvalidSpaceStarter {
		}
	}
	ret = append(ret, originalTextRunes[startIndex:])

	return ret
}

func ParseNumber(str string) (float64, error) {
	ret := float64(0)
	retErr := error(nil)

	str = strings.Trim(str, " \t\n")
	str = strings.ToLower(str)
	if strings.HasPrefix(str, "-") {
		// less than 0
		unsignStr := strings.Trim(str, " -\t")
		if tempRet, tempErr := ParseNumber(unsignStr); tempErr == nil {
			ret = -tempRet
		} else {
			retErr = tempErr
		}
	} else if strings.HasPrefix(str, "0x") {
		// hex
		if tempDec, err := strconv.ParseUint(str[2:], 16, 64); err == nil {
			ret = float64(tempDec)
		} else {
			retErr = err
		}
	} else if (strings.HasPrefix(str, "b\"") && strings.HasSuffix(str, "\"")) || (strings.HasPrefix(str, "b'") && strings.HasSuffix(str, "'")) {
		// binary
		str = str[2 : len(str)-1]
		if tempBin, err := strconv.ParseUint(str, 2, 64); err == nil {
			ret = float64(tempBin)
		} else {
			retErr = err
		}
	} else {
		// decimal
		ret, retErr = strconv.ParseFloat(str, 64)
	}

	return ret, retErr
}

type RangeValue struct {
	minValue *float64
	maxValue *float64

	allowMinEqual bool
	allowMaxEqual bool
}

func (rValue *RangeValue) In(targetValue float64) bool {
	allowedMin := false
	allowedMax := false

	if rValue.minValue != nil {
		if rValue.allowMinEqual {
			if targetValue >= *(rValue.minValue) {
				allowedMin = true
			}
		} else {
			if targetValue > *(rValue.minValue) {
				allowedMin = true
			}
		}
	} else {
		allowedMin = true
	}

	if rValue.maxValue != nil {
		if rValue.allowMaxEqual {
			if targetValue <= *(rValue.maxValue) {
				allowedMax = true
			}
		} else {
			if targetValue < *(rValue.maxValue) {
				allowedMax = true
			}
		}
	} else {
		allowedMax = true
	}

	return allowedMin && allowedMax
}

func (rValue *RangeValue) Out(targetValue float64) bool {
	return !rValue.In(targetValue)
}

func ParseNumberRange(str string) (value *RangeValue, retErr error) {
	// support following patterns
	// pat1. x-y(x, y: variable and contain same value)
	// pat2. x<=v<=y(v: fixed character, x, y: variable)

	value = &RangeValue{}
	retErr = nil

	str = strings.Trim(str, " \t\"'")
	str = strings.ToLower(str)

	if strings.Index(str, "v") >= 0 {
		minmaxSlice := strings.Split(str, "v")

		for _, minmaxValue := range minmaxSlice {
			position := 0
			if position = strings.Index(minmaxValue, ">"); position >= 0 {
				tempMinmaxValue := strings.Trim(minmaxValue, ">=")
				if strings.Index(minmaxValue, tempMinmaxValue) < position {
					// max
					if equalPosition := strings.Index(minmaxValue, "="); equalPosition == position+1 {
						value.allowMaxEqual = true
					}

					if tempValue, parseErr := ParseNumber(tempMinmaxValue); parseErr == nil {
						value.maxValue = &tempValue
					} else {
						retErr = parseErr
						break
					}
				} else {
					// min
					if equalPosition := strings.Index(minmaxValue, "="); equalPosition == position+1 {
						value.allowMinEqual = true
					}

					if tempValue, parseErr := ParseNumber(tempMinmaxValue); parseErr == nil {
						value.minValue = &tempValue
					} else {
						retErr = parseErr
						break
					}
				}
			} else if position = strings.Index(minmaxValue, "<"); position >= 0 {
				tempMinmaxValue := strings.Trim(minmaxValue, "<=")
				if strings.Index(minmaxValue, tempMinmaxValue) > position {
					// max
					if equalPosition := strings.Index(minmaxValue, "="); equalPosition == position+1 {
						value.allowMaxEqual = true
					}

					if tempValue, parseErr := ParseNumber(tempMinmaxValue); parseErr == nil {
						value.maxValue = &tempValue
					} else {
						retErr = parseErr
						break
					}
				} else {
					// min
					if equalPosition := strings.Index(minmaxValue, "="); equalPosition == position+1 {
						value.allowMinEqual = true
					}

					if tempValue, parseErr := ParseNumber(tempMinmaxValue); parseErr == nil {
						value.minValue = &tempValue
					} else {
						retErr = parseErr
						break
					}
				}
			} else {
				retErr = fmt.Errorf("unknown format: %s", minmaxValue)
				break
			}
		}
	} else if strings.Contains(str, "-") {
		startPosition := 0
		if strings.HasPrefix(str, "-") {
			// remove "-" for split range values
			startPosition = 1
		}

		minmaxSlice := strings.SplitN(str[startPosition:], "-", 2)
		if len(minmaxSlice) == 2 {
			// return "-" for comparation
			minmaxSlice[0] = "-" + minmaxSlice[0]
			if tempValue, parseErr := ParseNumber(minmaxSlice[0]); parseErr == nil {
				value.minValue = &tempValue

				if tempValue, parseErr := ParseNumber(minmaxSlice[1]); parseErr == nil {
					value.maxValue = &tempValue

					value.allowMinEqual = true
					value.allowMaxEqual = true

					if (*value.minValue) > (*value.maxValue) {
						value.minValue, value.maxValue = value.maxValue, value.minValue
					}
				} else {
					retErr = parseErr
				}
			} else {
				retErr = parseErr
			}
		} else {
			retErr = fmt.Errorf("unknown format: %s", str)
		}
	} else {
		retErr = fmt.Errorf("unknown format: %s", str)
	}

	return value, retErr
}
