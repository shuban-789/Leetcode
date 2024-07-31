import "strings"

func longestCommonPrefix(strs []string) string {
    var marked string
    var prefix string
    marked = string(strs[0])
    if len(strs) == 1 {
        return string(strs[0])
    }
    for i := 1; i < len(strs); i++ {
        if len(string(strs[i])) < len(marked) {
            marked = string(strs[i])
        }
    }
    for j := 0; j < len(marked); j++ {
        target := string(marked[j])
        index := 0
        for k := 0; k < len(strs); k++ {
            if j < len(strs[k]) && string(strs[k][j]) == target {
                index += 1
            }
        }
        if index == len(strs) {
            prefix += target
        } else {
            break
        }
    }
    return prefix
}

// Runtime beats 100% of submissions
