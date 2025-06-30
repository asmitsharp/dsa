func timeRequiredToBuy(tickets []int, k int) int {
   queue := make([]int, 0, len(tickets))
    for i := range tickets {
        queue = append(queue, i)
    }

    time := 0

    for tickets[k] > 0 {
        person := queue[0]
        queue  = queue[1:]

        tickets[person]--
        time++

        if tickets[person] > 0 {
            queue = append(queue, person)
        }
    }

    return time
   
}