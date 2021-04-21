package main

import (
	"encoding/json"
	"fmt"
)

//匿名嵌套Profile时序列化后的json串为单层的
type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
	Profile
}

//嵌套的json串，需要改为具名嵌套或定义字段tag
type User2 struct {
	Name    string   `json:"name"`
	Email   string   `json:"email,omitempty"`
	Hobby   []string `json:"hobby,omitempty"`
	Profile `json:"profile"`
}

type User3 struct {
	Name     string   `json:"name"`
	Email    string   `json:"email,omitempty"`
	Hobby    []string `json:"hobby,omitempty"`
	*Profile `json:"profile,omitempty"`
}

//

type Profile struct {
	Website string `json:"site"`
	Slogan  string `json:"slogan"`
}

func main() {
	u1 := User{
		Name:  "七米",
		Hobby: []string{"足球", "双色球"},
	}
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	//str:{"name":"七米","hobby":["足球","双色球"],"site":"","slogan":""}
	fmt.Printf("str:%s\n", b)

	u2 := User2{
		Name:  "七米",
		Hobby: []string{"足球", "双色球"},
	}
	b, err = json.Marshal(u2)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	//str:{"name":"七米","hobby":["足球","双色球"],"profile":{"site":"","slogan":""}}
	fmt.Printf("str:%s\n", b)

	u3 := User3{
		Name:  "七米",
		Hobby: []string{"足球", "双色球"},
	}
	b, err = json.Marshal(u3)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	//str:{"name":"七米","hobby":["足球","双色球"]}
	fmt.Printf("str:%s\n", b)
}
