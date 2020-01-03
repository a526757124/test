package service

import (
	"errors"
	"fmt"

	"github.com/a526757124/test/dottask/conn"
	"github.com/a526757124/test/dottask/models"
)

type TaskService struct{}

func (taskService *TaskService) Insert(model *models.Task) error {
	n, err := conn.GetMysqlClient().Insert("INSERT INTO `Tasks`(`TaskID`,`IsRun`,`TaskType`,`DueTime`,`Interval`,`Express`,`QueueSize`,`HandlerName`,`HandlerData`)	VALUES(?,?,?,?,?,?,?,?,?);",
		model.TaskID, model.IsRun, model.TaskType, model.DueTime, model.Interval, model.Express, model.QueueSize, model.HandlerName, model.HandlerData)
	if err != nil {
		return err
	}
	fmt.Println(n)
	if n == 0 {
		return errors.New("新增失败！")
	}
	return nil
}

// func (taskService *TaskService) GetAllList(queryParm *viewModel.TaskQueryParm) (*models.Task, error) {
// 	results := []*models.Task{}
// 	sql := "SELECT TaskID,`IsRun`,`TaskType`,`DueTime`,`Interval`,`Express`,`QueueSize`,`HandlerName`,`HandlerData` FROM `Tasks`"
// 	whereSql := " where 1=1 "
// 	err := conn.GetMysqlClient().FindList(&results, sql)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &results, nil
// }
