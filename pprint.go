// Based on https://github.com/k0kubun/pp by Takashi Kokubun (c) 2015 The MIT License (MIT)
// +build !windows

package tty

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	c "github.com/Joker/tty/color"
	bb "github.com/valyala/bytebufferpool"
)

const (
	Bool        = c.Cyan_h
	Integer     = c.Blue_h
	Float       = c.Magenta_h
	String      = c.Red
	StringQuote = c.Red_h
	EscapedChar = c.Magenta_h
	FieldName   = c.Yellow
	Pointer     = c.Blue_h
	Nil         = c.Cyan_h
	Time        = c.Blue_h
	StructName  = c.Green
	ObjectLen   = c.Blue

	printMapTypes   = true
	printBufferSize = 1024

	end  = c.Reset
	snil = Nil + "nil" + end
)

var (
	MaxDepth = -1
	TabSize  = 2
)

type printer struct {
	// *bytes.Buffer
	*bb.ByteBuffer
	tw *tabwriter.Writer

	depth   int
	value   reflect.Value
	visited map[uintptr]bool
}

func newPrinter(object interface{}) *printer {
	var (
		// buffer = bytes.NewBufferString("")
		buffer = bb.Get()
		writer = new(tabwriter.Writer)
	)
	writer.Init(buffer, TabSize, 0, 1, ' ', 0)

	return &printer{
		// Buffer: buffer,
		ByteBuffer: buffer,
		tw:         writer,

		depth:   0,
		value:   reflect.ValueOf(object),
		visited: map[uintptr]bool{},
	}
}

func (p *printer) format(object interface{}) string {
	pp := newPrinter(object)
	pp.depth = p.depth
	pp.visited = p.visited
	if value, ok := object.(reflect.Value); ok {
		pp.value = value
	}
	return pp.String()
}

func (p *printer) String() string {
	switch p.value.Kind() {
	case reflect.Bool:
		fmt.Fprintf(p.tw, "%s%s%s", Bool, p.raw(), end)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(p.tw, "%s%s%s", Integer, p.raw(), end)
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(p.tw, "%s%s%s", Float, p.raw(), end)
	case reflect.String:
		p.printString()
	case reflect.Map:
		p.printMap()
	case reflect.Struct:
		p.printStruct()
	case reflect.Array, reflect.Slice:
		p.printSlice()
	case reflect.Interface:
		p.printInterface()
	case reflect.Ptr:
		p.printPtr()
	case reflect.Func:
		fmt.Fprintf(p.tw, "%s {...}", p.typeString())
	case reflect.Chan:
		fmt.Fprintf(p.tw, "(%s)(%s%#v%s)", p.typeString(), Pointer, p.value.Pointer(), end)
	case reflect.UnsafePointer:
		fmt.Fprintf(p.tw, "%s(%s%#v%s)", p.typeString(), Pointer, p.value.Pointer(), end)
	case reflect.Invalid:
		fmt.Fprint(p.tw, snil)
	default:
		fmt.Fprint(p.tw, p.raw())
	}
	p.tw.Flush()
	// return p.Buffer.String()
	return p.ByteBuffer.String()
}

