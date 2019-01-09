package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	t := time.Now() //2018-07-11 15:07:51.8858085 +0800 CST m=+0.004000001
	fmt.Printf("时间戳（秒）：%v;\n", t.Unix())
	fmt.Printf("时间戳（纳秒）：%v;\n", t.UnixNano())
	fmt.Printf("时间戳（毫秒）：%v;\n", t.UnixNano()/1e6)
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", t.UnixNano()/1e9)
	fmt.Println(t)

	//获取当前时间戳
	fmt.Println(t.Unix()) //1531293019

	//时间 to 时间戳
	loc, _ := time.LoadLocation("Asia/Shanghai")                                     //设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-07-11 15:07:51", loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	fmt.Println(tt.Unix())                                                           //1531292871

	//时间戳 to 时间
	tm := time.Unix(1546985052, 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05")) //2018-07-11 15:10:19

	fmt.Println(TimeToInt64(t))
	fmt.Println(Int64ToTime(TimeToInt64(t)))
	fmt.Println(TimeToInt64Location(t, "America/Antigua"))
	fmt.Println(Int64ToTime(TimeToInt64Location(t, "America/Antigua")))
	fmt.Println(Int64ToTimeLocation(TimeToInt64Location(t, "America/Antigua"), "America/Antigua"))
}

func TimeToInt64(t time.Time) int64 {
	return t.Unix()
}
func Int64ToTime(tt int64) time.Time {
	return time.Unix(tt, 0)
}

//"Asia/Shanghai"
func TimeToInt64Location(t time.Time, locat string) int64 {
	timeStr := t.Format("2006-01-02 15:04:05")
	loc, _ := time.LoadLocation(locat)                                 //设置时区
	tt, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	return tt.Unix()
}
func Int64ToTimeLocation(tt int64, locat string) time.Time {
	loc, _ := time.LoadLocation(locat) //设置时区
	time.Local = loc
	return time.Unix(tt, 0)
}

func TimeToMSInt64(t time.Time) int64 {
	timeStr := t.Format("2006-01-02 15:04:05.000")
	loc, _ := time.LoadLocation("Asia/Shanghai")                       //设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	return tt.UnixNano() / 1e6
}
