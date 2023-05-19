package main

import (
	"bytes"
	"encoding/json"
	"os"
	"reflect"
	"strings"
	TextTempl "text/template"

	ThcompUtility "github.com/thcomp/GoLang_Utility"
)

type Parameter struct {
	ConfigFilepath                        string `cui_param:"name:config,init:'./tz.json',desc:'config file for timezone',expect:'file|exist'"`
	OutputTemplateFilepath                string `cui_param:"name:template,init:'./output.tpl',desc:'output template filepath'"`
	OutputTzTextTemplateFilepath          string `cui_param:"name:tztext_template,init:'./const_tz_text.tpl',desc:'output timezone text template filepath'"`
	OutputFieldTemplateFilepath           string `cui_param:"name:field_template,init:'./field.tpl',desc:'output field template filepath'"`
	OutputMethodTemplateFilepath          string `cui_param:"name:method_template,init:'./method.tpl',desc:'output field template filepath'"`
	OutputShortNameMethodTemplateFilepath string `cui_param:"name:sn_method_template,init:'./short_name_method.tpl',desc:'output field template filepath'"`
	OutputFilepath                        string `cui_param:"name:out,init:'../TimeHelper.go',desc:'output filepath'"`
	OutputPackage                         string `cui_param:"name:package,init:'utility',desc:'package of output file'"`
	OutputStructName                      string `cui_param:"name:struct,init:'TimeHelper',desc:'package of output file'"`
	Debug                                 bool   `cui_param:"name:debug,init:false,desc:'enable debug log'"`
}

func main() {
	params := &Parameter{}

	if errs := ThcompUtility.GetCUIParameter(params, false); errs == nil || len(errs) == 0 {
		if params.Debug {
			ThcompUtility.ChangeLogLevel(ThcompUtility.LogLevelD)
		} else {
			ThcompUtility.ChangeLogLevel(ThcompUtility.LogLevelI)
		}

		if configFile, openErr := os.Open(params.ConfigFilepath); openErr == nil {
			defer configFile.Close()

			tzMap := map[string]interface{}{}
			if decodeErr := json.NewDecoder(configFile).Decode(&tzMap); decodeErr == nil {
				generate(params, tzMap)
			} else {
				ThcompUtility.LogfE("fail to decode tz config: %s, %v", params.ConfigFilepath, decodeErr)
			}
		}
	} else {
		ThcompUtility.LogfE("fail to get cui parameters: %v", errs)
	}
}

func generate(params *Parameter, tzMap map[string]interface{}) {
	if outputFile, openErr := os.OpenFile(params.OutputFilepath, os.O_CREATE|os.O_WRONLY, 0644); openErr == nil {
		defer outputFile.Close()

		funcMap := TextTempl.FuncMap{
			"createTimezoneText":     func() string { return outputTimezoneText(params, tzMap) },
			"createStructureFields":  func() string { return outputStructureFields(params, tzMap) },
			"createStructureMethods": func() string { return outputStructureMethods(params, tzMap) },
		}
		if outputTemplate, parseErr := TextTempl.New("go_output").Funcs(funcMap).ParseFiles(params.OutputTemplateFilepath); parseErr == nil {
			values := map[string]interface{}{
				"Package":   params.OutputPackage,
				"Structure": params.OutputStructName,
			}
			if execErr := outputTemplate.ExecuteTemplate(outputFile, "go_output", values); execErr != nil {
				ThcompUtility.LogfE("fail to output file: %v", execErr)
			}
		} else {
			ThcompUtility.LogfE("fail to parse template file: %s, %v", params.OutputTemplateFilepath, parseErr)
		}
	} else {
		ThcompUtility.LogfE("fail to open output file: %s, %v", params.OutputFilepath, openErr)
	}
}

func outputTimezoneText(params *Parameter, tzMap map[string]interface{}) string {
	builder := ThcompUtility.StringBuilder{}

	if outputTemplate, parseErr := TextTempl.ParseFiles(params.OutputTzTextTemplateFilepath); parseErr == nil {
		tmplValues := map[string]interface{}{}

		for region, value := range tzMap {
			region = strings.ToUpper(region[0:1]) + strings.ToLower(region[1:])

			if mapValue, assertionOK := value.(map[string]interface{}); assertionOK {
				if len(mapValue) > 0 {
					for city, cityValue := range mapValue {
						if _, assertionOK := cityValue.(map[string]interface{}); assertionOK {
							city = strings.ToUpper(city[0:1]) + strings.ToLower(city[1:])
							tmplValues["Region"] = region
							tmplValues["City"] = city
							buffer := bytes.NewBuffer([]byte{})
							if execErr := outputTemplate.Execute(buffer, tmplValues); execErr == nil {
								builder.Append(buffer.String())
							} else {
								ThcompUtility.LogfE("fail to execute template: %s, %v", params, outputTimezoneText, execErr)
							}
						} else {
							// only region, no output(ex. UTC)
						}
					}
				}
			}
		}
	} else {
		ThcompUtility.LogfE("fail to parse template file: %s, %v", params.OutputTemplateFilepath, parseErr)
	}

	return builder.String()
}

