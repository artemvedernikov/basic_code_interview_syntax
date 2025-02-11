//https://leetcode.com/problems/remove-all-occurrences-of-a-substring
// stack to collect the bytes of string

func removeOccurrences(s string, part string) string {
    var stack []byte

    for i := 0; i < len(s); i++ {
        elem := s[i]
        stack = append(stack, elem)

        // cleanup stack
        if (len(stack) >= len(part)) {
            partInd := 0
            needToTrim := true
            for stackInd := len(stack) - len(part); stackInd < len(stack); stackInd++ {
                stackElem := stack[stackInd]
                partElem := part[partInd]
                
                if (stackElem != partElem) {
                    needToTrim = false
                    break
                } 
                partInd++
            }
            if needToTrim {
                stack = stack[: len(stack) - len(part)]
            }
        }
    }

    return string(stack)
}
