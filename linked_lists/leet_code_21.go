/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sort_by_start_list(l1 *ListNode, l2 *ListNode) (*ListNode, *ListNode) {
    if l1.Val < l2.Val {
        return l1, l2
    }
    return l2, l1
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    smaller, bigger := sort_by_start_list(l1, l2)
    temp_smaller := smaller
    temp_bigger := bigger
    var temp_previous *ListNode
    for temp_smaller != nil {
        for temp_smaller != nil && temp_bigger != nil && temp_smaller.Val <= temp_bigger.Val {
            temp_previous = temp_smaller
            temp_smaller = temp_smaller.Next
        }
        if temp_bigger == nil {
            break
        } else if temp_smaller != nil {
            temp := temp_bigger
            temp_bigger = temp_bigger.Next
            temp_previous.Next = temp
            temp.Next = temp_smaller
            temp_smaller = temp_previous
        }
    }
    if temp_smaller == nil {
        temp_previous.Next = temp_bigger
    }
    return smaller
    
}