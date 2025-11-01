package problems
// 9. Palindrome Number
// Easy
// Topics
// premium lock icon
// Companies
// Hint
// Given an integer x, return true if x is a palindrome, and false otherwise.

 

// Example 1:

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
// Example 2:

// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:

// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
 

// Constraints:

// -231 <= x <= 231 - 1
 

// Follow up: Could you solve it without converting the integer to a string?

func isPalindrome(x int) bool {
    temp := x
	tempResult := 0

	for i := 0; temp > 0; i++ {
		digit := temp % 10
		tempResult = (tempResult * 10) + digit
		temp = (temp / 10)
	}

	return tempResult == x
}
