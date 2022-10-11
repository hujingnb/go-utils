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
