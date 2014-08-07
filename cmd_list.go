package main

import (
	"bufio"
	"fmt"
	"github.com/gonuts/commander"
	"io"
	"os"
	"strings"
)

func make_cmd_list(filename string) *commander.Command {
	cmd_list := func(cmd *commander.Command, args []string) error {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		br := bufio.NewReader(f)
		n := 1
		for {
			b, _, err := br.ReadLine()
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
			line := string(b)
			if strings.HasPrefix(line, "-") {
				fmt.Printf("\u2611 %03d: %s\n", n, strings.TrimSpace(string(line[1:])))
			} else {
				fmt.Printf("\u2610 %03d: %s\n", n, strings.TrimSpace(line))
			}
			n++

		}
		return nil
	}

	return &commander.Command{
		Run:       cmd_list,
		UsageLine: "list [options]",
		Short:     "show list index",
	}
}
