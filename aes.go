package tutil

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

/**
    @date: 2022/11/22
**/
// 进行加密的方法 密钥是32位
func EncryptAES(key []byte, plaintext string) string {
	// create cipher
	c, err := aes.NewCipher(key)
	CheckError(err)
	out := make([]byte, len(plaintext))
	c.Encrypt(out, []byte(plaintext))
	return hex.EncodeToString(out)
}

// 使用AES算法进行解密
func DecryptAES(key []byte, ct string) {
	cipphertext, _ := hex.DecodeString(ct)
	c, err := aes.NewCipher(key)
	CheckError(err)
	pt := make([]byte, len(cipphertext))
	c.Decrypt(pt, cipphertext)
	s := string(pt[:])
	fmt.Println("DECRYPTED:", s)
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
