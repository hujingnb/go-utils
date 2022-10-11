package hnumber

import "fmt"

// ExampleHexInput 将字符串按照特定进制读取
func ExampleHexInput() {
	numStr := "1330"
	radix := 9

	num, err := HexInput(numStr, radix)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
	// OutPut: 999
}

// ExampleHexOutput 将数字转为目标进制
func ExampleHexOutput() {
	num := int64(1330)
	radix := 16

	numStr, err := HexOutput(num, radix)
	if err != nil {
		panic(err)
	}
	fmt.Println(numStr)
	// OutPut: 532
}

func ExampleRandRange() {
	ret, err := RandRange(2, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
	// Output: 2
}
