package golang

// **************************************************************
// 归并排序算法
// **************************************************************

// MergeSort 归并排序
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	left := MergeSort(arr[0:i])
	right := MergeSort(arr[i:])
	result := merge(left, right)
	return result
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	l, r := len(left), len(right)
	m, n := 0, 0 // left和right的index位置
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
		} else {
			result = append(result, left[m])
			m++
		}
	}
	result = append(result, right[n:]...)
	result = append(result, left[m:]...)
	return result
}
