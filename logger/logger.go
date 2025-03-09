package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func GetTimeDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Log in red color
func Red(message string) {
	message = fmt.Sprintf("[%s] [NEWS API] - %s", GetTimeDate(), message)
	color.New(color.FgHiRed).Println(message)
}

// Log in yellow color
func Yellow(message string) {
	message = fmt.Sprintf("[%s] [NEWS API] - %s", GetTimeDate(), message)
	color.New(color.FgHiYellow).Println(message)
}

// Log in green color
func Green(message string) {
	message = fmt.Sprintf("[%s] [NEWS API] - %s", GetTimeDate(), message)
	color.New(color.FgHiGreen).Println(message)
}
