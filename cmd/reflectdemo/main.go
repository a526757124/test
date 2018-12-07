package main

import (
	"fmt"
	"reflect"

	"github.com/devfeel/mapper/reflectx"
)

type Test struct {
	IsDel bool
	Name  string
}

func main() {
	test := new(Test)
	structScan(test)
}
func structScan(dest interface{}) {
	destValue := reflect.ValueOf(dest)
	t := reflect.TypeOf(dest) //反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
	fmt.Println("Type:", t)   //调用t.Name方法来获取这个类型的名称

	fmt.Println("destValue:", destValue) //调用t.Name方法来获取这个类型的名称
	fmt.Println(destValue.Field(0).Interface())

	for i := 0; i < t.NumField(); i++ { //通过索引来取得它的所有字段，这里通过t.NumField来获取它多拥有的字段数量，同时来决定循环的次数
		f := t.Field(i)                       //通过这个i作为它的索引，从0开始来取得它的字段
		val := destValue.Field(i).Interface() //通过interface方法来取出这个字段所对应的值
		fmt.Printf("%6s:%v =%v\n", f.Name, f.Type, val)
	}
	var isSlice bool
	_, errCheckSlice := reflectx.BaseType(destValue.Type(), reflect.Slice)
	if errCheckSlice != nil {
		isSlice = false
	} else {
		isSlice = true
	}
	if isSlice {
		fmt.Println("1")
	} else {
		fmt.Println("2")
		elemType := destValue.Type()
		fmt.Println(elemType.Kind())
		if elemType.Kind() != reflect.Ptr {
			fmt.Println("slice elem must ptr ")
		}

	}
}
