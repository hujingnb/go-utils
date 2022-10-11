// Package hnumber 数字处理
package hnumber

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// RandRange 获得随机数,范围: [min, max]
func RandRange(min int, max int) (int, error) {
	if min > max {
		return 0, errors.New(fmt.Sprintf("min %d need less max %d", min, max))
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := r.Intn(max + 1 - min)
	return randNum + min, nil
}
