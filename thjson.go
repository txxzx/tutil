package tutil

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/swxctx/xlog"
)

/**
    @date: 2022/8/15
**/

/*
第三方json转为 对应的结构体
*/
func ThJson() {
	jsonRes := `{
              "Id":101,
               "Name":"frank"
    }`
	data := new(User)
	if err := json.Unmarshal([]byte(jsonRes), &data); err != nil {
		log.Printf("err-> %v", err)
		return
	}
	fmt.Printf("data=%+v", data)
}

type User struct {
	Id   int
	Name string
}

// 第三方json转为map,无论map里面value是什么类型都会转化，不会报错
func JsonToMap() {
	jsonR := `{
	"Id":101,
     "Name":"frank"
}`
	data2 := make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonR), &data2); err != nil {
		xlog.Errorf("err-> %v", err)
		return
	}
	fmt.Printf("data2-> %v", data2)
}

// 第三方json转为map方便进行读取map里面的字段值,再将map转为结构体
func JsonToMapStr() {
	jsonRe := `{
	"Id":101,
	"Name":"frank",
"Details":{
"Gender":"man",
"Height":190,
"Age":189,
"Address":"北京"
}
}`
	// 先把第三方json转为Map
	data3 := make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonRe), &data3); err != nil {
		xlog.Errorf("err-> %v", err)
		return
	}
	// 定义一个结构
	data4 := new(Student)
	// 将map转为结构体
	if err := mapstructure.Decode(data3, data4); err != nil {
		xlog.Errorf("err-> %v", err)
		return
	}
	fmt.Printf("data4-> %v", data4)
	fmt.Printf("data4.Details.Age-> %d", data4.Details.Age)
}

type Student struct {
	Id      int
	Name    string
	Details Details
}
type Details struct {
	Gender  string
	Age     int
	Address string
}
