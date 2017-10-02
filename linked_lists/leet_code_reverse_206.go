/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
