package utility

import (
	"encoding/json"
)

const sCUIParamTag = "cui_param"
const sCUIParamTagName = "name"
const sCUIParamTagInit = "init"
const sCUIParamTagDescription = "desc"

type sParamTagInfo struct {
	Name        string      `json:"name"`
	InitValue   interface{} `json:"init"`
	Description string      `json:"desc"`
	Except      string      `json:"except"`
}

func GetCUIParameter(receiver interface{}) error {
	reflectHelper := NewReflectHelper(receiver)
	for i := 0; i < reflectHelper.NumField(); i++ {
		tagInfo := reflectHelper.Tag(i)
		if tagValue, exist := tagInfo.Lookup(sCUIParamTag); exist {
			paramInfo := sParamTagInfo{}
			if unmarshalErr := json.Unmarshal([]byte(tagValue), &paramInfo); unmarshalErr == nil {
				// json format
			} else {
				// name and value separated by ":" joint comma format
				// ex. name:"file",init:1,desc:test,except:"file|exist"
				//nameAndValueSlice := strings.Split(tagValue, ",")
			}
		}
	}
	return nil
}
