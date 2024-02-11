package config

import (
	"os"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var DBmongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "bukupedia",
}

var Mongoconn = atdb.MongoConnect(DBmongoinfo)
