package main

import (
    "fmt"
    "math"
    "sort"
    "strconv"
    "strings"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
    i := m - 1
    j := n - 1
    index := m + n - 1
    for index >= 0 {
        if i < 0 {
            //第一个数组已经遍历完毕
            nums1[index] = nums2[j]
            index--
            j--
        } else if j < 0 {
            //第二个数组已经遍历完毕
            nums1[index] = nums1[i]
            index--
            i--
        } else if (nums1[i] > nums2[j] && i >= 0) {
            //比较两个数组元素
            nums1[index] = nums1[i]
            index--
            i--
        } else if (nums1[i] <= nums2[j] && j >= 0) {
            //比较两个数组元素
            nums1[index] = nums2[j]
            index--
            j--
        }
    }
}


//寻找数组的中心索引
//给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。
//
//我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
//
//如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
//
//示例 1:
//
//输入:
//nums = [1, 7, 3, 6, 5, 6]
//输出: 3
//解释:
//索引3 (nums[3] = 6) 的左侧数之和(1 + 7 + 3 = 11)，与右侧数之和(5 + 6 = 11)相等。
//同时, 3 也是第一个符合要求的中心索引。
//示例 2:
//
//输入:
//nums = [1, 2, 3]
//输出: -1
//解释:
//数组中不存在满足此条件的中心索引。

//暴力方式
func pivotIndex1(nums []int) int {
    preSum := 0
    lastSum := 0
    for i:=0;i< len(nums); i ++ {
        lastSum = 0
        for j:= i + 1; j< len(nums); j ++ {
            lastSum += nums[j]
            //if lastSum > preSum {
            //    break
            //}
        }
        if preSum == lastSum {
            return i
        }
        preSum += nums[i]
    }
    return -1
}

func pivotIndex(nums []int) int {
    left := make([]int, len(nums))
    right := make([]int, len(nums))
    sum:=0
    for i, v:= range nums {
        sum += v
        left[i] = sum
    }
    sum=0
    for i:=len(nums)-1;i>=0; i-- {
        sum += nums[i]
        right[i] = sum
    }

    fmt.Println(left)
    fmt.Println(right)
    for i:=0; i< len(nums) ;i ++ {
        if left[i] == right[i] {
            return i
        }
    }

    return -1
}

func dominantIndex(nums []int) int {
    max := 0
    maxIndex := 0
    for i:=0; i < len(nums); i++ {
        if max < nums[i] {
            max = nums[i]
            maxIndex = i
        }
    }
    for i:=0; i< len(nums); i ++{
        if nums[i]>0 && maxIndex !=i && max/nums[i] < 2 {
               return -1
        }
    }
    return maxIndex
}

func findDiagonalOrder(matrix [][]int) []int {
    //索引和为偶数：
    //      元素在第一行，往右走
    //      元素在最后一列，往下走
    //      其他情况，往右上走
    //索引和为奇数：
    //      元素在第一列，往下走
    //      元素在最后一行，往右走
    //      其他情况，往左下走
    r := 0
    c := 0
    m := len(matrix)
    n := len(matrix[0])
    result := make([]int, m*n)
    for i := 0; i < m*n; i++ {
        result[i] = matrix[r][c]
        if (r+c)%2 == 0 {
            if c == n-1 {
                r++
            } else if r == 0 {
                c++
            } else {
                r--
                c++
            }
        } else {
            if r == m-1 {
                c++
            } else if c == 0 {
                r++
            } else {
                r++
                c--
            }
        }
    }
    return result
}

func spiralOrder(matrix [][]int) []int {
    var result []int
    m := len(matrix)
    n := len(matrix[0])
    beginX, endX, beginY, endY := 0, 0, 0, 0
    beginX = 0
    endX = n - 1
    beginY = 0
    endY = m - 1
    for {
        // 从左上到右上
        for j := beginX; j <= endX; j++ {
            result = append(result, matrix[beginY][j])
        }
        beginY++
        if beginY > endY {
            break
        }

        // 从右上到右下
        for i := beginY; i <= endY; i++ {
            result = append(result, matrix[i][endX])
        }
        endX--
        if endX < beginX {
            break
        }

        // 从右下到左下
        for j := endX; j >= beginX; j-- {
            result = append(result, matrix[endY][j])
        }
        endY--
        if endY < beginY {
            break
        }

        // 从左下到左上
        for i := endY; i >= beginY; i-- {
            result = append(result, matrix[i][beginX])
        }
        beginX++
        if beginX > endX {
            break
        }

    }
    return result
}

