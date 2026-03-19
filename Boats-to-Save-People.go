1func numRescueBoats(people []int, limit int) int {
2    sort.Ints(people)
3    left := 0
4    right := len(people) - 1
5    boat := 0
6
7    for left <= right {
8        if left == right{
9            boat++
10            break
11        }
12        sum := people[left] + people[right] 
13        if sum <= limit {
14            left++  
15            right--
16        } else if sum > limit {
17            right--
18        }
19        boat++
20    }
21    return boat
22}