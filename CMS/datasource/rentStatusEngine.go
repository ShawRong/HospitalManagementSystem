package datasource

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateRentStatus(db *gorm.DB, rentStatus RentStatus) error {
	err := db.Create(&rentStatus).Error
	return err
}

func FindRentStatusbyID(db *gorm.DB, ID int) (RentStatus, error) {
	var rentStatus RentStatus
	err := db.Where("id = ?", ID).First(&rentStatus).Error
	if err != nil {
		var invalid RentStatus
		return invalid, err
	}
	return rentStatus, nil
}

func FindRentStatusbyUserID(db *gorm.DB, UserID int) ([]RentStatus, error) {
	var rentStatuss []RentStatus
	db.Where("user_id = ?", UserID).Find(&rentStatuss)
	if len(rentStatuss) == 0 {
		return rentStatuss, errors.New("empty slice")
	} else {
		return rentStatuss, nil
	}
}

func FindRentStatusbyCarID(db *gorm.DB, CarID int) ([]RentStatus, error) {
	var rentStatuss []RentStatus
	db.Where("car_id = ?", CarID).Find(&rentStatuss)
	if len(rentStatuss) == 0 {
		return rentStatuss, errors.New("empty slice")
	} else {
		return rentStatuss, nil
	}
}

func UpdateRentStatus(db *gorm.DB, ID int, rentStatus RentStatus) error {
	foundRentStatus, err := FindRentStatusbyID(db, ID)
	if err != nil {
		return err
	}
	if foundRentStatus.ID != rentStatus.ID {
		return errors.New("rentStatus consistence wrong")
	}
	db.Model(&foundRentStatus).Updates(rentStatus)
	return nil
}

func DeleteRentStatus(db *gorm.DB, ID int) error {
	var rentStatus RentStatus
	err := db.Where("id = ?", ID).First(&rentStatus).Error
	if err != nil {
		return err
	}
	err1 := db.Delete(&rentStatus).Error
	if err1 != nil {
		return err1
	}
	return nil
}
