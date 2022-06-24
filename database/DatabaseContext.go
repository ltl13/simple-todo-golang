package contexts

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/go-sql-driver/mysql"
)

var dbInstance *sql.DB

var lock = &sync.Mutex{}

func GetDBInstance() *sql.DB {
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbInstance == nil {
			cfg := mysql.Config{
				User: "root",
				// Passwd: "130202001",
				Net:    "tcp",
				Addr:   "127.0.0.1:3306",
				DBName: "todo_manager",
				AllowNativePasswords: true,
			}

			var err error
			dbInstance, err = sql.Open("mysql", cfg.FormatDSN())
			if err != nil {
				log.Fatal(err)
				return nil
			}
			if pingErr := dbInstance.Ping(); pingErr != nil {
				log.Fatal(pingErr)
				return nil
			}
			fmt.Print("Connected!")
		}
	}
	return dbInstance
}
