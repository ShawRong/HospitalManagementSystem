package datasource

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateUser(db *gorm.DB, user User) error {
	err := db.Create(&user).Error
	return err
}

func FindUserbyID(db *gorm.DB, ID int) (User, error) {
	var user User
	err := db.Where("id = ?", ID).First(&user).Error
	if err != nil {
		var invalid User
		return invalid, err
	}
	return user, nil
}

func FindUserbyName(db *gorm.DB, Name string) ([]User, error) {
	var users []User
	db.Where("username = ?", Name).Find(&users)
	if len(users) == 0 {
		return users, errors.New("empty slice")
	} else {
		return users, nil
	}
}

func UpdateUser(db *gorm.DB, ID int, user User) error {
	foundUser, err := FindUserbyID(db, ID)
	if err != nil {
		return err
	}
	if foundUser.ID != user.ID {
		return errors.New("user consistence wrong")
	}
	db.Model(&foundUser).Updates(user)
	return nil
}

func DeleteUser(db *gorm.DB, ID int) error {
	var user User
	err := db.Where("id = ?", ID).First(&user).Error
	if err != nil {
		return err
	}
	err1 := db.Delete(&user).Error
	if err1 != nil {
		return err1
	}
	return nil
}
