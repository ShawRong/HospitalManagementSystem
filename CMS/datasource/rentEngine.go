package datasource

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateRent(db *gorm.DB, rent Rent) error {
	err := db.Create(&rent).Error
	return err
}

func CreateRentPointer(db *gorm.DB, rent *Rent) error {
	err := db.Create(rent).Error
	return err
}

func FindRentbyUserID(db *gorm.DB, UserID int) ([]Rent, error) {
	var rents []Rent
	db.Where("user_id = ?", UserID).Find(&rents)
	if len(rents) == 0 {
		return rents, errors.New("empty slice")
	} else {
		return rents, nil
	}
}

func FindRentbyUserIDCarID(db *gorm.DB, UserID int, CarID int) (Rent, error) {
	var rent Rent
	err := db.Where("user_id = ? AND car_id = ?", UserID, CarID).Find(&rent).Error
	if err != nil {
		var invalid Rent
		return invalid, err
	}
	return rent, nil
}

func DeleteRent(db *gorm.DB, UserID int, CarID int) error {
	var rent Rent
	err := db.Where("user_id = ? AND car_id = ?", UserID, CarID).First(&rent).Error
	if err != nil {
		return err
	}
	err1 := db.Delete(&rent).Error
	if err1 != nil {
		return err1
	}
	return nil
}
