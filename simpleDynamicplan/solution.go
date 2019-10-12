package main

import (
    "fmt"
    "math"
)

func climbStairs(n int) int {
    //     if n <=2 {
    //         return n
    //     }
    //     return climbStairs(n-1) + climbStairs(n-2)

    var m []int
    m = append(m,0)
    m = append(m,1)
    m = append(m,2)
    for i:=3; i<=n ; i++ {
        m = append(m, m[i-1]+m[i-2])
    }
    fmt.Println(m)
    return m[n]
}


func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    max, min := 0, prices[0]
    for i:=1 ;i < len(prices); i ++ {
        min = int(math.Min(float64(prices[i]), float64(min)))
        max = int(math.Max(float64(max),  float64(prices[i] - min)))
    }
    return max
}


func maxProfit1(prices []int) int {
    if len(prices)==0{
        return 0
    }
    dp := make([]int,len(prices))
    dp[0]=0
    mi := prices[0]
    for i:=1;i<len(prices);i++{
        dp[i]=max(prices[i]-mi,dp[i-1])
        if prices[i]<mi {
            mi = prices[i]
        }
    }
    return dp[len(dp)-1]
}

func min(a,b int)int {
    if a>b {
        return b
    }
    return a
}

func max(a,b int)int {
    if a<b {
        return b
    }
    return a
}



// 最大子序列和
func maxSubArray(nums []int) int {
    if len(nums) ==0 {
        return 0
    }
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    sum := dp[0]
    for i:=1; i<len(nums); i ++ {
        if dp[i-1] >= 0 {
           dp[i] = dp[i-1] + nums[i]
        } else {
          dp[i] = nums[i]
        }
        if sum < dp[i] {
            sum = dp[i]
        }
    }
    return sum
}

func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    max := dp[0]
    resut := dp[0]
    for i:=1;i < len(nums); i++ {
           if i - 2 < 0 {
               dp[i] = nums[i]
           } else {
               //dp[i] = nums[i] + dp[i-2]
               dp[i] = nums[i] + max
           }
           if max < dp[i-1] {
               max = dp[i-1]
           }
           if dp[i] > resut {
               resut = dp[i]
           }
    }
    fmt.Println(dp)
    return resut
}


func main() {
    //r := climbStairs(5)
    //fmt.Println(r)
    var prices = []int{2,1,1,1,4}
    r := rob(prices)
    fmt.Println(r)
}