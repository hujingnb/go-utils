package hsystem

import "fmt"

func ExamplePrintOutAndErrToFile() {
	err := PrintOutAndErrToFile("cache.log")
	if err != nil {
		fmt.Println(err)
	}
}
