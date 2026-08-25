// Package cmd wires up rnd-cli's subcommands. It stays thin: all actual
// generation logic lives in internal/generator so it can be tested without
// going through the CLI layer.
package cmd

import (
	"fmt"
	"os"
)

// command is a single rnd-cli subcommand.
type command struct {
	name  string
	short string
	run   func(args []string) error
}

var commands []command

// version is set at build time via -ldflags "-X .../cmd.version=...".
// GoReleaser fills it in with the git tag; local builds fall back to "dev".
var version = "dev"

// register adds a subcommand to the CLI. Subcommand files call this from an
// init() function.
func register(c command) {
	commands = append(commands, c)
}

// Execute parses os.Args and dispatches to the matching subcommand. It
// returns the process exit code.
func Execute() int {
	if len(os.Args) < 2 {
		printUsage()
		return 1
	}

	name := os.Args[1]
	if name == "-h" || name == "--help" {
		printUsage()
		return 0
	}
	if name == "-v" || name == "--version" {
		fmt.Println("rnd-cli version " + version)
		return 0
	}

	for _, c := range commands {
		if c.name == name {
			if err := c.run(os.Args[2:]); err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				return 1
			}
			return 0
		}
	}

	fmt.Fprintf(os.Stderr, "Error: unknown command %q\n\n", name)
	printUsage()
	return 1
}

func printUsage() {
	fmt.Println("rnd-cli - generate random things (UUIDs, strings, passwords, ...)")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  rnd-cli <command> [flags]")
	fmt.Println("  rnd-cli --version")
	fmt.Println()
	fmt.Println("Available Commands:")
	for _, c := range commands {
		fmt.Printf("  %-10s %s\n", c.name, c.short)
	}
}
