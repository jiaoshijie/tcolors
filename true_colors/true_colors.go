package true_colors

import (
	"fmt"

	cf "gitlab.com/jiaoshijie/tcolors/color_format"
)

type True_color struct {
	base  int
	reset string
	x     string
}

func New() *True_color {
	return &True_color{
		base:  38,
		reset: "\033[0m",
		x:     "\033[0m",
	}
}

func (c *True_color) fc(cn int, rgb *cf.Rgb) string {
	r, g, b := rgb.Get()
	return fmt.Sprintf("\033[%v;2;%v;%v;%vm", cn, r, g, b)
}

func (c *True_color) Scolor(rgb *cf.Rgb, str string, opt ...bool) string {
	cn := c.base
	if len(opt) > 0 && opt[0] {
		cn += 10
	}

	return fmt.Sprintf("%v%v%v", c.fc(cn, rgb), str, c.x)
}

func (c *True_color) Colored_line(p *cf.Parsed) string {
	var ret string
	if len(p.Index) == 0 {
		return p.Line
	}

	var s int = 0
	for i := range p.Index {
		ret += p.Line[s:p.Index[i]]
		s = p.Index[i] + 7
		ret += c.Scolor(p.Rgb[i], p.Matched[i], true)
	}
	ret += p.Line[s:]

	return ret
}
