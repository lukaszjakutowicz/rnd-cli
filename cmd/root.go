// Package cmd wires up rnd-cli's subcommands. It stays thin: all actual
// generation logic lives in internal/generator so it can be tested without
// going through the CLI layer.
package cmd

import (
	"flag"
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
var version = "dev"

// register adds a subcommand to the CLI. Subcommand files call this from an
// init() function.
func register(c command) {
	commands = append(commands, c)
}

// maxItems is the upper bound accepted by --items.
const maxItems = 20

// itemsFlag registers the shared --items flag, used by every generator
// command to control how many values are printed. Default is 1, max is
// maxItems.
func itemsFlag(fs *flag.FlagSet) *int {
	return fs.Int("items", 1, fmt.Sprintf("number of values to generate (max %d)", maxItems))
}

// validateItems checks the value parsed by itemsFlag.
func validateItems(items int) error {
	if items <= 0 {
		return fmt.Errorf("items must be greater than 0, got %d", items)
	}
	if items > maxItems {
		return fmt.Errorf("items must not be greater than %d, got %d", maxItems, items)
	}
	return nil
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
