1func maxProfit(prices []int) int {
2    left := 0
3    maxProfit := 0
4
5    for right := 0; right < len(prices); right++ {
6        currentProfit := prices[right] - prices[left]
7
8        if currentProfit > maxProfit {
9            maxProfit = currentProfit
10        }
11
12        if prices[left] > prices[right] {
13            left = right
14        }
15    }
16
17    return maxProfit
18}