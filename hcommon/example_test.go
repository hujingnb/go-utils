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

func ExampleCopyContextValue() {
	ctx := context.WithValue(context.Background(), "k1", "v1")
	fmt.Println(CopyContextValue(ctx).Value("k1"))
	// Output: v1
}

func ExampleReadChannelList() {
	ch := make(chan int, 3)
	ch <- 1
	fmt.Println(ReadChannelList([]chan int{ch}, true))
	// Output: 1 true
}
