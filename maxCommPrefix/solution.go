package main
import "fmt"
func longestCommonPrefix(strs []string) string {
    if len(strs) ==0 {
        return ""
    }
    if len(strs) == 1{
        return strs[0]
    }
    min_l := len(strs[0])
    min_i := 0
    for i, v := range strs {
        if min_l > len(v) {
            min_l = len(v)
            min_i = i
        }
    }
    for i:= 0 ;i < len(strs[min_i]);i++{
       for j := 0; j < len(strs); j++{
           if strs[min_i][i] != strs[j][i] {
               return string(strs[min_i][:i])
           }
       }
    }
    return strs[min_i]
}

func main(){
    strs := []string{"ca", "a" }
    res := longestCommonPrefix(strs)
    fmt.Println(res)
}
