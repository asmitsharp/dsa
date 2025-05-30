func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int)
\tfor _, num := range nums {
\t\tcount[num]++
\t}

\tunique := make([]int, 0, len(count))
\tfor key := range count {
\t\tunique = append(unique, key)
\t}

\tpartition := func(left, right, pivotIndex int) int {
\t\tpivotFrequency := count[unique[pivotIndex]]
\t\tunique[pivotIndex], unique[right] = unique[right], unique[pivotIndex]

\t\tstoreIndex := left
\t\tfor i := left; i < right; i++ {
\t\t\tif count[unique[i]] < pivotFrequency {
\t\t\t\tunique[storeIndex], unique[i] = unique[i], unique[storeIndex]
\t\t\t\tstoreIndex++
\t\t\t}
\t\t}

\t\tunique[right], unique[storeIndex] = unique[storeIndex], unique[right]

\t\treturn storeIndex
\t}

    var quickSelect func(int, int, int)
\tquickSelect = func(left, right, kSmallest int) {
\t\tif left == right {
\t\t\treturn
\t\t}

\t\tpivotIndex := rand.Intn(right-left+1) + left
\t\tpivotIndex = partition(left, right, pivotIndex)

\t\tif kSmallest == pivotIndex {
\t\t\treturn
\t\t} else if kSmallest < pivotIndex {
\t\t\tquickSelect(left, pivotIndex-1, kSmallest)
\t\t} else {
\t\t\tquickSelect(pivotIndex+1, right, kSmallest)
\t\t}
\t}

\tn := len(unique)
\tquickSelect(0, n-1, n-k)

\treturn unique[n-k:]
}