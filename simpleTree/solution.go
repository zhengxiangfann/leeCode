package main

import (
    "fmt"

)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 二叉树的最大深度
func maxDepth1(root *TreeNode) int {
    if root == nil {
        return 0
    }
    lm := maxDepth(root.Left)
    rm := maxDepth(root.Right)
    if lm >= rm {
        return lm+1
    }
    return rm+1
}


func maxDepth(root *TreeNode) int {
    n := 0
    last, nlast := root, root
    var var_l []*TreeNode
    last = root
    nlast = nil
    var_l = append(var_l, last)
    if root == nil {
        return 0
    }
    for len(var_l) > 0 {
        tn := var_l[0]
        var_l = var_l[1:]
        if tn.Left != nil {
            var_l = append(var_l, tn.Left)
            nlast = tn.Left
        }
        if tn.Left != nil {
            var_l = append(var_l, tn.Right)
            nlast = tn.Right
        }
        if (tn == last) {
            last = nlast
            n++
        }
    }
    return n
}


// 错误的
func isValidBST1(root *TreeNode) bool {
    if root == nil{
        return true
    }
    if root.Left != nil && root.Right != nil {
        if root.Val > root.Left.Val && root.Val < root.Right.Val {
            return true
        } else {
            return false
        }
    } else if root.Left ==nil && root.Right !=nil{
        if root.Val < root.Right.Val {
            return true
        } else {
            return false
        }
    } else if root.Left != nil && root.Right == nil {
        if root.Val > root.Left.Val {
            return true
        } else {
            return false
        }
    }
    return isValidBST(root.Left) && isValidBST(root.Right)
}


func isValidBST(root *TreeNode) bool {
    var res []int
    visitTree(root, &res)
    fmt.Println(res)
    for i:=0; i< len(res)-1; i ++ {
        if res[i] > res[i+1] {
            return false
        }
    }
    return true
}


// 中顺遍历二叉树
func visitTree(root *TreeNode, res *[]int){
    if root == nil {
        return
    }
    visitTree(root.Left, res)
    *res = append(*res, root.Val)
    visitTree(root.Right, res)
}

// 判断是否为镜象二叉树
func isMirror(l, r *TreeNode) bool{
    if l == nil && r == nil {
        return true
    }
    if l == nil || r == nil {
        return false
    }
    return l.Val == r.Val && isMirror(l.Right, r.Left) && isMirror(l.Left, r.Right)
}

func isSymmetric(root *TreeNode) bool {
    return isMirror(root, root)
}
//
////   二叉树的层次遍历
//func levelOrder(root *TreeNode) [][]int {
//
//
//
//
//
//
//}


func levelOrder1(root *TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    dfsHelper(root, 0, &result)
    fmt.Println(result)
    return result
}

func dfsHelper(node *TreeNode, level int, result *[][]int) {
    if node == nil {
        return
    }
    if len(*result) < level + 1 {
        *result = append(*result, make([]int, 0))
    }
    (*result)[level] = append((*result)[level], node.Val)
    dfsHelper(node.Left, level + 1, result)
    dfsHelper(node.Right, level + 1, result)
}
//
//func merge(nums1 []int, m int, nums2 []int, n int)  {
//
//    tmp := make([]int, m+n)
//    for i:=0; i <n ; i ++ {
//        for j:=i+1; j < m ; j ++ {
//            if nums2[i] <
//        }
//    }
//}

func createTree(nums []int,l, r int) *TreeNode {
    if l <= r {
        mid := (l + r ) / 2
        node := new(TreeNode)
        node.Val = nums[mid]
        node.Left = createTree(nums, l, mid-1)
        node.Right = createTree(nums, mid+1, r)
        return node
    } else {
        return nil
    }
}


func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    return createTree(nums, 0, len(nums)-1)
}





func main() {
    var nums = []int{-10, -3, 0, 5, 9}
    r := sortedArrayToBST(nums)
    var res []int
    visitTree(r, &res)
    fmt.Println(res)

}
