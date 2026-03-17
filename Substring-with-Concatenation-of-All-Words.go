1func findSubstring(s string, words []string) []int {
2    if len(s) == 0 || len(words) == 0 {
3		return []int{}
4	}
5
6	wordLen := len(words[0])
7	totalWords := len(words)
8	//windowSize := wordLen * totalWords
9
10	// Step 1: word frequency map
11	wordCount := make(map[string]int)
12	for _, w := range words {
13		wordCount[w]++
14	}
15
16	result := []int{}
17
18	// Step 2: multiple starting points
19	for i := 0; i < wordLen; i++ {
20
21		left := i
22		right := i
23		windowCount := make(map[string]int)
24		count := 0
25
26		for right+wordLen <= len(s) {
27
28			word := s[right : right+wordLen]
29			right += wordLen
30
31			// If word is valid
32			if _, exists := wordCount[word]; exists {
33				windowCount[word]++
34				count++
35
36				// If extra occurrence → shrink window
37				for windowCount[word] > wordCount[word] {
38					leftWord := s[left : left+wordLen]
39					windowCount[leftWord]--
40					left += wordLen
41					count--
42				}
43
44				// If all words matched
45				if count == totalWords {
46					result = append(result, left)
47				}
48
49			} else {
50				// Reset window
51				windowCount = make(map[string]int)
52				count = 0
53				left = right
54			}
55		}
56	}
57
58	return result
59}