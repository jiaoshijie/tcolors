package colors

import (
	"fmt"
)

var x string = "\033[0m"

func fc8(cn int) string {
	return fmt.Sprintf("\033[%vm", cn)
}

func fc(cn int, bg bool) string {
	if bg {
		return fmt.Sprintf("\033[48;5;%vm", cn)
	}
	return fmt.Sprintf("\033[38;5;%vm", cn)
}

func Show_8colors() {
	fmt.Println("  dark     bright")
	fmt.Printf("%v%v color%-2v%v %v%v color%-2v %v\n", fc8(37), fc8(40), 0, x, fc8(37), fc8(100), 8, x)
	for i := 1; i < 8; i++ {
		fmt.Printf("%v%v color%-2v%v %v%v color%-2v %v\n", fc8(30), fc8(40+i), i, x, fc8(30), fc8(100+i), 8+i, x)
	}
}

func Show_256colors() {
	fmt.Println("Standard colors")
	for i := 0; i < 8; i++ {
		fmt.Printf("%v%6v    %v", fc(i, true), i, x)
	}
	for i := 8; i < 16; i++ {
		fmt.Printf("%v%v%6v    %v", fc(0, false), fc(i, true), i, x)
	}
	fmt.Println()

	fmt.Println("216 colors")
	for b := 16; b < 232; b += 36 {
		for i := b; i < b+18; i++ {
			fmt.Printf("%v%4v %v", fc(i, true), i, x)
		}
		for i := b + 18; i < b+36; i++ {
			fmt.Printf("%v%v%4v %v", fc(0, false), fc(i, true), i, x)
		}
		fmt.Println()
	}

	fmt.Println("Grayscale colors")
	for i := 232; i < 244; i++ {
		fmt.Printf("%v  %v  %v", fc(i, true), i, x)
	}
	for i := 244; i < 256; i++ {
		fmt.Printf("%v%v  %v  %v", fc(0, false), fc(i, true), i, x)
	}
	fmt.Println()
}

func Show_style() {
	fmt.Printf("\033[1mbold%v\n", x)
	fmt.Printf("\033[2mdim%v\n", x)
	fmt.Printf("\033[3mitalic%v\n", x)
	fmt.Printf("\033[4munderline%v\n", x)
	fmt.Printf("\033[5mslow blink%v\n", x)
	fmt.Printf("\033[6mrapid blink%v\n", x)
	fmt.Printf("\033[7mreverse%v\n", x)
	fmt.Printf("\033[8mhide%v\n", x)
	fmt.Printf("\033[9mcrossed-out%v\n", x)
}
