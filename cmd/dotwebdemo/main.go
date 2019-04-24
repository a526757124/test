package main

import (
	"fmt"
	"strconv"
	"time"
	"github.com/devfeel/dotweb"
)

func main() {
	app := dotweb.New()
	app.SetProductionMode()
	app.HttpServer.SetEnabledAutoOPTIONS(false)
	InitRoute(app.HttpServer)
	fmt.Println("启动8888...");
	err := app.StartServer(8888)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func InitRoute(server *dotweb.HttpServer) {
	//server.SetEnabledGzip(true)
	//server.Router().ServerFile("/file/*filepath", "")
	//server.POST("/upload", Upload)
	server.GET("/home", Home)
	//server.Any("/application", GetApplicationList)
	server.GET("/a/:arg1", func(ctx dotweb.Context) error {
		fmt.Println(ctx.GetRouterName("arg1"))
		fmt.Println(1)
		return nil
	})
	server.GET("/a/:arg1/b/:arg2", func(ctx dotweb.Context) error {
		fmt.Println(2)
		fmt.Println(ctx.Request().Url())
		fmt.Println(server.Router().MatchPath(ctx, "/a/:arg1/b/:arg2"))
		fmt.Println(ctx.GetRouterName("arg1"))
		fmt.Println(ctx.GetRouterName("arg2"))
		fmt.Println(3)
		return nil
	})

}

// func NewCustomCROS() dotweb.Middleware {
// 	option := cors.NewConfig().UseDefault()
// 	return cors.Middleware(option)
// }

type Application struct {
	Name       string    `json:"name"`
	Status     int       `json:"status"`
	Desc       string    `json:"desc"`
	CreateDate time.Time `json:"createDate"`
}
type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func GetApplicationList(ctx dotweb.Context) error {
	fmt.Println("-------------------")
	applicationList := []*Application{}
	for i := 0; i < 1000; i++ {
		applicationList = append(applicationList, &Application{
			Name:       "应用" + strconv.Itoa(i),
			Status:     (i % 2),
			Desc:       "应用描述" + strconv.Itoa(i),
			CreateDate: time.Now(),
		})
	}
	result := Result{
		Code: 200,
		Data: &applicationList,
	}
	return ctx.WriteJson(result)
}

func Home(ctx dotweb.Context) error {
	return ctx.WriteString("123")
}
func Upload(ctx dotweb.Context) error {
	fmt.Println(ctx.File("updload"))
	files, err := ctx.Request().FormFiles()
	if err != nil {
		fmt.Println(err.Error())
	}
	for k, v := range files {
		fmt.Println(k, v)
		var filePath = v.FileName()
		fmt.Println(filePath + ",size:" + strconv.FormatInt(v.Size(), 10))

		v.SaveFile(filePath)

	}
	fmt.Println(files)
	return ctx.WriteString("ok")
}
