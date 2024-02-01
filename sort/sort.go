package sort

import (
    "context"
    "sync"
    "time"
)

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

// SleepSortNew sleep sort
//  - nums: the numbers will be sorted
func SleepSortNew(nums []int) (result []int) {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	channel := make(chan int)
	wg.Add(len(nums))
	for _, v := range nums {
		go func(v int) {
			time.Sleep(time.Duration(v) * time.Millisecond)
			channel <- v
			wg.Done()
		}(v)
	}
	go func() {
		for {
			select {
			case v, _ := <-channel:
				result = append(result, v)
			case <-ctx.Done():
				break
			}
		}
	}()
	wg.Wait()
	time.Sleep(time.Millisecond)
	cancel()
	return
}
