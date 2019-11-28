package service

import (
	"fmt"
	"github.com/Morzaka/01_interview/gorm_test/database"
	"github.com/Morzaka/01_interview/gorm_test/models"
	"github.com/Morzaka/01_interview/gorm_test/repositories"
	"github.com/google/uuid"
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
	if err != nil {
		log.Fatalln("Can't connect to db, ERROR:", err)
	}

	//db.AutoMigrate(&models.User{}, &models.Identity{}, &models.Friend{})
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	newIdentity := models.Identity{
		ID:       uuid.New(),
		Login:    "admin",
		Password: "admin",
		Salt:     "kd65ngf6jo9d4tso",
		Blocked:  false,
	}

	newIdentity.User = models.User{
		ID:         1,
		IdentityID: newIdentity.ID,
		FName:      "Jorj",
		LName:      "Kluni",
		Bio:        "actor and another things",
	}

	newIdentity.Admin = models.Admin{
		IdentityID: newIdentity.ID,
		Role:       models.CanBLockUsers,
	}

	err = repo.SaveIdentity(&newIdentity)
	if err != nil {
		log.Fatalln("Can't save user ERROR", err)
	}
}
