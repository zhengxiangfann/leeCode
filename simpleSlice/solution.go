package main

import (
    "fmt"
)

func removeDuplicates(nums []int) int {
    if len(nums) ==0 {
        return 0
    }
    cur := nums[0]
    for i:= 1; i< len(nums); i ++ {
        if cur == nums[i] {
            nums = append(nums[:i], nums[i+1:]...)
            i -=1
        }else{
            cur = nums[i]
        }
        fmt.Println(nums)
    }
    return len(nums)
}

func maxProfit(prices []int) int {
    return 0
}
//
//func rotate(nums []int, k int)  {
//    i := len(nums) - k
//    nums = append(nums[i:], nums[:i]...)
//    nums = nums[:]
//    fmt.Println(nums)
//}


func containsDuplicate(nums []int) bool {
    var_m := make(map[int]int)
    l := len(nums)
    s := 0
    for _,v := range nums {
        if tmp,ok := var_m[v]; ok {
            var_m[v] = tmp +1
            return true
        } else {
            var_m[v] =1
            s +=1
        }
    }
    if s == l {
        return false
    }
    return true
}


func intersect(nums1 []int, nums2 []int) []int {
    var result []int
    for i:=0; i< len(nums1); i++ {
        for j:=0; j < len(nums2); j ++ {
            if nums1[i] == nums2[j] {
                result = append(result, nums1[i])
                nums2[j] = -99
            }
        }
    }
    return result
}
func plusOne(digits []int) []int {
    l := len(digits)
    if digits[l-1] == 0 {
        digits[l-1] = 1
        return digits
    }
    //var result = []int
    carrybit := 1
    for i:= l-1; i>=0; i--{
        if digits[i] == 9 {
            if carrybit == 1{
                digits[i] = 0
                digits = append(digits[:i+1],digits[i+1:]...)
                carrybit = 1
            }
        } else {
            digits[i] = digits[i] + carrybit
            carrybit = 0
        }
    }
    if carrybit ==1 {
        result := make([]int, len(digits) +1)
        result[0] = 1
        result = append(result[:1], digits...)
        return result
    }
    return digits
}

func moveZeroes(nums []int) {
    lastIndex := len(nums)-1
    for i := 0; i < len(nums); {
        if nums[i] == 0 {
            for j := i; j < len(nums)-1; j++ {
                nums[j] = nums[j+1]
            }
            nums[len(nums)-1] = 0
            lastIndex --
        }
        if nums[i] != 0 {
            i++
        }
        if lastIndex <= i {
            break
        }
    }
}


func twoSum(nums []int, target int) []int {
    var result [2]int
    for i:=0; i< len(nums); i ++ {
        for j:=i+1; j < len(nums); j ++ {
            if nums[i] + nums[j] == target {
                result[0] = i
                result[1] = j
                return result[:]
            }
        }
    }
    return nil
}

func isValidSudoku(board [][]byte) bool {
    var row [9][10]bool
    var cow [9][10]bool
    var block [9][10]bool
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            var index = i/3*3 + j/3
            if board[i][j] != '.' {
                var tmp = board[i][j] - '0'
                if row[i][tmp] {
                    return false
                } else {
                    row[i][tmp] = true
                }
                if cow[j][tmp] {
                    return false
                } else {
                    cow[j][tmp] = true
                }
                if block[index][tmp] {
                    return false
                } else {
                    block[index][tmp] = true
                }

            }
        }

    }
    return true
}

func rotate(matrix [][]int) {
    n := len(matrix)
    for i := 0; i < len(matrix)/2; i++ {
        for j := i; j < n-1-i; j++ {
            //matrix[i][j] = matrix[n-1-j][i]
            //matrix[n-1-j][i] = matrix[n-1-j][n-1-j]
            //matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
            //matrix[j][n-1-i] = matrix[i][j]
            matrix[i][j],matrix[n-1-j][i],matrix[n-1-i][n-1-j],matrix[j][n-1-i]=matrix[n-1-j][i],matrix[n-1-i][n-1-j],matrix[j][n-1-i], matrix[i][j]
            //matrix[i][j], matrix[n-1-j][i], matrix[n-1-i][n-1-j], matrix[j][n-1-i] = matrix[n-1-j][i], matrix[n-1-j][n-1-j], matrix[j][n-1-i], matrix[i][j]
        }
    }
}


func main() {
    //var s = []int{-1,-2,-3,-4,-5}
    var s = [][]byte{
        {'5', '3', '.', '.', '7', '.', '.', '.', '.'},
        {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
        {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
        {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
        {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
        {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
        {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
        {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
        {'.', '.', '.', '.', '8', '.', '.', '7', '9'},
    }
    r := isValidSudoku(s)
    fmt.Println(r)

    var s1 = [][]int{{1,2,3},{4,5,6},{7,8,9}}
    rotate(s1)
    fmt.Println(s1)
}
