package main

import (
    "fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val int
    Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var (
        v1, v2 int
        tmp int
        NextI int
        Result *ListNode
    )
    Result = new(ListNode)
    CurP := Result
    for {
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        } else {
            v1 = 0
        }
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        } else {
            v2 = 0
        }
        tmp = v1 + v2
        if tmp >= 10 {
            CurP.Val = (tmp + NextI) % 10
            NextI = tmp / 10
        } else {
            if ( tmp + NextI) >= 10 {
                CurP.Val = (tmp + NextI) % 10
                NextI = (tmp + NextI) / 10
            } else {
                CurP.Val = tmp + NextI
                NextI = 0
            }
        }
        if l1 != nil || l2 != nil || NextI != 0 {
            CurP.Next = new(ListNode)
            CurP = CurP.Next
        } else {
            goto SUCCESS
        }
    }
    return nil
    SUCCESS:
        return Result
}


func main() {
    l1 := new(ListNode)
    l1.Val = 1
    //l1.Next = new(ListNode)
    //l1.Next.Val = 1
    //l1.Next.Next = new(ListNode)
    //l1.Next.Next.Val = 3

    l2 := new(ListNode)
    l2.Val = 9
    l2.Next = new(ListNode)
    l2.Next.Val = 9

    res := addTwoNumbers(l1, l2)

    for res != nil {
        fmt.Println(res.Val)
        res = res.Next
    }





}


