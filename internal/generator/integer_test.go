package generator

import "testing"

func TestNewInt_DefaultRange(t *testing.T) {
	for i := 0; i < 100; i++ {
		got, err := NewInt(IntOptions{Min: 0, Max: 100})
		if err != nil {
			t.Fatalf("NewInt() returned unexpected error: %v", err)
		}
		if got < 0 || got > 100 {
			t.Fatalf("NewInt() = %d, want value in [0, 100]", got)
		}
	}
}

func TestNewInt_CustomRange(t *testing.T) {
	min, max := int64(10), int64(20)
	for i := 0; i < 100; i++ {
		got, err := NewInt(IntOptions{Min: min, Max: max})
		if err != nil {
			t.Fatalf("NewInt() returned unexpected error: %v", err)
		}
		if got < min || got > max {
			t.Fatalf("NewInt() = %d, want value in [%d, %d]", got, min, max)
		}
	}
}

func TestNewInt_NegativeRange(t *testing.T) {
	min, max := int64(-50), int64(-10)
	for i := 0; i < 100; i++ {
		got, err := NewInt(IntOptions{Min: min, Max: max})
		if err != nil {
			t.Fatalf("NewInt() returned unexpected error: %v", err)
		}
		if got < min || got > max {
			t.Fatalf("NewInt() = %d, want value in [%d, %d]", got, min, max)
		}
	}
}

func TestNewInt_SingleValueRange(t *testing.T) {
	got, err := NewInt(IntOptions{Min: 5, Max: 5})
	if err != nil {
		t.Fatalf("NewInt() returned unexpected error: %v", err)
	}
	if got != 5 {
		t.Errorf("NewInt(Min: 5, Max: 5) = %d, want 5", got)
	}
}

func TestNewInt_InvalidRange(t *testing.T) {
	if _, err := NewInt(IntOptions{Min: 10, Max: 5}); err == nil {
		t.Error("NewInt(Min: 10, Max: 5) expected an error, got nil")
	}
}
