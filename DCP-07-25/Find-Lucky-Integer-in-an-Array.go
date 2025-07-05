func findLucky(arr []int) int {
    freq := make(map[int]int, 0)
    for _, num := range arr {
        freq[num]++
    }

    max := -1

    for num, count := range freq {
        if num == count {
            if num > max {
                max = num
            }
        }
    }

    return max
}