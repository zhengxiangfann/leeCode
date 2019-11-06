package main

import "fmt"

func isValid(s string) bool {
    if s == "" {
        return true
    }
    kuo := map[uint8]uint8{']':'[', '}':'{',')':'('}
    var tmp []uint8
    for i:=0;i < len(s); i ++ {
        if len(tmp) == 0 {
            tmp = append(tmp, s[i])
        } else {
            if tmp[len(tmp)-1] == kuo[s[i]] {
                tmp = tmp[:len(tmp)-1]
            } else {
                tmp = append(tmp, s[i])
            }
        }
    }
    if len(tmp) == 0 {
        return true
    } else {
        return false
    }
}


func missingNumber(nums []int) int {
    m := make([]int, len(nums) +1)
    for i:=0; i< len(nums); i++ {
        m[nums[i]] = 1
    }
    for i:=0 ; i< len(m) ;i ++ {
        if m[i] ==0 {
            return i
        }
    }
    return -1
}


func main() {
    var m = []int{3,0,1}
   r:=  missingNumber(m)
   fmt.Println(r)
}
