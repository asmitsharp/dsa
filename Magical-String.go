func magicalString(n int) int {
  if  n == 0 {
    return 0
  }

  if n <= 3 {
    return 1
  }

  s := make([]byte, n)
  s[0] = '1'
  s[1] = '2'
  s[2] = '2'

  i, j := 2, 3
  count := 1
  currentChar := byte('1')

  for j < n {
    num := int(s[i] - '0')
    for k := 0; k < num && j < n; k++ {
        s[j] = currentChar
        if currentChar == '1' {
            count++
        }
        j++
    }

    if currentChar == '1' {
        currentChar = '2'
    } else {
        currentChar = '1'
    }
    i++
  }
  return count
}