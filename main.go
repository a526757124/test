package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	t := time.Now() //2018-07-11 15:07:51.8858085 +0800 CST m=+0.004000001
	fmt.Println(t)

	//获取当前时间戳
	fmt.Println(t.Unix()) //1531293019

	//时间 to 时间戳
	loc, _ := time.LoadLocation("Asia/Shanghai")                                     //设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-07-11 15:07:51", loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	fmt.Println(tt.Unix())                                                           //1531292871

	//时间戳 to 时间
	tm := time.Unix(1531293019, 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05")) //2018-07-11 15:10:19

	//获取当前年月日,时分秒
	y := t.Year()                 //年
	m := t.Month()                //月
	d := t.Day()                  //日
	h := t.Hour()                 //小时
	i := t.Minute()               //分钟
	s := t.Second()               //秒
	fmt.Println(y, m, d, h, i, s) //2018 July 11 15 24 59

	// for i := 0; i < 1000; i++ {
	// 	httpHandler()
	// }

	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3(0))
	fmt.Println(test4(0))

}

func test1() int {
	i := 0
	defer func() {
		i = i + 1
	}()
	return i
}

func test2() (i int) {
	i = 0
	defer func() {
		i = i + 1
	}()
	return i
}

func test3(i int) int {
	defer func() {
		i = i + 1
	}()
	return i
}
func test4(i int) (ii int) {
	ii = i
	defer func() {
		ii = ii + 1
	}()
	return 10
}

func httpHandler() {
	errCh := make(chan error, 1)
	resultCh := make(chan error, 1)

	go func() {
		defer close(errCh)
		defer close(resultCh)
		errCh <- fmt.Errorf("shit")
	}()

	select {
	case <-errCh:
	case <-resultCh:
		fmt.Println("this shall not happen dude!")
	}
}
