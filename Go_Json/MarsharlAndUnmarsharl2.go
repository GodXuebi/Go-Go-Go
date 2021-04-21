package main

import (
	"encoding/json"
	"fmt"
)

//序列化与反序列化默认情况下使用结构体的字段名，我们可以通过给结构体字段添加tag来指定json序列化生成的字段名。
type Person struct {
	Name   string `json:"name"` // 指定json序列化/反序列化时使用小写name
	Age    int64
	Weight float64
}

//如果你想在json序列化/反序列化的时候忽略掉结构体中的某个字段，可以按如下方式在tag中添加-
type Person2 struct {
	Name   string `json:"name"` // 指定json序列化/反序列化时使用小写name
	Age    int64
	Weight float64 `json:"-"` // 指定json序列化/反序列化时忽略此字段
}

func main() {
	p1 := Person2{
		Name:   "Jane",
		Age:    18,
		Weight: 71.6,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}

	fmt.Printf("str:%s\n", b) //str:{"name":"Jane"}

	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	//p2:main.Person{Name:"Jane", Age:0, Weight:0}
	fmt.Printf("p2:%#v\n", p2)
}
