package cmd

import (
	"flag"
	"fmt"

	"rnd-cli/internal/generator"
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
	if err := fs.Parse(args); err != nil {
		return err
	}

	p, err := generator.NewPassword(generator.PasswordOptions{
		Length:       *length,
		MixedLetters: *mixedLetters,
		OnlyDigits:   *onlyDigits,
	})
	if err != nil {
		return fmt.Errorf("generating password: %w", err)
	}

	fmt.Println(p)
	return nil
}
