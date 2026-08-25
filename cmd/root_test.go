package cmd

import "testing"

func TestValidateItems(t *testing.T) {
	for _, items := range []int{1, 2, 10, 20} {
		if err := validateItems(items); err != nil {
			t.Errorf("validateItems(%d) returned unexpected error: %v", items, err)
		}
	}
}

func TestValidateItems_Invalid(t *testing.T) {
	for _, items := range []int{0, -1, -10, 21, 100} {
		if err := validateItems(items); err == nil {
			t.Errorf("validateItems(%d) expected an error, got nil", items)
		}
	}
}
