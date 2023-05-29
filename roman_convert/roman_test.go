package roman_convert

import (
	"log"
	"testing"
)

/**
    @date: 2023/3/28
**/

func TestRoman(t *testing.T) {
	sr := IntToRoman(2909)
	log.Println(sr)
}

func TestRomanToInt2(t *testing.T) {
	i := RomanToInt2("MMCMIX")
	log.Println(i)
}
