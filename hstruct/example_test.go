package hstruct

import (
	"fmt"
	"reflect"
)

func ExampleToMap() {
	s := struct {
		Data string `json:"data"`
	}{
		Data: "1234",
	}
	m, err := ToMap(s, "json")
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
	// Output: map[data:1234]
}

func ExampleName() {
	s := reflect.Value{}
	name, err := Name(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(name)
	// Output: Value
}
