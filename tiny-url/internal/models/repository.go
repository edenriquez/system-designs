package db

import (
	"reflect"
	"time"

	"github.com/go-xorm/xorm"
)

// Urls will store short url's
type Urls struct {
	ShortURLID  string
	OriginalURL string
	Created     time.Time `xorm:"created"`
	Updated     time.Time `xorm:"updated"`
	userID      *Users    `xorm:"user_id bigint" `
}

// Users will store account information
type Users struct {
	ID       int64
	Username string
	Token    string
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

// migrationGroup is the generi struct to map all the models
func migrationGroup() []interface{} {
	var tables []interface{}
	tables = append(tables, &Urls{})
	tables = append(tables, &Users{})
	return tables
}

// tableMigrationProcess iterats models and creates them
func tableMigrationProcess(newdriver *xorm.Engine) {
	for _, table := range migrationGroup() {
		err := newdriver.Sync2(table)
		if err != nil {
			tname := reflect.TypeOf(table).String()
			println("There is an error while creating: ", tname)
		}
	}

}
