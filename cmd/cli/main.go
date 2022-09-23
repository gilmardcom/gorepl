package main

import (
	"fmt"
	"log"

	"github.com/gilmardcom/gorepl/util/printer"
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
	log.Println("started...")
	printer.Clear()
	printer.Disonnected()
	printer.Connecting("TCP", "192.168.10.31", "8080")
	printer.Connected("TCP", "192.168.10.31", "8080")
	printer.Path("/this/is/the/path/to/the/data")
	printer.Prompt("|   prompt $ ")
	fmt.Print("|     text     |")
	fmt.Println("|     text     |")
	fmt.Println("|     text     |")
	alaki := Alaki{"Hani", true, false, 12345, 88.234, "something"}
	fmt.Println(alaki)
	printer.Json(alaki)

	// // Print with default helper functions
	// color.Cyan("Prints text in cyan.")

	// // A newline will be appended automatically
	// color.Blue("Prints %s in blue.", "text")

	// // These are using the default foreground colors
	// color.Red("We have red")
	// color.Magenta("And many others ..")

	// // Create a new color object
	// c := color.New(color.FgCyan).Add(color.Underline)
	// c.Println("Prints cyan text with an underline.")

	// // Or just add them to New()
	// d := color.New(color.FgCyan, color.Bold)
	// d.Printf("This prints bold cyan %s\n", "too!.")

	// // Mix up foreground and background colors, create new mixes!
	// red := color.New(color.FgRed)

	// boldRed := red.Add(color.Bold)
	// boldRed.Println("This will print text in bold red.")

	// whiteBackground := red.Add(color.BgWhite)
	// whiteBackground.Println("Red text with white background.")

	log.Println("... finished.")
}
