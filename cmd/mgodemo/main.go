package main

import (
	"fmt"

	"gitee.com/a526757124/database/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	testMongoConn       = "mongodb://118.31.32.168:27017/"
	testMongoDbName     = "dbtest"
	testObjectID        = "5a068700d3f05e12e844df86"
	collectionName_Demo = "User" //Demo表名
)

type User struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
}

func main() {
	db := mongo.NewMongoDBContext(testMongoConn, testMongoDbName)
	db.DefaultCollectionName = collectionName_Demo
	user := &User{
		Id:   mongo.NewObjectId(),
		Name: "test1",
	}
	err := db.InsertBlob(user)
	if err != nil {
		fmt.Println(err.Error())
	}

	// session, err := mgo.Dial("118.31.32.168:27017")
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()
	// db := session.DB("dbtest")
	// c := db.C("user")
	// c.RemoveAll(nil)
	// for i := 0; i < 100; i++ {
	// 	b := strings.Builder{}
	// 	b.WriteString("name")
	// 	b.WriteString(strconv.Itoa(i))
	// 	name1 := b.String()
	// 	err = c.Insert(&User{
	// 		Id_:  bson.NewObjectId(),
	// 		Name: name1,
	// 	})
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
}
