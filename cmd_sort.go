package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gonuts/commander"
)

func make_cmd_sort(filename string) *commander.Command {
	cmd_sort := func(cmd *commander.Command, args []string) error {
		if len(args) != 0 {
			cmd.Usage()
			return nil
		}
		var bottom bytes.Buffer
		w, err := os.Create(filename + "_")
		if err != nil {
			return err
		}
		defer w.Close()
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		br := bufio.NewReader(f)
		for {
			b, _, err := br.ReadLine()
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
			line := string(b)
			if !strings.HasPrefix(line, "-") {
				_, err = fmt.Fprintf(&bottom, "%s\n", line)
				if err != nil {
					return err
				}
			} else {
				_, err = fmt.Fprintf(w, "%s\n", line)
				if err != nil {
					return err
				}
			}
		}
		bottom.WriteTo(w)
		err = os.Remove(filename)
		if err != nil {
			return err
		}
		return os.Rename(filename+"_", filename)
	}

	return &commander.Command{
		Run:       cmd_sort,
		UsageLine: "sort",
		Short:     "sorts done to the top and undone to the bottom",
	}
}
