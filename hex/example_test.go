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
