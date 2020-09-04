package main

import (
	// 辅助库
	"github.com/golang/protobuf/proto"
	// test.pb.go 的路径
	"./packet"
)

func main() {

	bodyData := "shenzhen/nanshan/tencent/company"

	p := &packet.StringMessage{
		Body: proto.String(bodyData),
		Header: &packet.Header{
			MessageId: proto.String("1234567890"),
			Topic:     proto.String("golang"),
		},
	}

	pData, err := proto.Marshal(p)

	if err != nil {
		println(err.Error())
	}
	println(string(pData))

	p2 := &packet.StringMessage{}
	proto.Unmarshal(pData, p2)

	println(p2.GetHeader().GetTopic())
}


