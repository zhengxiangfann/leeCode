package main

import (
    "fmt"
    "strings"
)

func reverseString(s []byte)  {
    n := len(s)
    for i:=0; i< n/2; i ++{
        s[i],s[n-1-i] = s[n-1-i], s[i]
    }
}

func firstUniqChar(s string) int {
    if s== ""{
        return -1
    }
    var_m :=make(map[byte]int)
    for i:=0 ; i < len(s); i ++ {
        if v, ok := var_m[s[i]];ok {
            var_m[s[i]] = v +1
        } else {
            var_m[s[i]] = 1
        }
    }
    fmt.Println(var_m)
    for i:=0; i< len(s); i ++ {
        if v,ok := var_m[s[i]];ok && v==1 {
            return i
        }
    }
    return 0
}

func isAnagram(s string, t string) bool {
    if s == t {
        return true
    }
    if len(s) != len(t) {
        return false
    }
    l := len(s)
    s_map := make(map[byte]int)
    t_map := make(map[byte]int)
    for i:=0; i< l ;i ++ {
        if v, ok := s_map[s[i]]; ok {
            s_map[s[i]] = v +1
        } else {
            s_map[s[i]] = 1
        }
        if v, ok := t_map[t[i]]; ok {
            t_map[t[i]] = v +1
        } else {
            t_map[t[i]] = 1
        }
    }
    for k, v:= range s_map{
        if tmp_v, ok := t_map[k]; !ok {
            return false
        } else if v != tmp_v {
            return false
        }
    }
    for k, v:= range t_map{
        if tmp_v, ok := s_map[k]; !ok {
            return false
        }else if v != tmp_v {
            return false
        }
    }

    return true
}


func isPalindrome(s string) bool {
    if s == ""{
        return true
    }
    s = strings.ToLower(s)
    var s1 = []int32{}
    for _, v := range s{
        if (v >= '0' && v <= '9' ) || (v >= 'a' && v <= 'z') {
            s1 = append(s1, v)
        }
    }
    for i:= 0; i< len(s1)/2; i ++ {
        if s1[i] != s1[len(s1)-1-i] {
            return false
        }
    }
    return true
}


func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {
        return -1
    }
    if needle == haystack {
        return 0
    }
    if needle == ""{
        return 0
    }
    for i:=0; i< len(haystack)-len(needle)+1; i ++ {
        if string(haystack[i:i+len(needle)]) == needle {
            return i
        }
    }
    return -1
}


func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    bs := []byte{'1'}
    for i := 2; i <= n; i++ {
        bs = say(bs)
    }
    return string(bs)
}

func say(bs []byte) []byte {
    result := make([]byte, 0)
    x, y := 0, 1
    for x < len(bs) {
        for y < len(bs) && bs[x] == bs[y] {
            y++
        }
        result = append(result, byte(y-x+'0'), bs[x])
        x = y
    }
    return result
}






func main() {
    //var s  = []byte{'h', 'e', 'l', 'l', 'o','a', 'b'}
    //reverseString(s)
    //fmt.Println(string(s))

    r:=countAndSay(2)
    fmt.Println(r)
}
