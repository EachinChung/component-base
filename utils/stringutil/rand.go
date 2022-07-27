package stringutil

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// RandCode 随机验证码
func RandCode() string {
	b, err := rand.Int(rand.Reader, big.NewInt(10000))
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%04d", b)
}
