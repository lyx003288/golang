package golang

// **************************************************************
// 快速排序算法
// **************************************************************

// QuickSort 快速排序
// l: source起始下标
// u: 元素个数
func QuickSort(source []int, l, u int) {
	if l < u {
		m := partition(source, l, u)
		QuickSort(source, l, m)
		QuickSort(source, m+1, u)
	}
}

func partition(source []int, l, u int) int { //划分
	var (
		pivot = source[l]
		left  = l
		right = l + 1
	)
	for ; right < u; right++ {
		if source[right] <= pivot {
			left++
			source[left], source[right] = source[right], source[left]
		}
	}
	source[l], source[left] = source[left], source[l]
	return left
}
