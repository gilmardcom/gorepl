package printer

import (
	"fmt"
	"strings"

	tm "github.com/buger/goterm"
	"github.com/fatih/color"
	"github.com/gilmardcom/gorepl/config"
	json "github.com/nwidger/jsoncolor"
)

func Clear() {
	tm.MoveCursor(1, 1)
	tm.Clear()
	tm.Flush()
}

func Prompt(data string) {
	tm.Print(tm.Background(tm.Color(tm.Bold(data), tm.WHITE), tm.RED))
	tm.Flush()
}

func Promptln(data string) {
	tm.Println(tm.Background(tm.Color(tm.Bold(data), tm.WHITE), tm.RED))
	tm.Flush()
}

func Path(data string) {
	tm.Print(tm.Background(tm.Color(tm.Bold("Path: "), tm.WHITE), tm.BLACK))
	path := strings.Split(data, config.PathDelimiter)
	str := []string{}
	for i, item := range path {
		if i%2 == 0 {
			str = append(str, tm.Background(tm.Color(item, tm.YELLOW), tm.BLACK))
		} else {
			str = append(str, tm.Background(tm.Color(item, tm.GREEN), tm.BLACK))
		}
	}
	tm.Println(strings.Join(str, tm.Background(tm.Color(""+config.PathDelimiter+"", tm.WHITE), tm.BLACK)))
	tm.Flush()
}

func Connecting(protocol string, host string, port string) {
	tm.Print(tm.Background(tm.Color(" Connecting... ", tm.BLACK), tm.YELLOW))
	tm.Println(" " + host + " @ " + port + " (" + protocol + ") ")
	tm.Flush()

}

func Connected(protocol string, host string, port string) {
	tm.Print(tm.Background(tm.Color(" Connected ", tm.BLACK), tm.GREEN))
	tm.Println(" " + host + " @ " + port + " (" + protocol + ") ")
	tm.Flush()

}

func Disonnected() {
	tm.Println(tm.Background(tm.Color(" Disonnected ", tm.BLACK), tm.RED))
	tm.Flush()
}

func Json(data interface{}) {
	f := json.NewFormatter()
	f.StringColor = color.New(color.FgHiWhite)
	f.TrueColor = color.New(color.FgGreen)
	f.FalseColor = color.New(color.FgRed)
	f.NumberColor = color.New(color.FgWhite)
	f.NullColor = color.New(color.FgWhite, color.Bold)
	result, err := json.MarshalIndentWithFormatter(data, "", "    ", f)
	if err != nil {
		fmt.Println(err)
	} else {
		tm.Println(string(result))
		tm.Println("OK")
	}
	// newJson, err := json.MarshalIndent(data, "", "    ")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(string(newJson))
	// 	fmt.Println("OK")
	// }
	tm.Flush()
}
