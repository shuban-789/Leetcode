import "slices"

func removeByValue(s []int, value int) []int {
	for i, v := range s {
		if v == value {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func singleNumber(nums []int) int {
    for i := 0; i < len(nums); i++ {
        numscpy := make([]int, len(nums))
        copy(numscpy, nums)
        mod := removeByValue(numscpy, int(nums[i]))
        if slices.Contains(mod, int(nums[i])) {
            continue
        } else {
            return int(nums[i])
        }
    }
    return 0
}
