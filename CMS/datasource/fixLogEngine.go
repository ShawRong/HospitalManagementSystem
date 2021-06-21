package datasource

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateFixLog(db *gorm.DB, fixLog FixLog) error {
	err := db.Create(&fixLog).Error
	return err
}

func FindFixLogbyCarID(db *gorm.DB, CarID int) ([]FixLog, error) {
	var fixLogs []FixLog
	db.Where("car_id = ?", CarID).Find(&fixLogs)
	if len(fixLogs) == 0 {
		return fixLogs, errors.New("empty slice")
	} else {
		return fixLogs, nil
	}
}

func DeleteFixLog(db *gorm.DB, CarID int) error {
	var fixLog FixLog
	err := db.Where("car_id = ?", CarID).First(&fixLog).Error
	if err != nil {
		return err
	}
	err1 := db.Delete(&fixLog).Error
	if err1 != nil {
		return err1
	}
	return nil
}
