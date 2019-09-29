package main

import "fmt"

func romanToInt(s string) int {
    m := map[string]int{"I":1, "V":5, "X":10, "L":50, "C":100, "D":500, "M":1000}
    //fmt.Println(m)
    var result = 0
    for i := 0 ;i < len(s) -1 ; i ++ {
        if  m[string(s[i])] < m[string(s[i+1])] {
            result -= m[string(s[i])]
        } else {
            result += m[string(s[i])]
        }
    }
    result += m[string(s[len(s)-1])]
    return result
}

func main() {
    res := romanToInt("MCMXCIV")
    fmt.Println(res)
}
