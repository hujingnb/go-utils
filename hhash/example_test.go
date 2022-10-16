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

func ExampleSha256() {
	s := "hello"
	ret := Sha256(s)
	fmt.Println(ret)
	// Output: 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
}

func ExampleSha512() {
	s := "hello"
	ret := Sha512(s)
	fmt.Println(ret)
	// Output: 9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043
}

func ExampleCrc32() {
	s := "hello"
	ret := Crc32(s)
	fmt.Println(ret)
	// Output: 907060870
}
