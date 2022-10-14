package parse_test

import (
	"fmt"

	"gitlab.com/jiaoshijie/tcolors/parse"
)

func Example_parse_one_line() {
	fmt.Println(parse.ParseHex("#ff00000 is red, and #00ff00 is green."))

	// Output:
	// balabal
}
