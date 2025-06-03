package logger

import (
	"fmt"
	"time"
)

const (
	Green = "\033[32m"
	Reset = "\033[0m"
)

func Log(msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	formattedMsg := fmt.Sprintf("%s[+]%s %s : %s\n", Green, Reset, timestamp, msg)
	fmt.Print(formattedMsg)
}
