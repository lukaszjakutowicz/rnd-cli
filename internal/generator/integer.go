package generator

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// IntOptions controls the range for NewInt.
type IntOptions struct {
	Min int64
	Max int64
}

// NewInt returns a random integer in the inclusive range [Min, Max].
func NewInt(opts IntOptions) (int64, error) {
	if opts.Min > opts.Max {
		return 0, fmt.Errorf("min must not be greater than max, got min=%d max=%d", opts.Min, opts.Max)
	}

	span := big.NewInt(opts.Max - opts.Min + 1)
	n, err := rand.Int(rand.Reader, span)
	if err != nil {
		return 0, fmt.Errorf("generating random integer: %w", err)
	}

	return opts.Min + n.Int64(), nil
}
