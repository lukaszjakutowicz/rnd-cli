package generator

import "testing"

func TestNewPassword_Length(t *testing.T) {
	for _, length := range []int{1, 8, 16, 64} {
		got, err := NewPassword(PasswordOptions{Length: length})
		if err != nil {
			t.Fatalf("NewPassword(Length: %d) returned unexpected error: %v", length, err)
		}
		if len(got) != length {
			t.Errorf("NewPassword(Length: %d) = %q, want length %d, got %d", length, got, length, len(got))
		}
	}
}

func TestNewPassword_InvalidLength(t *testing.T) {
	for _, length := range []int{0, -1, -10} {
		if _, err := NewPassword(PasswordOptions{Length: length}); err == nil {
			t.Errorf("NewPassword(Length: %d) expected an error, got nil", length)
		}
	}
}

func TestNewPassword_DefaultCharsetIsLettersAndDigits(t *testing.T) {
	got, err := NewPassword(PasswordOptions{Length: 200})
	if err != nil {
		t.Fatalf("NewPassword() returned unexpected error: %v", err)
	}

	var sawLetter, sawDigit bool
	for _, r := range got {
		switch {
		case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z'):
			sawLetter = true
		case r >= '0' && r <= '9':
			sawDigit = true
		default:
			t.Fatalf("NewPassword() = %q, contains unexpected character %q", got, r)
		}
	}

	if !sawLetter || !sawDigit {
		t.Errorf("NewPassword() = %q, expected a mix of letters and digits", got)
	}
}

func TestNewPassword_MixedLetters(t *testing.T) {
	got, err := NewPassword(PasswordOptions{Length: 100, MixedLetters: true})
	if err != nil {
		t.Fatalf("NewPassword() returned unexpected error: %v", err)
	}

	for _, r := range got {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			t.Fatalf("NewPassword(MixedLetters: true) = %q, contains non-letter character %q", got, r)
		}
	}
}

func TestNewPassword_OnlyDigits(t *testing.T) {
	got, err := NewPassword(PasswordOptions{Length: 100, OnlyDigits: true})
	if err != nil {
		t.Fatalf("NewPassword() returned unexpected error: %v", err)
	}

	for _, r := range got {
		if r < '0' || r > '9' {
			t.Fatalf("NewPassword(OnlyDigits: true) = %q, contains non-digit character %q", got, r)
		}
	}
}

func TestNewPassword_MixedLettersAndOnlyDigitsMutuallyExclusive(t *testing.T) {
	if _, err := NewPassword(PasswordOptions{Length: 10, MixedLetters: true, OnlyDigits: true}); err == nil {
		t.Error("NewPassword(MixedLetters: true, OnlyDigits: true) expected an error, got nil")
	}
}

func TestNewPassword_Unique(t *testing.T) {
	first, err := NewPassword(PasswordOptions{Length: 32})
	if err != nil {
		t.Fatalf("NewPassword() returned unexpected error: %v", err)
	}

	second, err := NewPassword(PasswordOptions{Length: 32})
	if err != nil {
		t.Fatalf("NewPassword() returned unexpected error: %v", err)
	}

	if first == second {
		t.Errorf("NewPassword() returned the same value twice: %q", first)
	}
}
