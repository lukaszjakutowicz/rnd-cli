package generator

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const digits = "0123456789"
const specialChars = "!@#%^&*()-_=+[]{};:,.<>?"

// PasswordOptions controls how NewPassword generates its output.
type PasswordOptions struct {
	Length       int
	MixedLetters bool
	OnlyDigits   bool
	NoSpecial    bool
}

// NewPassword returns a random password of Length characters.
// By default it draws from mixed-case letters, digits, and special characters.
func NewPassword(opts PasswordOptions) (string, error) {
	if opts.Length <= 0 {
		return "", fmt.Errorf("length must be greater than 0, got %d", opts.Length)
	}
	if opts.MixedLetters && opts.OnlyDigits {
		return "", fmt.Errorf("mixed-letters and only-digits are mutually exclusive")
	}

	charset := lowercaseLetters + uppercaseLetters + digits + specialChars
	switch {
	case opts.MixedLetters:
		charset = lowercaseLetters + uppercaseLetters
	case opts.OnlyDigits:
		charset = digits
	case opts.NoSpecial:
		charset = lowercaseLetters + uppercaseLetters + digits
	}

	max := big.NewInt(int64(len(charset)))
	result := make([]byte, opts.Length)
	for i := range result {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", fmt.Errorf("generating random password: %w", err)
		}
		result[i] = charset[n.Int64()]
	}

	return string(result), nil
}
