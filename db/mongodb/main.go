package main

import (
	"github.com/devfeel/database/mongo"
)

const (
	testMongoConn       = "mongodb://118.31.32.168:27017/"
	testMongoDbName     = "test"
	testObjectID        = "5a068700d3f05e12e844df86"
	collectionName_Demo = "Demo" //Demo表名
)

var db = mongo.NewMongoDBContext(testMongoConn, testMongoDbName)

func init() {
	db.DefaultCollectionName = collectionName_Demo
}
func main() {
	db.FindList()
}
