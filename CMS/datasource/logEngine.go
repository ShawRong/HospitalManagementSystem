package datasource

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateLog(db *gorm.DB, log Log) error {
	err := db.Create(&log).Error
	return err
}

func FindLogbyUserID(db *gorm.DB, UserID int) ([]Log, error) {
	var logs []Log
	db.Where("user_id = ?", UserID).Find(&logs)
	if len(logs) == 0 {
		return logs, errors.New("empty slice")
	} else {
		return logs, nil
	}
}

func DeleteLog(db *gorm.DB, UserID int, StatusID int) error {
	var log Log
	err := db.Where("user_id = ? AND status_id = ?", UserID, StatusID).First(&log).Error
	if err != nil {
		return err
	}
	err1 := db.Delete(&log).Error
	if err1 != nil {
		return err1
	}
	return nil
}
