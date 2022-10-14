package parse

import (
	"regexp"
	"strconv"

	cf "gitlab.com/jiaoshijie/tcolors/color_format"
)

var re_hex *regexp.Regexp = regexp.MustCompile("#[0-9a-fA-F]{6}")

func ParseHex(str string) *cf.Parsed {
	values := re_hex.FindAllStringSubmatch(str, -1)
	all_index := re_hex.FindAllStringSubmatchIndex(str, -1)
	var matched []string
	var rgb []*cf.Rgb
	var index []int
	for i := range values {
		s := values[i][0]
		matched = append(matched, s)
		index = append(index, all_index[i][0])
		r, _ := strconv.ParseUint(s[1:3], 16, 8)
		g, _ := strconv.ParseUint(s[3:5], 16, 8)
		b, _ := strconv.ParseUint(s[5:7], 16, 8)
		rgb = append(rgb, cf.Rgb_new(uint8(r), uint8(g), uint8(b)))
	}

	return cf.Parsed_new(str, matched, rgb, index)
}
