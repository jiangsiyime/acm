package sort

// QuickSort quick sort function
//  - nums: the numbers will be sorted
//  - left: left index of the slice
//  - right: right index of the slice
func QuickSort(nums []int, left int, right int) {
    if left >= right {
        return
    }
    flag := left
    for i := left + 1; i <= right; i++ {
        if nums[i] < nums[left] {
            flag++
            nums[i], nums[flag] = nums[flag], nums[i]
        }
    }
    nums[flag], nums[left] = nums[left], nums[flag]
    QuickSort(nums, left, flag-1)
    QuickSort(nums, flag+1, right)
}

// MergeSort merge sort function
//  - nums: the numbers will be sorted
func MergeSort(nums []int) (result []int) {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	return Merge(left, right)
}

// Merge merge two slices
//  - left: left slice
//  - right: right slice
func Merge(left []int, right []int) (result []int) {
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return
}
