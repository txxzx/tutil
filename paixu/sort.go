package paixu

import "sort"

/**
    @date: 2022/8/8
**/

type student struct {
	name string
	age int
}

func sortPaiXu(){
	stu := []student{{"h",20},{"a",23},{"h",21}}
	sort.Slice(stu,func(i,j int)bool{
		if stu[i].name == stu[j].name {
			return  stu[i].age  > stu[j].age
		}
		return  stu[i].name < stu[j].name
	})
}

