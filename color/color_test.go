package color

import (
	"fmt"
	"testing"
)

func TestColors(t *testing.T) {
	fmt.Printf("\n%s%sBlack      %s", Blue_lb, Black, Reset)
	fmt.Printf("%sRed        %s", Red, Reset)
	fmt.Printf("%sGreen      %s", Green, Reset)
	fmt.Printf("%sYellow     %s", Yellow, Reset)
	fmt.Printf("%sBlue       %s", Blue, Reset)
	fmt.Printf("%sMagenta    %s", Magenta, Reset)
	fmt.Printf("%sCyan       %s", Cyan, Reset)
	fmt.Printf("%sWhite      %s\n", White, Reset)
	fmt.Printf("%sBlack_h    %s", Black_h, Reset)
	fmt.Printf("%sRed_h      %s", Red_h, Reset)
	fmt.Printf("%sGreen_h    %s", Green_h, Reset)
	fmt.Printf("%sYellow_h   %s", Yellow_h, Reset)
	fmt.Printf("%sBlue_h     %s", Blue_h, Reset)
	fmt.Printf("%sMagenta_h  %s", Magenta_h, Reset)
	fmt.Printf("%sCyan_h     %s", Cyan_h, Reset)
	fmt.Printf("%sWhite_h    %s\n", White_h, Reset)
	fmt.Printf("%sBlack_l    %s", Black_l, Reset)
	fmt.Printf("%sRed_l      %s", Red_l, Reset)
	fmt.Printf("%sGreen_l    %s", Green_l, Reset)
	fmt.Printf("%sYellow_l   %s", Yellow_l, Reset)
	fmt.Printf("%sBlue_l     %s", Blue_l, Reset)
	fmt.Printf("%sMagenta_l  %s", Magenta_l, Reset)
	fmt.Printf("%sCyan_l     %s", Cyan_l, Reset)
	fmt.Printf("%sWhite_l    %s\n", White_l, Reset)
	fmt.Printf("%sBlack_b    %s", Black_b, Reset)
	fmt.Printf("%sRed_b      %s", Red_b, Reset)
	fmt.Printf("%sGreen_b    %s", Green_b, Reset)
	fmt.Printf("%sYellow_b   %s", Yellow_b, Reset)
	fmt.Printf("%sBlue_b     %s", Blue_b, Reset)
	fmt.Printf("%sMagenta_b  %s", Magenta_b, Reset)
	fmt.Printf("%sCyan_b     %s", Cyan_b, Reset)
	fmt.Printf("%sWhite_b    %s\n", White_b, Reset)
	fmt.Printf("%sBlack_hb   %s", Black_hb, Reset)
	fmt.Printf("%sRed_hb     %s", Red_hb, Reset)
	fmt.Printf("%sGreen_hb   %s", Green_hb, Reset)
	fmt.Printf("%sYellow_hb  %s", Yellow_hb, Reset)
	fmt.Printf("%sBlue_hb    %s", Blue_hb, Reset)
	fmt.Printf("%sMagenta_hb %s", Magenta_hb, Reset)
	fmt.Printf("%sCyan_hb    %s", Cyan_hb, Reset)
	fmt.Printf("%sWhite_hb   %s\n", White_hb, Reset)
	fmt.Printf("%sBlack_lb   %s", Black_lb, Reset)
	fmt.Printf("%sRed_lb     %s", Red_lb, Reset)
	fmt.Printf("%sGreen_lb   %s", Green_lb, Reset)
	fmt.Printf("%sYellow_lb  %s", Yellow_lb, Reset)
	fmt.Printf("%sBlue_lb    %s", Blue_lb, Reset)
	fmt.Printf("%sMagenta_lb %s", Magenta_lb, Reset)
	fmt.Printf("%sCyan_lb    %s", Cyan_lb, Reset)
	fmt.Printf("%sWhite_lb   %s\n\n", White_lb, Reset)
}

func TestColorFunc(t *testing.T) {
	NowBlack()
	fmt.Println("NowBlack")
	End()
	NowRed()
	fmt.Println("NowRed")
	End()
	NowGreen()
	fmt.Println("NowGreen")
	End()
	NowYellow()
	fmt.Println("NowYellow")
	End()
	NowBlue()
	fmt.Println("NowBlue")
	End()
	NowMagenta()
	fmt.Println("NowMagenta")
	End()
	NowCyan()
	fmt.Println("NowCyan")
	End()
	NowWhite()
	fmt.Println("NowWhite")
	End()
	fmt.Print("\n\n")
}

func TestColorNum(t *testing.T) {
	for i := 1; i < 256; i++ {
		fmt.Printf("%s   %d\t%s", N(i), i, Reset)
		if d := i % 8; d == 0 {
			fmt.Print("\n")
		}
	}
	End()
	fmt.Print("\n\n")
	for i := 1; i < 257; i++ {
		fmt.Printf("%s %d \t%s", Nb(i), i, Reset)
		if d := i % 8; d == 0 {
			fmt.Print("\n")
		}
	}
	End()
	fmt.Print("\n\n")
	for r := 0; r <= 255; r += 22 {
		for g := 0; g <= 255; g += 22 {
			for b := 0; b <= 255; b += 22 {
				fmt.Printf("%s%s.", RGBb(r, g, b), RGB(r, b, g))
			}
		}
	}
	End()
	fmt.Print(" \n\n")
}
