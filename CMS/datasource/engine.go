package datasource

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Car struct {
	ID       int `gorm:"primary_key;AUTO_INCREMENT"`
	Deposite int
}

type User struct {
	Username string
	IsMember bool
	ID       int `gorm:"primary_key;AUTO_INCREMENT"`
	Honesty  int
}

type Employee struct {
	ID   int `gorm:"primary_key;AUTO_INCREMEN"`
	Name string
}

type Log struct {
	UserID   int
	StatusID int
}

type Rent struct {
	UserID   int `gorm:"primary_key;AUTO_INCREMENT:false"`
	CarId    int `gorm:"primary_key;AUTO_INCREMENT:false"`
	StatusID int
}

type RentStatus struct {
	ID        int `gorm:"primary_key;AUTO_INCREMENT"`
	UserID    int
	CarID     int
	Fine      int
	Cost      int
	ChargerID int
	Date      time.Time
}

type FixLog struct {
	ID    int `gorm:"primary_key;AUTO_INCREMENT"`
	CarID int
	Fee   int
	Time  time.Time
}

var CMSdb *gorm.DB

func DBConnect() error {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/"+"CMSdb"+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	db.AutoMigrate(&Car{}, &User{}, &Employee{}, &Log{}, &Rent{}, &RentStatus{}, &FixLog{})
	db.Model(&Log{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Log{}).AddForeignKey("status_id", "rent_statuses(id)", "CASCADE", "CASCADE")
	db.Model(&Rent{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Rent{}).AddForeignKey("car_id", "cars(id)", "CASCADE", "CASCADE")
	db.Model(&Rent{}).AddForeignKey("status_id", "rent_statuses(id)", "CASCADE", "CASCADE")
	db.Model(&RentStatus{}).AddForeignKey("charger_id", "employees(id)", "CASCADE", "CASCADE")
	db.Model(&RentStatus{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&RentStatus{}).AddForeignKey("car_id", "cars(id)", "CASCADE", "CASCADE")
	db.Model(&FixLog{}).AddForeignKey("car_id", "cars(id)", "CASCADE", "CASCADE")
	CMSdb = db
	return nil
}

func DBClose(db *gorm.DB) {
	db.Close()
}
