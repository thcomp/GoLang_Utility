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
