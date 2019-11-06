package main

import "fmt"

type MinStack struct {
    min int
    length int
    Val []int
    HisMinIdx []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        min:0,
        length:0,
        Val:make([]int,0),
        HisMinIdx:make([]int,0), // 历史最小值的栈
    }
}

func (this *MinStack) Push(x int)  {
    if this.length ==0 {
        this.HisMinIdx = append(this.HisMinIdx, 0)
        this.min = 0
    }
    this.Val = append(this.Val, x)
    this.length ++
    if  this.Val[this.min]  > x {
        this.min = this.length-1
        this.HisMinIdx = append(this.HisMinIdx, this.length -1)
    }
}

func (this *MinStack) Pop() {
    if this.length == 0{
        return
    }
    if  this.length-1  == this.HisMinIdx[len(this.HisMinIdx)-1]  && len(this.HisMinIdx)>1{
        this.HisMinIdx = this.HisMinIdx[:len(this.HisMinIdx)-1]
        this.min = this.HisMinIdx[len(this.HisMinIdx)-1]
    }
    this.Val = this.Val[:this.length-1]
    this.length --
}


func (this *MinStack) Top() int {
    if this.length ==0 {
        return -1
    }
    return this.Val[this.length-1]
}


func (this *MinStack) GetMin() int {
    if this.length == 0{
        return -1
    }
    return this.Val[this.min]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//
//["MinStack","push","push","push","push","getMin","pop","getMin","pop","getMin","pop","getMin"]
//[[],[2],[0],[3],[0],[],[],[],[],[],[],[]]
//
//["MinStack","push","push","push",
// "top","pop","getMin",
// "pop","getMin",
// "pop","push",
// "top","getMin",
// "push","top","getMin","pop","getMin"]
//[[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]


func main() {
    obj := Constructor();
    obj.Push(2147483646)
    obj.Push(2147483646)
    obj.Push(2147483647)
    //obj.Pop()
    fmt.Println(obj.Top())
    obj.Pop()
    obj.GetMin()

    obj.Pop()
    obj.GetMin()

    obj.Pop()
    obj.Push(2147483647)
    fmt.Println(obj.Top())
    obj.GetMin()

    obj.Push(-2147483648)
    fmt.Println(obj.Top())
    obj.GetMin()
    obj.Pop()
    obj.GetMin()
}