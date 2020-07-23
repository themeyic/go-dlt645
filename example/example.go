package main

import (
	"fmt"
	dlt "github.com/themeyic/go-dlt645"
	"github.com/themeyic/go-dlt645/dltcon"
	"time"
)

func main() {
	//调用ClientProvider的构造函数,返回结构体指针
	p := dlt.NewClientProvider()
	//windows 下面就是 com开头的，比如说 com3
	//mac OS 下面就是 /dev/下面的，比如说 dev/tty.usbserial-14320
	p.Address = "com3"
	p.BaudRate = 2400
	p.DataBits = 8
	p.Parity = "E"
	p.StopBits = 1
	p.Timeout = 100 * time.Millisecond

	client := dltcon.NewClient(p)
	client.LogMode(true)
	err := client.Start()
	if err != nil {
		fmt.Println("start err,", err)
		return
	}
	//MeterNumber是表号 005223440001
	//DataMarker是数据标识别 02010300
	test := &dlt.Dlt645ConfigClient{"005223440001", "02010300"}
	for {
		value, err := test.SendMessageToSerial(client)
		if err != nil {
			fmt.Println("readHoldErr,", err)
		} else {
			fmt.Printf("%#v\n", value)
		}
		time.Sleep(time.Second * 3)
	}
}
