1type NumArray struct {
2    prefix []int
3}
4
5
6func Constructor(nums []int) NumArray {
7    prefix := make([]int, len(nums)+1)
8    for i, v := range nums {
9        prefix[i+1] = prefix[i] + v
10    }
11    return NumArray{prefix : prefix}
12}
13
14
15func (this *NumArray) SumRange(left int, right int) int {
16    return this.prefix[right+1] - this.prefix[left]
17}
18
19
20/**
21 * Your NumArray object will be instantiated and called as such:
22 * obj := Constructor(nums);
23 * param_1 := obj.SumRange(left,right);
24 */