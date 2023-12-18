package main

import (
	"go-app/pkg/router"
	"os"

	"github.com/BoyYangZai/go-server-lib/pkg/config_reader"
	"github.com/BoyYangZai/go-server-lib/pkg/database"
)

func main() {
	// Init config
	absolute_path, _ := os.Getwd()
	config_reader.InitConfig(absolute_path)
	print(config_reader.GetConfigByKey("name"))

	// Init database
	database.InitDatabase(database.DatabaseConfig{
		Host:     config_reader.GetConfigByKey("database.host"),
		Port:     config_reader.GetConfigByKey("database.port"),
		User:     config_reader.GetConfigByKey("database.username"),
		Password: config_reader.GetConfigByKey("database.password"),
		DBName:   config_reader.GetConfigByKey("database.database"),
	})

	// 业务代码

	router.CreateRouter()

	print("\napp end")

}
