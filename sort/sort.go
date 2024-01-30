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
