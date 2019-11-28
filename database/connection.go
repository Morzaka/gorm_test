package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func ConnectToDB(DBType, connectArg string) (*gorm.DB, error) {
	db, err := gorm.Open(DBType, connectArg)

	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		log.Println("ping ERROR")
		return nil, err
	}

	log.Println("ping Success")
	return db, err
}
