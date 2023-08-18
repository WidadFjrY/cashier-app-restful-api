package helper

import (
	"fmt"
	"github.com/fatih/color"
)

func TextSuccess(status string, message interface{}) {
	OK := color.New(color.FgBlack).Add(color.BgGreen).SprintfFunc()

	fmt.Printf("%s %s\n", OK(status), message)
}

func TextFailed(status string, message interface{}) {
	OK := color.New(color.FgBlack).Add(color.BgRed).SprintfFunc()

	fmt.Printf("%s %s\n", OK(status), message)
}
