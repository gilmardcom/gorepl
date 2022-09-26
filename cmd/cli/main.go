package main

import (
	// "fmt"
	"errors"
	"fmt"
	"log"

	"github.com/gilmardcom/gorepl"
)

type Alaki struct {
	Name     string  `json:"Name"`
	BoolT    bool    `json:"BoolT"`
	BoolF    bool    `json:"BoolF"`
	IntNum   int     `json:"Integer"`
	FloatNum float32 `json:"Flaot"`
	Nulldata string  `json:"Null"`
}

func main() {

	gorepl.Clear()

	log.Println("started...")

	gorepl.Config.PromotSeparateLine = false
	gorepl.Prompt()
	fmt.Println("what ever comes after prompt")
	gorepl.Connection()
	gorepl.Config.ConStatus = gorepl.CsConnecting
	gorepl.Connection()
	gorepl.Config.ConStatus = gorepl.CsConnected
	gorepl.Connection()
	gorepl.Config.ConStatus = gorepl.CsDisonnected
	gorepl.Connection()
	gorepl.Config.ConStatus = gorepl.CsNoConnection
	gorepl.Connection()
	gorepl.Path("/this/is/the/path/to/the/data")
	fmt.Println()
	gorepl.Path("/this/is/the/path/to/some/more/data")
	fmt.Println()
	alaki := Alaki{"Hani", true, false, 12345, 88.234, "something"}
	fmt.Println(alaki)
	gorepl.Json(alaki)

	gorepl.AddCommand("help", func(args []string) (string, error) {
		fmt.Println("this is the help for this repl: empty for now!")
		if len(args) > 0 {
			return "", errors.New("No args expected!")
		}
		return "", nil
	})

	gorepl.GoRepl()

	log.Println("... finished.")
}
