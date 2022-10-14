package colors

import (
	"fmt"
	"log"
	"strings"
)

var g_base int = 30
var g_reset string = "\033[0m"
var g_x string = "\033[0m"
var g_8colors map[string]int = map[string]int{
	"black":   0,
	"red":     1,
	"green":   2,
	"yellow":  3,
	"blue":    4,
	"magenta": 5,
	"cyan":    6,
	"white":   7,
}

func Scolor(color string, str string, opt ...bool) string {
	// opt[0] -> bright
	// opt[1] -> background
	c_index, ok := g_8colors[strings.ToLower(color)]
	if !ok {
		log.Printf("%v is not in 8 colors.\n", color)
		return str
	}
	cn := g_base + c_index
	if len(opt) > 0 && opt[0] {
		cn += 60
	}
	if len(opt) > 1 && opt[1] {
		cn += 10
	}
	return fmt.Sprintf("%v%v%v", fc(cn), str, g_x)
}

func fc(cn int) string {
	return fmt.Sprintf("\033[%vm", cn)
}

func Show_all_colors() {
	fmt.Println("dark colors:")
	for i := 0; i < 8; i++ {
		fmt.Printf("\t%v\\033[%vm%v\n", fc(g_base+i), g_base+i, g_x)
	}
	fmt.Println("bright colors:")
	for i := 0; i < 8; i++ {
		fmt.Printf("\t%v\\033[%vm%v\n", fc(g_base+i+60), g_base+i+60, g_x)
	}
}
