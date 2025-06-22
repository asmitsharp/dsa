func finalPrices(prices []int) []int {
    n := len(prices)
    stack := []int{}

    for i := 0; i < n; i++ {
        for len(stack) > 0 && prices[i] <= prices[stack[len(stack) - 1]] {
            top := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            prices[top] -= prices[i]
        }
        stack = append(stack, i)
    }

    return prices
}