//i=0
//  dp[0][0] = 1
//
//i = 1
//  dp[1][0] = dp[0][0] = 1
//  dp[1][1] = 1
//
//i = 2
//  dp[2][0] = 1
//  dp[2][1] = dp[1][0] + 1 = 2
//  dp[2][2] = 1
//
//i = 3
//  dp[3][0] = 1
//  dp[3][1] = dp[2][0] + dp[2][1] = 3
//  dp[3][2] = dp[2][1] + dp[2][2] = 3
//  dp[3][3] = 1

// 杨辉三角
func generate(numRows int) [][]int {
    if numRows <=0 {
        return nil
    }
    dp := make([][]int, numRows)
    for i:= 0;i < numRows ;i ++ {
        dp[i] = make([]int, i+1)
        dp[i][0] = 1
        for j:=1 ; j <= i ; j ++ {
            if j == i {
                dp[i][j] = 1
            } else {
                dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
            }
        }
    }
    return dp
}

func addBinary(a string, b string) string {
    if a ==  ""{
        return b
    }
    if b == ""{
        return a
    }

    if len(a) > len(b) {
        cha := len(a) - len(b)
        var c []byte
        for i:=0; i< cha; i++{
            c = append(c, '0')
        }
        b = fmt.Sprintf("%s%s", string(c), b)
    } else {
        cha := len(b) - len(a)
        var c []byte
        for i:=0; i< cha; i++{
            c = append(c, '0')
        }
        a = fmt.Sprintf("%s%s", string(c), a)
    }
    carrybit := 0
    sum_value := ""
    for i:= len(a)-1; i>=0; i-- {
        a1,_ := strconv.Atoi(string(a[i]))
        b1,_ := strconv.Atoi(string(b[i]))
        sum1 := a1 + b1 + carrybit
        if sum1 >=2 {
            carrybit = sum1/2
            sum_value = fmt.Sprintf( "%d%s",  sum1%2,  sum_value)
        } else {
            carrybit = 0
            sum_value= fmt.Sprintf("%d%s",sum1, sum_value)
        }
    }
    if carrybit ==1{
        sum_value= fmt.Sprintf("%s%s","1", sum_value)
    }
    return sum_value
}



func arrayPairSum(nums []int) int {
    //for i:= 0;i < len(nums);i++ {
    //    for j:=i+1; j<len(nums);j++{
    //        if nums[i] > nums[j] {
    //            nums[i], nums[j] = nums[j], nums[i]
    //        }
    //    }
    //}
    sort.Ints(nums)
    sum := 0
    for i:=0;i< len(nums);i += 2 {
        sum += nums[i]
    }
    return sum
}

func twoSum(numbers []int, target int) []int {
    var res []int
    for i, j := 0, len(numbers)-1; i < j; {
        if numbers[i]+numbers[j] == target {
            res = append(res, i+1)
            res = append(res, j+1)
        } else if numbers[i]+numbers[j] > target {
            j--
        } else if numbers[i]+numbers[j] < target {
            i++
        }
    }
    return res
}


func removeElement(nums []int, val int) int {
    i,j:=0,len(nums)-1
    if len(nums) == 1{
        if nums[0] == val {
            return 0
        } else {
            return 1
        }
    }

    for ; i < j; {
        if nums[i] == val {
            if nums[j] != val {
                nums[i], nums[j] = nums[j], nums[i]
                j--
                i++
            } else if nums[j] == val {
                j--
            }
        } else {
            i++
        }
    }
    return i+1
}

func removeElement11(nums []int, val int) int {

    if len(nums) == 0 {
        return 0
    }
    index := 0
    for ; index < len(nums); {
        if nums[index] == val {
            nums = append(nums[:index], nums[index+1:]...)
            continue
        }
        index++
    }
    return len(nums)
}

func findMaxConsecutiveOnes(nums []int) int {
    //[1 1 1 0 0 1 1] 求最长连续的1的个数
    maxSum := 0
    curSum := 0
    for i:= 0 ;i <len(nums) ; i ++ {
        if nums[i] == 1 {
            curSum +=1
        } else if nums[i] == 0{
            if maxSum < curSum {
                maxSum = curSum
            }
            curSum = 0
        }
    }
    if maxSum < curSum {
        maxSum = curSum
    }
    return maxSum
}


//
func minSubArrayLen(s int, nums []int) int {
    //给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。
    flag :=true
    minOfLength := s
    sum:= 0
    left := 0
    for right:=0; right < len(nums); right ++ {
        sum  = sum + nums[right]
        for sum >= s {
            flag = false
            if right - left +1 < minOfLength {
                minOfLength = right - left +1
            }
            sum = sum - nums[left]
            left ++
        }
    }
    if flag {
        return 0
    }
    return minOfLength
}

func yh(numRows int) [][]int {
    if numRows<=0 {
        return nil
    }
    dp := make([][]int, numRows)
    for i:= 0;i < numRows ;i ++ {
        dp[i] = make([]int, i+1)
        dp[i][0] = 1
        for j:=1 ; j <= i ; j ++ {
            if j == i {
                dp[i][j] = 1
            } else {
                dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
            }
        }
    }
    return dp
}

