package tutil

import (
	"fmt"
	"testing"
)

/**
    @date: 2022/11/22
**/

func TestAES(t *testing.T) {
	// cipher key
	key := "thisis32bitlongpassphraseimusing"
	// plaintext
	pt := "This is a secret"
	c := EncryptAES([]byte(key), pt)
	// plaintext
	fmt.Println(pt)
	// ciphertext
	fmt.Println(c)
	// decrypt
	DecryptAES([]byte(key), c)
}
