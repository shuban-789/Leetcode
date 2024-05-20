import (
    "sort"
)

func joinArray(array1 []int, array2 []int) []int {
    for i := 0; i < len(array2); i++ {
        array1 = append(array1, array2[i])
    }
    return array1
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    joinedArray := joinArray(nums1, nums2)
    sort.Ints(joinedArray)
    var median float64
    if len(joinedArray)%2 == 0 {
        median = (float64(joinedArray[len(joinedArray)/2]) + float64(joinedArray[(len(joinedArray)/2)-1])) / 2.0
    } else {
        median = float64(joinedArray[(len(joinedArray)-1)/2])
    }
    return median
}
