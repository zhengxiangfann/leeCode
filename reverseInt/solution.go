package main

import (
    "fmt"
    "math"
    "strconv"
)

func reverse(x int) int {
    var f bool
    if x > 0 {
        f = true
    } else {
        f = false
    }
    str := strconv.Itoa(int(math.Abs(float64(x))))
    var rstr []uint8
    for i := len(str) - 1; i >= 0; i-- {
        rstr = append(rstr, str[i])
    }
    n, _ := strconv.Atoi(string(rstr))
    if n > 1 << 31 -1 {
        return 0
    }
    if f {
        return n
    }
    return -n
}


func main(){
    fmt.Println(math.MaxInt32)
    fmt.Println(reverse(1534236469))

}
