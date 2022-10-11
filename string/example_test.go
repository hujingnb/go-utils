package hstring

import "fmt"

func ExampleReverse() {
	s := "abc"

	ret := Reverse(s)
	fmt.Println(ret)
	/*
		Output:
			"cba"
	*/
}

func ExampleCamelToSnake() {
	s := "helloWord"
	ret := CamelToSnake(s)
	fmt.Println(ret)
	/*
		Output:
			"hello_word"
	*/
}

func ExampleSnakeToCamel() {
	s := "hello_word"
	ret := SnakeToCamel(s)
	fmt.Println(ret)
	/*
		Output:
			"helloWord"
	*/
}

func ExampleToString() {
	i := 5
	ret := ToString(i)
	fmt.Println(ret)
	/*
		Output:
			"5"
	*/
}

func ExampleGet32Md5() {
	s := "hello"
	ret := Get32Md5(s)
	fmt.Println(ret)
	/*
		Output:
			5d41402abc4b2a76b9719d911017c592
	*/
}

func ExampleGet16Md5() {
	s := "hello"
	ret := Get16Md5(s)
	fmt.Println(ret)
	/*
		Output:
			bc4b2a76b9719d91
	*/
}
