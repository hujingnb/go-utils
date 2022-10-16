package hsystem

import (
	"fmt"
	"os"
)

// 将标准输出和错误输出写入到文件中
func ExamplePrintOutAndErrToFile() {
	err := PrintOutAndErrToFile("cache.log")
	if err != nil {
		fmt.Println(err)
	}
	os.Remove("cache.log")
	// Output:
}
