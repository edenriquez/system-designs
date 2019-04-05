package db

import (
	"fmt"
	"time"

	"github.com/go-xorm/xorm"
)

type Urls struct {
	ID          int64
	ShortURL    string
	OriginalURL string
	Created     time.Time `xorm:"created"`
	Updated     time.Time `xorm:"updated"`
}

type Users struct {
	ID       int64
	Username string
	Token    string
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

func tableMigrationProcess(newdriver *xorm.Engine) {
	err := newdriver.Sync2(new(Urls))
	fmt.Println(err)
	err = newdriver.Sync2(new(Users))
	fmt.Println(err)
}
