package generator

import "testing"

func TestNewString_Length(t *testing.T) {
	for _, length := range []int{1, 8, 16, 64} {
		got, err := NewString(StringOptions{Length: length})
		if err != nil {
			t.Fatalf("NewString(Length: %d) returned unexpected error: %v", length, err)
		}
		if len(got) != length {
			t.Errorf("NewString(Length: %d) = %q, want length %d, got %d", length, got, length, len(got))
		}
	}
}

func TestNewString_InvalidLength(t *testing.T) {
	for _, length := range []int{0, -1, -10} {
		if _, err := NewString(StringOptions{Length: length}); err == nil {
			t.Errorf("NewString(Length: %d) expected an error, got nil", length)
		}
	}
}

func TestNewString_LowercaseOnly(t *testing.T) {
	got, err := NewString(StringOptions{Length: 100, LowercaseOnly: true})
	if err != nil {
		t.Fatalf("NewString() returned unexpected error: %v", err)
	}

	for _, r := range got {
		if r < 'a' || r > 'z' {
			t.Fatalf("NewString(LowercaseOnly: true) = %q, contains non-lowercase character %q", got, r)
		}
	}
}

func TestNewString_UppercaseOnly(t *testing.T) {
	got, err := NewString(StringOptions{Length: 100, UppercaseOnly: true})
	if err != nil {
		t.Fatalf("NewString() returned unexpected error: %v", err)
	}

	for _, r := range got {
		if r < 'A' || r > 'Z' {
			t.Fatalf("NewString(UppercaseOnly: true) = %q, contains non-uppercase character %q", got, r)
		}
	}
}

func TestNewString_DefaultCharsetIsMixedCase(t *testing.T) {
	// Generate a long string so both cases are overwhelmingly likely to
	// appear at least once, and confirm no unexpected characters slip in.
	got, err := NewString(StringOptions{Length: 200})
	if err != nil {
		t.Fatalf("NewString() returned unexpected error: %v", err)
	}

	var sawLower, sawUpper bool
	for _, r := range got {
		switch {
		case r >= 'a' && r <= 'z':
			sawLower = true
		case r >= 'A' && r <= 'Z':
			sawUpper = true
		default:
			t.Fatalf("NewString() = %q, contains unexpected character %q", got, r)
		}
	}

	if !sawLower || !sawUpper {
		t.Errorf("NewString() = %q, expected a mix of lowercase and uppercase characters", got)
	}
}

func TestNewString_LowercaseAndUppercaseMutuallyExclusive(t *testing.T) {
	if _, err := NewString(StringOptions{Length: 10, LowercaseOnly: true, UppercaseOnly: true}); err == nil {
		t.Error("NewString(LowercaseOnly: true, UppercaseOnly: true) expected an error, got nil")
	}
}

func TestNewString_Unique(t *testing.T) {
	first, err := NewString(StringOptions{Length: 32})
	if err != nil {
		t.Fatalf("NewString() returned unexpected error: %v", err)
	}

	second, err := NewString(StringOptions{Length: 32})
	if err != nil {
		t.Fatalf("NewString() returned unexpected error: %v", err)
	}

	if first == second {
		t.Errorf("NewString() returned the same value twice: %q", first)
	}
}
