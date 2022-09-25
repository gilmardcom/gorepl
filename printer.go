package gorepl

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

// Prompt

func Prompt(data string) {
	Config.ColorPrompt.Print(data + Config.PromptSign)
}

func Promptln(data string) {
	Config.ColorPrompt.Println(data + Config.PromptSign)
}

// Connection status
func Connecting(protocol string, host string, port string) {
	Config.ColorConnecting.Print(" Connecting... ")
	fmt.Println(" " + host + " @ " + port + " (" + protocol + ") ")
}

func Connected(protocol string, host string, port string) {
	Config.ColorConnected.Print(" Connected ")
	fmt.Println(" " + host + " @ " + port + " (" + protocol + ") ")
}

func Disonnected() {
	Config.ColorDisconnected.Println(" Disconnected ")
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
