package tty

import (
	"fmt"
	"testing"
)

func TestMsg(t *testing.T) {
	fmt.Println("")

	Info("Info")
	Note("Note")
	Warn("Warn")
	Error("Error")
	Fatal("Fatal")

	//

	fmt.Println("")
	fmt.Println(`SetFlags(0):`)

	SetFlags(0)

	Infof("Infof(\"%%d\", %d)", 42)
	Notef("Notef(\"%%d\", %d)", 42)

	//

	fmt.Println("")
	fmt.Println(`SetFlags(LstdFlags | Lshortfile):`)

	SetFlags(LstdFlags | Lshortfile)

	Warnf("Warnf(\"%%d\", %d)", 42)
	Errorf("Errorf(\"%%d\", %d)", 42)
	Fatalf("Fatalf(\"%%d\", %d)", 42)

	fmt.Println("")
}
