package gorepl

import (
	fcolor "github.com/fatih/color"
)

type AppConfig struct {
	PathDelimiter      string
	PromptSign         string
	ColorPrompt        *fcolor.Color
	ColorConnecting    *fcolor.Color
	ColorConnected     *fcolor.Color
	ColorDisconnected  *fcolor.Color
	ColorPathDelimiter *fcolor.Color
	ColorPathOdd       *fcolor.Color
	ColorPathEven      *fcolor.Color
	ColorJsonString    *fcolor.Color
	ColorJsonTrue      *fcolor.Color
	ColorJsonFalse     *fcolor.Color
	ColorJsonNumber    *fcolor.Color
	ColorJsonNull      *fcolor.Color
}

var Config = AppConfig{
	PathDelimiter:      "/",
	PromptSign:         "$",
	ColorPrompt:        fcolor.New(fcolor.FgHiCyan, fcolor.BgBlue),
	ColorConnecting:    fcolor.New(fcolor.FgBlack, fcolor.BgYellow),
	ColorConnected:     fcolor.New(fcolor.FgBlack, fcolor.BgGreen),
	ColorDisconnected:  fcolor.New(fcolor.FgBlack, fcolor.BgRed),
	ColorPathDelimiter: fcolor.New(fcolor.BgHiBlack),
	ColorPathOdd:       fcolor.New(fcolor.FgHiWhite, fcolor.BgHiBlack),
	ColorPathEven:      fcolor.New(fcolor.FgHiCyan, fcolor.BgHiBlack),
	ColorJsonString:    fcolor.New(fcolor.FgHiWhite),
	ColorJsonTrue:      fcolor.New(fcolor.FgGreen),
	ColorJsonFalse:     fcolor.New(fcolor.FgRed),
	ColorJsonNumber:    fcolor.New(fcolor.FgWhite),
	ColorJsonNull:      fcolor.New(fcolor.FgWhite, fcolor.Bold),
}
