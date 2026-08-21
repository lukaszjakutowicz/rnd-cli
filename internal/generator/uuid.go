// Package generator contains the core, CLI-agnostic logic for producing
// random values (UUIDs, strings, passwords, ...). Keeping this logic here
// (rather than in cmd/) makes it independently unit-testable.
package generator

import "github.com/google/uuid"

// NewUUID returns a new random (version 4) UUID as a string, e.g.
// "f47ac10b-58cc-4372-a567-0e02b2c3d479".
//
// It uses github.com/google/uuid, which sources its randomness from
// crypto/rand.
func NewUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
