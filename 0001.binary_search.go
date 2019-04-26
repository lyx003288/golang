package golang

// **************************************************************
// 二分查找算法 与 二分查找算法变种算法
// **************************************************************

//BinarySearch 二分查找算法
func BinarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		switch {
		case nums[mid] > target:
			r = mid - 1
		case nums[mid] < target:
			l = mid + 1
		default:
			return mid
		}
	}

	return -1
}

// 查找第一个与target相等的元素
func findFirstEqual(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	if i < len(nums) && nums[i] == target {
		return i
	}

	return -1
}

// 查找第一个等于或者大于key的元素
func findFirstEqualLarger(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return i
}

// 查找第一个大于key的元素
func findFirstLarger(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] <= target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return i
}

// 查找最后一个与target相等的元素
func findLastEqual(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] <= target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	if j >= 0 && nums[j] == target {
		return j
	}

	return -1
}

// 查找最后一个等于或者小于target的元素
func findLastEqualSmaller(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] <= target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return j
}

// 查找最后一个小于target的元素
func findLastSmaller(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return j
}
