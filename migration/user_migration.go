package migration

import (
	"log"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"

	"github.com/Morzaka/01_interview/gorm_test/models"
)

func UserMigrate(db *gorm.DB) {
	db.LogMode(true)

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "300000000004",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("users").Error
			},
		},
		{
			ID: "300000000005",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Identity{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table("identities").Error
			},
		},
		{
			ID: "300000000006",
			Migrate: func(tx *gorm.DB) error {
				if err := db.Create(&models.Admin{
					Role: models.AllPermission,
				}).Error; err != nil {
					return err
				}
				if err := db.Create(&models.Admin{
					Role: models.CanBLockUsers,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("admins").Error
			},
		},
		{
			ID: "300000000007",
			Migrate: func(tx *gorm.DB) error {
				if err := db.Create(&models.Hobby{
					HobbyType: models.Sport,
				}).Error; err != nil {
					return err
				}
				if err := db.Create(&models.Hobby{
					HobbyType: models.Embroidery,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("hobbies").Error
			},
		},
		{
			ID: "300000000008",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Friends{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("friends").Error
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v\n", err)
	}
	log.Printf("Migratiion did run successfully")

	m.InitSchema(func(tx *gorm.DB) error {
		if err := tx.Model(models.User{}).AddForeignKey("hobbies", "hobbies (id)", "RESTRICT", "RESTRINCT").Error; err != nil {
			return err
		}
		if err := tx.Model(models.User{}).AddForeignKey("identity", "hobbies (id)", "RESTRICT", "RESTRINCT").Error; err != nil {
			return err
		}
		return nil
	})
}
