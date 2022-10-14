package colors

import (
	"fmt"
	"log"
	"strings"
)

type Eight_color struct {
	base     int
	_8colors map[string]int
	reset    string
	x        string
}

func New() *Eight_color {
	return &Eight_color{
		base: 30,
		_8colors: map[string]int{
			"black":   0,
			"red":     1,
			"green":   2,
			"yellow":  3,
			"blue":    4,
			"magenta": 5,
			"cyan":    6,
			"white":   7,
		},
		reset: "\033[0m",
		x:     "\033[0m",
	}
}

func (c *Eight_color) Scolor(color string, str string, opt ...bool) string {
	// opt[0] -> bright
	// opt[1] -> background
	c_index, ok := c._8colors[strings.ToLower(color)]
	if !ok {
		log.Printf("%v is not in 8 colors.\n", color)
		return str
	}
	cn := c.base + c_index
	if len(opt) > 0 && opt[0] {
		cn += 60
	}
	if len(opt) > 1 && opt[1] {
		cn += 10
	}
	return fmt.Sprintf("%v%v%v", c.fc(cn), str, c.x)
}

func (c *Eight_color) fc(cn int) string {
	return fmt.Sprintf("\033[%vm", cn)
}

func (c *Eight_color) Show_all_colors() {
	fmt.Println("dark colors:")
	for i := 0; i < 8; i++ {
		fmt.Printf("\t%v\\033[%vm%v\n", c.fc(c.base+i), c.base+i, c.x)
	}
	fmt.Println("bright colors:")
	for i := 0; i < 8; i++ {
		fmt.Printf("\t%v\\033[%vm%v\n", c.fc(c.base+i+60), c.base+i+60, c.x)
	}
}
