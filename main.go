package main

import (
	"os"

	"github.com/lukaszjakutowicz/rnd-cli/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
