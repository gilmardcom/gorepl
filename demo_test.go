package gorepl

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

type SampleJson struct {
	Name     string  `json:"Name"`
	BoolT    bool    `json:"BoolT"`
	BoolF    bool    `json:"BoolF"`
	IntNum   int     `json:"Integer"`
	FloatNum float32 `json:"Flaot"`
	Nulldata string  `json:"Null"`
}

var value = 0
var sampleJson = SampleJson{"Hani", true, false, 12345, 88.234, "something"}

func TestDemo(t *testing.T)() {
	
	fmt.Println("test started...")

	Config.ConStatus = CsDisonnected

	AddCommand("quit", "q", "quits the repl", func(args []string) (string, error) {
		StopRepl()
		return "", nil
	})

	AddCommand("clear", "c", "clears the screen", func(args []string) (string, error) {
		Clear()
		return "", nil
	})

	AddCommand("help", "h", "shows default help", func(args []string) (string, error) {
		DefaultHelp()
		return "", nil
	})

	AddCommand("print", "p", "prints the current value", func(args []string) (string, error) {
		fmt.Println("Current value:", value)
		return "", nil
	})

	AddCommand("reset", "r", "resets the current value", func(args []string) (string, error) {
		value = 0
		return "", nil
	})

	AddCommand("add", "a", "adds the given value to the current value, e.g. add 10", func(args []string) (string, error) {
		if len(args) < 1 {
			return "At least one arg needed", errors.New("invalid number of args")
		}
		for _, item := range args {
			givenValue, _ := strconv.Atoi(item)
			value += givenValue
		}
		return "", nil
	})

	AddCommand("sub", "s", "substracts the given value from the current value, e.g. sub 10", func(args []string) (string, error) {
		if len(args) < 1 {
			return "At least one arg needed", errors.New("invalid number of args")
		}
		for _, item := range args {
			givenValue, _ := strconv.Atoi(item)
			value -= givenValue
		}
		return "", nil
	})

	AddCommand("json", "j", "prints a sample json", func(args []string) (string, error) {
		Json(sampleJson)
		return "", nil
	})

	GoRepl()
}
