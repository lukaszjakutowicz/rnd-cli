package cmd

import (
	"flag"
	"fmt"

	"github.com/lukaszjakutowicz/rnd-cli/internal/generator"
)

func init() {
	register(command{
		name:  "password",
		short: "Generate a random password",
		run:   runPassword,
	})
}

func runPassword(args []string) error {
	fs := flag.NewFlagSet("password", flag.ContinueOnError)
	length := fs.Int("length", 16, "length of the generated password")
	mixedLetters := fs.Bool("mixed-letters", false, "use letters only (a-zA-Z)")
	onlyDigits := fs.Bool("only-digits", false, "use digits only (0-9)")
	noSpecial := fs.Bool("no-special", false, "exclude special characters from the default charset")
	items := itemsFlag(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	if err := validateItems(*items); err != nil {
		return err
	}

	for i := 0; i < *items; i++ {
		p, err := generator.NewPassword(generator.PasswordOptions{
			Length:       *length,
			MixedLetters: *mixedLetters,
			OnlyDigits:   *onlyDigits,
			NoSpecial:    *noSpecial,
		})
		if err != nil {
			return fmt.Errorf("generating password: %w", err)
		}
		fmt.Println(p)
	}
	return nil
}
