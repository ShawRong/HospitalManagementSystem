package datasource

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Car struct {
	ID       int `gorm:"primary_key"`
	Deposite int
}

type User struct {
	Username string
	IsMember bool
	ID       int `gorm:"primary_key"`
	Honesty  int
}

type Employee struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type Log struct {
	UserID     int
	StatusID   int
	User       User       `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	RentStatus RentStatus `gorm:"foreignkey:StatusID;association_foreignkey:ID"`
}

type Rent struct {
	UserID     int
	CarId      int
	StatusID   int
	User       User       `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Car        Car        `gorm:"foreignkey:CarID;association_foreignkey:ID"`
	RentStatus RentStatus `gorm:"foreignkey:StatusID;association_foreignkey:ID"`
}

type RentStatus struct {
	ID     int `gorm:"primary_key"`
	UserID int
	CarID  int
	Fine   int
	Cost   int
	Date   time.Time
	User   User `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Car    Car  `gorm:"foreignkey:CarID;association_foreignkey:ID"`
}

type FixLog struct {
	CarID int
	Log   string
	Car   Car `gorm:"foreignkey:CarID;association_foreignkey:ID"`
}

var CMSdb *gorm.DB

func DBConnect() error {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/"+"CMSdb"+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	db.AutoMigrate(&Car{}, &User{}, &Employee{}, &Log{}, &Rent{}, &RentStatus{}, &FixLog{})
	CMSdb = db
	return nil
}

func DBClose(db *gorm.DB) {
	db.Close()
}
