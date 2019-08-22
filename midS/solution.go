package main

import (
    "fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var mid float64
    //if len(nums1) == 0 {
    //    if len(nums2) %2 ==0 {
    //        mid = float64(nums2[len(nums2)/2] + nums2[len(nums2)/2-1]) / 2.0
    //    } else {
    //        mid = float64(nums2[len(nums2)/2])
    //    }
    //    return mid
    //}
    //if len(nums2) == 0 {
    //    if len(nums1) %2 ==0 {
    //        mid = float64(nums1[len(nums1)/2] + nums1[len(nums1)/2-1]) / 2.0
    //    } else {
    //        mid = float64(nums1[len(nums1)/2])
    //    }
    //    return mid
    //}

    var newS []int
    var iFirst, iSecond int
    for {
        if iFirst < len(nums1) && iSecond < len(nums2) {
            if nums1[iFirst] <= nums2[iSecond] {
                newS = append(newS, nums1[iFirst])
                iFirst++
            } else {
                newS = append(newS, nums2[iSecond])
                iSecond++
            }
        } else {
            break
        }
    }

    for {
        if iFirst < len(nums1) {
            newS = append(newS, nums1[iFirst])
            iFirst++
        } else {
            break
        }
    }
    for {
        if iSecond < len(nums2) {
            newS = append(newS, nums2[iSecond])
            iSecond++
        } else {
            break
        }
    }
    fmt.Println(newS)
    if len(newS)%2 == 0 {
        mid = float64(newS[len(newS)/2]+newS[len(newS)/2-1]) / 2.0
    } else {
        mid = float64(newS[len(newS)/2])
    }
    return mid
}

func main() {

    s1 := []int{1, 2}
    s2 := []int{3, 4}

    r := findMedianSortedArrays(s2, s1)
    fmt.Println(r)
}
