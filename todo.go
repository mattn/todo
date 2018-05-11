package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	todo_filename = ".todo"
)
func get_tty_width() int {
	cmd := exec.Command("stty","size")
	cmd.Stdin = os.Stdin
	o,err := cmd.Output()
	if err != nil {
		fmt.Printf("err: %s",err);
		os.Exit(-1);
	}
	parts := strings.Split(strings.Trim(string(o),"\n")," ")
	cp,_ := strconv.Atoi(parts[1])
	return cp;
}
func print_header() {
	h := "\n Todo - %s\n";
	if os.Getenv("TODO_HEADER") != "" {
		h = os.Getenv("TODO_HEADER")
	}
  cp := get_tty_width();	
	for i:=0; i < cp - 2; i++ {
		h = "━" + h + "━";
	}
	h += "\n\n";
	dir,_ := os.Getwd();
	dir = strings.Replace(dir,os.Getenv("GOPATH"),"$GOPATH",1);
	dir = strings.Replace(dir,os.Getenv("HOME"),"~",1);
	fmt.Printf(h,dir);
}
func print_footer() {
	h := "\n";
	if os.Getenv("TODO_FOOTER") != "" {
		h = os.Getenv("TODO_FOOTER")
	}
  cp := get_tty_width();
	for i:=0; i < cp - 2; i++ {
		h += "━";
	}
	h += "\n\n";
	fmt.Printf(h);
}
func main() {
	filename := ""
	existCurTodo := false
	curDir, err := os.Getwd()
	if err == nil {
		filename = filepath.Join(curDir, todo_filename)
		_, err = os.Stat(filename)
		if err == nil {
			existCurTodo = true
		}
	}
	if !existCurTodo {
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		filename = filepath.Join(home, todo_filename)
		_,err = os.Stat(filename);
		if err != nil {
			_,err = exec.Command("touch",filename).Output();
			if err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(-1);
			}
		}
	}
	command := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "todo for cli",
	}
	command.Subcommands = []*commander.Command{
		make_cmd_list(filename),
		make_cmd_here(),
		make_cmd_add(filename),
		make_cmd_delete(filename),
		make_cmd_done(filename),
		make_cmd_undone(filename),
		make_cmd_clean(filename),
	}
	print_header()
	err = command.Dispatch(os.Args[1:])
	print_footer()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
