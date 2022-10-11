package hstruct

import "fmt"

func ExampleToMap() {
	s := struct {
		Data string `json:"data"`
	}{
		Data: "1234",
	}
	m := ToMap(s, "json")
	fmt.Println(m)
	/*
		Output:
			{"data": "1234}
	*/
}
