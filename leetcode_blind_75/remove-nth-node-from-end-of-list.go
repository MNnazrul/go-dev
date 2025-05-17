/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    sz := 0
    cur := head 
    for cur != nil {
        sz++
        cur = cur.Next
    }
    if n == sz {
        head = head.Next
        return head
    }

    sz -= n
    sz--
    ref := head
    for sz > 0 {
        sz--
        ref = ref.Next
    }
    ref.Next = ref.Next.Next
    return head
}

func init() {
    debug.SetMemoryLimit(1)
}