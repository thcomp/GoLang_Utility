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

func Split(originalText string, sepTextSlice []string, invalidSpaceStarters ...string) []string {
	ret := []string{}
	startPosition := 0
	invalidSpaceMap := map[string]int{}

	for index := 0; index < len(originalText); {
		incrementSize := 1
		matchedInvalidSpaceStarter := false

		for _, invalidSpaceStarter := range invalidSpaceStarters {
			if index+len(invalidSpaceStarter) < len(originalText) {
				if originalText[index:index+len(invalidSpaceStarter)] == invalidSpaceStarter {
					if _, exist := invalidSpaceMap[invalidSpaceStarter]; exist {
						delete(invalidSpaceMap, invalidSpaceStarter)
					} else {
						invalidSpaceMap[invalidSpaceStarter] = index
					}

					incrementSize = len(invalidSpaceStarter)
					matchedInvalidSpaceStarter = true
					break
				}
			}
		}

		if !matchedInvalidSpaceStarter && len(invalidSpaceMap) == 0 {
			for _, sepText := range sepTextSlice {
				if index+len(sepText) < len(originalText) {
					if originalText[index:index+len(sepText)] == sepText {
						ret = append(ret, originalText[startPosition:index])
						startPosition = index + len(sepText)
						break
					}
				}
			}
		}

		index += incrementSize
	}

	if startPosition < len(originalText) {
		ret = append(ret, originalText[startPosition:])
	}

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

var sHalfKanaToKana = map[string]string{
	"ｧ":  "ァ",
	"ｨ":  "ィ",
	"ｩ":  "ゥ",
	"ｪ":  "ェ",
	"ｫ":  "ォ",
	"ｬ":  "ャ",
	"ｭ":  "ュ",
	"ｮ":  "ョ",
	"ｯ":  "ッ",
	"ｱ":  "ア",
	"ｲ":  "イ",
	"ｳ":  "ウ",
	"ｴ":  "エ",
	"ｵ":  "オ",
	"ｶ":  "カ",
	"ｷ":  "キ",
	"ｸ":  "ク",
	"ｹ":  "ケ",
	"ｺ":  "コ",
	"ｻ":  "サ",
	"ｼ":  "シ",
	"ｽ":  "ス",
	"ｾ":  "セ",
	"ｿ":  "ソ",
	"ﾀ":  "タ",
	"ﾁ":  "チ",
	"ﾂ":  "ツ",
	"ﾃ":  "テ",
	"ﾄ":  "ト",
	"ﾅ":  "ナ",
	"ﾆ":  "ニ",
	"ﾇ":  "ヌ",
	"ﾈ":  "ネ",
	"ﾉ":  "ノ",
	"ﾊ":  "ハ",
	"ﾋ":  "ヒ",
	"ﾌ":  "フ",
	"ﾍ":  "ヘ",
	"ﾎ":  "ホ",
	"ﾏ":  "マ",
	"ﾐ":  "ミ",
	"ﾑ":  "ム",
	"ﾒ":  "メ",
	"ﾓ":  "モ",
	"ﾔ":  "ヤ",
	"ﾕ":  "ユ",
	"ﾖ":  "ヨ",
	"ﾗ":  "ラ",
	"ﾘ":  "リ",
	"ﾙ":  "ル",
	"ﾚ":  "レ",
	"ﾛ":  "ロ",
	"ﾜ":  "ワ",
	"ｦ":  "ヲ",
	"ﾝ":  "ン",
	"ｶﾞ": "ガ",
	"ｷﾞ": "ギ",
	"ｸﾞ": "グ",
	"ｹﾞ": "ゲ",
	"ｺﾞ": "ゴ",
	"ｻﾞ": "ザ",
	"ｼﾞ": "ジ",
	"ｽﾞ": "ズ",
	"ｾﾞ": "ゼ",
	"ｿﾞ": "ゾ",
	"ﾀﾞ": "ダ",
	"ﾁﾞ": "ヂ",
	"ﾂﾞ": "ヅ",
	"ﾃﾞ": "デ",
	"ﾄﾞ": "ド",
	"ﾊﾞ": "バ",
	"ﾋﾞ": "ビ",
	"ﾌﾞ": "ブ",
	"ﾍﾞ": "ベ",
	"ﾎﾞ": "ボ",
	"ﾊﾟ": "パ",
	"ﾋﾟ": "ピ",
	"ﾌﾟ": "プ",
	"ﾍﾟ": "ペ",
	"ﾎﾟ": "ポ",
}

var sNFDtoNFC = map[string]string{
	"カ゛": "ガ",
	"キ゛": "ギ",
	"ク゛": "グ",
	"ケ゛": "ゲ",
	"コ゛": "ゴ",
	"サ゛": "ザ",
	"シ゛": "ジ",
	"ス゛": "ズ",
	"セ゛": "ゼ",
	"ソ゛": "ゾ",
	"タ゛": "ダ",
	"チ゛": "ヂ",
	"ツ゛": "ヅ",
	"テ゛": "デ",
	"ト゛": "ド",
	"ハ゛": "バ",
	"ヒ゛": "ビ",
	"フ゛": "ブ",
	"ヘ゛": "ベ",
	"ホ゛": "ボ",
	"ハ゜": "パ",
	"ヒ゜": "ピ",
	"フ゜": "プ",
	"ヘ゜": "ペ",
	"ホ゜": "ポ",
}

var sNFCtoNFD = map[string]string{
	"ガ": "カ゛",
	"ギ": "キ゛",
	"グ": "ク゛",
	"ゲ": "ケ゛",
	"ゴ": "コ゛",
	"ザ": "サ゛",
	"ジ": "シ゛",
	"ズ": "ス゛",
	"ゼ": "セ゛",
	"ゾ": "ソ゛",
	"ダ": "タ゛",
	"ヂ": "チ゛",
	"ヅ": "ツ゛",
	"デ": "テ゛",
	"ド": "ト゛",
	"バ": "ハ゛",
	"ビ": "ヒ゛",
	"ブ": "フ゛",
	"ベ": "ヘ゛",
	"ボ": "ホ゛",
	"パ": "ハ゜",
	"ピ": "ヒ゜",
	"プ": "フ゜",
	"ペ": "ヘ゜",
	"ポ": "ホ゜",
}

var sHalfAlphabetToAlphabet = map[string]string{
	"A": "Ａ",
	"B": "Ｂ",
	"C": "Ｃ",
	"D": "Ｄ",
	"E": "Ｅ",
	"F": "Ｆ",
	"G": "Ｇ",
	"H": "Ｈ",
	"I": "Ｉ",
	"J": "Ｊ",
	"K": "Ｋ",
	"L": "Ｌ",
	"M": "Ｍ",
	"N": "Ｎ",
	"O": "Ｏ",
	"P": "Ｐ",
	"Q": "Ｑ",
	"R": "Ｒ",
	"S": "Ｓ",
	"T": "Ｔ",
	"U": "Ｕ",
	"V": "Ｖ",
	"W": "Ｗ",
	"X": "Ｘ",
	"Y": "Ｙ",
	"Z": "Ｚ",
	"a": "ａ",
	"b": "ｂ",
	"c": "ｃ",
	"d": "ｄ",
	"e": "ｅ",
	"f": "ｆ",
	"g": "ｇ",
	"h": "ｈ",
	"i": "ｉ",
	"j": "ｊ",
	"k": "ｋ",
	"l": "ｌ",
	"m": "ｍ",
	"n": "ｎ",
	"o": "ｏ",
	"p": "ｐ",
	"q": "ｑ",
	"r": "ｒ",
	"s": "ｓ",
	"t": "ｔ",
	"u": "ｕ",
	"v": "ｖ",
	"w": "ｗ",
	"x": "ｘ",
	"y": "ｙ",
	"z": "ｚ",
}

var sAlphabetToHalfAlphabet = map[string]string{
	"Ａ": "A",
	"Ｂ": "B",
	"Ｃ": "C",
	"Ｄ": "D",
	"Ｅ": "E",
	"Ｆ": "F",
	"Ｇ": "G",
	"Ｈ": "H",
	"Ｉ": "I",
	"Ｊ": "J",
	"Ｋ": "K",
	"Ｌ": "L",
	"Ｍ": "M",
	"Ｎ": "N",
	"Ｏ": "O",
	"Ｐ": "P",
	"Ｑ": "Q",
	"Ｒ": "R",
	"Ｓ": "S",
	"Ｔ": "T",
	"Ｕ": "U",
	"Ｖ": "V",
	"Ｗ": "W",
	"Ｘ": "X",
	"Ｙ": "Y",
	"Ｚ": "Z",
	"ａ": "a",
	"ｂ": "b",
	"ｃ": "c",
	"ｄ": "d",
	"ｅ": "e",
	"ｆ": "f",
	"ｇ": "g",
	"ｈ": "h",
	"ｉ": "i",
	"ｊ": "j",
	"ｋ": "k",
	"ｌ": "l",
	"ｍ": "m",
	"ｎ": "n",
	"ｏ": "o",
	"ｐ": "p",
	"ｑ": "q",
	"ｒ": "r",
	"ｓ": "s",
	"ｔ": "t",
	"ｕ": "u",
	"ｖ": "v",
	"ｗ": "w",
	"ｘ": "x",
	"ｙ": "y",
	"ｚ": "z",
}

var sHalfNumberToNumber = map[string]string{
	"0": "０",
	"1": "１",
	"2": "２",
	"3": "３",
	"4": "４",
	"5": "５",
	"6": "６",
	"7": "７",
	"8": "８",
	"9": "９",
}

var sNumberToHalfNumber = map[string]string{
	"０": "0",
	"１": "1",
	"２": "2",
	"３": "3",
	"４": "4",
	"５": "5",
	"６": "6",
	"７": "7",
	"８": "8",
	"９": "9",
}

var sHalfSymbolToSymbol = map[string]string{
	"!":  "！",
	"\"": "“",
	"#":  "＃",
	"$":  "＄",
	"%":  "％",
	"&":  "＆",
	"'":  "’",
	"(":  "（",
	")":  "）",
	"=":  "＝",
	"-":  "—",
	"^":  "＾",
	"~":  "～",
	"|":  "｜",
	"\\": "￥",
	"`":  "‘",
	"@":  "＠",
	"[":  "［",
	"{":  "｛",
	"+":  "＋",
	";":  "；",
	"*":  "＊",
	":":  "：",
	"]":  "］",
	"}":  "｝",
	",":  "，",
	"<":  "＜",
	".":  "．",
	">":  "＞",
	"/":  "／",
	"?":  "？",
	"_":  "＿",
}

var sSymbolToHalfSymbol = map[string]string{
	"！": "!",
	"“": "\"",
	"”": "\"",
	"＃": "#",
	"＄": "$",
	"％": "%",
	"＆": "&",
	"’": "'",
	"（": "(",
	"）": ")",
	"＝": "=",
	"—": "-",
	"＾": "^",
	"～": "~",
	"｜": "|",
	"￥": "\\",
	"‘": "`",
	"＠": "@",
	"「": "[",
	"［": "[",
	"【": "[",
	"『": "[",
	"｛": "{",
	"＋": "+",
	"；": ";",
	"＊": "*",
	"：": ":",
	"」": "]",
	"］": "]",
	"】": "]",
	"』": "]",
	"｝": "}",
	"，": ",",
	"＜": "<",
	"．": ".",
	"＞": ">",
	"／": "/",
	"？": "?",
	"＿": "_",
}

func ExchangeHalfKanaToKana(originalText string) string {
	ret := originalText

	for halfKana, kana := range sHalfKanaToKana {
		ret = strings.Replace(ret, halfKana, kana, -1)
	}

	return ret
}

func ExchangeNFDtoNFC(originalText string) string {
	ret := originalText

	for nfdStr, nfcStr := range sNFDtoNFC {
		ret = strings.Replace(ret, nfdStr, nfcStr, -1)
	}

	return ret
}

func ExchangeNFCtoNFD(originalText string) string {
	ret := originalText

	for nfcStr, nfdStr := range sNFCtoNFD {
		ret = strings.Replace(ret, nfcStr, nfdStr, -1)
	}

	return ret
}

func ExchangeHalfNumberToNumber(originalText string) string {
	ret := originalText

	for halfNum, num := range sHalfNumberToNumber {
		ret = strings.Replace(ret, halfNum, num, -1)
	}

	return ret
}

func ExchangeNumberToHalfNumber(originalText string) string {
	ret := originalText

	for num, halfNum := range sNumberToHalfNumber {
		ret = strings.Replace(ret, num, halfNum, -1)
	}

	return ret
}

func ExchangeHalfAlphabetToAlphabet(originalText string) string {
	ret := originalText

	for halfNum, num := range sHalfAlphabetToAlphabet {
		ret = strings.Replace(ret, halfNum, num, -1)
	}

	return ret
}

func ExchangeAlphabetToHalfAlphabet(originalText string) string {
	ret := originalText

	for num, halfNum := range sAlphabetToHalfAlphabet {
		ret = strings.Replace(ret, num, halfNum, -1)
	}

	return ret
}

func ExchangeHalfSpaceToSpace(originalText string) string {
	return strings.Replace(originalText, " ", "　", -1)
}

func ExchangeSpaceToHalfSpace(originalText string) string {
	return strings.Replace(originalText, "　", " ", -1)
}

func ExchangeHalfSymbolToSymbol(originalText string) string {
	ret := originalText

	for halfSymbol, symbol := range sHalfSymbolToSymbol {
		ret = strings.Replace(ret, halfSymbol, symbol, -1)
	}

	return ret
}

func ExchangeSymbolToHalfSymbol(originalText string) string {
	ret := originalText

	for symbol, halfSymbol := range sSymbolToHalfSymbol {
		ret = strings.Replace(ret, symbol, halfSymbol, -1)
	}

	return ret
}
