package paixu

import (
	"sort"
	"testing"
)

/**
    @date: 2022/8/8
**/

type stu struct {
	name string
	age int
}

type student1 []stu

func (s student1) Len() int {
	return len(s)
}

func (s student1) Less(i, j int) bool {
	if s[i].name == s[j].name{
		return s[i].age > s[j].age
	}
	return  s[i].name < s[j].name
}

func (s student1) Swap (i,j int){
	s[i] ,s[j] = s[j] ,s[i]
}

func Test_Paixu (t *testing.T){
	stu1 := student1{{"h",20},{"a",23},{"h",45}}
	sort.Sort(stu1)
}