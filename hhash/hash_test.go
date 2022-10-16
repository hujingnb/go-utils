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

func TestSha256(t *testing.T) {
	testList := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "hello",
			Output: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
	}
	for _, test := range testList {
		t.Run(test.Input, func(t *testing.T) {
			if Sha256(test.Input) != test.Output {
				t.Error("sha256 fail")
			}
		})
	}
}

func TestGetSha512(t *testing.T) {
	testList := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "hello",
			Output: "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043",
		},
	}
	for _, test := range testList {
		t.Run(test.Input, func(t *testing.T) {
			if Sha512(test.Input) != test.Output {
				t.Error("sha512 fail")
			}
		})
	}
}
