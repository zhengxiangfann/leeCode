package main

import (
    "fmt"
    "math"
)
//
//func myAtoi(s string) int {
//
//    index := 0
//
//
//
//
//    s_to_i := map[uint8]int{
//        48:0,
//        49:1,
//        50:2,
//        51:3,
//        52:4,
//        53:5,
//        54:6,
//        55:7,
//        56:8,
//        57:9,
//    }
//
//    s = strings.Trim(s, " ")
//    if s == "" {
//        return 0
//    }
//    if !(s[0] == '-' || s[0] == '+' ) && (s[0] > 57 || s[0] < 48 ) {
//        return 0
//    }
//
//
//
//
//    if string(s[0]) == "-" {
//        f = false
//        if_f = true
//    } else if string(s[0]) == "+" {
//        f = true
//        if_f = true
//    }
//    for i:= 0 ; i < len(s) ; i ++ {
//        if find_head {
//            if s[i] >= 48 && s[i] <= 57  {
//                end = i+1
//            } else {
//                break
//            }
//        } else {
//            if s[i] > 48 && s[i] <= 57{
//                find_head = true
//                //if_f = true
//                start = i
//                end = i+1
//            } else if i !=0 && if_f && s[i] != 48{
//                break
//            }
//        }
//    }
//
//    var res int
//    rs := s[start: end]
//    fmt.Println("rs", rs)
//    for j:= len(rs) -1 ; j >=0 ; j -- {
//        px := 0
//        for k := 0 ; k < len(rs) - j  ; k ++ {
//            if px ==0 {
//                px = 1
//            } else {
//                px = px * 10
//            }
//        }
//        res = res + s_to_i[rs[j]] * px
//    }
//    if res >= - math.MaxInt32  && res <= math.MaxInt32 {
//        if f {
//            return res
//        }
//        return -res
//    } else {
//        if f {
//            return math.MaxInt32
//        } else {
//            return -math.MaxInt32-1
//        }
//    }
//    return 0
//}


func myAtoi1(str string) int {

    INT_Max := math.MaxInt32
    INT_Min := math.MinInt32
    x := 0
    base := 0
    sign := 1
    if str == "" {
        return 0
    }
    for x < len(str) && str[x] == ' ' {
        x++
    }
    if x < len(str) {
        if str[x] == '-' {
            sign = -1
            x++
        } else if str[x] == '+' {
            x++
        }

    }
    for ; x < len(str) && str[x] > 47 && str[x] < 58; {
        if base > INT_Max/10 || (base == INT_Max/10 && str[x]-'0' > 7) {
            if sign == 1 {
                return INT_Max
            } else {
                return INT_Min
            }
        }
        base = 10*base + int(str[x]-'0')
        x++
    }
    return base * sign
}


func main(){
   fmt.Println( myAtoi1("0-1"))
}