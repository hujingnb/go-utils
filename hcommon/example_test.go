package hcommon

import "fmt"

func ExampleCopy() {
	arr := []int{1, 2, 3}
	ret := Copy(arr)
	arr[0] = 5
	fmt.Println(ret)
	// Output: [1 2 3]
}
