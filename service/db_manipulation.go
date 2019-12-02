package service

import (
	"fmt"
	"github.com/Morzaka/01_interview/gorm_test/database"
	"github.com/Morzaka/01_interview/gorm_test/db_query"
	"github.com/Morzaka/01_interview/gorm_test/models"
	"github.com/Morzaka/01_interview/gorm_test/repositories"
	"log"
)

const (
	dbType   = "postgres"
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	dbName   = "users"
	password = "postgres"
)

func InitializeData() {

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbName, password)
	db, err := database.ConnectToDB(dbType, dbInfo)
	defer db.Close()
	if err != nil {
		log.Fatalln("Can't connect to db, ERROR:", err)
	}

	if !db.HasTable("users") {
		t := db_query.Newtables(db)
		if err := t.CreateSequence(); err != nil {
			log.Fatalln(err)
		}
		if err := t.CreateUserTable(); err != nil {
			log.Fatalln("CreateUserTable", err)
		}
		if err := t.CreateIdentityTable(); err != nil {
			log.Fatalln("CreateIdentity", err)
		}
		if err := t.CreateAdminTable(); err != nil {
			log.Fatalln("CreateAdmin", err)
		}
		if err := t.CreateFriendsTable(); err != nil {
			log.Fatalln("CreateFriends", err)
		}
		if err := t.CreateHobbiesTable(); err != nil {
			log.Fatalln("CreateHobbies", err)
		}
	}

	repo := repositories.NewUserRepository(db)

	newUser := models.User{
		FName: "Jhon",
		LName: "McCloud",
		Bio:   "Some guy from West Virginia",
	}

	err = repo.SaveUser(&newUser)
	if err != nil {
		log.Fatalln("Can't save user ERROR", err)
	}
}
