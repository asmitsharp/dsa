1func findMaximumXOR(nums []int) int {
2    maxXOR := 0
3    
4    for bit := 31; bit >= 0; bit-- {
5        candidate := maxXOR | (1 << bit)
6        
7        prefixes := make(map[int]bool)
8        mask := ^((1 << bit) - 1)
9        
10        for _, num := range nums {
11            prefix := num & mask
12            prefixes[prefix] = true
13            
14            needed := candidate ^ prefix
15            if prefixes[needed] {
16                maxXOR = candidate
17                break
18            }
19        }
20    }
21    return maxXOR
22}