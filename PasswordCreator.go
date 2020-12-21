package utility

import (
	"bytes"
	"math/rand"
	"time"
)

const PwValidSmallAlpha = 1
const PwValidLargeAlpha = 2
const PwValidNumeric = 4
const PwValidSymbol = 8

func CreatePassword(length int, validTypes ...int) string {
	ret := bytes.NewBuffer([]byte{})
	validCharacters := ""
	validCharMap := map[int]string{}

	if validTypes != nil && len(validTypes) > 0 {
		for _, validType := range validTypes {
			if (validType & PwValidSmallAlpha) == PwValidSmallAlpha {
				validCharMap[PwValidSmallAlpha] = "abcdefghijklmnopqrstuvwxyz"
			}
			if (validType & PwValidLargeAlpha) == PwValidLargeAlpha {
				validCharMap[PwValidLargeAlpha] = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
			}
			if (validType & PwValidNumeric) == PwValidNumeric {
				validCharMap[PwValidNumeric] = "0123456789"
			}
			if (validType & PwValidSymbol) == PwValidSymbol {
				validCharMap[PwValidSymbol] = "!\"#$%&'()=-~^\\|`@{[+;:*]}<,>./?_"
			}
		}
	} else {
		validCharMap[PwValidSmallAlpha] = "abcdefghijklmnopqrstuvwxyz"
		validCharMap[PwValidLargeAlpha] = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		validCharMap[PwValidNumeric] = "0123456789"
	}

	for _, value := range validCharMap {
		validCharacters = validCharacters + value
	}

	if length == 0 {
		length = 12
	}

	source := rand.NewSource(time.Now().UnixNano())
	validCharLen := len(validCharacters)
	for i := 0; i < length; i++ {
		ret.WriteByte(validCharacters[int(source.Int63()%int64(validCharLen))])
	}

	return ret.String()
}
