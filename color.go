// +build !windows

// Control Sequence Introducer  (SGR - Select Graphic Rendition)
// '\033[%d;%dm' 			where %d is SGR(0..107)
// '\033[38;2;%d;%d;%d;m' 	where %d is r;g;b (Cyan = 0;205;205)
// '\033[38;5;%dm' 			where %d is color index (0..255)
package ioterm

import (
	"fmt"
	"log"
	"os"
)

const (
	Escape = "\033["   // \x1b
	Reset  = "\033[0m" // \x1b[39;49m

	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	Black_h   = "\033[30;1m"
	Red_h     = "\033[31;1m"
	Green_h   = "\033[32;1m"
	Yellow_h  = "\033[33;1m"
	Blue_h    = "\033[34;1m"
	Magenta_h = "\033[35;1m"
	Cyan_h    = "\033[36;1m"
	White_h   = "\033[37;1m"

	Black_l   = "\033[90m"
	Red_l     = "\033[91m"
	Green_l   = "\033[92m"
	Yellow_l  = "\033[93m"
	Blue_l    = "\033[94m"
	Magenta_l = "\033[95m"
	Cyan_l    = "\033[96m"
	White_l   = "\033[97m"

	Black_b   = "\033[40m"
	Red_b     = "\033[41m"
	Green_b   = "\033[42m"
	Yellow_b  = "\033[43m"
	Blue_b    = "\033[44m"
	Magenta_b = "\033[45m"
	Cyan_b    = "\033[46m"
	White_b   = "\033[47m"

	Black_hb   = "\033[40;1m"
	Red_hb     = "\033[41;1m"
	Green_hb   = "\033[42;1m"
	Yellow_hb  = "\033[43;1m"
	Blue_hb    = "\033[44;1m"
	Magenta_hb = "\033[45;1m"
	Cyan_hb    = "\033[46;1m"
	White_hb   = "\033[47;1m"

	Black_lb   = "\033[100m"
	Red_lb     = "\033[101m"
	Green_lb   = "\033[102m"
	Yellow_lb  = "\033[103m"
	Blue_lb    = "\033[104m"
	Magenta_lb = "\033[105m"
	Cyan_lb    = "\033[106m"
	White_lb   = "\033[107m"
)

var std = log.New(os.Stderr, "", log.LstdFlags)

func Info(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[INFO]%s  %s\n", Green_h, Reset, fmt.Sprint(v...)))
}
func Note(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[NOTE]%s  %s\n", Blue_h, Reset, fmt.Sprint(v...)))
}
func Warn(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[WARN]%s  %s\n", Yellow_h, Reset, fmt.Sprint(v...)))
}
func Error(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[ERRO]%s  %s\n", Red_h, Reset, fmt.Sprint(v...)))
}
func Fatal(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[FATA]%s  %s\n", Red, Reset, fmt.Sprint(v...)))
}
func Infof(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[INFO]%s  %s\n", Green_h, Reset, fmt.Sprintf(format, v...)))
}
func Notef(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[NOTE]%s  %s\n", Blue_h, Reset, fmt.Sprintf(format, v...)))
}
func Warnf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[WARN]%s  %s\n", Yellow_h, Reset, fmt.Sprintf(format, v...)))
}
func Errorf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[ERRO]%s  %s\n", Red_h, Reset, fmt.Sprintf(format, v...)))
}
func Fatalf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[FATA]%s  %s\n", Red, Reset, fmt.Sprintf(format, v...)))
}

func N(n int) string {
	return fmt.Sprintf("\033[38;5;%dm", n)
}
func Nb(n int) string {
	return fmt.Sprintf("\033[48;5;%dm", n)
}
func RGB(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}
func RGBb(r, g, b int) string {
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

func Black_f() {
	fmt.Print("\033[30m")
}
func Red_f() {
	fmt.Print("\033[31m")
}
func Green_f() {
	fmt.Print("\033[32m")
}
func Yellow_f() {
	fmt.Print("\033[33m")
}
func Blue_f() {
	fmt.Print("\033[34m")
}
func Magenta_f() {
	fmt.Print("\033[35m")
}
func Cyan_f() {
	fmt.Print("\033[36m")
}
func White_f() {
	fmt.Print("\033[37m")
}
func End() {
	fmt.Print("\033[0m")
}

func Esc(e string) {
	fmt.Print("\033[" + e)
}
