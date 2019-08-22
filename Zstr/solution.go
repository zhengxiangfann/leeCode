package main

import "fmt"

func convert(s string, numRows int) string {
    var resStr []uint8
    var i, j= 0, 0
    if len(s) <= numRows || numRows == 1{
        return s
    }

    for i < numRows {
        j = i
        if i == 0 || i == numRows-1 {
            for j < len(s) {
                resStr = append(resStr, s[j])
                j += numRows + (numRows - 2)
                if numRows+(numRows-2) == 0 {
                    break
                }
            }
        } else {
            for j < len(s) {
                resStr = append(resStr, s[j])
                j += 2*(numRows-i) - 2
                if j >= len(s) {
                    break
                } else {
                    resStr = append(resStr, s[j])
                    j += 2 * i
                }
            }
        }
        i++
    }
    return string(resStr)
}

func main() {
    r := convert("as", 1)
    fmt.Println(r)
}

