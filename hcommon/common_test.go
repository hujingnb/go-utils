package hcommon

import (
	"context"
	"fmt"
	"github.com/hujingnb/go-utils/harray"
	"github.com/hujingnb/go-utils/hmap"
	"testing"
	"time"
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

func TestGetContextKeys(t *testing.T) {
	// 多个KV
	multKv := context.WithValue(
		context.WithValue(context.Background(), "k1", "v1"),
		"k2",
		"v2",
	)
	// 最后一个为非 KV
	endNotValueKv, _ := context.WithCancel(multKv)
	// 中间存在非 KV
	middleNotValueKv := context.WithValue(endNotValueKv, "k3", "v3")
	// 定义结构体
	user := struct {
		Name string
		Age  int
	}{"xiaoming", 2}
	testList := []struct {
		ctx  context.Context
		keys []interface{}
	}{
		{ // 正常情况
			ctx:  context.WithValue(context.Background(), "k1", "v1"),
			keys: []interface{}{"k1"},
		},
		{ // 无KV
			ctx:  context.Background(),
			keys: nil,
		},
		// 多个KV, 最后元素非 value
		{
			ctx:  endNotValueKv,
			keys: []interface{}{"k2", "k1"},
		},
		{ // 多个KV, 最后元素为 KV
			ctx:  middleNotValueKv,
			keys: []interface{}{"k3", "k2", "k1"},
		},
		{ // 存在重复K
			ctx: context.WithValue(
				context.WithValue(context.Background(), "k1", "v1"),
				"k1",
				"v1",
			),
			keys: []interface{}{"k1"},
		},
		{ // K 为结构体
			ctx:  context.WithValue(context.Background(), user, "v1"),
			keys: []interface{}{user},
		},
	}
	for index, test := range testList {
		t.Run(fmt.Sprintf("%d", index), func(t *testing.T) {
			retKeys := GetContextKeys(test.ctx)
			// 判断与预期结果是否一致
			if len(retKeys) != len(test.keys) {
				t.Error("error")
				return
			} else {
				for i := 0; i < len(retKeys); i++ {
					if retKeys[i] != test.keys[i] {
						t.Error("error")
						return
					}
				}
			}
		})
	}
}

func TestReadChannelList(t *testing.T) {
	// 阻塞测试
	t.Run("block", func(t *testing.T) {
		ch := make(chan int)
		go func() {
			time.Sleep(time.Second)
			ch <- 1
		}()
		ret, ok := ReadChannelList([]chan int{ch}, true)
		if !ok || ret != 1 {
			t.Error("fail")
		}
	})
	// 非阻塞测试
	t.Run("not_block", func(t *testing.T) {
		wait := make(chan int)
		ch := make(chan int)
		go func() {
			time.Sleep(time.Second)
			ch <- 1
			wait <- 1
		}()
		// 首次读取失败
		ret, ok := ReadChannelList([]chan int{ch}, false)
		if ok || ret != 0 {
			t.Error("fail")
		}
		// 等待发送完毕再次读取
		time.Sleep(time.Second)
		ret, ok = ReadChannelList([]chan int{ch}, false)
		if !ok || ret != 1 {
			t.Error("fail")
		}
	})
}
