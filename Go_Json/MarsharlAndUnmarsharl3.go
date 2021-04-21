package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Hobby []string `json:"hobby"`
}

//当 struct 中的字段没有值时， json.Marshal() 序列化的时候不会忽略这些字段，
//而是默认输出字段的类型零值（例如int和float类型零值是 0，string类型零值是""，对象类型零值是 nil）。
//如果想要在序列序列化时忽略这些没有值的字段时，可以在对应字段添加omitempty tag。
type User2 struct {
	Name  string   `json:"name"`
	Email string   `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
}

func main() {
	u1 := User{
		Name: "七米",
	}
	// struct -> json string
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)

	u2 := User2{
		Name: "七米",
	}
	// struct -> json string
	b, err = json.Marshal(u2)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)

	a := User2{}
	err = json.Unmarshal(b, &a)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("a:%#v\n", a)
}
