package gorepl

import (
	"bufio"
	"fmt"
	// "fmt"
	"os"
	"strings"
)

var stop = false

func StopRepl() {
	stop = true
}

type CommandHandler func([]string) (string, error)

var commandsMap = map[string]CommandHandler{}

func AddCommand(cmdName string, cmdHanlder CommandHandler) {
	commandsMap[cmdName] = cmdHanlder
}

func GoRepl() {

	reader := bufio.NewScanner(os.Stdin)
	for !stop {
		Prompt()
		for reader.Scan() && !stop {
			userInput := cleanInput(reader.Text())
			cmd, args := extractCommand(userInput)
			if handler, ok := commandsMap[cmd]; ok {
				result, err := handler(args)
				if err != nil {
					Errorln(err.Error())
				}
				if result != "" {
					fmt.Println(result)
				}
			} else {
				hanldeInvalidCommand(cmd, args)
			}
			Prompt()
		}
	}
}

func hanldeInvalidCommand(cmd string, args []string) {
	if len(args) > 0 {
		Errorln(fmt.Sprintf("cmd '%s' with args=%s is not supported", cmd, args))
	} else {
		Errorln(fmt.Sprintf("cmd '%s' is not supported", cmd))
	}
}

func cleanInput(userInput string) string {
	output := strings.TrimSpace(userInput)
	output = strings.ToLower(output)
	return output
}

func extractCommand(userInput string) (string, []string) {
	fields := strings.Fields(userInput)
	end := len(fields)
	return fields[0], fields[1:end]
}
