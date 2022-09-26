package gorepl

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	json "github.com/nwidger/jsoncolor"
)

func Clear() error {
	var command *exec.Cmd
	if operatingSystem := runtime.GOOS; operatingSystem == "linux" || operatingSystem == "darwin" || operatingSystem == "android" {
		command = exec.Command("clear")
	} else if operatingSystem == "windows" {
		command = exec.Command("cmd", "/c", "cls")
	} else {
		panic("Unsupported OS: " + runtime.GOOS)
	}
	command.Stdout = os.Stdout
	return command.Run()
}

func Error(err string) {
	Config.ColorError.Print("Error: " + err)
}

func Errorln(err string) {
	Config.ColorError.Println("Error: " + err)
}

// Prompt

func Prompt() {
	if Config.EmptyLineBeforePrompt {
		fmt.Println()
	}
	Config.ColorPrompt.Print(Config.PromptValue + Config.PromptSign)
	if Config.PromotSeparateLine {
		fmt.Println()
	} else {
		fmt.Print(" ")
	}
}

// Connection status
func Connection() {
	switch Config.ConStatus {
	case CsConnecting:
		Config.ColorConnecting.Print(" Connecting... ")
		fmt.Println(" " + Config.ConHost + " @ " + strconv.Itoa(Config.ConPort) + " (" + Config.ConProtocol + ") ")
	case CsConnected:
		Config.ColorConnected.Print(" Connected ")
		fmt.Println(" " + Config.ConHost + " @ " + strconv.Itoa(Config.ConPort) + " (" + Config.ConProtocol + ") ")
	case CsDisonnected:
		Config.ColorDisconnected.Println(" Disconnected ")
	case CsNoConnection:
		// print nothing
	}
}

// Path
func Path(data string) {
	path := strings.Split(data, Config.PathDelimiter)
	for i, item := range path {
		if i%2 == 0 {
			Config.ColorPathEven.Print(item)
		} else {
			Config.ColorPathOdd.Print(item)
		}
		if i < len(path)-1 {
			Config.ColorPathDelimiter.Print(Config.PathDelimiter)
		}
	}
}

// JSON pretty data
const noPrefix = ""
const indent4 = "    "

func Json(data interface{}) {
	f := json.NewFormatter()
	f.StringColor = Config.ColorJsonString
	f.TrueColor = Config.ColorJsonTrue
	f.FalseColor = Config.ColorJsonFalse
	f.NumberColor = Config.ColorJsonNumber
	f.NullColor = Config.ColorJsonNull
	result, err := json.MarshalIndentWithFormatter(data, noPrefix, indent4, f)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(result))
	}
}
