package gorepl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

var stop = false

func StopRepl() {
	stop = true
}

type CommandHandler func([]string) (string, error)

var commandsMap = map[string]CommandHandler{}

type commandInfo struct {
	commandName  string
	commandAlias string
	helpComment  string
}

var commandsInfo = map[string]commandInfo{}

func AddCommand(cmdName string, cmdAlias string, helpComment string, cmdHanlder CommandHandler) {
	if cmdName == "" {
		fmt.Println("Command cannot be empty")
		return
	}
	_, ok := commandsMap[cmdName]
	if ok {
		fmt.Println("Command ", cmdName, "already exists")
		return
	}
	commandsMap[cmdName] = cmdHanlder
	if cmdAlias != "" {
		_, ok = commandsMap[cmdAlias]
		if ok {
			fmt.Println("Command alias", cmdAlias, "already exists")
			return
		}
		commandsMap[cmdAlias] = cmdHanlder
	}
	commandsInfo[cmdName] = commandInfo{
		commandName:  cmdName,
		commandAlias: cmdAlias,
		helpComment:  helpComment,
	}
}

func GoRepl() {
	reader := bufio.NewScanner(os.Stdin)
	Prompt()
	for !stop && reader.Scan() {
		userInput := cleanInput(reader.Text())
		if userInput != "" {
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
		}
		if !stop {
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
	return output
}

func extractCommand(userInput string) (string, []string) {
	fields := strings.Fields(userInput)
	end := len(fields)
	return fields[0], fields[1:end]
}

func DefaultHelp() {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', 0)
	fmt.Println()
	fmt.Fprintln(w, "Name:\tAlias:\tDescription:")
	fmt.Fprintln(w, "-----\t------\t------------")
	for _, cmdInfo := range commandsInfo {
		helpLine := fmt.Sprintf("%s\t%s\t%s", cmdInfo.commandName, cmdInfo.commandAlias, cmdInfo.helpComment)
		fmt.Fprintln(w, helpLine)
	}
	w.Flush()
}
