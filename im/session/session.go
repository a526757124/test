package session

import (
	"sync"

	"github.com/devfeel/dotweb/framework/crypto/uuid"
	"golang.org/x/net/websocket"
)

var sessionMap sync.Map

func init() {

}

type (
	SessionManager interface {
		Add(session *Session) string
		Update(uuid string, session *Session)
		Remove(sid string)
	}
	Session struct {
		Conn *websocket.Conn
		//session id
		Id              string
		Sourece         int
		DeviceId        string
		Host            string
		UserId          string
		Platform        string
		PlatformVersion string
		appkey          string
		bindTime        int64
		updateTime      int64
		sign            string
		status          int
		//用于dwr websocket存储多开页面创建的session
		//sessions map[string,*Session]
	}
)

func Add(session *Session) string {
	uuid := uuid.NewV4().String()
	sessionMap.Store(uuid, session)
	return uuid
}
func Update(uuid string, session *Session) {
	sessionMap.Store(uuid, session)
}
func Remove(sid string) {
	sessionMap.Delete(sid)
}
