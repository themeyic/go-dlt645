package main

import (
	"fmt"
	dlt "go-dlt645"
	"go-dlt645/dltcon"
	"time"
)

func main(){
	//调用ClientProvider的构造函数,返回结构体指针
	p := dlt.NewClientProvider()
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
	test := dlt.Dlt645ConfigClient{"005223440001","02010300"}
	for {
		value,err := test.SendMessageToSerial(client)
		if err != nil {
			fmt.Println("readHoldErr,", err)
		} else {
			fmt.Printf("%#v\n", value)
		}

		time.Sleep(time.Second * 3)
	}
}