package main

import (
	"fmt"

	"./example"
	"github.com/golang/protobuf/proto"
)

func main() {
	// 建立一個 User 格式，並在其中放入資料。
	data := example.User{
		Id:       12345,
		Username: "Yami Odymel",
		Password: "test",
	}

	// 將資料編碼成 Protocol Buffer 格式（請注意是傳入 Pointer）。
	dataBuffer, _ := proto.Marshal(&data)

	// 將已經編碼的資料解碼成 protobuf.User 格式。
	var user example.User
	proto.Unmarshal(dataBuffer, &user)

	// 輸出解碼結果。
	fmt.Println(user.Id, " ", user.Username, " ", user.Password)
}
