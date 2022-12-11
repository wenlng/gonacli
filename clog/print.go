package clog

import (
	"fmt"
	"time"
)

const (
	color_red = uint8(iota + 91)
	color_green
	color_yellow
	color_blue
	color_magenta

	info = "[INFO]"
	trac = "[TRAC]"
	erro = "[ERRO]"
	warn = "[WARN]"
	succ = "[SUCC]"
)

func Trace(a ...interface{}) {
	prefix := yellow(trac)
	fmt.Println(formatLog(prefix), a)
}

func Info(a ...interface{}) {
	prefix := blue(info)
	fmt.Println(formatLog(prefix), a)
}

func Success(a ...interface{}) {
	prefix := green(succ)
	fmt.Println(formatLog(prefix), a)
}

func Warning(a ...interface{}) {
	prefix := magenta(warn)
	fmt.Println(formatLog(prefix), a)
}

func Error(a ...interface{}) {
	prefix := red(erro)
	fmt.Println(formatLog(prefix), a)
}

func red(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_red, s)
}

func green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_green, s)
}

func yellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_yellow, s)
}

func blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_blue, s)
}

func magenta(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_magenta, s)
}

func formatLog(prefix string) string {
	return time.Now().Format("2006/01/02 15:04:05") + " " + prefix + " "
}
