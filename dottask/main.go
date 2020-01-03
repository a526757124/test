package main

import (
	"fmt"

	"github.com/a526757124/test/dottask/service"

	"github.com/a526757124/test/dottask/models"
	"github.com/devfeel/database/mysql"
)

var taskService *service.TaskService

func init() {
	taskService = &service.TaskService{}

}
func main() {
	mysqlClient := mysql.NewMySqlDBContext("root:123456@tcp(118.31.32.168:3306)/test?charset=utf8&allowOldPasswords=1")
	if mysqlClient == nil {
		panic("mysql connection failed!")
	}

	model := models.Task{}
	model.TaskID = "test1"
	err := taskService.Insert(&model)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("OK")
}
