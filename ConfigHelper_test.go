package utility

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

type ConfigDataSub struct {
	Value1 string  `json:"val1"`
	Value2 int     `json:"val2"`
	Value3 float32 `json:"val3"`
}

type ConfigData struct {
	Value1 string         `json:"val1"`
	Value2 int            `json:"val2"`
	Value3 float32        `json:"val3"`
	Value4 ConfigDataSub  `json:"val4"`
	Value5 *ConfigDataSub `json:"val5,omitempty"`
}

func TestExpandConfigData(t *testing.T) {
	testJsonFile := `TestData/test.json`
	jsonDataArray := []string{
		`{"val1": "data1","val2": 1234567,"val3": 1.12345,"val4": {"val1": "data2","val2": 56789,"val3": 6.789}}`,
		`{"val1": "data3","val2": 6789,"val3": 4.124565,"val4": {"val1": "data2","val2": 56789,"val3": 6.789},"val5": {"val1": "data4","val2": 8901,"val3": 8.9123}}`,
	}

	configData := new(ConfigData)
	configHelper := NewConfigHelper(configData, testJsonFile)

	for _, jsonData := range jsonDataArray {
		if testFile, openError := os.OpenFile(testJsonFile, os.O_CREATE, os.ModePerm); openError == nil {
			// テストデータを作成
			var originalConfigData ConfigData
			json.Unmarshal([]byte(jsonData), &originalConfigData)
			testFile.Write([]byte(jsonData))
			testFile.Close()

			configHelper.ExpandConfigData()

			if configData.Value1 != originalConfigData.Value1 {
				t.Errorf("Value1 is not matched: %s vs %s", configData.Value1, originalConfigData.Value1)
			}

			if configData.Value2 != originalConfigData.Value2 {
				t.Errorf("Value2 is not matched: %d vs %d", configData.Value2, originalConfigData.Value2)
			}

			if configData.Value3 != originalConfigData.Value3 {
				t.Errorf("Value3 is not matched: %f vs %f", configData.Value3, originalConfigData.Value3)
			}

			if configData.Value4.Value1 != originalConfigData.Value4.Value1 {
				t.Errorf("Value1 of Value4 is not matched: %s vs %s", configData.Value4.Value1, originalConfigData.Value4.Value1)
			}

			if configData.Value4.Value2 != originalConfigData.Value4.Value2 {
				t.Errorf("Value2 of Value4 is not matched: %d vs %d", configData.Value4.Value2, originalConfigData.Value4.Value2)
			}

			if configData.Value4.Value3 != originalConfigData.Value4.Value3 {
				t.Errorf("Value3 of Value4 is not matched: %f vs %f", configData.Value4.Value3, originalConfigData.Value4.Value3)
			}

			if (configData.Value5 == nil && originalConfigData.Value5 != nil) || (configData.Value5 != nil && originalConfigData.Value5 == nil) {
				t.Errorf("Value5 is not matched: %v vs %v", configData.Value5, originalConfigData.Value5)
			} else if configData.Value5 != nil {
				if configData.Value5.Value1 != originalConfigData.Value5.Value1 {
					t.Errorf("Value1 of Value5 is not matched: %s vs %s", configData.Value5.Value1, originalConfigData.Value5.Value1)
				}

				if configData.Value5.Value2 != originalConfigData.Value5.Value2 {
					t.Errorf("Value2 of Value5 is not matched: %d vs %d", configData.Value5.Value2, originalConfigData.Value5.Value2)
				}

				if configData.Value4.Value3 != originalConfigData.Value4.Value3 {
					t.Errorf("Value3 of Value4 is not matched: %f vs %f", configData.Value4.Value3, originalConfigData.Value4.Value3)
				}
			}

			os.Remove(testJsonFile)
			duration, _ := time.ParseDuration(`1s`)
			time.Sleep(duration)
		}
	}

	var lastUpdateTimeUt = configHelper.lastUpdateTimeUt
	configHelper.ExpandConfigData()

	if lastUpdateTimeUt != configHelper.lastUpdateTimeUt {
		t.Errorf("it should not update: %d vs %d", lastUpdateTimeUt, configHelper.lastUpdateTimeUt)
	}
}
