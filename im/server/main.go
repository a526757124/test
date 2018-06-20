package main

import (
	"fmt"

	"github.com/devfeel/dotweb"
)

func main() {
	app := dotweb.New()
	InitRoute(app.HttpServer)
	err := app.StartServer(9999)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func InitRoute(server *dotweb.HttpServer) {
	server.GET("/", Index)
	server.WebSocket("/ws/onsocket", OnWebSocket)
}
func Index(ctx dotweb.Context) error {
	ctx.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	flag := ctx.HttpServer().Router().MatchPath(ctx, "/d/:x/y")

	return ctx.WriteString("index - " + ctx.Request().Method + " - " + fmt.Sprint(flag))
}
func OnWebSocket(ctx dotweb.Context) error {
	fmt.Println(ctx.Request().URL)
	// appId := ctx.QueryString("appid")
	// groupId := ctx.QueryString("groupid")
	// groupIds := ctx.QueryString("groupids")
	// userId := ctx.QueryString("userid")
	// from := ctx.QueryString("from")
	// token := ctx.QueryString("token")
	ctx.WebSocket().SendMessage("1")

	for {

	}

	return nil
}
