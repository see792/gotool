package database

import (
	"github.com/see792/gotool/config"
	"github.com/see792/gotool/database/mongodb"
	"github.com/see792/gotool/database/mysql"
	"github.com/see792/gotool/database/redis"
)

type DataBase struct {
	MysqlDB *mysql.MysqlDB
	MongoDB *mongodb.MongoDB
	RedisDB *redis.RedisDB
}

func New(config *config.Config) *DataBase {
	newDateBase := new(DataBase)
	newDateBase.MysqlDB = mysql.New(&config.MySql)
	newDateBase.RedisDB = redis.New(&config.Redis)
	newDateBase.MongoDB = mongodb.New(&config.MongoDB)
	return newDateBase
}
