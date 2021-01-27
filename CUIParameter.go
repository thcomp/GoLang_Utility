package utility

import (
	"encoding/json"
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const sCUIParamTag = "cui_param"
const sCUIParamTagName = "name"
const sCUIParamTagInit = "init"
const sCUIParamTagDescription = "desc"
const sCUIParamTagExpect = "expect"

const CUIParamExpectIn = "in("
const CUIParamExpectFolder = "folder"
const CUIParamExpectFile = "file"
const CUIParamExpectExist = "exist"
const CUIParamExpectNotExist = "not_exist"

type sParamTagInfo struct {
	Name           string `json:"name"`
	targetField    reflect.Value
	targetFieldInf interface{}
	inputValue     interface{}
	InitValue      interface{} `json:"init"`
	Description    string      `json:"desc"`
	Expect         string      `json:"expect"`
}

func GetCUIParameter(receiver interface{}, debug bool) []error {
	currentLogLevel := GetLogLevel()
	defer ChangeLogLevel(currentLogLevel)
	if debug {
		ChangeLogLevel(LogLevelV)
	}

	paramInfoSlice, ret := parseParameter(receiver)
	if len(paramInfoSlice) > 0 {
		for _, paramInfo := range paramInfoSlice {
			infHelper := NewInterfaceHelper(paramInfo.targetFieldInf)
			targetKind := infHelper.GetKind()

			switch targetKind {
			case reflect.Bool:
				defBoolValue, _ := paramInfo.InitValue.(bool)
				paramInfo.inputValue = flag.Bool(paramInfo.Name, defBoolValue, paramInfo.Description)
				LogfV("register %s, %v, %s", paramInfo.Name, defBoolValue, paramInfo.Description)
				break
			case reflect.Int:
				fallthrough
			case reflect.Int8:
				fallthrough
			case reflect.Int16:
				fallthrough
			case reflect.Int32:
				fallthrough
			case reflect.Int64:
				initInfHelper := NewInterfaceHelper(paramInfo.InitValue)
				defIntValue, _ := initInfHelper.GetNumber()
				paramInfo.inputValue = flag.Int64(paramInfo.Name, int64(defIntValue), paramInfo.Description)
				LogfV("register %s, %v, %s", paramInfo.Name, defIntValue, paramInfo.Description)
				break
			case reflect.Uint:
				fallthrough
			case reflect.Uint8:
				fallthrough
			case reflect.Uint16:
				fallthrough
			case reflect.Uint32:
				fallthrough
			case reflect.Uint64:
				initInfHelper := NewInterfaceHelper(paramInfo.InitValue)
				defUintValue, _ := initInfHelper.GetNumber()
				paramInfo.inputValue = flag.Uint64(paramInfo.Name, uint64(defUintValue), paramInfo.Description)
				LogfV("register %s, %v, %s", paramInfo.Name, defUintValue, paramInfo.Description)
				break
			case reflect.Float32:
				fallthrough
			case reflect.Float64:
				defFloatValue, _ := paramInfo.InitValue.(float64)
				paramInfo.inputValue = flag.Float64(paramInfo.Name, defFloatValue, paramInfo.Description)
				LogfV("register %s, %v, %s", paramInfo.Name, defFloatValue, paramInfo.Description)
				break
			case reflect.String:
				defStringValue, _ := paramInfo.InitValue.(string)
				paramInfo.inputValue = flag.String(paramInfo.Name, defStringValue, paramInfo.Description)
				LogfV("register %s, %v, %s", paramInfo.Name, defStringValue, paramInfo.Description)
				break
			}
		}
		LogV("parse args")
		flag.Parse()

		for _, paramInfo := range paramInfoSlice {
			if paramInfo.inputValue != nil {
				infHelper := NewInterfaceHelper(paramInfo.targetField)

				inputValue := reflect.ValueOf(paramInfo.inputValue)
				inputValueIndirect := reflect.Indirect(inputValue)
				LogfV("input value: %v, indirect: %v", inputValue, inputValueIndirect)

				infHelper.Set(inputValueIndirect)

				if validErr := isValidParameter(infHelper, paramInfo); validErr != nil {
					ret = append(ret, validErr)
					break
				}
			}
		}
	}

	return ret
}

func parseParameter(receiver interface{}) ([](*sParamTagInfo), []error) {
	paramInfoSlice := [](*sParamTagInfo){}
	ret := []error{}
	receiverValue := reflect.ValueOf(receiver)

	if receiverValue.Kind() == reflect.Ptr {
		reflectHelper := NewReflectHelper(receiver)

		for i := 0; i < reflectHelper.NumField(); i++ {
			if reflectHelper.ValueKind(i) == reflect.Struct {
				tempParamInfoSlice, tempErrors := parseParameter(reflectHelper.GetAddrByIndex(i).Interface())
				if tempParamInfoSlice != nil && len(tempParamInfoSlice) > 0 {
					for _, tempParamInfo := range tempParamInfoSlice {
						paramInfoSlice = append(paramInfoSlice, tempParamInfo)
					}
				}
				if tempErrors != nil && len(tempErrors) > 0 {
					for _, tempError := range tempErrors {
						ret = append(ret, tempError)
					}
				}
				continue
			} else {
				tempError := error(nil)
				tagInfo := reflectHelper.Tag(i)
				if tagValue, exist := tagInfo.Lookup(sCUIParamTag); exist {
					paramInfo := &sParamTagInfo{targetField: reflectHelper.GetPtrByIndex(i), targetFieldInf: reflectHelper.GetByIndex(i)}
					if unmarshalErr := json.Unmarshal([]byte(tagValue), &paramInfo); unmarshalErr == nil {
						// json format
					} else {
						// name and value separated by ":" joint comma format
						// ex. name:"file",init:1,desc:test,expect:"file|exist"
						nameAndValueSlice := strings.Split(tagValue, ",")
						for _, nameAndValue := range nameAndValueSlice {
							nameAndValue = strings.Trim(nameAndValue, " \t")
							nameValueSlice := strings.Split(nameAndValue, ":")

							if len(nameValueSlice) >= 2 {
								switch nameValueSlice[0] {
								case sCUIParamTagName:
									paramInfo.Name = strings.Trim(nameValueSlice[1], " \t\"'")
									break
								case sCUIParamTagInit:
									nameValueSlice[1] = strings.Trim(nameValueSlice[1], " \t")
									if (strings.HasPrefix(nameValueSlice[1], "\"") && strings.HasSuffix(nameValueSlice[1], "\"")) ||
										(strings.HasPrefix(nameValueSlice[1], "'") && strings.HasSuffix(nameValueSlice[1], "'")) {
										paramInfo.InitValue = nameValueSlice[1][1 : len(nameValueSlice[1])-1]
									} else {
										nameValueSlice[1] = strings.ToLower(nameValueSlice[1])

										if nameValueSlice[1] == "true" {
											paramInfo.InitValue = true
										} else if nameValueSlice[1] == "false" {
											paramInfo.InitValue = false
										} else if strings.HasPrefix(nameValueSlice[1], "0x") {
											// hex
											if tempValue, parseErr := strconv.ParseUint(nameValueSlice[1], 16, 64); parseErr == nil {
												paramInfo.InitValue = tempValue
											}
										} else if strings.HasPrefix(nameValueSlice[1], "-") {
											// under 0 value
											if strings.Index(nameValueSlice[1], "0x") > 0 {
												if tempValue, parseErr := strconv.ParseInt(nameValueSlice[1], 16, 64); parseErr == nil {
													paramInfo.InitValue = tempValue
												} else {
													tempError = parseErr
												}
											} else {
												if tempValue, parseErr := strconv.ParseInt(nameValueSlice[1], 10, 64); parseErr == nil {
													paramInfo.InitValue = tempValue
												} else {
													tempError = parseErr
												}
											}
										} else {
											if tempValue, parseErr := strconv.ParseFloat(nameValueSlice[1], 64); parseErr == nil {
												paramInfo.InitValue = tempValue
											} else {
												tempError = parseErr
											}
										}
									}
									break
								case sCUIParamTagDescription:
									paramInfo.Description = strings.Trim(nameValueSlice[1], " \t\"'")
									break
								case sCUIParamTagExpect:
									paramInfo.Expect = strings.Trim(nameValueSlice[1], " \t\"'")
									break
								}
							}
						}
					}

					if len(paramInfo.Name) == 0 {
						paramInfo.Name = reflectHelper.Name(i)
					}

					if tempError == nil {
						paramInfoSlice = append(paramInfoSlice, paramInfo)
					} else {
						ret = append(ret, tempError)
					}
				}
			}
		}

	} else {
		ret = append(ret, fmt.Errorf("receiver need pointer"))
	}

	return paramInfoSlice, ret
}

func isValidParameter(fieldInfHelper *InterfaceHelper, paramInfo *sParamTagInfo) error {
	ret := error(nil)

	if len(paramInfo.Expect) > 0 {
		if fieldInfHelper.IsString() {
			tempValue, _ := fieldInfHelper.GetString()
			ret = isValidStringParameter(tempValue, paramInfo.Expect)
		} else if fieldInfHelper.IsNumber() {
			tempValue, _ := fieldInfHelper.GetNumber()
			ret = isValidNumberParameter(tempValue, paramInfo.Expect)
		}
	}

	return ret
}

func isValidStringParameter(strValue string, expect string) error {
	ret := error(nil)
	expectFileAndNotExist := false
	expectFolderAndNotExist := false
	needExistButNotExist := false
	needNotExistButExist := false
	inItem := false

	expectItems := strings.Split(expect, "|")
	for _, expectItem := range expectItems {
		expectItem = strings.ToLower(strings.Trim(expectItem, " \t"))
		if expectItem == CUIParamExpectFile {
			if !IsExist(strValue) {
				expectFileAndNotExist = true
			} else if IsDir(strValue) {
				ret = fmt.Errorf("expect file, but it is folder: '%s'", strValue)
				break
			}
		} else if expectItem == CUIParamExpectFolder {
			if !IsExist(strValue) {
				expectFolderAndNotExist = true
			} else if IsFile(strValue) {
				ret = fmt.Errorf("expect folder, but it is file: '%s'", strValue)
				break
			}
		} else if expectItem == CUIParamExpectExist {
			if !IsExist(strValue) {
				needExistButNotExist = true
			}
		} else if expectItem == CUIParamExpectNotExist {
			if IsExist(strValue) {
				needNotExistButExist = true
			}
		} else if strings.HasPrefix(expectItem, CUIParamExpectIn) && strings.HasSuffix(expectItem, ")") {
			expectItem = strings.Trim(expectItem[len(CUIParamExpectIn):len(expectItem)-1], " \t\"'")
			splitExpectItem := strings.Split(expectItem, ",")

			for _, uniSplitExpectItem := range splitExpectItem {
				if uniSplitExpectItem == strValue {
					inItem = true
					break
				}
			}
		}
	}

	if inItem {
		// no-op
	} else {
		if expectFileAndNotExist || expectFolderAndNotExist || needExistButNotExist || needNotExistButExist {
			if (expectFileAndNotExist || expectFolderAndNotExist) && needExistButNotExist {
				ret = fmt.Errorf("not found: '%s'", strValue)
			} else if needNotExistButExist {
				ret = fmt.Errorf("expect not exist, but exist: '%s'", strValue)
			}
		}
	}

	return ret
}

func isValidNumberParameter(floatValue float64, expect string) error {
	ret := error(nil)
	splitExpect := strings.Split(expect, ",")

	if len(splitExpect) == 1 {
		// uni value or uni range
		splitExpect[0] = strings.Trim(splitExpect[0], " \t")

		if rangeValue, parseErr := ParseNumberRange(splitExpect[0]); parseErr == nil {
			// range
			if rangeValue.Out(floatValue) {
				ret = fmt.Errorf("%f is out of %s", floatValue, expect)
			}
		} else {
			// value
			if tempValue, err := ParseNumber(splitExpect[0]); err == nil {
				if floatValue != tempValue {
					ret = fmt.Errorf("not expected value: %f not in %s", floatValue, expect)
				}
			} else {
				ret = err
			}
		}
	} else {
		// multi value or range
		valid := false
		for _, uniSplitExpect := range splitExpect {
			if err := isValidNumberParameter(floatValue, uniSplitExpect); err == nil {
				valid = true
				break
			}
		}

		if !valid {
			ret = fmt.Errorf("not expected value: %f not in %s", floatValue, expect)
		}
	}

	return ret
}
