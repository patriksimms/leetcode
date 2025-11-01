package problems

import "testing"

func TestExampleOne(t *testing.T) {
	list1 := &ListNode{Val: 2}
	list1.Next = &ListNode{Val: 4}
	list1.Next.Next = &ListNode{Val: 3}

	list2 := &ListNode{Val: 5}
	list2.Next = &ListNode{Val: 6}
	list2.Next.Next = &ListNode{Val: 4}

	exp := &ListNode{Val: 7}
	exp.Next = &ListNode{Val: 0}
	exp.Next.Next = &ListNode{Val: 8}

	res := addTwoNumbers(list1, list2)

	if res.Val != exp.Val || res.Next.Val != exp.Next.Val || res.Next.Next.Val != exp.Next.Next.Val {
		t.Errorf("addTwoNumbers did return the wrong result! exp: %d %d %d act: %d %d %d", exp.Val, exp.Next.Val, exp.Next.Next.Val, res.Val, res.Next.Val, res.Next.Next.Val)
	}
}

func TestExampleTwo(t *testing.T) {
	list1 := &ListNode{Val: 0}

	list2 := &ListNode{Val: 0}

	exp := &ListNode{Val: 0}

	res := addTwoNumbers(list1, list2)

	if res.Val != exp.Val {
		t.Errorf("addTwoNumbers did return the wrong result! exp: %d  act: %d ", exp.Val, res.Val)
	}
}

func TestExampleThree(t *testing.T) {
	l1 := &ListNode{Val: 9}
	l1.Next = &ListNode{Val: 9}
	l1.Next.Next = &ListNode{Val: 9}
	l1.Next.Next.Next = &ListNode{Val: 9}
	l1.Next.Next.Next.Next = &ListNode{Val: 9}
	l1.Next.Next.Next.Next.Next = &ListNode{Val: 9}
	l1.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9}

	l2 := &ListNode{Val: 9}
	l2.Next = &ListNode{Val: 9}
	l2.Next.Next = &ListNode{Val: 9}
	l2.Next.Next.Next = &ListNode{Val: 9}

	exp := &ListNode{Val: 8}
	exp.Next = &ListNode{Val: 9}
	exp.Next.Next = &ListNode{Val: 9}
	exp.Next.Next.Next = &ListNode{Val: 9}
	exp.Next.Next.Next.Next = &ListNode{Val: 0}
	exp.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	exp.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	exp.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 1}

	res := addTwoNumbers(l1, l2)

	for exp != nil && res != nil {
		if exp.Val != res.Val {
			t.Errorf("addTwoNumbers returned wrong result: expected %d but got %d", exp.Val, res.Val)
		}
		exp = exp.Next
		res = res.Next
	}

	if exp != nil || res != nil {
		t.Errorf("addTwoNumbers produced lists of different lengths")
	}
}
