package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		x := 0
		y := 0

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		result := x + y + carry
		newNode := &ListNode{
			Val: (result % 10),
		}
		carry = (result / 10)

		current.Next = newNode
		current = current.Next
	}
	return head
}

func createListWithArray(arr []int) *ListNode {
    if len(arr) == 0 {
        return nil
    }
    head := &ListNode{Val: arr[0]}
    current := head
    for _, val := range arr[1:] {
        new_node := &ListNode{Val: val}
        current.Next = new_node
        current = current.Next
    }
    return head
}

func main() {
	l1_arr := []int{2, 4, 3}
	l1 := createListWithArray(l1_arr)
	l2_arr := []int{5, 6, 4}
	l2 := createListWithArray(l2_arr)
	teste := addTwoNumbers(l1, l2)
	fmt.Println(teste)
}