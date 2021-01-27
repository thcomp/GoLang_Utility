package utility

import (
	"fmt"
	"strings"
)

type InputHandler struct {
	Key     string
	Handler func(string) bool
}

func GetUserInput(prefix string, caseSensitive bool, inputHandler InputHandler, otherInputHandlers ...InputHandler) (userInput string, ret bool) {
	userInput = ""
	ret = false

	targetInputMap := map[string](*InputHandler){}
	othersInputHandler := (*InputHandler)(nil)

	if inputHandler.Key == "" {
		othersInputHandler = &inputHandler
	} else {
		if caseSensitive {
			targetInputMap[inputHandler.Key] = &inputHandler
		} else {
			targetInputMap[strings.ToLower(inputHandler.Key)] = &inputHandler
		}
	}

	if len(otherInputHandlers) > 0 {
		for _, otherInputHandler := range otherInputHandlers {
			if otherInputHandler.Key == "" {
				othersInputHandler = &otherInputHandler
			} else {
				if caseSensitive {
					targetInputMap[otherInputHandler.Key] = &otherInputHandler
				} else {
					targetInputMap[strings.ToLower(otherInputHandler.Key)] = &otherInputHandler
				}
			}
		}
	}

	for {
		fmt.Printf(prefix)

		inputText := ``
		if _, scanErr := fmt.Scanln(&inputText); scanErr == nil {
			if !caseSensitive {
				inputText = strings.ToLower(inputText)
			}

			if inputHandlerData, exist := targetInputMap[inputText]; exist {
				ret = inputHandlerData.Handler(inputText)
			} else {
				if othersInputHandler != nil {
					ret = othersInputHandler.Handler(inputText)
				}
			}

			if ret {
				break
			}
		} else {
			fmt.Println(scanErr.Error())
			break
		}
	}

	return
}
