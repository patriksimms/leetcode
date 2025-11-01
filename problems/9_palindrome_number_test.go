package problems

import "testing"

func TestIsPalindromeExample1(t *testing.T) {
	x := 121
	expected := true

	result := isPalindrome(x)

	if result != expected {
		t.Errorf("isPalindrome(%d) = %v; want %v", x, result, expected)
	}
}

func TestIsPalindromeExample2(t *testing.T) {
	x := -121
	expected := false

	result := isPalindrome(x)

	if result != expected {
		t.Errorf("isPalindrome(%d) = %v; want %v", x, result, expected)
	}
}

func TestIsPalindromeExample3(t *testing.T) {
	x := 10
	expected := false

	result := isPalindrome(x)

	if result != expected {
		t.Errorf("isPalindrome(%d) = %v; want %v", x, result, expected)
	}
}
