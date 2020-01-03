package conn

import (
	"github.com/devfeel/database"
	"github.com/devfeel/database/mysql"
)

//var mysqlClient *mysql.MySqlDBContext

type MysqlConn struct{}

func init() {
	//	mysqlClient = mysql.NewMySqlDBContext("root:123456@tcp(118.31.32.168:3306)/test?charset=utf8&allowOldPasswords=1")
}

//get mysqlClient conn
func GetMysqlClient() database.DBContext {
	mysqlClient := mysql.NewMySqlDBContext("root:123456@tcp(118.31.32.168:3306)/test?charset=utf8&allowOldPasswords=1")
	if mysqlClient == nil {
		panic("mysql connection failed!")
	}
	return mysqlClient
}
