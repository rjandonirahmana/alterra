package database

import (
	"errors"
	"mvcApi/middleware"
	model "mvcApi/models"

	"gorm.io/gorm"
)

type Db struct {
	db      *gorm.DB
	service middleware.ServiceIn
}

type Repository interface {
	GetUsers() (interface{}, error)
	CreateUser(user model.Users) (model.Users, error)
	UpdateUser(id int, Newuser *model.Users) (interface{}, error)
	DeleteUser(id int, user model.Users) error
	GetUser(id int) (model.Users, error)
	LoginUser(user *model.Users) error
}

func NewRepo(db *gorm.DB, service middleware.ServiceIn) *Db {
	return &Db{db: db, service: service}
}

func (d *Db) GetUsers() (interface{}, error) {
	var users []model.Users

	err := d.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (d *Db) CreateUser(user model.Users) (model.Users, error) {

	if err := d.db.Save(&user).Error; err != nil {
		return model.Users{}, err
	}

	token, err := d.service.GenerateToken(int(user.ID))

	if err != nil {
		return model.Users{}, err
	}

	user.Token = token

	return user, nil

}

func (d *Db) UpdateUser(id int, Newuser *model.Users) (interface{}, error) {

	user := model.Users{}
	tx := d.db.Find(&user, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected > 0 {
		tx = d.db.Model(&user).Updates(Newuser)

		if tx.Error != nil {
			return nil, tx.Error
		}

		return user, nil
	}

	return nil, tx.Error
}

func (d *Db) DeleteUser(id int, user model.Users) error {

	err := d.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}

func (d *Db) GetUser(id int) (model.Users, error) {
	var user model.Users

	if err := d.db.Find(&user, id).Error; err != nil {
		return model.Users{}, errors.New("cant find user with id")
	}

	return user, nil
}

func (d *Db) LoginUser(user *model.Users) error {

	err := d.db.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error
	if err != nil {
		return errors.New("failed find id")
	}

	user.Token, err = d.service.GenerateToken(int(user.ID))
	if err != nil {
		return errors.New("failed find id")
	}

	err = d.db.Save(user).Error
	if err != nil {
		return errors.New("failed find id")
	}

	return nil
}
