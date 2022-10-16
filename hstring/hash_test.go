package hstring

import (
	"testing"
)

func TestGet32Md5(t *testing.T) {
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
			ret := Get32Md5(test.Input)
			if ret != test.Output {
				t.Error("output error")
			}
		})
	}
}

func TestGet16Md5(t *testing.T) {
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
			ret := Get16Md5(test.Input)
			if ret != test.Output {
				t.Error("output error")
			}
		})
	}
}

func TestGetSha1(t *testing.T) {
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
			ret := GetSha1(test.Input)
			if ret != test.Output {
				t.Error("output error")
			}
		})
	}
}
