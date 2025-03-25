package main

import (
	"os"

	"github.com/biograph-health/biograph-platform-interview/cmd"
)

func main() {
	cmd.Execute(os.Args[1:])
}
