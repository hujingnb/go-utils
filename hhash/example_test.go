package hstring

import "fmt"

// 32位md5
func ExampleMd5By32() {
	s := "hello"
	ret := Md5By32(s)
	fmt.Println(ret)
	// Output: 5d41402abc4b2a76b9719d911017c592
}

// 16位 md5
func ExampleMd5By16() {
	s := "hello"
	ret := Md5By16(s)
	fmt.Println(ret)
	// Output: bc4b2a76b9719d91
}

// sha1
func ExampleSha1() {
	s := "hello"
	ret := Sha1(s)
	fmt.Println(ret)
	// Output: aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
}
