package main

import (
	"os"

	"github.com/bon-ami/eztools"
)

func main() {
	logger, err := os.OpenFile("xls2xml.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err == nil {
		defer logger.Close()
		eztools.InitLogger(logger)
	} else {
		eztools.ShowStrln("Failed to open log file")
	}
	if len(os.Args) > 1 {
		rd(os.Args[1])
	} else {
		rd("D:\\Translator.xlsx")
	}
	return
}
