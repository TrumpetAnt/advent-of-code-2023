package utils

import (
	"fmt"
	"os"
)

func IfErrorPrintAndExit(msg string, err error) {
	if err != nil {
		fmt.Println(msg)
		fmt.Println("    Reason: " + err.Error())
		os.Exit(1)
	}
}
