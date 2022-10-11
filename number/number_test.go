package hnumber

import (
	"fmt"
	"testing"
)

func FuzzRandRange(f *testing.F) {
	f.Add(0, 0)
	f.Fuzz(func(t *testing.T, min, max int) {
		ret, err := RandRange(min, max)
		if min > max {
			if err == nil {
				t.Error("can't check min less then max")
			}
		} else {
			if ret < min || ret > max {
				errStr := fmt.Sprintf("value out of range. min %d, max %d, ret %d", min, max, ret)
				t.Error(errStr)
			}
		}
	})
}
