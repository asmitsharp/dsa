1func letterCombinations(digits string) []string {
2    if len(digits) == 0 {
3        return []string{}
4    }
5
6     mapping := map[byte]string{
7        '2': "abc",
8        '3': "def",
9        '4': "ghi",
10        '5': "jkl",
11        '6': "mno",
12        '7': "pqrs",
13        '8': "tuv",
14        '9': "wxyz",
15    }
16    result := []string{}
17    backtrack(0,mapping, "", &result, digits)
18    return result
19}
20
21func backtrack(index int, mapping map[byte]string, curr string, result *[]string, digits string) {
22    if index == len(digits) {
23        *result = append(*result, curr)
24        return
25    }
26
27    letters := mapping[digits[index]]
28    for i := 0; i < len(letters); i++ {
29        backtrack(index+1, mapping, curr + string(letters[i]), result, digits)
30    }
31}