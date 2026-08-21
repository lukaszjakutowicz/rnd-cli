package cmd

import (
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
	id, err := generator.NewUUID()
	if err != nil {
		return fmt.Errorf("generating uuid: %w", err)
	}

	fmt.Println(id)
	return nil
}
