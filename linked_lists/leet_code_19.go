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
        counter++
        temp = temp.Next
    }
    return counter
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    list_length := length(head)
    if list_length == n {
        head = head.Next
        return head
    }
    from_start_index := list_length - n
    counter := 0
    temp := head
    for counter != from_start_index {
        counter++
        if counter != 1 {
            temp = temp.Next
        }
    }
    temp.Next = temp.Next.Next
    return head
}