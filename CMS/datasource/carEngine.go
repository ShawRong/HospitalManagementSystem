package datasource

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateCar(db *gorm.DB, car Car) error {
	err := db.Create(&car).Error
	return err
}

func FindCarbyID(db *gorm.DB, ID int) (Car, error) {
	var car Car
	err := db.Where("id = ?", ID).First(&car).Error
	if err != nil {
		var invalid Car
		return invalid, err
	}
	return car, nil
}

func FindCarInRange(db *gorm.DB, min int, max int) ([]Car, error) {
	var cars []Car
	db.Where("id BETWEEN ? AND ?", min, max).Find(&cars)
	if len(cars) == 0 {
		return cars, errors.New("empty slice")
	} else {
		return cars, nil
	}
}

func UpdateCar(db *gorm.DB, ID int, car Car) error {
	foundCar, err := FindCarbyID(db, ID)
	if err != nil {
		return err
	}
	if foundCar.ID != car.ID {
		return errors.New("user consistence wrong")
	}
	db.Model(&foundCar).Updates(car)
	return nil
}

func DeleteCar(db *gorm.DB, ID int) error {
	var car Car
	err := db.Where("id = ?", ID).First(&car).Error
	if err != nil {
		return err
	}
	err1 := db.Delete(&car).Error
	if err1 != nil {
		return err1
	}
	return nil
}
