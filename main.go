package main

import (
	"os"

	"github.com/cvescan/cvescan/cmd"
)

func main() {
	defer os.Exit(0)
	cmd.Execute()
}
