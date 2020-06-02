package tty

import (
	"testing"
)

func TestMsg(t *testing.T) {
	Println("")

	Info("Info")
	Note("Note")
	Warn("Warn")
	Error("Error")
	Fatal("Fatal")

	//

	Println("")
	Println(`SetFlags(0):`)

	SetFlags(0)

	Infof("Infof(\"%%d\", %d)", 42)
	Notef("Notef(\"%%d\", %d)", 42)

	//

	Println("")
	Println(`SetFlags(LstdFlags | Lshortfile):`)

	SetFlags(LstdFlags | Lshortfile)

	Warnf("Warnf(\"%%d\", %d)", 42)
	Errorf("Errorf(\"%%d\", %d)", 42)
	Fatalf("Fatalf(\"%%d\", %d)", 42)

	//

	Println("")

	JSON(`{"name":  {"first":"Tom","last":"Anderson"}, "heve id":true,  "age":37, "children": 
	["Sara","Alex","Jack"], "fav.movie": "Deer Hunter", "movie ID":[1462126142327524116,1462126151243273148,1462126159038469292,1462126167757990750,1462126174877941853,1462126183389466506,1462126190275371345,1462126197339694532,1462126205865246332,1462126214076983985,1462126221652502387,1462126230673792530,1462126239850196408,1462126248585289564, 1462126041996725853,1462126050268282258,1462126059509103606,1462126067142189827,1462126075203195168,1462126083847111149,1462126092031197531,1462126101266954638,1462126108409799353,1462126117233898681,
	1462126124761412722,1462126134569626749], "friends": [
	{"first": "Janet", "last": 
	"Murphy", "age": 44} ]}`)

	Println("")
}
