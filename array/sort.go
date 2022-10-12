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
