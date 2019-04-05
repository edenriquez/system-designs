package db

import (
	"log"
	"os"
	"strconv"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

// SetUpDB should instance db engine and migrate tables
func SetUpDB() {
	// get engine
	orm, err := xorm.NewEngine("postgres", connectionString())
	if err != nil {
		log.Println("database error: ", err.Error())
	}
	// debug
	show, _ := strconv.ParseBool(os.Getenv("SHOWTESTSQL"))
	orm.ShowSQL(show)

	tableMigrationProcess(orm)
}

func connectionString() string {
	return "dbname=" + os.Getenv("DB_NAME") + " user=postgres sslmode=disable"
}
