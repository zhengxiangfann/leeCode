package main

import "fmt"

//Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    hP, curP, nP := head, head, head
    for curP != nil {
        curP = curP.Next
        n--
        if n <= 0 {
            nP = nP.Next
        }
    }
    nP.Val = nP.Next.Val
    nP.Next = nP.Next.Next
    return hP
}

func main(){
    var l = new(ListNode)
    l.Val=1
    l.Next = new(ListNode)
    l.Next.Val = 2
    l.Next.Next = new(ListNode)
    l.Next.Next.Val = 3
    l.Next.Next.Next =nil
    //for l != nil {
    //    fmt.Print(l.Val, " ")
    //    l = l.Next
    //}
    //fmt.Println()
    r := removeNthFromEnd(l,2)
    for r != nil {
        fmt.Print(r.Val, " ")
        r = r.Next
    }
}
