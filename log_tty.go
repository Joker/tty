// +build !windows

package tty

import (
	"fmt"
	"log"
	"os"

	c "github.com/Joker/tty/color"
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

func Print(in interface{}) {
	std.Output(2, newPrinter(in).String())
	std.Output(2, "\n")
}
func Println(v ...interface{}) {
	var out = make([]interface{}, len(v))
	for i, object := range v {
		out[i] = newPrinter(object).String()
	}
	std.Output(2, fmt.Sprintln(out...))
}
