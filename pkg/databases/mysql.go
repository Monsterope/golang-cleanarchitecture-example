package databases

import (
	"cleanarchitecture-example/configs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (con *DatabaseConfig) MysqlConnect() {
	db_host := configs.GetEnv("db.host")
	db_port := configs.GetEnv("db.port")
	db_name := configs.GetEnv("db.database")
	db_user := configs.GetEnv("db.username")
	db_pass := configs.GetEnv("db.password")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db_user, db_pass, db_host, db_port, db_name)

	var err error
	con.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("cannot connect database please")
	}
}
