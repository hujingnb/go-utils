package hstruct

import (
	"fmt"
	"testing"
)

func TestToMap(t *testing.T) {
	// 未指定tag
	s1 := struct {
		V1 int
		V2 string
		V3 struct {
			Data string
		}
	}{
		V1: 20,
		V2: "test",
		V3: struct {
			Data string
		}{
			Data: "v3",
		},
	}
	m1 := ToMap(s1, "")
	if s1.V1 != m1["V1"] ||
		s1.V2 != m1["V2"] ||
		s1.V3 != m1["V3"] ||
		fmt.Sprintf("%T", m1["V1"]) != fmt.Sprintf("%T", s1.V1) ||
		fmt.Sprintf("%T", m1["V2"]) != fmt.Sprintf("%T", s1.V2) ||
		fmt.Sprintf("%T", m1["V3"]) != fmt.Sprintf("%T", s1.V3) {
		t.Error("fail")
	}

	// 指定tag
	s2 := struct {
		V1 int `tt:"v1"`
		V2 int `tt:"-"`
		V3 int `tt:""`
		V4 int
		V5 int `tt:"vv"`
	}{
		V1: 1,
		V2: 2,
		V3: 3,
		V4: 4,
		V5: 5,
	}
	m2 := ToMap(s2, "tt")
	if m2["v1"] != s2.V1 ||
		m2["v2"] != nil ||
		m2["V2"] != nil ||
		m2["v3"] != nil ||
		m2["v4"] != nil ||
		m2["v5"] != nil ||
		m2["vv"] != s2.V5 {
		t.Error("fail")
	}
	// 指针
	s3 := &struct {
		V1 int
	}{
		V1: 1,
	}
	m3 := ToMap(s3, "")
	if m3["V1"] != s3.V1 {
		t.Error("fail")
	}
}
