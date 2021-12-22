package connection

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db        *sql.DB
	goquDB    *goqu.Database
	mutexDB   sync.Mutex
	mutexGoqu sync.Mutex
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

func GoquDB(db *sql.DB) (*goqu.Database, error) {
	if goquDB == nil {
		mutexGoqu.Lock()
		defer mutexGoqu.Unlock()
		if goquDB == nil {
			dialect := goqu.Dialect("mysql")
			goquDB = dialect.DB(db)
		}
	}
	return goquDB, nil
}
