1func readBinaryWatch(turnedOn int) []string {
2    res := []string{}
3
4    var backtrack func(index int, hours int, min int, rem int)
5    backtrack = func(index int, hours int, min int, rem int) {
6        if rem == 0 {
7            if hours < 12 && min < 60 {
8                res = append(res, fmt.Sprintf("%d:%02d", hours, min))
9            }
10            return
11        }
12
13        if index == 10 {
14            return
15        }
16
17        if index < 4 {
18            backtrack(index+1, hours+(1<<(3-index)), min, rem-1)
19        } else {
20            backtrack(index+1, hours, min+(1<<(9-index)), rem-1)
21        }
22
23        backtrack(index+1,hours, min, rem)
24    }
25
26    backtrack(0, 0, 0, turnedOn)
27    return res
28}