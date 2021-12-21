package connection

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db        *sql.DB
	gormDB    *gorm.DB
	mutexDB   sync.Mutex
	mutexGorm sync.Mutex
)

//DataBaseConnection fields for database connection
type DataBaseConnection struct {
	User string
	Pass string
	Host string
	Port int
	DB   string
}

//StringConnection
func StringConnection(dbFields DataBaseConnection) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		dbFields.User,
		dbFields.Pass,
		dbFields.Host,
		dbFields.Port,
		dbFields.DB,
	)
}

//DBConnection
func DBConnection(dbFields DataBaseConnection) (*sql.DB, error) {
	var err error
	if db == nil {
		mutexDB.Lock()
		defer mutexDB.Unlock()
		//TODO: freddy use config for connection variables Fri 17 Dec 2021 11:44:27 PM -05
		if db == nil {
			fmt.Printf("strin conn: %s \n", StringConnection(dbFields))
			db, err = sql.Open("mysql", StringConnection(dbFields))
		}
	}
	return db, err
}

func GormDB(db *sql.DB) (*gorm.DB, error) {
	var err error
	if gormDB == nil {
		mutexGorm.Lock()
		defer mutexGorm.Unlock()
		if gormDB == nil {
			gormDB, err = gorm.Open(mysql.New(mysql.Config{
				Conn: db,
			}), &gorm.Config{})
		}
	}
	return gormDB, err
}
