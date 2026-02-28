1type MyCalendarTwo struct {
2    events [][]int
3    overlaps [][]int
4}
5
6
7func Constructor() MyCalendarTwo {
8    return MyCalendarTwo{
9        events:   make([][]int, 0),
10        overlaps: make([][]int, 0),
11    }
12}
13
14
15func (this *MyCalendarTwo) Book(startTime int, endTime int) bool {
16    // check agains overlaps
17    for _, overlap := range this.overlaps {
18        if startTime < overlap[1] && endTime > overlap[0] {
19            return false // triple booking
20        }
21    }
22
23    // check for all booking to find new overlaps
24    for _, event := range this.events {
25        if startTime < event[1] && endTime > event[0] {
26            newOverlap := []int{max(startTime, event[0]), min(endTime, event[1])}
27            this.overlaps = append(this.overlaps, newOverlap)
28        }
29    }
30
31    this.events = append(this.events, []int{startTime, endTime})
32    return true
33}
34
35func max(a, b int) int {
36    if a > b {
37        return a
38    }
39    return b
40}
41func min(a, b int) int {
42    if a < b {
43        return a
44    }
45    return b
46}
47
48/**
49 * Your MyCalendarTwo object will be instantiated and called as such:
50 * obj := Constructor();
51 * param_1 := obj.Book(startTime,endTime);
52 */