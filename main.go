package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/PatrikOlin/monkey_interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language REPL!\n", user.Username)
	fmt.Printf("Feel free to type in commands and fuck about for a bit\n")
	repl.Start(os.Stdin, os.Stdout)

}
