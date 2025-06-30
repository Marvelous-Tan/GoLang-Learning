package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{
		ID: 1,
		Name: "John Doe",
		Email: "john.doe@example.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))

	var newUser User
	err = json.Unmarshal(jsonData, &newUser)
	if err != nil {
		fmt.Println("Error unmarshalling from JSON:", err)
		return
	}
}
/*
  json.Marshal() 将结构体转换为json字符串
  json.Unmarshal() 将json字符串转换为结构体
  json.MarshalIndent() 将结构体转换为json字符串，并添加缩进
  json.UnmarshalIndent() 将json字符串转换为结构体，并添加缩进
  json.MarshalIndent() 将结构体转换为json字符串，并添加缩进
*/