package utility

import (
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
