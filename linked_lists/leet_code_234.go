/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func length(head *ListNode) int {
    temp := head
    counter := 0
    for temp != nil {
        temp = temp.Next
        counter++
    }
    return counter
}

func split_list(head *ListNode) (*ListNode, *ListNode) {
    list_length := length(head)
    counter := 0
    mid := list_length/2
    temp := head
    for counter != mid {
        counter++
        if counter != 1 {
            temp = temp.Next
        }
    }
    var head1, head2 *ListNode
    if list_length == 1 {
        head1 = head
        head2 = head
    } else if list_length%2 == 0 {
        head1 = temp
        head2 = temp.Next
        head1.Next = nil
    } else if list_length%2 == 1 {
        head1 = temp
        head2 = temp.Next.Next
        head1.Next.Next = nil
        head1.Next = nil
    }
    return head1, head2
}

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    } else {
        var p1, p2, p3 *ListNode
        p1 = head
        p2 = head
        p3 = head
        p2 = p1.Next
        p3 = p2.Next
        for p3 != nil {
            p2.Next = p1
            p1 = p2
            p2 = p3
            p3 = p3.Next
            
        }
        p2.Next = p1
        head.Next = nil
        head = p2
    }
    return head
    
}

func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }
    _, head2 := split_list(head)
    head2 = reverseList(head2)
    temp1 := head
    temp2 := head2
    for temp1 != nil {
        if temp1.Val != temp2.Val {
            return false
        }
        temp1 = temp1.Next
        temp2 = temp2.Next
    }
    return true
}