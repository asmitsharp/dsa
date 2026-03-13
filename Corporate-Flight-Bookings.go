1func corpFlightBookings(bookings [][]int, n int) []int {
2    result :=  make([]int, n+1)
3
4    for _, booking := range bookings {
5        start, end, seat := booking[0], booking[1], booking[2]
6
7        result[start-1] += seat
8        result[end] -= seat
9    }
10
11    for i := 1; i < n; i++ {
12        result[i] += result[i-1]
13    }
14
15    return result[:n]
16}