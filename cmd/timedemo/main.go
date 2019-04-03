package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	t := time.Now() //2018-07-11 15:07:51.8858085 +0800 CST m=+0.004000001
	fmt.Printf("时间戳（秒）：%v;\n", TimeToInt64(t))
	fmt.Printf("时间戳（纳秒）：%v;\n", TimeToUnixNano(t))
	fmt.Printf("时间戳（毫秒）：%v;\n", TimeToUnixMs(t))
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", t.UnixNano()/1e9)
	fmt.Printf("时间（秒）：%v;\n", Int64ToTime(TimeToInt64(t)))
	//fmt.Println(Unix(TimeToInt64(t), TimeToUnixNano(t)))
	//fmt.Println(Int64ToTime(TimeToUnixMs(t)))
	//fmt.Println(Int64ToTime(TimeToUnixNano(t)))
	//fmt.Printf("时间（纳秒）：%v;\n", Int64ToTime(TimeToUnixNano(t)))
	fmt.Printf("时间（毫秒）：%v;\n", Int64ToTime(TimeToUnixMs(t)))
	fmt.Printf("时间（纳秒）：%v;\n", Int64ToTime(TimeToUnixNano(t)))

	fmt.Println(t)

	// //获取当前时间戳
	// fmt.Println(t.Unix()) //1531293019

	// //时间 to 时间戳
	// loc, _ := time.LoadLocation("Asia/Shanghai")                                     //设置时区
	// tt, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-07-11 15:07:51", loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	// fmt.Println(tt.Unix())                                                           //1531292871

	// //时间戳 to 时间
	// tm := time.Unix(1546985052, 0)
	// fmt.Println(tm.Format("2006-01-02 15:04:05")) //2018-07-11 15:10:19

	// fmt.Println(TimeToInt64(t))
	// fmt.Println(Int64ToTime(TimeToInt64(t)))
	// fmt.Println(TimeToInt64Location(t, "America/Antigua"))
	// fmt.Println(Int64ToTime(TimeToInt64Location(t, "America/Antigua")))
	// fmt.Println(Int64ToTimeLocation(TimeToInt64Location(t, "America/Antigua"), "America/Antigua"))
}
func Unix(nsec int64) int64 {
	if nsec < 0 || nsec >= 1e9 {
		n := nsec / 1e9
		nsec -= n * 1e9
		if nsec < 0 {
			nsec += 1e9
		}
	}
	return nsec
}
func TimeToInt64(t time.Time) int64 {
	return t.Unix()
}

// TimeToUnixMs transform time to Unix time, the number of ms elapsed
func TimeToUnixMs(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// TimeToUnixNano transform time to Unix time, the number of nanosecond elapsed
func TimeToUnixNano(t time.Time) int64 {
	return t.UnixNano()
}
func Int64ToTimeMs1(ss int64) time.Time {
	nsec := int64(0)
	sec := int64(0)
	constNum := int64(1e3)
	// if ss > 0 || ss <= constNum {
	// 	sec = ss / constNum
	// 	nsec = Unix00(sec, ss)
	// }
	nsec = ss / constNum
	if nsec < 0 || nsec >= constNum {
		n := nsec / constNum
		sec += n
		nsec -= n * constNum
		if nsec < 0 {
			nsec += constNum
			sec--
		}
	}
	fmt.Printf("ss: %v,sec: %v,nsec: %v", ss, sec, nsec)
	fmt.Println()
	return time.Unix(sec, nsec)
}
func Int64ToTimeNsec(ss int64) time.Time {
	nsec := int64(0)
	sec := int64(0)
	constNum := int64(1e9)
	if ss > 0 || ss <= constNum {
		sec = ss / constNum
		nsec = Unix11(sec, ss)
	}
	fmt.Printf("ss: %v,sec: %v,nsec: %v", ss, sec, nsec)
	fmt.Println()
	return time.Unix(sec, nsec)
}
func Unix00(sec int64, nsec int64) int64 {
	if nsec < 0 || nsec >= 1e3 {
		n := nsec / 1e3
		sec += n
		nsec -= n * 1e3
		if nsec < 0 {
			nsec += 1e3
			sec--
		}
	}
	return nsec
}
func Unix11(sec int64, nsec int64) int64 {
	if nsec < 0 || nsec >= 1e9 {
		n := nsec / 1e9
		sec += n
		nsec -= n * 1e9
		if nsec < 0 {
			nsec += 1e9
			sec--
		}
	}
	return nsec
}

func Int64ToTime(ss int64) time.Time {
	//10位时间戳
	if ss < 1e10 {
		return time.Unix(ss, 0)
	}
	//13位时间戳
	if ss >= 1e12 && ss < 1e13 {
		ss = ss * 1e6
		return time.Unix(0, ss)
	} else { //19位时间戳
		return time.Unix(0, ss)
	}
}
func Int64ToTimeMs(tt int64) time.Time {
	return time.Unix(tt, 13)
}
func Int64ToTimeNano(tt int64) time.Time {
	return time.Unix(tt, 19)
}

//"Asia/Shanghai"
func TimeToInt64Location(t time.Time, locat string) int64 {
	timeStr := t.Format("2006-01-02 15:04:05")
	loc, _ := time.LoadLocation(locat)                                 //设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
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
