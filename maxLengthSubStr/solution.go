package main

import "fmt"
func lengthOfLongestSubstring(s string) int {
    if s == "" {
        return 0
    }
    var maxStr []int32
    var maxL  int
    for _, v := range s {
        ll := len(maxStr)
       for j:= 0 ; j < ll ; j ++ {
           if v == maxStr[j] {
               if len(maxStr) > maxL {
                   maxL = len(maxStr)
               }
               maxStr = maxStr[j+1:]
               break
               //maxStr = append(maxStr, v)
           }
       }
       maxStr = append(maxStr, v)
    }
    if len(maxStr) > maxL {
        maxL = len(maxStr)
    }
    return maxL
}

func main() {
    r:= lengthOfLongestSubstring("abcb")
    fmt.Printf("solution: %d", r)
}
