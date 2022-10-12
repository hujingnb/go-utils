// Package harray 数组操作
package harray

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

// Chunk 数组切割
func Chunk[T any](array []T, size int) ([][]T, error) {
	if size < 1 {
		return nil, errors.New("size min 1")
	}
	l := len(array)
	result := make([][]T, 0, int(math.Ceil(float64(l)/float64(size))))
	for i, j := 0, 0+size; i < l; i, j = i+size, j+size {
		if j > l {
			j = l
		}
		result = append(result, array[i:j])
	}
	return result, nil
}

// Shuffle 打乱数组顺序
func Shuffle[T any](array []T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	/*
		思路:
			从后向前遍历数组
			每个元素都和他前面的任意一个进行交换
	*/
	for i := len(array) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

// InArray 检查元素是否在数组中
func InArray[T comparable](array []T, searchData T) bool {
	for _, item := range array {
		if item == searchData {
			return true
		}
	}
	return false
}

// Unique 数组去重
func Unique[T comparable](array []T) []T {
	size := len(array)
	result := make([]T, 0, size)
	tmpMap := make(map[T]bool)
	for _, item := range array {
		if _, ok := tmpMap[item]; !ok {
			tmpMap[item] = true
			result = append(result, item)
		}
	}
	return result
}

// Equal 比较两数组是否相等
func Equal[T comparable](arr1, arr2 []T) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	size := len(arr1)
	for i := 0; i < size; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// Intersect 多个数组取交集
func Intersect[T comparable](arrList ...[]T) []T {
	arrLen := len(arrList)
	if arrLen <= 0 {
		return nil
	}
	if arrLen == 1 {
		return arrList[1]
	}
	// 统计第一个数组中存在的数据
	dataMap := make(map[T]bool)
	for _, item := range arrList[0] {
		dataMap[item] = true
	}
	// 依次遍历后面的所有数组, 将在后续数组中不存在的数据剔除
	for i := 1; i < arrLen; i++ {
		tmpDataMap := make(map[T]bool)
		for _, item := range arrList[i] {
			if _, ok := dataMap[item]; ok {
				tmpDataMap[item] = true
			}
		}
		dataMap = tmpDataMap
	}
	// 生成返回结果
	result := make([]T, 0, len(dataMap))
	for k, _ := range dataMap {
		result = append(result, k)
	}
	return result
}

// Diff 计算数组的差集
// 即数据在 arr 中, 但不在其他数组中
func Diff[T comparable](arr []T, arrList ...[]T) []T {
	if len(arr) <= 0 {
		return nil
	}
	if len(arrList) <= 0 {
		return arr
	}
	// 统计数组中的数据
	dataMap := make(map[T]bool)
	for _, item := range arr {
		dataMap[item] = true
	}
	// 遍历其他数组, 将其他数组中存在的数据去掉
	for _, itemArr := range arrList {
		for _, item := range itemArr {
			if _, ok := dataMap[item]; ok {
				dataMap[item] = false
			}
		}
	}
	// 返回结果
	result := make([]T, 0, len(dataMap))
	for k, v := range dataMap {
		if v {
			result = append(result, k)
		}
	}
	return result
}

// BinarySearch 从数组中查找指定数据(二分查找). 请自行保证数组有序.
//
// 返回目标值得数组下标, 若未找到, 则返回-1
//
// comparator 返回值说明(分别适用于升序和降序):
// 	0 : 相等, 目标数据
//  >0: 目标在当前元素的右侧
//  <0: 目标在当前元素的左侧
func BinarySearch[T comparable](arr []T, comparator func(T) int) int {
	tail := 0
	head := len(arr) - 1
	for tail <= head {
		mid := (head + tail) / 2
		midValue := arr[mid]
		tmpPar := comparator(midValue)
		if tmpPar == 0 { // 相等, 找到了
			return mid
		} else if tmpPar > 0 {
			head = mid - 1
		} else {
			tail = mid + 1
		}
	}
	return -1
}
