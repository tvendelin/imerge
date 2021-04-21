package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/tvendelin/imerge"
)

var n *imerge.Node
var spl = regexp.MustCompile(`^(\d+),\s*(\d+)$`)

const usage = `
Enter interval start and end separated by comma.
's' will dump merged intervals entered so far.
'x' will do the same as 's' and exit.
`

func dump() {
	if n == nil {
		fmt.Println("No intervals so far")
		return
	}
	fmt.Println(n.Intervals())
}

func merge(s string) error {
	var err error
	if !spl.MatchString(s) {
		return fmt.Errorf("Input %s was not undestood", s)
	}
	ss := spl.FindStringSubmatch(s)[1:]
	start, err := strconv.Atoi(ss[0])
	if err != nil {
		return err
	}
	end, err := strconv.Atoi(ss[1])
	if err != nil {
		return err
	}
	if n == nil {
		n, err = imerge.NewNode(start, end)
		return err
	}
	err = n.Merge(start, end)
	return err
}

func main() {
	fmt.Println(usage)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		ui := strings.TrimSpace(strings.ToLower(s.Text()))
		switch ui {
		case "s":
			dump()
		case "x":
			dump()
			os.Exit(0)
		default:
			err := merge(ui)
			if err != nil {
				fmt.Println(err)
				fmt.Println(usage)
			}
		}
	}
}
