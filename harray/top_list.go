package harray

import (
	"sort"
)

// NewTopList 创建一个 top n 的队列
// 参数:
// size: 队列容量
// canRepeat: 元素是否可重
// comparator: 元素比较函数. 相等为0
func NewTopList[T any](size int, canRepeat bool, comparator func(a, b T) int) *TopList[T] {
	return &TopList[T]{
		size:       size,
		data:       make([]T, size),
		len:        0,
		repeat:     canRepeat,
		comparator: comparator,
	}
}

// TopList 用来维护一个最大值队列
// 内部通过二分查找, 故查找的时间复杂度为 O(lgn)
// 数据更新时, 会发生数组数据复制, 故更新时时间复杂度为 O(n)
type TopList[T any] struct {
	size       int              // 当前数组大小
	data       []T              // 存放数据的数组. 从大到小排序
	len        int              // 数组中当前存放的数据大小
	repeat     bool             // 数据是否可重
	comparator func(a, b T) int // 比较数据大小
}

// Add 添加新值
func (t *TopList[T]) Add(i ...T) {
	// 获取当前值插入的位置
	for _, v := range i {
		t.addOne(v)
	}
}

func (t *TopList[T]) addOne(i T) {
	insertIndex := t.getInsertIndex(i)
	if insertIndex < 0 || insertIndex >= t.size {
		return
	}
	// 值已经存在
	oldValue := t.data[insertIndex]
	if !t.repeat && t.comparator(i, oldValue) == 0 {
		return
	}
	// 将值插入
	t.data[insertIndex] = i
	// 这个位置原来是空的, 将数据放入直接返回即可
	if t.len <= insertIndex {
		t.len++
		return
	}
	// 若这个位置原来有值, 向后顺延
	// 若超过长度就丢弃
	lastValue := oldValue
	for j := insertIndex + 1; j < t.size; j++ {
		t.data[j], lastValue = lastValue, t.data[j]
	}
	if t.len < t.size {
		t.len++
	}
}

func (t *TopList[T]) getInsertIndex(i T) int {
	if t.len <= 0 { // 当前还没有数据, 直接放入第一个
		return 0
	}
	if t.comparator(i, t.data[0]) >= 0 { // 最大, 放入第一个
		return 0
	}
	if t.comparator(i, t.data[t.len-1]) <= 0 { // 最小, 放入最后一个
		return t.len
	}
	// 二分查找数据的位置
	return sort.Search(t.len, func(index int) bool {
		return t.comparator(t.data[index], i) <= 0
	})
}

// GetData 获取当前内容
func (t *TopList[T]) GetData() []T {
	return t.data[:t.len]
}
