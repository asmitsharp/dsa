1func addOperators(num string, target int) []string {
2    result := []string{}
3
4    var backtrack func(index int, curr int, prev int, path string)
5    backtrack = func(index int, curr int, prev int, path string) {
6        if index == len(num) {
7            if curr == target {
8                result = append(result, path)
9            }
10            return
11        }
12
13        for i := index; i < len(num); i++ {
14            if i != index && num[index] == '0' {
15                break
16            }
17
18            val := 0
19            for j := index; j <= i; j++ {
20                val = val*10 + int(num[j]-'0')
21            }
22
23            if index == 0 {
24                backtrack(i+1, val, val, path+num[index:i+1])
25            } else {
26                backtrack(i+1, curr+val, val, path+"+"+num[index:i+1])
27                backtrack(i+1, curr-val, -val, path+"-"+num[index:i+1])
28                backtrack(i+1, curr-prev+prev*val, prev*val, path+"*"+num[index:i+1])
29            }
30        }
31        
32    }
33    backtrack(0, 0, 0, "")
34    return result
35}