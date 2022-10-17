package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.com/jiaoshijie/tcolors/parse"
	tc "gitlab.com/jiaoshijie/tcolors/true_colors"

	"golang.org/x/term"
)

func main() {
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		log.Println("Please run `tcolors` in an interactive terminal")
		return
	}

	if len(os.Args) < 2 {
		handle_file(os.Stdin)
	} else {
		if file, err := os.Open(os.Args[1]); err == nil {
			handle_file(file)
		} else {
			log.Println(err)
		}
	}
}

func handle_file(fh *os.File) {
	reader := bufio.NewReader(fh)
	var line string
	var err error

	for true {
		if line, err = reader.ReadString('\n'); err != nil {
			break
		}
		line = strings.TrimSuffix(line, "\n")

		cf := parse.ParseHex(line)
		fmt.Println(tc.Colored_line(cf))
	}
}
