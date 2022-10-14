package color_format

type Rgb struct {
	r uint8
	g uint8
	b uint8
}

func Rgb_new(r uint8, g uint8, b uint8) *Rgb {
	return &Rgb{r: r, g: g, b: b}
}

func (rgb *Rgb) Get() (uint8, uint8, uint8) {
	return rgb.r, rgb.g, rgb.b
}

type Parsed struct {
	Line string
	Matched []string
	Rgb []*Rgb
	Index []int
}

func Parsed_new(line string, matched []string, rgb []*Rgb, index []int) *Parsed {
	return &Parsed{
		Line: line,
		Matched: matched,
		Rgb: rgb,
		Index: index,
	}
}
