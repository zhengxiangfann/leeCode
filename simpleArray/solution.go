package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
    i := m - 1
    j := n - 1
    index := m + n - 1
    for index >= 0 {
        if i < 0 {
            //第一个数组已经遍历完毕
            nums1[index] = nums2[j]
            index--
            j--
        } else if j < 0 {
            //第二个数组已经遍历完毕
            nums1[index] = nums1[i]
            index--
            i--
        } else if (nums1[i] > nums2[j] && i >= 0) {
            //比较两个数组元素
            nums1[index] = nums1[i]
            index--
            i--
        } else if (nums1[i] <= nums2[j] && j >= 0) {
            //比较两个数组元素
            nums1[index] = nums2[j]
            index--
            j--
        }
    }
}





func main() {
    //
    var num1 = []int{1,2,3,0,0,0}
    var num2 = []int{4,5,6}

    //var num1 = []int{0}
    //var num2 = []int{1}
    merge(num1, 3, num2, 3)
    fmt.Println(num1)
}
