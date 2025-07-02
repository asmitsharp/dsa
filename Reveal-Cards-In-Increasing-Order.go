func deckRevealedIncreasing(deck []int) []int {
    n := len(deck)
    sort.Ints(deck)

    deque := make([]int, 0)

    for i := n - 1; i >= 0; i-- {
        if len(deque) > 0 {
            last := deque[len(deque) - 1]
            deque = deque[:len(deque) - 1]

            deque = append([]int{last}, deque...)
        }
        deque = append([]int{deck[i]}, deque...)
    }

    return deque
}