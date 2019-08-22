package main

import "fmt"

//
//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
//示例 1：
//
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//示例 2：
//
//输入: "cbbd"
//输出: "bb"

func longestPalindrome(s string) string {
    var length = len(s)
    var lps  [][]bool
    for k:= 0 ; k < length ; k ++ {
        tmp := make([]bool, length)
        lps = append(lps, tmp)
    }
    var maxLen int
    var start int
    for i := 0 ; i< length ; i ++ {
        for j:= 0 ; j <= i ; j ++ {
            if i - j < 2 {
                lps[j][i] = s[i] == s[j]
            } else {
                lps[j][i] = lps[j + 1][i - 1] && (s[i] == s[j])
            }

            if lps[j][i] && (i - j + 1) > maxLen {
                maxLen = i - j + 1;
                start = j;
            }
        }
    }
    return s[start: start + maxLen]
}


func main() {
    r := longestPalindrome("abab")
    fmt.Println(r)
}