package tutil

/**
    @date: 2022/9/18
**/

import (
	"testing"
)
func TestRand(t *testing.T){
	str := RandString(10)
	t.Logf("string -> %s",str)

	by := RandByte(10)
	t.Logf("tr-> %s",by)

	i := GenPiecesCount(3,10)
	t.Logf("%d",i)

	e := Reverse(4)
	t.Logf("%d",e)
}