// func topKFrequent(nums []int, k int) []int {
//     count := make(map[int]int)
// \tfor _, num := range nums {
// \t\tcount[num]++
// \t}

// \tunique := make([]int, 0, len(count))
// \tfor key := range count {
// \t\tunique = append(unique, key)
// \t}

// \tpartition := func(left, right, pivotIndex int) int {
// \t\tpivotFrequency := count[unique[pivotIndex]]
// \t\tunique[pivotIndex], unique[right] = unique[right], unique[pivotIndex]

// \t\tstoreIndex := left
// \t\tfor i := left; i < right; i++ {
// \t\t\tif count[unique[i]] < pivotFrequency {
// \t\t\t\tunique[storeIndex], unique[i] = unique[i], unique[storeIndex]
// \t\t\t\tstoreIndex++
// \t\t\t}
// \t\t}

// \t\tunique[right], unique[storeIndex] = unique[storeIndex], unique[right]

// \t\treturn storeIndex
// \t}

//     var quickSelect func(int, int, int)
// \tquickSelect = func(left, right, kSmallest int) {
// \t\tif left == right {
// \t\t\treturn
// \t\t}

// \t\tpivotIndex := rand.Intn(right-left+1) + left
// \t\tpivotIndex = partition(left, right, pivotIndex)

// \t\tif kSmallest == pivotIndex {
// \t\t\treturn
// \t\t} else if kSmallest < pivotIndex {
// \t\t\tquickSelect(left, pivotIndex-1, kSmallest)
// \t\t} else {
// \t\t\tquickSelect(pivotIndex+1, right, kSmallest)
// \t\t}
// \t}

// \tn := len(unique)
// \tquickSelect(0, n-1, n-k)

// \treturn unique[n-k:]
// }

func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
    }

    pairs := make([][2]int, 0, len(freq))
    for num, f := range freq {
        pairs = append(pairs, [2]int{f, num})
    }

    return heapsort(pairs, k)
}

func heapsort(pairs [][2]int, k int) []int {
    n := len(pairs)

    for i := n/2 - 1; i >= 0; i-- {
        heapify(pairs, n, i)
    }

    result := make([]int, k)
    for i := 0; i < k; i++ {
        result[i] = pairs[0][1]
        pairs[0], pairs[n-1-i] = pairs[n-1-i], pairs[0]
        heapify(pairs, n-1-i, 0)
    }

    return result
}

func heapify(pairs [][2]int, n, i int) {
    largest := i
    left := 2*i + 1
    right := 2*i + 2

    if left < n && pairs[left][0] > pairs[largest][0] {
        largest = left
    }

    if right < n && pairs[right][0] > pairs[largest][0] {
        largest = right
    }

    if largest != i {
        pairs[i], pairs[largest] = pairs[largest], pairs[i]
        heapify(pairs, n, largest)
    }
}