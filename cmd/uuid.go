package cmd

import (
	"flag"
	"fmt"

	"github.com/lukaszjakutowicz/rnd-cli/internal/generator"
)

func init() {
	register(command{
		name:  "uuid",
		short: "Generate a random UUID (v4)",
		run:   runUUID,
	})
}

func runUUID(args []string) error {
	fs := flag.NewFlagSet("uuid", flag.ContinueOnError)
	items := itemsFlag(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	if err := validateItems(*items); err != nil {
		return err
	}

	for i := 0; i < *items; i++ {
		id, err := generator.NewUUID()
		if err != nil {
			return fmt.Errorf("generating uuid: %w", err)
		}
		fmt.Println(id)
	}
	return nil
}
