package cmd

import (
	"flag"
	"fmt"

	"github.com/lukaszjakutowicz/rnd-cli/internal/generator"
)

func init() {
	register(command{
		name:  "integer",
		short: "Generate a random integer in a range",
		run:   runInteger,
	})
}

func runInteger(args []string) error {
	fs := flag.NewFlagSet("integer", flag.ContinueOnError)
	min := fs.Int64("min", 0, "minimum value (inclusive)")
	max := fs.Int64("max", 100, "maximum value (inclusive)")
	items := itemsFlag(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	if err := validateItems(*items); err != nil {
		return err
	}

	for i := 0; i < *items; i++ {
		n, err := generator.NewInt(generator.IntOptions{
			Min: *min,
			Max: *max,
		})
		if err != nil {
			return fmt.Errorf("generating integer: %w", err)
		}
		fmt.Println(n)
	}
	return nil
}