func getRow(rowIndex int) []int {
    if rowIndex<=0 {
        return nil
    }
    dp := make([][]int, rowIndex)
    for i:= 0;i < rowIndex ;i ++ {
        dp[i] = make([]int, i+1)
        dp[i][0] = 1
        for j:=1 ; j <= i ; j ++ {
            if j == i {
                dp[i][j] = 1
            } else {
                dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
            }
        }
    }
    return dp[rowIndex]
}

func reverseWords11(s string) string {
    var result []string
    res := strings.Split(s, " ")
    for i:=len(res) -1 ;i >=0; i -- {
        if strings.Trim(res[i]," ") != "" {
            result = append(result, strings.Trim(res[i], " "))
        }
    }
    return strings.Join(result, " ")
}

func reverseWords1111(s string) string {
    var result []uint8
    for startIndex, i := len(s)-1, len(s)-1; i >= 0; i-- {
        if s[i] == ' ' {
            if i < len(s)-1 {
                if s[i+1] == ' ' {
                    startIndex = i
                } else {

                    if startIndex == len(s)-1 {
                        startIndex = startIndex +1
                    }
                    result = append(result, s[i+1:startIndex]...)
                    result = append(result, ' ')
                    startIndex = i
                }
            }
        }
        if i == 0 {
           result = append(result, s[i:startIndex]...)
        }
    }
    return strings.Trim(string(result), " ")
}


func reverseWords(s string) string {
    sList := strings.Split(s, " ")
    var result []string
    for _, v := range sList {
        if tV := strings.Trim(v, " "); tV != "" {
            tmprever := make([]uint8, len(tV))
            for j := len(tV) -1 ; j>=0; j-- {
                tmprever[len(tV)-1-j] = tV[j]
            }
            result = append(result, string(tmprever))
        }
    }
    return strings.Join(result, " ")
}


func isPowerOfThree1111(n int) bool {
     //r := math.Log10(float64(n)) / math.Log10(3)
     //fmt.Println(r)
     //fmt.Println(int(r))
     //return false
    if n<=0 {
        return false
    }
    for n > 1 {
        if n %3 !=0 {
            return false
        }
        n = n/3
    }
    return true
}




func fizzBuzz(n int) []string {
    //res := make([]string, n)
    var res []string
    for i:=1; i<=n ;i++ {
        if i%3 ==0 && i%5 != 0 {
            //res[i-1] = "Fizz"
            res = append(res, "Fizz")
        } else if i%5==0 && i%3 != 0 {
            //res[i-1] = "Buzz"
            res = append(res, "Buzz")
        } else if i%3==0 && i%5 ==0 {
            //res[i-1] = "FizzBuzz"
            res = append(res, "FizzBuzz")
        } else {
            //res[i-1] = string(i)
            res = append(res, strconv.Itoa(i))
        }
    }
    return res
}


func countPrimes(n int) int {
    if n <= 1 {
        return 0
    }
    dp := make([]byte, n+1)
    count := 0
    for i := 2; i < n; i++ {
        if dp[i] == 0 {
            count++;
            k := 2
            for i*k <= n {
                dp[i*k] = 1
                k++
            }
        } else {
            continue
        }
    }
    return count
}

func hammingWeight(num uint32) int {
    count := 0
    for num > 0 {
        if (num%2 == 1) {
            count++
            num = (num - 1) / 2;
        } else {
            num = num / 2;
        }
    }
    return count
}

func hammingDistance(x int, y int) int {
     r := x & y
     return r

}



func reverseBits(num uint32) uint32 {
    var i uint32
    i, c:= 1, 31
    sum := 0
    for c >= 0 {
        if num&i == 0 {
            sum +=0 * int(math.Pow(2, float64(c)))
        } else {
            sum += 1 * int(math.Pow(2, float64(c)))
        }
        i = i <<1
        c --
    }
    fmt.Println(sum)
    return uint32(sum)
}

func main() {
    //var num1 = [][]int{{1}}
    //var num2 = []int{4,5,6}
    ////var num1 = []int{0}
    ////var num2 = []int{1}
    //merge(num1, 3, num2, 3)
    //fmt.Println(num1)
    //var arr = []int{2,3,1,2,4,3}
    //r := reverseWords("Let's take LeetCode contest")
    //r := isPowerOfThree(27)
    //r := countPrimes(10)
    //fmt.Println(r)

    //r := hammingWeight(20)
    //fmt.Println(r)

    //r := hammingDistance(1,2)
    //fmt.Println(r)

    r := reverseBits(30)
    fmt.Println(r)
}
