package main

import (
    "fmt"
)

//Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil || n <= 0 {
        return head
    }
    hP, curP, nP := head, head, head
    for i := 0; i < n; i++ {
        if curP == nil {
            return head
        } else {
            curP = curP.Next
        }
    }

    if curP == nil {
        return hP.Next
    }

    for curP.Next != nil {
        curP = curP.Next
        nP = nP.Next
    }
    //nP.Val = nP.Next.Val
    nP.Next = nP.Next.Next
    return hP
}

func reverseList1(head *ListNode) *ListNode {
   if head == nil || head.Next == nil {
       return head
   }
   n := reverseList(head.Next)
   head.Next.Next = head
   head.Next = nil
   return n
}

func reverseList(head *ListNode) *ListNode {
    var pre, cur, next  *ListNode
    cur = head
    for cur !=nil {
        next = cur.Next
        cur.Next =pre
        pre = cur
        cur =next
    }
    return pre
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    if l1.Val < l2.Val {
        l1.Next = mergeTwoLists(l1.Next, l2)
        return l1
    } else {
        l2.Next = mergeTwoLists(l1, l2.Next)
        return l2
    }
}

// 回文链表
func isPalindrome(head *ListNode) bool {
    var stack []int
    mid, cur := head, head
    n := 0
    for cur.Next != nil {
        if cur.Next.Next != nil {
            cur = cur.Next.Next
            mid = mid.Next
        } else {
            cur = cur.Next
        }
        n++
    }

    for mid.Next != nil {
        mid = mid.Next
        stack = append(stack, mid.Val)
    }
    i := len(stack)-1
    for i>=0 {
        if head.Val != stack[i] {
            return false
        }
        i--
    }
    return true
}



// 判断链表是否有环形
func hasCycle(head *ListNode) bool {
    fast, slow := head, head
    for slow != nil &&fast !=nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}

func hasCycle1(head *ListNode) bool {
    pm := map[*ListNode]int{}
    pos := 0
    for head != nil {
        if _, ok := pm[head]; ok {
            return true
        }
        pm[head] = pos
        head = head.Next
        pos++
    }
    return false
}

func main(){
    var l1 = new(ListNode)
    l1.Val=1
    l1.Next = nil
    l1.Next = new(ListNode)
    l1.Next.Val = 2
    l1.Next.Next = nil

    var l2 = new(ListNode)
    l2.Val=1
    l2.Next = new(ListNode)
    l2.Next.Val = 2
    l2.Next.Next = new(ListNode)
    l2.Next.Next.Val=3
    l2.Next.Next.Next= new(ListNode)
    l2.Next.Next.Next.Val=4
    l2.Next.Next.Next.Next=nil
    //l2.Next.Next.Next.Next=new(ListNode)
    //l2.Next.Next.Next.Next.Val=5
    //l2.Next.Next.Next.Next.Next=nil

    r := hasCycle(l1)
    fmt.Println(r)
}
