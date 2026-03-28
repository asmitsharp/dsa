1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val int
5 *     Left *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9
10type Codec struct {
11    
12}
13
14func Constructor() Codec {
15    return Codec{}
16}
17
18// Serializes a tree to a single string.
19func (this *Codec) serialize(root *TreeNode) string {
20    res := []string{}
21
22    var dfs func(node *TreeNode)
23    dfs = func(node *TreeNode) {
24        if node == nil {
25            res = append(res, "null")
26            return
27        }
28
29        res = append(res, strconv.Itoa(node.Val))
30
31        dfs(node.Left)
32        dfs(node.Right)
33    }
34    dfs(root)
35    return strings.Join(res, ",")
36}
37
38// Deserializes your encoded data to tree.
39func (this *Codec) deserialize(data string) *TreeNode {    
40    vals := strings.Split(data, ",")
41    i := 0
42
43    var dfs func() *TreeNode 
44    dfs = func() *TreeNode {
45        if vals[i] == "null" {
46            i++
47            return nil
48        }
49
50        val, _ := strconv.Atoi(vals[i])
51        i++
52
53        node := &TreeNode{Val: val}
54
55        node.Left = dfs()
56        node.Right = dfs()
57
58        return node
59    }
60    return dfs()
61}
62
63
64/**
65 * Your Codec object will be instantiated and called as such:
66 * ser := Constructor();
67 * deser := Constructor();
68 * data := ser.serialize(root);
69 * ans := deser.deserialize(data);
70 */