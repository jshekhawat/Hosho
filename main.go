package main

import (
	"os"

	"github.com/jshekhawat/hosho/repl"
)

func main() {
	r := repl.New(os.Stdin, os.Stdout)
	r.Cor()
}
