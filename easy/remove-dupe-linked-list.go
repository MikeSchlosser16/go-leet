/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    initial := head
    for head != nil && head.Next != nil{
        for head.Next != nil && (head.Val == head.Next.Val){
            head.Next = head.Next.Next // 'Skip' the next element
        }
        head = head.Next
    }
    return initial
}


// Maybe a more clear solution?
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }

    list := head
    for list.Next != nil{
        if list.Val == list.Next.Val{
            list.Next = list.Next.Next
        }else{
            list = list.Next
        }
    }
    return head
}
