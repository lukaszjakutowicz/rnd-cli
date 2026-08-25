package cmd

import (
	"flag"
	"fmt"

	"github.com/lukaszjakutowicz/rnd-cli/internal/generator"
)

func init() {
	register(command{
		name:  "string",
		short: "Generate a random string",
		run:   runString,
	})
}

func runString(args []string) error {
	fs := flag.NewFlagSet("string", flag.ContinueOnError)
	length := fs.Int("length", 16, "length of the generated string")
	lowercase := fs.Bool("lowercase", false, "use lowercase letters only")
	uppercase := fs.Bool("uppercase", false, "use uppercase letters only")
	items := itemsFlag(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	if err := validateItems(*items); err != nil {
		return err
	}

	for i := 0; i < *items; i++ {
		s, err := generator.NewString(generator.StringOptions{
			Length:        *length,
			LowercaseOnly: *lowercase,
			UppercaseOnly: *uppercase,
		})
		if err != nil {
			return fmt.Errorf("generating string: %w", err)
		}
		fmt.Println(s)
	}
	return nil
}
