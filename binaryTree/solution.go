package main

import "fmt"

//* Definition for a binary tree node.
 type TreeNode struct {
     Val   int
     Left  *TreeNode
     Right *TreeNode
 }

func preorderTraversal(root *TreeNode) []int {
    var result []int
    if root == nil {
        return result
    }

    result = append(result, root.Val)
    left_result := preorderTraversal(root.Left)
    result = append(result, left_result...)

    right_result := preorderTraversal(root.Right)
    result = append(result,right_result...)
    return result
}


func inorderTraversal(root *TreeNode) []int {
    var result []int
    if root == nil {
        return result
    }
    left_result := inorderTraversal(root.Left)
    result = append(result, left_result...)
    result = append(result, root.Val)
    right_result := inorderTraversal(root.Right)
    result = append(result,right_result...)
    return result
}

func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    } else if root.Left == nil && root.Right == nil {
        return sum == root.Val
    } else if root.Left != nil && root.Right == nil {
        sum -= root.Val
        return hasPathSum(root.Left, sum)
    } else if root.Left == nil && root.Right != nl {
        sum -= root.Val
        return hasPathSum(root.Right, sum)
    } else {
        sum -= root.Val
        return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
    }
}

func main() {
    root := new(TreeNode)
    root.Val =1
    root.Left = new(TreeNode)
    root.Left.Val = 2
    root.Right = nil
    r := inorderTraversal(root)
    fmt.Println(r)
}