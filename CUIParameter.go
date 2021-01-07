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

type sParamTagInfo struct {
	Name           string `json:"name"`
	targetField    reflect.Value
	targetFieldInf interface{}
	inputValue     interface{}
	InitValue      interface{} `json:"init"`
	Description    string      `json:"desc"`
	Expect         string      `json:"expect"`
}

func GetCUIParameter(receiver interface{}) []error {
	ret := []error{}
	receiverValue := reflect.ValueOf(receiver)

	if receiverValue.Kind() == reflect.Ptr {
		reflectHelper := NewReflectHelper(receiver)
		paramInfoSlice := [](*sParamTagInfo){}

		for i := 0; i < reflectHelper.NumField(); i++ {
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
							case "name":
								paramInfo.Name = strings.Trim(nameValueSlice[1], " \t\"'")
								break
							case "init":
								nameValueSlice[1] = strings.Trim(nameValueSlice[1], " \t")
								if (strings.HasPrefix(nameValueSlice[1], "\"") && strings.HasSuffix(nameValueSlice[1], "\"")) ||
									(strings.HasPrefix(nameValueSlice[1], "'") && strings.HasSuffix(nameValueSlice[1], "'")) {
									paramInfo.InitValue = nameValueSlice[1]
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
							case "desc":
								paramInfo.Description = strings.Trim(nameValueSlice[1], " \t\"'")
								break
							case "expect":
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

		if len(paramInfoSlice) > 0 {
			for _, paramInfo := range paramInfoSlice {
				infHelper := NewInterfaceHelper(paramInfo.targetFieldInf)
				targetKind := infHelper.GetKind()

				switch targetKind {
				case reflect.Bool:
					defBoolValue, _ := paramInfo.InitValue.(bool)
					paramInfo.inputValue = flag.Bool(paramInfo.Name, defBoolValue, paramInfo.Description)
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
					break
				case reflect.Float32:
					fallthrough
				case reflect.Float64:
					defFloatValue, _ := paramInfo.InitValue.(float64)
					paramInfo.inputValue = flag.Float64(paramInfo.Name, defFloatValue, paramInfo.Description)
					break
				case reflect.String:
					defStringValue, _ := paramInfo.InitValue.(string)
					paramInfo.inputValue = flag.String(paramInfo.Name, defStringValue, paramInfo.Description)
					break
				}
			}
			flag.Parse()

			for _, paramInfo := range paramInfoSlice {
				if paramInfo.inputValue != nil {
					infHelper := NewInterfaceHelper(paramInfo.targetField)
					infHelper.Set(reflect.Indirect(reflect.ValueOf(paramInfo.inputValue)))
				}
			}
		}
	} else {
		ret = append(ret, fmt.Errorf("receiver need pointer"))
	}

	return ret
}
