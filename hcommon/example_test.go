package hcommon

import (
	"context"
	"fmt"
)

func ExampleCopy() {
	arr := []int{1, 2, 3}
	ret := Copy(arr)
	arr[0] = 5
	fmt.Println(ret)
	// Output: [1 2 3]
}

func ExampleGetContextKeys() {
	ctx := context.WithValue(context.Background(), "k1", "v1")
	fmt.Println(GetContextKeys(ctx))
	// Output: [k1]
}
