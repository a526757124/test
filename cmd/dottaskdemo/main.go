package main

import (
	"fmt"
	"runtime"
	"time"

	. "github.com/devfeel/dottask"
)

var service *TaskService

func Job_Test(ctx *TaskContext) error {
	fmt.Println(time.Now().String(), " => Job_Test")
	//time.Sleep(time.Second * 3)
	return nil
}
func Job_SysInfo(ctx *TaskContext) error {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(time.Now().String(), " => Job_SysInfo")
	return nil
}

func Loop_Test(ctx *TaskContext) error {
	fmt.Println(time.Now().String(), " => Loop_Test")
	time.Sleep(time.Second * 3)
	return nil
}

func main() {

	//step 1: init new task service
	service = StartNewService()

	//step 2: register task handler
	_, err := service.CreateCronTask("testcron", true, "*/1 * * * * *", Job_SysInfo, nil)
	if err != nil {
		fmt.Println("service.CreateCronTask error! => ", err.Error())
	}

	//step 3: start all task
	service.StartAllTask()

	fmt.Println(service.PrintAllTasks())

	for {
		time.Sleep(time.Hour)
	}

}
