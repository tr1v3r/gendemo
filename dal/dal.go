package dal

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gendemo/config"
)

var DB *gorm.DB
var once sync.Once

func init() {
	once.Do(func() {
		DB = ConnectDB()
		// DB = DB.Debug()
	})
}

func ConnectDB() (conn *gorm.DB) {
	var err error
	if os.Getenv("DB_CONNECT_MODE") == "DSN" {
		conn, err = gorm.Open(mysql.Open(config.MySQLDSN))
	}
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	return conn
}
