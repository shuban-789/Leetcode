import (
"strings"
)

func reversePrefix(word string, ch byte) string {
    var listword []string
    for h := 0; h < len(word); h++ {
        listword = append(listword, string(word[h]))
    }
    index := 0
    for i := 0; i < len(listword); i++ {
        if string(ch) == string(listword[i]) {
            index = i
            break
        }
    }
    for j, k := 0, index; j < k; j, k = j+1, k-1 {
        listword[j], listword[k] = listword[k], listword[j]
    }
    retstr := strings.Join(listword, "")
    return retstr
}
