package hcommon

import (
	"github.com/hujingnb/go-utils/harray"
	"github.com/hujingnb/go-utils/hmap"
	"testing"
)

func TestCopy(t *testing.T) {
	// 结构体
	t.Run("struct", func(t *testing.T) {
		tmpInt := 10
		s := struct {
			V1 string // 基础类型
			V2 int
			V3 struct { // 结构体
				V4 string
			}
			V5 *struct { // 结构体指针
				V6 int
			}
			V7 interface{} // 接口
			V8 *struct {   // 空指针
			}
			V9  interface{}    // 空指针
			V10 *int           // 基础类型指针
			V11 []int          // 数组
			V12 []int          // 空数组
			V13 []int          // 数组空指针
			V14 map[int]string // map
			V15 map[int]string // 空map
			V16 map[int]string // map 空指针
		}{
			V1: "v1",
			V2: 5,
			V3: struct {
				V4 string
			}{
				V4: "v4",
			},
			V5: &struct {
				V6 int
			}{
				V6: 9,
			},
			V7:  "v7",
			V10: &tmpInt,
			V11: []int{1, 2, 3, 4},
			V12: []int{},
			V14: map[int]string{
				5: "aa",
			},
			V15: make(map[int]string),
		}
		ret := Copy(s)
		if ret.V1 != s.V1 ||
			ret.V2 != s.V2 ||
			ret.V3.V4 != s.V3.V4 ||
			ret.V5 == nil ||
			ret.V5.V6 != s.V5.V6 ||
			ret.V7 != s.V7 ||
			ret.V8 != nil ||
			ret.V9 != nil ||
			*ret.V10 != *s.V10 ||
			!harray.Equal(ret.V11, s.V11) ||
			!harray.Equal(ret.V12, s.V12) ||
			!harray.Equal(ret.V13, s.V13) ||
			!hmap.Equal(ret.V14, s.V14) ||
			!hmap.Equal(ret.V15, s.V15) ||
			!hmap.Equal(ret.V16, s.V16) {
			t.Error("fail")
		}
		s.V2++
		s.V5.V6++
		*s.V10++
		s.V11 = append(s.V11, 5)
		if ret.V2 == s.V2 ||
			ret.V5.V6 == s.V5.V6 ||
			*ret.V10 == *s.V10 ||
			harray.Equal(s.V11, ret.V11) { // 非深度 copy
			t.Error("not depth copy")
		}
	})
	// 结构体指针
	t.Run("struct_ptr", func(t *testing.T) {
		s := &struct {
			V1 int
		}{
			V1: 5,
		}
		ret := Copy(s)
		if ret.V1 != s.V1 {
			t.Error("fail")
		}
		s.V1++
		if ret.V1 == s.V1 {
			t.Error("not depth copy")
		}
	})
	// 基础类型, 整形
	t.Run("int", func(t *testing.T) {
		i := 5
		ret := Copy(i)
		if ret != i {
			t.Error("fail")
		}
		i++
		if ret == i {
			t.Error("not depth copy")
		}
	})
	// 基础类型, 整形, 指针
	t.Run("int_ptr", func(t *testing.T) {
		tmp := 5
		i := &tmp
		ret := Copy(i)
		if ret == i ||
			*ret != *i {
			t.Error("fail")
		}
		*i++
		if *ret == *i {
			t.Error("not depth copy")
		}
	})
	// 数组
	t.Run("array", func(t *testing.T) {
		arr := [3]int{1, 2, 3}
		ret := Copy(arr)
		for i := 0; i < len(arr); i++ {
			if arr[i] != ret[i] {
				t.Error("fail")
			}
		}
		// 修改数值
		arr[2]++
		if ret[2] == arr[2] {
			t.Error("not depth copy")
		}
	})
	// 切片
	t.Run("slice", func(t *testing.T) {
		arr := []int{1, 2, 3}
		ret := Copy(arr)
		if !harray.Equal(arr, ret) {
			t.Error("fail")
		}
		// 修改单个数值
		arr[0]++
		if harray.Equal(arr, ret) {
			t.Error("not depth copy")
		}
		// 增加元素
		arr = append(arr, 5)
		if harray.Equal(arr, ret) {
			t.Error("not depth copy")
		}
	})
	// 空切片
	t.Run("slice_empty", func(t *testing.T) {
		arr := []int{}
		ret := Copy(arr)
		if !harray.Equal(arr, ret) {
			t.Error("fail")
		}
		arr = append(arr, 2)
		if harray.Equal(arr, ret) {
			t.Error("not depth copy")
		}
	})
	// 结构体切片
	t.Run("slice_struct", func(t *testing.T) {
		arr := []struct {
			V1 int
		}{
			{
				V1: 2,
			},
		}
		ret := Copy(arr)
		if !harray.Equal(arr, ret) {
			t.Error("fail")
		}
		arr[0].V1++
		if harray.Equal(arr, ret) {
			t.Error("not depth copy")
		}
	})
	// 结构体指针切片
	t.Run("slice_struct_ptr", func(t *testing.T) {
		arr := []*struct {
			V1 int
		}{
			{
				V1: 2,
			},
		}
		ret := Copy(arr)
		if len(arr) != len(ret) {
			t.Error("fail")
		}
		for i := 0; i < len(arr); i++ {
			if arr[i].V1 != ret[i].V1 {
				t.Error("fail")
			}
		}
		// 修改值
		arr[0].V1++
		if arr[0].V1 == ret[0].V1 {
			t.Error("not depth copy")
		}

	})
	// map
	t.Run("map", func(t *testing.T) {
		m := map[string]int{
			"a": 1,
			"b": 2,
		}
		ret := Copy(m)
		if !hmap.Equal(m, ret) {
			t.Error("fail")
		}
		m["a"]++
		if m["a"] == ret["a"] {
			t.Error("not depth copy")
		}
	})
	// map struct
	t.Run("map_struct", func(t *testing.T) {
		m := map[string]struct {
			V1 int
		}{
			"a": {
				V1: 2,
			},
		}
		ret := Copy(m)
		if !hmap.Equal(m, ret) {
			t.Error("fail")
		}
	})
	// 布尔
	t.Run("bool", func(t *testing.T) {
		b := true
		ret := Copy(b)
		if ret != b {
			t.Error("fail")
		}
		b = false
		if ret == b {
			t.Error("not depth copy")
		}
	})
	// 管道
	t.Run("channel", func(t *testing.T) {
		c := make(chan int)
		ret := Copy(c)
		if c == ret { // 比较地址
			t.Error("fail")
		}
	})
	// 字符串
	t.Run("string", func(t *testing.T) {
		s := "hello"
		ret := Copy(s)
		if ret != s || &s == &ret {
			t.Error("fail")
		}
	})
	// 字符串指针
	t.Run("string_ptr", func(t *testing.T) {
		tmpS := "hello"
		s := &tmpS
		ret := Copy(s)
		if *ret != *s || ret == s {
			t.Error("fail")
		}
	})
	// 空指针
	t.Run("empty_ptr", func(t *testing.T) {
		var input interface{}
		ret := Copy(input)
		if ret != input {
			t.Error("fail")
		}
	})
	// function
	t.Run("function", func(t *testing.T) {
		f := func(data int) int {
			return data + 1
		}
		ret := Copy(f)
		if f(1) != ret(1) {
			t.Error("fail")
		}
		if &f == &ret {
			t.Error("not depth copy")
		}
	})
}
