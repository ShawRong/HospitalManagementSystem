package datasource

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateEmployee(db *gorm.DB, employee Employee) error {
	err := db.Create(&employee).Error
	return err
}

func FindEmployeebyID(db *gorm.DB, ID int) (Employee, error) {
	var employee Employee
	err := db.Where("id = ?", ID).First(&employee).Error
	if err != nil {
		var invalid Employee
		return invalid, err
	}
	return employee, nil
}

func FindEmployeebyName(db *gorm.DB, Name string) ([]Employee, error) {
	var employees []Employee
	db.Where("username = ?", Name).Find(&employees)
	if len(employees) == 0 {
		return employees, errors.New("empty slice")
	} else {
		return employees, nil
	}
}

func UpdateEmployee(db *gorm.DB, ID int, employee Employee) error {
	foundEmployee, err := FindCarbyID(db, ID)
	if err != nil {
		return err
	}
	if foundEmployee.ID != employee.ID {
		return errors.New("user consistence wrong")
	}
	db.Model(&foundEmployee).Updates(employee)
	return nil
}

func DeleteEmployee(db *gorm.DB, ID int) error {
	var employee Employee
	err := db.Where("id = ?", ID).First(&employee).Error
	if err != nil {
		return err
	}
	err1 := db.Delete(&employee).Error
	if err1 != nil {
		return err1
	}
	return nil
}
