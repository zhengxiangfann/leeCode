package main

import (
    "fmt"
    "math"
)

/*动态规划*/

func fib(n int) int {
    if n <=2 {
        return 1
    }
    return fib(n-1) + fib(n-2)
}

func fib_dy(n int) int {
    if n < 2 {
        return n
    }
    prev := 0
    curr := 1
    for  i := 0; i < n - 1; i++ {
        sum := prev + curr
        prev = curr
        curr = sum
    }
    return curr
}


func fib_dp(N int) int{

    var dp []int
    dp = append(dp, 0,1,1)
    //dp[1], dp[2] = 1,1
    for i := 3; i <= N; i++ {
        dp= append(dp, dp[i-1] + dp[i-2])
        //dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[N]
}

// 这个算法基本不能用
func coinChange(coins []int, amount int) int{
    if amount == 0 {
        return 0
    }
    ans := math.MaxInt32
    for _, c :=  range coins {
        if amount - c  < 0 {
            continue
        }
        subProb := coinChange(coins, amount - c)
        if (subProb == -1) {
            continue
        }
        //ans = min(ans, subProb + 1)
        if  ans > subProb+1 {
            ans = subProb +1
        }
    }
    if ans == math.MaxInt32 {
        return -1
    }else {
        return ans
    }
}



func coinChange_dy(coins []int, amount int) int {
    var dp []int
    dp = append(dp, 0)
    for i:=0; i<= amount +1; i ++{
        dp = append(dp, amount +1)
    }

    for i := 0; i < amount+1; i++ {
        // 内层 for 在求所有子问题 + 1 的最小值
        for _, coin := range coins {
            if i-coin < 0 {
                continue
            }

            if dp[i] > dp[i-coin]+1 {
                dp[i] = dp[i-coin] + 1
                //dp = append(dp,  dp[i-coin] +1)
            }
        }
    }
    if dp[amount] == amount+1 {
        return -1
    } else {
        return dp[amount]
    }
}


func main(){
    //v := fib_dp(3)
    //fmt.Println(v)

    coins := []int{1,2,5,10}
    r := coinChange_dy(coins[:], 3)
    fmt.Println(r)
}
