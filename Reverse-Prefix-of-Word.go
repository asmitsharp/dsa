// func reversePrefix(word string, ch byte) string {
//      stack := make([]rune, 0)

//     idx := -1
//     for i,char := range word {
//         if byte(char) == ch {
//             idx = i
//             stack = append(stack, char)
//             break
//         } else {
//             stack = append(stack, char)
//         }
//     }

//     if idx == -1 {
//         return word
//     }

//     result := make([]rune, 0, len(word))
//     for len(stack) > 0 {
//         result = append(result, stack[len(stack) - 1])
//         stack = stack[:len(stack) - 1]
//     }

//     for i := idx + 1; i < len(word); i++ {
//         result = append(result, rune(word[i]))
//     }

//     return string(result)
// }

// Two Pointer

func reversePrefix(word string, ch byte) string {
    runes := []byte(word)
   idx := -1
   for i, char := range runes {
    if char == ch {
        idx = i
        break
    }
   }

   if idx == -1 {
    return word
   }

   l := 0
   r := idx

   for l < r {
    runes[l], runes[r] = runes[r], runes[l]
    l++
    r--
   }

   return string(runes)
}