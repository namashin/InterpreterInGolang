package main

import (
	"Monkey/repl"
	"os"
)

func main() {
	repl.Start(os.Stdout, os.Stdout)
}
