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
