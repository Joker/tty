// +build !windows

package tty

import (
	"fmt"
	"log"
	"os"

	c "github.com/Joker/tty/color"
	json "github.com/tidwall/pretty"
)

const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = Ldate | Ltime // initial values for the standard logger

	rs = c.Reset
)

var std = log.New(os.Stderr, "", log.Lshortfile)

// log.SetFlags(log.LstdFlags | log.Lshortfile)
func SetFlags(flag int) {
	std = log.New(os.Stderr, "", flag)
}

func Info(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[INFO]%s  %s%s\n", c.Green_h, rs, fmt.Sprint(v...), rs))
}
func Note(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[NOTE]%s  %s%s\n", c.Blue_h, rs, fmt.Sprint(v...), rs))
}
func Warn(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[WARN]%s  %s%s\n", c.Yellow_h, rs, fmt.Sprint(v...), rs))
}
func Error(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[ERRO]%s  %s%s\n", c.Red_h, rs, fmt.Sprint(v...), rs))
}
func Fatal(v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[FATA]%s  %s%s\n", c.Red, rs, fmt.Sprint(v...), rs))
}
func Infof(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[INFO]%s  %s%s\n", c.Green_h, rs, fmt.Sprintf(format, v...), rs))
}
func Notef(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[NOTE]%s  %s%s\n", c.Blue_h, rs, fmt.Sprintf(format, v...), rs))
}
func Warnf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[WARN]%s  %s%s\n", c.Yellow_h, rs, fmt.Sprintf(format, v...), rs))
}
func Errorf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[ERRO]%s  %s%s\n", c.Red_h, rs, fmt.Sprintf(format, v...), rs))
}
func Fatalf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(" %s[FATA]%s  %s%s\n", c.Red, rs, fmt.Sprintf(format, v...), rs))
}

//

func PPrint(in interface{}) {
	std.Output(2, newPrinter(in).String())
}
func PPrintln(v ...interface{}) {
	var out = make([]interface{}, len(v))
	for i, object := range v {
		out[i] = newPrinter(object).String()
	}
	std.Output(2, fmt.Sprintln(out...))
}
func PPstr(in interface{}) string {
	return newPrinter(in).String()
}

//

func JSON(in string) {
	std.Output(2, string(json.Color(json.Pretty([]byte(in)), nil)))
}
func ByJSON(in []byte) {
	std.Output(2, string(json.Color(json.Pretty(in), nil)))
}
func JSONstr(in string) string {
	return string(json.Color(json.Pretty([]byte(in)), nil))
}
func JSONbt(in []byte) string {
	return string(json.Color(json.Pretty(in), nil))
}

//

func Print(v ...interface{}) {
	fmt.Print(v...)
}
func Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
func Println(v ...interface{}) {
	fmt.Println(v...)
}
