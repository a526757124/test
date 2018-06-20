package main

import (
	"fmt"
	"strconv"

	"github.com/devfeel/dotweb"
)

func main() {
	app := dotweb.New()
	InitRoute(app.HttpServer)
	err := app.StartServer(8888)
	if err != nil {
		fmt.Print(err.Error())
	}
}
func InitRoute(server *dotweb.HttpServer) {
	server.Router().ServerFile("/file/*filepath", "")
	server.POST("/upload", Upload)
	server.POST("/home", Home)

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
