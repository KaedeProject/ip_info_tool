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

func LogVIp(msg string) {
	formattedMsg := fmt.Sprintf("%s[IP]%s %s\n", Green, Reset, msg)
	fmt.Print(formattedMsg)
}

func LogVCity(msg string) {
	formattedMsg := fmt.Sprintf("%s[City]%s %s\n", Green, Reset, msg)
	fmt.Print(formattedMsg)
}

func LogVnRegion(msg string) {
	formattedMsg := fmt.Sprintf("%s[Region]%s %s\n", Green, Reset, msg)
	fmt.Print(formattedMsg)
}

func LogVCountry(msg string) {
	formattedMsg := fmt.Sprintf("%s[Country]%s %s\n", Green, Reset, msg)
	fmt.Print(formattedMsg)
}

func LogVOrg(msg string) {
	formattedMsg := fmt.Sprintf("%s[Org]%s %s\n", Green, Reset, msg)
	fmt.Print(formattedMsg)
}
