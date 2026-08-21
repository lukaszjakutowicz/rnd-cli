package generator

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewUUID(t *testing.T) {
	got, err := NewUUID()
	if err != nil {
		t.Fatalf("NewUUID() returned unexpected error: %v", err)
	}

	parsed, err := uuid.Parse(got)
	if err != nil {
		t.Fatalf("NewUUID() = %q, which is not a valid UUID: %v", got, err)
	}

	if parsed.Version() != 4 {
		t.Errorf("NewUUID() version = %d, want 4", parsed.Version())
	}
}

func TestNewUUID_Unique(t *testing.T) {
	first, err := NewUUID()
	if err != nil {
		t.Fatalf("NewUUID() returned unexpected error: %v", err)
	}

	second, err := NewUUID()
	if err != nil {
		t.Fatalf("NewUUID() returned unexpected error: %v", err)
	}

	if first == second {
		t.Errorf("NewUUID() returned the same value twice: %q", first)
	}
}
