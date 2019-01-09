package mongodb

import (
	"github.com/devfeel/database/mongo"
)

type MongoDatabsse struct {
}

const (
	testMongoConn       = "mongodb://118.31.32.168:27017/"
	testMongoDbName     = "dbtest"
	testObjectID        = "5a068700d3f05e12e844df86"
	collectionName_Demo = "User" //Demo表名
)

var db = mongo.NewMongoDBContext(testMongoConn, testMongoDbName)

func init() {
	db.DefaultCollectionName = collectionName_Demo
}

//insert
func (*MongoDatabsse) Insert(args ...interface{}) (n int64, err error) {
	db.
	return 0, nil
}
