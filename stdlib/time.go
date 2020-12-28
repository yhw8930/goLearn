package main

import (
	"fmt"
	"goLearn/consts"
	"time"
)

func main() {
	timeDemo()
	timestampDemo()
	timestampDemo2()
	stringToTime()
	timeToString()
	timeArithmetic()
}

//stdlib.Time类型表示时间
func timeDemo() {
	now := time.Now() //获取当前时间
	fmt.Printf("current stdlib:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数
func timestampDemo() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

//时间戳转为时间格式
func timestampDemo2() {
	timestamp := time.Now().Unix()
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	utcTime := timeObj.UTC()           // 获取UTC时区的时间
	//时间戳转化为日期
	datetime := time.Unix(timestamp, 0).Format(consts.DefaultTimeFormat)
	fmt.Printf("stdlib :%+v\nutctime :%+v\ndate :%+v\n", timeObj, utcTime, datetime)
}

//字符串转time
func stringToTime() {
	dateTime := "2015-01-01 12:03:00"
	//location, _ := stdlib.LoadLocation("Local")
	date, _ := time.ParseInLocation(consts.DefaultTimeFormat, dateTime, time.Local)
	fmt.Printf("date: %+v\n", date)
}

//time转字符串
func timeToString() {
	now := time.Now()
	date := now.Format(consts.DefaultTimeFormat)
	fmt.Printf("date: %+v\n", date)
}

func timeArithmetic() {
	now := time.Now()
	add := now.Add(time.Hour)
	fmt.Printf("nextHour: %+v\n", add)
	fmt.Println(now.Sub(add))
	fmt.Println(now.Before(add))
}
