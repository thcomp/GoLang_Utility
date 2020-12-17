package utility

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type ConfigHelper struct {
	configData       interface{}
	configFilePath   string
	lastUpdateTimeUt int64
	isCreatedByFunc  bool
}

func NewConfigHelper(configData interface{}, configFilePath string) *ConfigHelper {
	var ret *ConfigHelper = nil

	ret = new(ConfigHelper)
	ret.configData = configData
	ret.configFilePath = configFilePath
	ret.isCreatedByFunc = true

	return ret
}

func (helper *ConfigHelper) ExpandConfigData() error {
	var ret error = nil

	ret, _ = helper.ExpandConfigData2()

	return ret
}

func (helper *ConfigHelper) ExpandConfigData2() (error, bool) {
	var ret error = nil
	var modified = false

	if helper.isCreatedByFunc {
		if configFileStat, statError := os.Stat(helper.configFilePath); statError == nil {
			var currentConfigFileTimeUt int64 = configFileStat.ModTime().Unix()
			if currentConfigFileTimeUt > helper.lastUpdateTimeUt {
				helper.lastUpdateTimeUt = currentConfigFileTimeUt

				if jsonBytes, readFileErr := ioutil.ReadFile(helper.configFilePath); readFileErr == nil {
					if unmarshalErr := json.Unmarshal(jsonBytes, &helper.configData); unmarshalErr != nil {
						LogE(unmarshalErr.Error())
					} else {
						modified = true
					}
				} else {
					LogE(readFileErr.Error())
				}
			}
		} else {
			LogE(statError.Error())
		}
	} else {
		ret = errors.New(`helper is not created by NewConfigHelper`)
	}

	return ret, modified
}

func (helper *ConfigHelper) StoreConfigData() error {
	var ret error = nil

	if helper.isCreatedByFunc {
		mode := os.FileMode(0644)
		if configFileStat, statError := os.Stat(helper.configFilePath); statError == nil {
			mode = configFileStat.Mode()
		}

		if jsonBytes, marshalErr := json.Marshal(helper.configData); marshalErr == nil {
			if ret = ioutil.WriteFile(helper.configFilePath, jsonBytes, mode); ret == nil {
				if configFileStat, statError := os.Stat(helper.configFilePath); statError == nil {
					helper.lastUpdateTimeUt = configFileStat.ModTime().Unix()
				}
			}
		} else {
			ret = marshalErr
		}
	} else {
		ret = errors.New(`helper is not created by NewConfigHelper`)
	}

	return ret
}
