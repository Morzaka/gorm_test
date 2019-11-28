package repositories

import (
	"github.com/Morzaka/01_interview/gorm_test/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) SaveIdentity(identity *models.Identity) error {
	if err := r.db.Save(identity).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) SaveAdmin(admin *models.Admin) error {
	if err := r.db.Save(admin).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) SaveHobby(hobby *models.Hobby) error {
	if err := r.db.Save(hobby).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) SaveFriends(friend *models.Friend) error {
	if err := r.db.Save(friend).Error; err != nil {
		return err
	}
	return nil
}
