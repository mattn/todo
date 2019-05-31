package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

const (
	doneMark1 = "\u2610"
	doneMark2 = "\u2611"
)

func makeCmdList(filename string) *commander.Command {
	cmdList := func(cmd *commander.Command, args []string) error {
		nflag := cmd.Flag.Lookup("n").Value.Get().(bool)
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
				if !nflag {
					fmt.Printf("%s %03d: %s\n", doneMark2, n, strings.TrimSpace(line[1:]))
				}
			} else {
				fmt.Printf("%s %03d: %s\n", doneMark1, n, strings.TrimSpace(line))
			}
			n++

		}
		return nil
	}

	flg := *flag.NewFlagSet("list", flag.ExitOnError)
	flg.Bool("n", false, "only not done")

	return &commander.Command{
		Run:       cmdList,
		UsageLine: "list [options]",
		Short:     "show list index",
		Flag:      flg,
	}
}
