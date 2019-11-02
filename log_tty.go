// +build !windows

// Control Sequence Introducer  (SGR - Select Graphic Rendition)
// '\033[%d;%dm' 			where %d is SGR(0..107)
// '\033[38;2;%d;%d;%d;m' 	where %d is r;g;b (Cyan = 0;205;205)
// '\033[38;5;%dm' 			where %d is color index (0..255)

// bash:
// for COLOR in {1..255}; do echo -en "\e[38;5;${COLOR}m${COLOR} "; done; echo;

package ioterm

import (
	"fmt"
	"log"
	"os"

	c "github.com/Joker/ioterm/color"
)

const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

var std = log.New(os.Stderr, "", log.Lshortfile)

// log.SetFlags(log.LstdFlags | log.Lshortfile)
func SetFlags(flag int) {
	std = log.New(os.Stderr, "", flag)
}

func Info(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[INFO]%s  %s%s\n", c.Green_h, c.Reset, fmt.Sprint(v...), c.Reset))
}
func Note(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[NOTE]%s  %s%s\n", c.Blue_h, c.Reset, fmt.Sprint(v...), c.Reset))
}
func Warn(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[WARN]%s  %s%s\n", c.Yellow_h, c.Reset, fmt.Sprint(v...), c.Reset))
}
func Error(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[ERRO]%s  %s%s\n", c.Red_h, c.Reset, fmt.Sprint(v...), c.Reset))
}
func Fatal(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[FATA]%s  %s%s\n", c.Red, c.Reset, fmt.Sprint(v...), c.Reset))
}
func Infof(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[INFO]%s  %s%s\n", c.Green_h, c.Reset, fmt.Sprintf(format, v...), c.Reset))
}
func Notef(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[NOTE]%s  %s%s\n", c.Blue_h, c.Reset, fmt.Sprintf(format, v...), c.Reset))
}
func Warnf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[WARN]%s  %s%s\n", c.Yellow_h, c.Reset, fmt.Sprintf(format, v...), c.Reset))
}
func Errorf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[ERRO]%s  %s%s\n", c.Red_h, c.Reset, fmt.Sprintf(format, v...), c.Reset))
}
func Fatalf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[FATA]%s  %s%s\n", c.Red, c.Reset, fmt.Sprintf(format, v...), c.Reset))
}

// N for 256 colors, where n is color index (0..255)
func N(n int) string {
	return fmt.Sprintf("\033[38;5;%dm", n)
}

// Nb for 256 background colors, where n is color index (0..255)
func Nb(n int) string {
	return fmt.Sprintf("\033[48;5;%dm", n)
}
func RGB(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}
func RGBb(r, g, b int) string {
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

func NowBlack() {
	fmt.Print("\033[30m")
}
func NowRed() {
	fmt.Print("\033[31m")
}
func NowGreen() {
	fmt.Print("\033[32m")
}
func NowYellow() {
	fmt.Print("\033[33m")
}
func NowBlue() {
	fmt.Print("\033[34m")
}
func NowMagenta() {
	fmt.Print("\033[35m")
}
func NowCyan() {
	fmt.Print("\033[36m")
}
func NowWhite() {
	fmt.Print("\033[37m")
}
func End() {
	fmt.Print("\033[0m")
}

func Esc(e string) {
	fmt.Print("\033[" + e)
}
