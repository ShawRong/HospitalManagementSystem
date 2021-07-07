package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Item struct {
	Earn int
	Time int
}

func GetReportViewYear(db *gorm.DB) ([]Item, error) {
	var Results []Item
	rows, err := db.Raw(" select sum(earn) as earn, DATE_FORMAT(time,'%Y') as year from report group by DATE_FORMAT(time,'%Y');").Rows()
	defer rows.Close()
	var earn int
	var time int
	for rows.Next() {
		rows.Scan(&earn, &time)
		var result Item
		result.Earn = earn
		result.Time = time
		Results = append(Results, result)
	}
	return Results, err
}

func GetReportViewQuarter(db *gorm.DB) ([]Item, error) {
	var Results []Item
	rows, err := db.Raw(" select sum(earn) as earn, QUARTER(time) as quarter from report group by YEAR(time), QUARTER(time) order by YEAR(time), QUARTER(time);").Rows()
	defer rows.Close()
	var earn int
	var time int
	for rows.Next() {
		rows.Scan(&earn, &time)
		var result Item
		result.Earn = earn
		result.Time = time
		Results = append(Results, result)
	}
	return Results, err
}

func GetReportViewMonth(db *gorm.DB) ([]Item, error) {
	var Results []Item
	rows, err := db.Raw(" select sum(earn) as earn, DATE_FORMAT(time,'%Y%m') as month from report group by DATE_FORMAT(time,'%Y%m') order by DATE_FORMAT(time,'%Y%m');").Rows()
	defer rows.Close()
	var earn int
	var time int
	for rows.Next() {
		rows.Scan(&earn, &time)
		var result Item
		result.Earn = earn
		result.Time = time
		Results = append(Results, result)
	}
	return Results, err
}

func GetReportViewDay(db *gorm.DB) ([]Item, error) {
	var Results []Item
	rows, err := db.Raw(" select sum(earn) as earn, DATE_FORMAT(time,'%Y%m%d') as day from report group by DATE_FORMAT(time,'%Y%m%d') order by DATE_FORMAT(time,'%Y%m%d') ;").Rows()
	defer rows.Close()
	var earn int
	var time int
	for rows.Next() {
		rows.Scan(&earn, &time)
		var result Item
		result.Earn = earn
		result.Time = time
		Results = append(Results, result)
	}
	return Results, err
}
