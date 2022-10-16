package hstring

import (
	"testing"
)

func TestMd5By32(t *testing.T) {
	testList := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "hello",
			Output: "5d41402abc4b2a76b9719d911017c592",
		},
	}
	for _, test := range testList {
		t.Run(test.Input, func(t *testing.T) {
			ret := Md5By32(test.Input)
			if ret != test.Output {
				t.Error("output error")
			}
		})
	}
}

func TestMd5By16(t *testing.T) {
	testList := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "hello",
			Output: "bc4b2a76b9719d91",
		},
	}
	for _, test := range testList {
		t.Run(test.Input, func(t *testing.T) {
			ret := Md5By16(test.Input)
			if ret != test.Output {
				t.Error("output error")
			}
		})
	}
}

func TestSha1(t *testing.T) {
	testList := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "hello",
			Output: "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d",
		},
	}
	for _, test := range testList {
		t.Run(test.Input, func(t *testing.T) {
			ret := Sha1(test.Input)
			if ret != test.Output {
				t.Error("output error")
			}
		})
	}
}
