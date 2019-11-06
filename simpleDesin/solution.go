package main1

import (
    "fmt"
    "math/rand"
    "time"
)

type Solution struct {
    Nums *[]int
}

func Constructor(nums []int) Solution {
    if len(nums) >0 {
        return Solution{
            Nums:&nums,
        }
    } else {
        return Solution{
            Nums:nil,
        }
    }
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    if this.Nums != nil {
        return *this.Nums
    }
    return nil
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    var slice []int
    if this.Nums == nil{
       return slice
    }
    slice = append(slice, *this.Nums...)
    //copy(slice, *this.Nums)
    //r := rand.New(rand.NewSource(time.Now().Unix()))
    for i:=0;i<len(slice); i++ {
        n := len(slice)
        randIndex := r.Intn(n)
        slice[i], slice[randIndex] = slice[randIndex], slice[i]
        //slice = slice[:n-1]
    }
    return slice

    //for i := len(slice) - 1; i > 0; i-- {
    //    randIndex := r.Intn(i)
    //    slice[i-1], slice[randIndex] = slice[randIndex], slice[i-1]
    //    slice = slice[:i-1]
    //}
    //return slice
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func Shuffle(slice []interface{}) {
    r := rand.New(rand.NewSource(time.Now().Unix()))
    for len(slice) > 0 {
        n := len(slice)
        randIndex := r.Intn(n)
        slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
        slice = slice[:n-1]
    }

}


func main() {

    //slice := []interface{}{"a", "b", "c", "d", "e", "f"}
    //Shuffle(slice)
    //fmt.Println(slice)
    //for _, v := range slice {
    //    fmt.Println(v.(string))
    //}


    var a = []int{1, 2, 3}
    obj1 := Constructor(a)

    for i:=0;i<len(a); i ++ {
        r := obj1.Shuffle()
        fmt.Print(r,)
    }

    fmt.Println(a)
}