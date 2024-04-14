package main

import (
	"fmt"
	"log"
	"my-protobuf/basic"
	"time"
)

type logWritter struct {
}

func (writer logWritter) Write(byte []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05" + " " + string(byte)))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWritter))

	//basic.BasicHello()
	basic.BasicUser()
	//basic.ProtoToJsonUser()
	//basic.JsonToProtoUser()
}
