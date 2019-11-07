package main

import (
	"database/sql"
	"fmt"

	"github.com/a526757124/database/postgressql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func sqlOpen() {
	var err error
	//db, err = sql.Open("postgres", "port=5433 user=postgres password=123456 dbname=mytest sslmode=disable")
	//port是数据库的端口号，默认是5432，如果改了，这里一定要自定义；
	//user就是你数据库的登录帐号;
	//dbname就是你在数据库里面建立的数据库的名字;
	//sslmode就是安全验证模式;

	//还可以是这种方式打开
	db, err := sql.Open("postgres", "postgres://mytest:123456@118.31.32.168/mytest?sslmode=verify-full")
	if db != nil {
		fmt.Println("open....")
	}
	checkErr(err)
}
func sqlInsert() {
	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec("ficow", "软件开发部门", "2017-03-09")
	//这里的三个参数就是对应上面的$1,$2,$3了

	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
func sqlDelete() {
	//删除数据
	stmt, err := db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err := stmt.Exec(1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
func sqlSelect() {
	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	println("-----------")
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid = ", uid, "\nname = ", username, "\ndep = ", department, "\ncreated = ", created, "\n-----------")
	}
}
func sqlUpdate() {
	//更新数据
	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err := stmt.Exec("ficow", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
func sqlClose() {
	if db != nil {
		fmt.Println("close....")
		db.Close()
	}

}
func checkErr(err error) {
	if err != nil {

		panic(err)
	}
}

func sqlTest() {

	sep := "----------\n"
	sqlOpen()
	println(sep, "*sqlOpen")

	// sqlSelect()
	// println(sep, "*sqlSelect")

	// sqlInsert()
	// sqlSelect()
	// println(sep, "*sqlInsert")

	// sqlUpdate()
	// sqlSelect()
	// println(sep, "*sqlUpdate")

	// sqlDelete()
	// sqlSelect()
	// println(sep, "*sqlDelete")

	sqlClose()
	println(sep, "*sqlClose")
}

const (
	host     = "118.31.32.168"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "mytest"
)

type UserInfo struct {
	ID   int    `mapper:"id"`
	Name string `mapper:"name"`
}

func UsePostgresSqlDBContext() {
    db := postgressql.NewPostgresSqlDBContext("host=118.31.32.168 port=5432 user=postgres password=123456 dbname=mytest sslmode=disable")
    
    // resultProc, err := db.ExecProc("add", 3, 2)
    // if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(resultProc)
	// }
	// n, err := db.Insert("INSERT INTO userinfo(id,name) VALUES($1,$2)", 5, "李四")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(n)
    // }
    
    sum, err := db.Scalar("select id,name from userinfo")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
    }
    
    max, err := db.Scalar("select max(id) from userinfo")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(max)
	}
	// for i := 0; i < 10; i++ {
	// 	n, err := db.Insert("INSERT INTO userinfo(id,name) VALUES($1,$2)", i, i)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(n)
	// 	}
	// }

	result := new(UserInfo)
	err = db.FindOne(result, "SELECT * FROM userinfo where id=$1 limit 1", 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	count, err := db.Scalar("SELECT * FROM userinfo")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(count)
	}

	// n, err = db.Update("update userinfo set name=$1 where id=$2", "王五", 1)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(n)
	// }
	// err := db.FindOne(result, "SELECT * FROM userinfo where id=$1 limit 1", 1)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	// n, err = db.Update("delete from userinfo where id=$1", 1)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(n)
	// }

	// err = db.FindOne(result, "SELECT * FROM userinfo where id=$1 limit 1", 1)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	var results []*UserInfo
	err = db.FindList(&results, "SELECT * FROM userinfo")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(results)
	}

	// resultMap, err := db.FindListMap("SELECT * FROM userinfo")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(resultMap)
	// }

	err = db.FindListByPage(&results, "userinfo", "*", "id>$1", "id ASC", 10, 0,0)
	if err != nil {
		fmt.Println(err)
	} else {
        fmt.Println(results)
        for i,v:= range results{
            fmt.Println(i)
            fmt.Println(v)
        }
	}
}
func main() {
	UsePostgresSqlDBContext()
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// db, err := sql.Open("postgres", "host=118.31.32.168 port=5432 user=postgres password=123456 dbname=mytest sslmode=disable")
	// if db != nil {
	// 	fmt.Println("open....")
	// }
	// defer db.Close()
	// err = db.Ping()
	// checkErr(err)

	// rows, err := db.Query("SELECT * FROM userinfo")
	// checkErr(err)
	// println("-----------")
	// for rows.Next() {
	// 	var id int
	// 	var username string
	// 	err = rows.Scan(&id, &username)
	// 	checkErr(err)
	// 	fmt.Println("id = ", id, "\nname = ", username, "\ndep = ", "\n-----------")
	// }
	// checkErr(err)
	// db.Close()
}
