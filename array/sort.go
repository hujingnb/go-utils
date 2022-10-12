package harray

// 数组的排序操作

// SortBubble 冒泡排序
func SortBubble[T any](arr []T, comparator func(a, b T) int) {
	length := len(arr)
	for i := length; i > 0; i-- {
		flag := false // 判断本次是否进行了冒泡操作
		for j := 0; j < i-1; j++ {
			if comparator(arr[j], arr[j+1]) > 0 {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

// SortInsert 插入排序
func SortInsert[T any](arr []T, comparator func(a, b T) int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		// 保存当前元素
		tmpData := arr[i]
		// 将左侧所有大于当前元素的值向右运动
		j := i - 1
		for ; j >= 0 && comparator(arr[j], tmpData) > 0; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = tmpData
	}
}

// SortSelect 选择排序
func SortSelect[T any](arr []T, comparator func(a, b T) int) {
	minIndex := 0
	length := len(arr)
	// 对数组进行遍历, 不需要遍历最后一个元素, 因为后面没有元素进行比较了
	for i := 0; i < length-1; i++ {
		minIndex = i // 记录最小元素下标
		// 遍历后面的所有元素. 找到其中最小的元素
		for j := i + 1; j < length; j++ {
			if comparator(arr[j], arr[minIndex]) <= 0 {
				minIndex = j
			}
		}
		// 将最小的元素拿过来, 互换位置
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// SortQuick 快排
func SortQuick[T any](arr []T, comparator func(a, b T) int) {
	sortQuick(arr, 0, len(arr)-1, comparator)
}

func sortQuick[T any](arr []T, head, tail int, comparator func(a, b T) int) {
	if head < 0 || tail < 0 || head >= tail ||
		head >= len(arr) || tail >= len(arr) {
		return
	}
	// 选择基准值
	middle := arr[head]
	i, j := head, tail
	// 将数组拆分
	for i < j {
		// 找到后面小于 middle 的值
		for ; i < j && comparator(arr[j], middle) >= 0; j-- {
		}
		arr[i] = arr[j]
		for ; i < j && comparator(arr[i], middle) <= 0; i++ {
		}
		arr[j] = arr[i]
	}
	arr[i] = middle
	// 对两侧数组再次排序
	sortQuick(arr, head, i-1, comparator)
	sortQuick(arr, i+1, tail, comparator)
}

// SortMerge 归并排序
func SortMerge[T any](arr []T, comparator func(a, b T) int) {
	length := len(arr)
	var size int
	for size = 2; size < length; size *= 2 {
		// 当前步长进行归并
		i := 0
		for ; i < length; i += size {
			tmpTail := i + size - 1
			if tmpTail >= length { // 范围超出
				break
			}
			middle := (i + tmpTail) / 2
			sortMerge(arr, i, tmpTail, middle, comparator)
		}
		// 步长遍历后, 尾部仍存在部分数据数据
		if i < length {
			middle := (length + i) / 2
			sortMerge(arr, i, length-1, middle, comparator)
		}
	}
	// 对最后的两个数组进行合并
	sortMerge(arr, 0, length-1, size/2-1, comparator)
}

// 对两个有序数组进行合并
func sortMerge[T any](arr []T, start, tail, middle int, comparator func(a, b T) int) {
	tmp := make([]T, 0, tail-start+1)
	// 数组合并
	var i, j int
	for i, j = start, middle+1; i <= middle && j <= tail; {
		// 前一个数组较小
		if comparator(arr[i], arr[j]) < 0 {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}
	// 写入剩余数据
	for ; i <= middle; i++ {
		tmp = append(tmp, arr[i])
	}
	for ; j <= tail; j++ {
		tmp = append(tmp, arr[j])
	}
	// 将归并后的数组赋值回原数组中
	tmpLength := len(tmp)
	for x, y := 0, start; x < tmpLength; x, y = x+1, y+1 {
		arr[y] = tmp[x]
	}
}
