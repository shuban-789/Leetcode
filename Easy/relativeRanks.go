import (
    "sort"
    "strconv"
)

func findRelativeRanks(score []int) []string {
    var output []string
    var sorted[]int
    var gold string
    var silver string
    var bronze string
    for j := 0; j < len(score); j++ {
        sorted = append(sorted, score[j])
    }
    sort.Ints(sorted)
    gold = string(sorted[len(sorted)-1])
    if len(sorted) > 1 {
        silver = string(sorted[len(sorted)-2])
    }
    if len(sorted) > 2 {
        bronze = string(sorted[len(sorted)-3])
    }
    for i := 0; i < len(score); i++ {
        if string(score[i]) == gold {
            output = append(output, "Gold Medal")
        } else if string(score[i]) == silver {
            output = append(output, "Silver Medal")
        } else if string(score[i]) == bronze {
            output = append(output, "Bronze Medal")
        } else {
            for k := 0; k < len(sorted); k++ {
                if string(sorted[k]) == string(score[i]) {
                    output = append(output, strconv.Itoa(len(sorted)-k))
                }
            }
        }
    }
    return output
}

// TODO: Achieve a better runtime 😭
