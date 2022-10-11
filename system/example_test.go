package hsystem

import "fmt"

// 将标准输出和错误输出写入到文件中
func ExamplePrintOutAndErrToFile() {
	err := PrintOutAndErrToFile("cache.log")
	if err != nil {
		fmt.Println(err)
	}
}
