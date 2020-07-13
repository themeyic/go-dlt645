package dlt645

import (
	"fmt"
	"strconv"
	"strings"
)


type Dlt645ConfigClient struct {
	MeterNumber string
	DataMarker string
}

func (dltconfig *Dlt645ConfigClient) SendMessageToSerial(dlt Client) (response string, err error){
	//表号
	meterNumberHandle := HexStringToBytes(dltconfig.MeterNumber)
	meterNumberHandleX := fmt.Sprintf("% x",meterNumberHandle)
	meterNumberHandleReverse := strings.Split(meterNumberHandleX," ")
	for i :=0;i<len(meterNumberHandleReverse)/2;i++{
		mid := meterNumberHandleReverse[i]
		meterNumberHandleReverse[i] = meterNumberHandleReverse[len(meterNumberHandleReverse)-1-i]
		meterNumberHandleReverse[len(meterNumberHandleReverse)-1-i] =mid
	}
	midMeterNumberHandle := fmt.Sprintf("% s",meterNumberHandleReverse)
	meterNumberHandleReverseFinished := strings.Replace(midMeterNumberHandle,"[","",-1)
	meterNumberHandleReverseFinished = strings.Replace(meterNumberHandleReverseFinished,"]","",-1)
	//数据标识
	DataMarkerHandle := HexStringToBytes(dltconfig.DataMarker)
	DataMarkerHandleX := fmt.Sprintf("% x",DataMarkerHandle)
	DataMarkerHandleReverse := strings.Split(DataMarkerHandleX," ")
	for i :=0;i<len(DataMarkerHandleReverse)/2;i++{
		mid := DataMarkerHandleReverse[i]
		DataMarkerForEnd,_ := strconv.Atoi(DataMarkerHandleReverse[len(DataMarkerHandleReverse)-1-i])
		DataMarkerHandleReverse[i] =  strconv.Itoa(DataMarkerForEnd+33)
		DataMarkerForStrt,_ := strconv.Atoi(mid)
		DataMarkerHandleReverse[len(DataMarkerHandleReverse)-1-i] =strconv.Itoa(DataMarkerForStrt+33)
	}
	midDataMarkerHandle := fmt.Sprintf("% s",DataMarkerHandleReverse)
	DataMarkerHandleReverseFinished := strings.Replace(midDataMarkerHandle,"[","",-1)
	DataMarkerHandleReverseFinished = strings.Replace(DataMarkerHandleReverseFinished,"]","",-1)


	messageFinshed :="68 "+meterNumberHandleReverseFinished+" 68"+" 11 "+"04 "+DataMarkerHandleReverseFinished+" 71 16"
	value,err :=dlt.SendRawFrame(messageFinshed)
	return value,err
}

func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from + 1, to - 1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}