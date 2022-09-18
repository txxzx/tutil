package tutil

import (
	"crypto/rand"
	"math/big"
	mrand "math/rand"
)

/**
    @date: 2022/9/18
**/

var (
	chars = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// 生成随机字符串
func RandString(l int) string {
	bs := []byte{}
	for i := 0; i < l; i++ {
		bs = append(bs, chars[mrand.Intn(len(chars))])
	}
	return string(bs)
}

// RandByte
func RandByte(l int) []byte {
	bs := []byte{}
	for i := 0; i < l; i++ {
		bs = append(bs, chars[mrand.Intn(len(chars))])
	}
	return bs
}

// 生成指定范围内的数
func GenPiecesCount(min, max int64) int32 {
	maxBigInt := big.NewInt(max)
	for {
		i, _ := rand.Int(rand.Reader, maxBigInt)
		if i.Int64() >= min {
			return int32(i.Int64())
		}
	}
}
