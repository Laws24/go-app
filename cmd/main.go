package main

import (
	"go-app/pkg/router"

	"github.com/BoyYangZai/go-server-lib/pkg/config_reader"
	"github.com/BoyYangZai/go-server-lib/pkg/database"
)

func main() {
	// 初始化 config
	config_reader.InitConfig("C:\\Users\\RuFeng\\Desktop\\go-app")

	// 初始化 database
	database.InitDatabase(database.DatabaseConfig{User: config_reader.GetConfigByKey("database.User"), Password: config_reader.GetConfigByKey("database.Password"),
		Host: config_reader.GetConfigByKey("database.Host"), Port: config_reader.GetConfigByKey("database.Port"), DBName: config_reader.GetConfigByKey("database.DBName")})

	router.CreateRouter()
}
