package main

import (
    "fmt"
)
/*
递归解决
*/

func isMatch(s string, p string) bool {
    if s == "" {
        return true
    }
    var first_match bool
    //if s != "" && (p[0] == s[0] || p[0] == '.') {
    //    first_match = true
    //} else {
    //    first_match = false
    //}
    first_match = s != "" && (p[0] == s[0] || p[0] == '.')
    if len(p) >= 2 && p[1] == '*'{
         return isMatch(s, p[2:]) || (first_match && isMatch(s[1:], p))
    } else  {
        return first_match && isMatch(s[1:], p[1:])
    }
}

func main() {
    s := "ab"
    p := ".*c"
    r := isMatch(s, p)
    fmt.Println("是否匹配:", r)
}
