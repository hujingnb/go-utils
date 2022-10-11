package hex

import "fmt"

// ExampleInput 将字符串按照特定进制读取
func ExampleInput() {
	numStr := "1330"
	radix := 9

	num, err := Input(numStr, radix)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
	/*
		OutPut:
			999, nil
	*/
}

// ExampleOutput 将数字转为目标进制
func ExampleOutput() {
	num := int64(1330)
	radix := 16

	numStr, err := Output(num, radix)
	if err != nil {
		panic(err)
	}
	fmt.Println(numStr)
	/*
		OutPut:
			"1330", nil
	*/
}
