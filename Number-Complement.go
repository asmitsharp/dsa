1func findComplement(num int) int {
2    mask := (1 << bits.Len(uint(num))) - 1
3    return num ^ mask
4}