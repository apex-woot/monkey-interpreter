package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/apex-woot/monkey-interpreter/cmd"
	"github.com/apex-woot/monkey-interpreter/dlog"
	"github.com/apex-woot/monkey-interpreter/repl"
)

func main() {
	cfg := cmd.ParseCFG()
	fmt.Printf("%+v\n", cfg)
	dlog.Configure(cfg.Dlog)

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s! THis is a Monkey programming language!\n", user.Username)
	fmt.Print("Feel free to type in commands!\n")
	repl.Start(os.Stdin, os.Stdout)
}
