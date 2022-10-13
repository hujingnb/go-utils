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

// 32位md5
func ExampleGet32Md5() {
	s := "hello"
	ret := Get32Md5(s)
	fmt.Println(ret)
	// Output: 5d41402abc4b2a76b9719d911017c592
}

// 16位 md5
func ExampleGet16Md5() {
	s := "hello"
	ret := Get16Md5(s)
	fmt.Println(ret)
	// Output: bc4b2a76b9719d91
}

// sha1
func ExampleGetSha1() {
	s := "hello"
	ret := GetSha1(s)
	fmt.Println(ret)
	// Output: aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
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
