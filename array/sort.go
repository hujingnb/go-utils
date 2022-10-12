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
