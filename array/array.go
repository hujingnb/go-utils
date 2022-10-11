// Package harray 数组操作
package harray

import (
	"errors"
	"math"
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
