package main

import (
    "github.com/golang/protobuf/proto"
    "protobuf用法/protobuf/packet"
    "io/ioutil"
    "os"
    "fmt"
)

func write() {
    packetContent_data := &packet.Packet_Content {
        CToID:   "WQ",
        CFromID: "FJJ",
        CContext: "Hello, WQ!",
        CFlag: 1,
        COther: "",
    };
    packetBody_data := &packet.Packet_Body {
        BType:   1,
        BAction: "hello",
        BContentType: packet.Packet_TEXT,
        BContent: packetContent_data,
        BTimestamp: 1586502239103,
    };

    //编码数据
    data, _ := proto.Marshal(packetBody_data)
    //把数据写入文件
    ioutil.WriteFile("./test.txt", data, os.ModePerm)
}

func read() {
    //读取文件数据
    data, _ := ioutil.ReadFile("./test.txt")
    packetBody_data := &packet.Packet_Body{}
    //解码数据
    proto.Unmarshal(data, packetBody_data)
    fmt.Printf("%v\n", packetBody_data)
}

func main() {
    write()
    read()
}
