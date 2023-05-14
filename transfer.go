package tutil

import (
	"bytes"
	"encoding/binary"
)

/**
    @date: 2023/5/14
**/

// 整型转为字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// 字节转为整型
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

/*
无符号整型表示 表示数字只能是正数不能是负数
int8 -127到127
uint8 0到255
int16 -32768到32767
uint16 0到65535
int32 -2147483648到2147483647
int64 -9223372036854775808到9223372036854775807
uint64 0到18446744073709551615
go里面没有字符类型（char），而是使用byte（uint8）和rune（int32）来代表字符。
我们声明一个字符时，默认是rune类型，除非特别定义。 一个string变量既可以被拆分为字符，也可以被拆分为字节；前者使用rune[]切片表示，后者使用byte[]切片表示
一个rune值就是代表一个字符，在输入输出中经常看到类似’\U0001F3A8’，’\u2665’的就是一个rune字符（unicode字符），其中的每位都是一个16进制数

*/
