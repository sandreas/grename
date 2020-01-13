package database

import (
	"errors"
	"github.com/jinzhu/gorm"
	"os"
)

//import _ "github.com/jinzhu/gorm/dialects/mysql"
// import _ "github.com/jinzhu/gorm/dialects/postgres"
import _ "github.com/jinzhu/gorm/dialects/sqlite"

// import _ "github.com/jinzhu/gorm/dialects/mssql"

//type DatabaseConnection struct {
//	Db *gorm.DB
//
//}

type Credentials struct {
	Host     string
	Port     uint
	Username string
	Password string
	Driver   string
	Database string
}

func InitDatabase(u *Credentials) (*gorm.DB, error) {
	if u.Driver == "sqlite3" {
		err := touchFile(u.Database)
		if err != nil {
			return nil, err
		}
		db, err := gorm.Open(u.Driver, u.Database)
		if err != nil {
			return nil, err
		}
		db.AutoMigrate(&File{}, &Tag{}, &FileTag{})
		return db, nil
	}
	return nil, errors.New("")
}

func touchFile(name string) error {
	file, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	return file.Close()
}
