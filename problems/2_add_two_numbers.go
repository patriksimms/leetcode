package problems

import (
	"fmt"
	"strconv"
	"math/big"
)

// 2. Add Two Numbers
// Medium
// Topics
// premium lock icon
// Companies
// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example 1:

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
// Example 2:

// Input: l1 = [0], l2 = [0]
// Output: [0]
// Example 3:

// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]

// Constraints:

// The number of nodes in each linked list is in the range [1, 100].
// 0 <= Node.val <= 9
// It is guaranteed that the list represents a number that does not have leading zeros.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	firstNumString := ""
	secondNumString := ""

	for i := l1; i != nil; i = i.Next {
		firstNumString = fmt.Sprintf("%d%s", i.Val, firstNumString)
	}

	for i := l2; i != nil; i = i.Next {
		secondNumString = fmt.Sprintf("%d%s", i.Val, secondNumString)
	}

	// we can assume only numbers are in the string because we are putting them there

	firstInt := new(big.Int)
	firstInt.SetString(firstNumString, 10)
	secondInt := new(big.Int)
	secondInt.SetString(secondNumString, 10)

	sum := new(big.Int)

	sum = sum.Add(firstInt, secondInt)

	resString := sum.String()

	var resList *ListNode
	for i := 0; i < len(resString); i++ {
		digitInt, _ := strconv.Atoi(string(resString[i]))
		newNode := &ListNode{Val: digitInt, Next: resList}
		resList = newNode
	}

	return resList

}