func outputStructureFields(params *Parameter, tzMap map[string]interface{}) string {
	builder := ThcompUtility.StringBuilder{}

	if outputTemplate, parseErr := TextTempl.ParseFiles(params.OutputFieldTemplateFilepath); parseErr == nil {
		tmplValues := map[string]interface{}{}

		for region, value := range tzMap {
			region = strings.ToUpper(region[0:1]) + strings.ToLower(region[1:])

			if mapValue, assertionOK := value.(map[string]interface{}); assertionOK {
				if len(mapValue) > 0 {
					for city, cityValue := range mapValue {
						if _, assertionOK := cityValue.(map[string]interface{}); assertionOK {
							city = strings.ToUpper(city[0:1]) + strings.ToLower(city[1:])
							tmplValues["Region"] = region
							tmplValues["City"] = city
							buffer := bytes.NewBuffer([]byte{})
							if execErr := outputTemplate.Execute(buffer, tmplValues); execErr == nil {
								builder.Append(buffer.String())
							} else {
								ThcompUtility.LogfE("fail to execute template: %s, %v", params, outputTimezoneText, execErr)
							}
						} else {
							// only region, no output(ex. UTC)
						}
					}
				}
			}
		}
	} else {
		ThcompUtility.LogfE("fail to parse template file: %s, %v", params.OutputTemplateFilepath, parseErr)
	}

	return builder.String()
}

func outputStructureMethods(params *Parameter, tzMap map[string]interface{}) string {
	builder := ThcompUtility.StringBuilder{}

	if outputTemplate, parseErr := TextTempl.ParseFiles(params.OutputMethodTemplateFilepath); parseErr == nil {
		tmplValues := map[string]interface{}{"Structure": params.OutputStructName}

		outputShortNameMethodTemplate, _ := TextTempl.ParseFiles(params.OutputShortNameMethodTemplateFilepath)
		for region, value := range tzMap {
			ThcompUtility.LogfD("region: %s, type of value: %v", region, reflect.TypeOf(value))

			if mapValue, assertionOK := value.(map[string]interface{}); assertionOK {
				tempRegion := strings.ToUpper(region[0:1]) + strings.ToLower(region[1:])

				if len(mapValue) > 0 {
					for city, cityValue := range mapValue {
						if cityValueMap, assertionOK := cityValue.(map[string]interface{}); assertionOK {
							city = strings.ToUpper(city[0:1]) + strings.ToLower(city[1:])
							tmplValues["Region"] = tempRegion
							tmplValues["City"] = city
							buffer := bytes.NewBuffer([]byte{})
							ThcompUtility.LogfD("city: %s", city)
							if execErr := outputTemplate.Execute(buffer, tmplValues); execErr == nil {
								builder.Appendf("%s\n", buffer.String())
							} else {
								ThcompUtility.LogfE("fail to execute template: %s, %v", params, outputTimezoneText, execErr)
							}

							if outputShortNameMethodTemplate != nil {
								for paramName, paramValue := range cityValueMap {
									if paramName == "short_name" {
										if shortName, assertionOK := paramValue.(string); assertionOK {
											tmplValues["ShortName"] = shortName
											buffer := bytes.NewBuffer([]byte{})
											ThcompUtility.LogfD("short name: %s", shortName)
											if execErr := outputShortNameMethodTemplate.Execute(buffer, tmplValues); execErr == nil {
												builder.Appendf("%s\n", buffer.String())
											} else {
												ThcompUtility.LogfE("fail to execute template: %s, %v", params, outputTimezoneText, execErr)
											}
										}
									}
								}
							}
						} else {
							// only region, no output(ex. UTC)
							tmplValues["Region"] = region
							tmplValues["City"] = ""
							buffer := bytes.NewBuffer([]byte{})
							ThcompUtility.LogfD("city: %s", "")
							if execErr := outputTemplate.Execute(buffer, tmplValues); execErr == nil {
								builder.Appendf("%s\n", buffer.String())
							} else {
								ThcompUtility.LogfE("fail to execute template: %s, %v", params, outputTimezoneText, execErr)
							}
						}
					}
				}
			} else {
				// only region, no output(ex. UTC)
				tmplValues["Region"] = region
				tmplValues["City"] = ""
				buffer := bytes.NewBuffer([]byte{})
				ThcompUtility.LogfD("city: %s", "")
				if execErr := outputTemplate.Execute(buffer, tmplValues); execErr == nil {
					builder.Appendf("%s\n", buffer.String())
				} else {
					ThcompUtility.LogfE("fail to execute template: %s, %v", params, outputTimezoneText, execErr)
				}
			}
		}
	} else {
		ThcompUtility.LogfE("fail to parse template file: %s, %v", params.OutputTemplateFilepath, parseErr)
	}

	return builder.String()
}
