package logmsg

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	red    = color.New(color.FgRed).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
	purple = color.New(color.FgMagenta).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
	gray   = color.New(color.FgHiBlack).SprintFunc()
)

const (
	warningPrefix = "[WARNING]"
	errorPrefix   = "[ERROR]"
	successPrefix = "[SUCCESS]"
	infoPrefix    = "[INFO]"
)

func format(color func(...interface{}) string, prefix string, message interface{}) string {
	return color(prefix + " " + fmt.Sprint(message))
}

func Warning(message interface{}) string {
	return format(yellow, warningPrefix, message)
}

func Error(message interface{}) string {
	return format(red, errorPrefix, message)
}

func Success(message interface{}) string {
	return format(green, successPrefix, message)
}

func Info(message interface{}) string {
	return format(blue, infoPrefix, message)
}
