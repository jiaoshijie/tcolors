package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gitlab.com/jiaoshijie/tcolors/parse"
	tc "gitlab.com/jiaoshijie/tcolors/true_colors"
)

func main() {
	ec := tc.New()
	reader := bufio.NewReader(os.Stdin)
	var line string
	var err error

	for true {
		if line, err = reader.ReadString('\n'); err != nil {
			break
		}
		line = strings.TrimSpace(line)

		cf := parse.ParseHex(line)
		fmt.Println(ec.Colored_line(cf))
	}
}
