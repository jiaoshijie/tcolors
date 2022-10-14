package true_colors

import (
	"fmt"

	cf "gitlab.com/jiaoshijie/tcolors/color_format"
)

const g_base int = 38
const g_reset string = "\033[0m"
const g_x string = "\033[0m"

func fc(cn int, rgb *cf.Rgb) string {
	var gg string
	r, g, b := rgb.Get()
	if g_base < cn {
		if g > 0x7f {
			gg = fmt.Sprintf("\033[38;2;0;0;0m")
		} else {
			gg = fmt.Sprintf("\033[38;2;255;255;255m")
		}
	}
	return fmt.Sprintf("%v\033[%v;2;%v;%v;%vm", gg, cn, r, g, b)
}

func Scolor(rgb *cf.Rgb, str string, opt ...bool) string {
	cn := g_base
	if len(opt) > 0 && opt[0] {
		cn += 10
	}

	return fmt.Sprintf("%v%v%v", fc(cn, rgb), str, g_x)
}

func Colored_line(p *cf.Parsed) string {
	var ret string
	if len(p.Index) == 0 {
		return p.Line
	}

	var s int = 0
	for i := range p.Index {
		ret += p.Line[s:p.Index[i]]
		s = p.Index[i] + 7
		ret += Scolor(p.Rgb[i], p.Matched[i], true)
	}
	ret += p.Line[s:]

	return ret
}
