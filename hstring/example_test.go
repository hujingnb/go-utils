package hstring

import "fmt"

// 字符串反转
func ExampleReverse() {
	s := "abc"

	ret := Reverse(s)
	fmt.Println(ret)
	// Output: cba
}

// 驼峰转蛇形
func ExampleCamelToSnake() {
	s := "helloWord"
	ret := CamelToSnake(s)
	fmt.Println(ret)
	// Output: hello_word
}

// 蛇形转驼峰
func ExampleSnakeToCamel() {
	s := "hello_word"
	ret := SnakeToCamel(s)
	fmt.Println(ret)
	// Output: helloWord
}

// 转字符串
func ExampleToString() {
	i := 5
	ret := ToString(i)
	fmt.Println(ret)
	// Output: 5
}

// 字符串填充
func ExamplePad() {
	ret := Pad(1, 5, "0", PadLeft)
	fmt.Println(ret)
	// Output: 00001
}

func ExampleJoin() {
	arr := []int{1, 2, 3}
	ret := Join(arr, ",")
	fmt.Println(ret)
	// Output: 1,2,3
}
