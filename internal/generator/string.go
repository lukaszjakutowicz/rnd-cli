package generator

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// StringOptions controls how NewString generates its output.
type StringOptions struct {
	// Length is the number of characters to generate. Must be greater than 0.
	Length int
	// LowercaseOnly restricts the character set to a-z.
	LowercaseOnly bool
	// UppercaseOnly restricts the character set to A-Z.
	UppercaseOnly bool
}

// NewString returns a random string of opts.Length characters, sourced from
// crypto/rand. By default it draws from both lowercase and uppercase
// letters; LowercaseOnly or UppercaseOnly narrow the character set to just
// one case. Setting both LowercaseOnly and UppercaseOnly is an error, since
// it would leave no characters to choose from.
func NewString(opts StringOptions) (string, error) {
	if opts.Length <= 0 {
		return "", fmt.Errorf("length must be greater than 0, got %d", opts.Length)
	}
	if opts.LowercaseOnly && opts.UppercaseOnly {
		return "", fmt.Errorf("lowercase-only and uppercase-only are mutually exclusive")
	}

	charset := lowercaseLetters + uppercaseLetters
	switch {
	case opts.LowercaseOnly:
		charset = lowercaseLetters
	case opts.UppercaseOnly:
		charset = uppercaseLetters
	}

	max := big.NewInt(int64(len(charset)))
	result := make([]byte, opts.Length)
	for i := range result {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", fmt.Errorf("generating random string: %w", err)
		}
		result[i] = charset[n.Int64()]
	}

	return string(result), nil
}
