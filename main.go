package main

import (
	"os"

	"rnd-cli/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
