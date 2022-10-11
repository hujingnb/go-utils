package hnumber

import "fmt"

func ExampleRandRange() {
	ret, err := RandRange(2, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
	// Output: 2
}
