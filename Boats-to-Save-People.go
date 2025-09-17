func numRescueBoats(people []int, limit int) int {
    sort.Ints(people)
    l, r := 0, len(people) - 1
    count := 0
    for l <= r {
        if people[l] + people[r]  <= limit {
            l++
        } 
        r--
        count++

    }

    return count
}