func (p *printer) raw() string {
	// Some value causes panic when Interface() is called.
	switch p.value.Kind() {
	case reflect.Bool:
		return fmt.Sprintf("%#v", p.value.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%#v", p.value.Int())
	case reflect.Uint, reflect.Uintptr:
		return fmt.Sprintf("%#v", p.value.Uint())
	case reflect.Uint8:
		return fmt.Sprintf("0x%02x", p.value.Uint())
	case reflect.Uint16:
		return fmt.Sprintf("0x%04x", p.value.Uint())
	case reflect.Uint32:
		return fmt.Sprintf("0x%08x", p.value.Uint())
	case reflect.Uint64:
		return fmt.Sprintf("0x%016x", p.value.Uint())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", p.value.Float())
	case reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("%#v", p.value.Complex())
	default:
		return fmt.Sprintf("%#v", p.value.Interface())
	}
}

//

func (p *printer) printString() {
	quoted := strconv.Quote(p.value.String())
	quoted = quoted[1 : len(quoted)-1]

	fmt.Fprintf(p.tw, "%s\"%s", StringQuote, end)
	for len(quoted) > 0 {
		pos := strings.IndexByte(quoted, '\\')
		if pos == -1 {
			fmt.Fprintf(p.tw, "%s%s%s", String, quoted, end)
			break
		}
		if pos != 0 {
			fmt.Fprintf(p.tw, "%s%s%s", String, quoted[0:pos], end)
		}

		n := 1
		switch quoted[pos+1] {
		case 'x': // "\x00"
			n = 3
		case 'u': // "\u0000"
			n = 5
		case 'U': // "\U00000000"
			n = 9
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9': // "\000"
			n = 3
		}
		fmt.Fprintf(p.tw, "%s%s%s", EscapedChar, quoted[pos:pos+n+1], end)
		quoted = quoted[pos+n+1:]
	}
	fmt.Fprintf(p.tw, "%s\"%s", StringQuote, end)
}

func (p *printer) printMap() {
	if p.value.Len() == 0 {
		fmt.Fprintf(p.tw, "%s{}", p.typeString())
		return
	}

	if p.visited[p.value.Pointer()] {
		fmt.Fprintf(p.tw, "%s{...}", p.typeString())
		return
	}
	p.visited[p.value.Pointer()] = true

	if printMapTypes {
		fmt.Fprintf(p.tw, "%s{\n", p.typeString())
	} else {
		fmt.Fprintln(p.tw, "{")
	}
	p.indented(func() {
		keys := p.value.MapKeys()
		for i := 0; i < p.value.Len(); i++ {
			value := p.value.MapIndex(keys[i])
			fmt.Fprintf(p.tw, "%s%s:\t%s,\n", p.indent(), p.format(keys[i]), p.format(value))
		}
	})
	fmt.Fprintf(p.tw, "%s}", p.indent())
}

func (p *printer) printStruct() {
	if p.value.Type().String() == "time.Time" {
		p.printTime()
		return
	}

	if p.value.NumField() == 0 {
		fmt.Fprint(p.tw, p.typeString()+"{}")
		return
	}

	fmt.Fprintln(p.tw, p.typeString()+"{")
	p.indented(func() {
		for i := 0; i < p.value.NumField(); i++ {
			field := fmt.Sprintf("%s%s%s", FieldName, p.value.Type().Field(i).Name, end)
			value := p.value.Field(i)
			fmt.Fprintf(p.tw, "%s%s:\t%s,\n", p.indent(), field, p.format(value))
		}
	})
	fmt.Fprintf(p.tw, "%s}", p.indent())
}

func (p *printer) printTime() {
	if !p.value.CanInterface() {
		fmt.Fprint(p.tw, "(unexported time.Time)")
		return
	}

	tm := p.value.Interface().(time.Time)
	fmt.Fprintf(p.tw,
		"%s%d%s-%s%02d%s-%s%02d%s %s%02d%s:%s%02d%s:%s%02d%s %s%s%s",
		Time, tm.Year(), end,
		Time, tm.Month(), end,
		Time, tm.Day(), end,
		Time, tm.Hour(), end,
		Time, tm.Minute(), end,
		Time, tm.Second(), end,
		Time, tm.Location().String(), end,
	)
}

func (p *printer) printSlice() {
	if p.value.Kind() == reflect.Slice && p.value.IsNil() {
		fmt.Fprintf(p.tw, "%s(%s)", p.typeString(), snil)
		return
	}
	if p.value.Len() == 0 {
		fmt.Fprintf(p.tw, "%s{}", p.typeString())
		return
	}

	if p.value.Kind() == reflect.Slice {
		if p.visited[p.value.Pointer()] {
			// Stop travarsing cyclic reference
			fmt.Fprintf(p.tw, "%s{...}", p.typeString())
			return
		}
		p.visited[p.value.Pointer()] = true
	}

	// Fold a large buffer
	if p.value.Len() > printBufferSize {
		fmt.Fprintf(p.tw, "%s{...}", p.typeString())
		return
	}

	fmt.Fprintln(p.tw, p.typeString()+"{")
	p.indented(func() {
		groupsize := 0
		switch p.value.Type().Elem().Kind() {
		case reflect.Uint8:
			groupsize = 16
		case reflect.Uint16:
			groupsize = 8
		case reflect.Uint32:
			groupsize = 8
		case reflect.Uint64:
			groupsize = 4
		}

		if groupsize > 0 {
			for i := 0; i < p.value.Len(); i++ {
				// indent for new group
				if i%groupsize == 0 {
					fmt.Fprint(p.tw, p.indent())
				}
				// slice element
				fmt.Fprintf(p.tw, "%s,", p.format(p.value.Index(i)))
				// space or newline
				if (i+1)%groupsize == 0 || i+1 == p.value.Len() {
					fmt.Fprint(p.tw, "\n")
				} else {
					fmt.Fprint(p.tw, " ")
				}
			}
		} else {
			for i := 0; i < p.value.Len(); i++ {
				fmt.Fprintf(p.tw, "%s%s,\n", p.indent(), p.format(p.value.Index(i)))
			}
		}
	})
	fmt.Fprintf(p.tw, "%s}", p.indent())
}

func (p *printer) printInterface() {
	e := p.value.Elem()
	if e.Kind() == reflect.Invalid {
		fmt.Fprint(p.tw, snil)
	} else if e.IsValid() {
		fmt.Fprint(p.tw, p.format(e))
	} else {
		fmt.Fprintf(p.tw, "%s(%s)", p.typeString(), snil)
	}
}

func (p *printer) printPtr() {
	if p.visited[p.value.Pointer()] {
		fmt.Fprintf(p.tw, "&%s{...}", p.elemTypeString())
		return
	}
	if p.value.Pointer() != 0 {
		p.visited[p.value.Pointer()] = true
	}

	if p.value.Elem().IsValid() {
		fmt.Fprintf(p.tw, "&%s", p.format(p.value.Elem()))
	} else {
		fmt.Fprintf(p.tw, "(%s)(%s)", p.typeString(), snil)
	}
}

//

func (p *printer) typeString() string {
	return p.colorizeType(p.value.Type().String())
}

func (p *printer) elemTypeString() string {
	return p.colorizeType(p.value.Elem().Type().String())
}

func (p *printer) colorizeType(t string) string {
	prefix := ""

	if p.match(t, `^\[\].+$`) {
		prefix = "[]"
		t = t[2:]
	}

	if p.match(t, `^\[\d+\].+$`) {
		num := regexp.MustCompile(`\d+`).FindString(t)
		prefix = fmt.Sprintf("[%s%s%s]", ObjectLen, num, end)
		t = t[2+len(num):]
	}

	if p.match(t, `^[^\.]+\.[^\.]+$`) {
		ts := strings.Split(t, ".")
		t = fmt.Sprintf("%s.%s%s%s", ts[0], StructName, ts[1], end)
	} else {
		t = fmt.Sprintf("%s%s%s", StructName, t, end)
	}
	return prefix + t
}

func (p *printer) match(text, exp string) bool {
	return regexp.MustCompile(exp).MatchString(text)
}

func (p *printer) indented(proc func()) {
	p.depth++
	if MaxDepth == -1 || p.depth <= MaxDepth {
		proc()
	}
	p.depth--
}

func (p *printer) indent() string {
	return strings.Repeat("\t", p.depth)
}
