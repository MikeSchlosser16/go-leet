/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{} // Begin building big list
    head := result // Points to head of list we are building, keeps track for return
    for l1 != nil && l2 != nil{
        // compare first elements
        if l1.Val < l2.Val {
            // L1 is smaller, append and move its next
            result.Next = l1
            l1 = l1.Next
        }else{
            // L2 is smaller, append and move its next
            result.Next = l2
            l2 = l2.Next
        }
        result = result.Next
    }

    // Handle case where element lists dont match, append the rest of the elements left onto the end
    // ie -- 1,4,5
    //    -- 1,3,7,8,9
    //    -- loop gives us: 1,3,4,5, then need to append 7,8,9
    if l1 == nil {
        result.Next = l2
    } else {
        result.Next = l1
    }

    return head.Next
}
