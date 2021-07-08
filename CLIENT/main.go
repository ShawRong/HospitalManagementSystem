package main

import (
	"fmt"
	"os"
	"os/exec"

	"rong.com/client/routerTest"
)

func opt2() bool {
	clean()
	var option int
	fmt.Printf("--USER------------------------\n")
	fmt.Printf("1.New User With Username IsMember Honesty\n")
	fmt.Printf("2.Get User By ID\n")
	fmt.Printf("3.Get User By Name\n")
	fmt.Printf("4.Delete User By ID\n")
	fmt.Printf("--CAR-------------------------\n")
	fmt.Printf("5.New Car With Deposite\n")
	fmt.Printf("6.Get Car By ID\n")
	fmt.Printf("7.Get Car In Range(min, max)\n")
	fmt.Printf("8.Delete Car By ID\n")
	fmt.Printf("--Rent------------------------\n")
	fmt.Printf("9.New Rent With User ID AND Car ID\n")
	fmt.Printf("10.Get Rent With User ID\n")
	fmt.Printf("11.Get PayBack With User ID AND Car ID\n")
	fmt.Printf("12.Delete Rent With UserID AND CarID\n")
	fmt.Printf("13.Get RentStatus With ID\n")
	fmt.Printf("14.Set RentStatus By ID With Fine, Cost, Chager ID\n")
	fmt.Printf("--FIX LOG---------------------\n")
	fmt.Printf("15.New Fix Log With Car ID AND the Fee\n")
	fmt.Printf("16.Get Fix Log In Range(min, max)\n")
	fmt.Printf("17.Delete Fix Log By ID\n")
	fmt.Printf("--REPORT----------------------\n")
	fmt.Printf("18.Report In Year OR Quarter OR Month OR Day\n")
	fmt.Printf("--BACK------------------------\n")
	fmt.Printf("19.Back\n")
	fmt.Scanf("%d\n", &option)
	if option == 1 {
		var username string
		var IsMember string
		var honesty int

		fmt.Scanf("%s %s %d\n", &username, &IsMember, &honesty)
		var temp bool
		if IsMember == "true" {
			temp = true
		} else if IsMember == "false" {
			temp = false
		}
		routerTest.NewUserTest(username, temp, honesty)
		pause()
	} else if option == 2 {
		var ID int
		fmt.Scanf("%d\n", &ID)
		routerTest.GetUserTest(ID, "")
		pause()
	} else if option == 3 {
		var username string
		fmt.Scanf("%s\n", &username)
		routerTest.GetUserTest(-1, username)
		pause()
	} else if option == 4 {
		var ID int
		fmt.Scanf("%d\n", &ID)
		routerTest.DeleteUserTest(ID)
		pause()
	} else if option == 5 {
		var ID int
		fmt.Scanf("%d\n", &ID)
		routerTest.NewCarTest(ID)
		pause()
	} else if option == 6 {
		var ID int
		fmt.Scanf("%d\n", &ID)
		routerTest.GetCarTest(ID)
		pause()
	} else if option == 7 {
		var min, max int
		fmt.Scanf("%d %d\n", &min, &max)
		routerTest.GetCarInRangeTest(min, max)
		pause()
	} else if option == 8 {
		var ID int
		fmt.Scanf("%d\n", &ID)
		routerTest.DeleteCarTest(ID)
		pause()
	} else if option == 9 {
		var UserID, CarID int
		fmt.Scanf("%d %d\n", &UserID, &CarID)
		routerTest.NewRentTest(UserID, CarID)
		pause()
	} else if option == 10 {
		var ID int
		fmt.Scanf("%d\n", &ID)
		routerTest.GetRentTest(ID)
		pause()
	} else if option == 11 {
		var UserID, CarID int
		fmt.Scanf("%d %d\n", &UserID, &CarID)
		routerTest.GetPayBackTest(UserID, CarID)
		pause()
	} else if option == 12 {
		var UserID, CarID int
		fmt.Scanf("%d %d\n", &UserID, &CarID)
		routerTest.DeleteRentTest(UserID, CarID)
		pause()
	} else if option == 13 {
		var ID int
		fmt.Scanf("%d\n", &ID)
		routerTest.GetRentStatusTest(ID)
		pause()
	} else if option == 14 {
		var ID, Fine, Cost, ChargerID int
		fmt.Scanf("%d %d %d %d\n", &ID, &Fine, &Cost, &ChargerID)
		routerTest.SetRentStatusTest(ID, Fine, Cost, ChargerID)
		pause()
	} else if option == 15 {
		var CarID, Fee int
		fmt.Scanf("%d %d\n", &CarID, &Fee)
		routerTest.NewFixLogTest(CarID, Fee)
		pause()
	} else if option == 16 {
		var min, max int
		fmt.Scanf("%d %d\n", &min, &max)
		routerTest.GetFixLogInRangeTest(min, max)
		pause()
	} else if option == 17 {
		var ID int
		fmt.Scanf("%d\n", &ID)
		routerTest.DeleteFixLogTest(ID)
		pause()
	} else if option == 18 {
		var opt string
		fmt.Scanf("%s\n", &opt)
		routerTest.ReportTest(opt)
		pause()
	} else if option == 19 {
		return false
	}
	return true
}

func opt1() bool {
	clean()
	var option int
	fmt.Printf("1.My Information\n")
	fmt.Printf("2.Find Car Information By id\n")
	fmt.Printf("3.Back\n")
	fmt.Scanf("%d\n", &option)
	if option == 1 {
		var ID, password int
		fmt.Printf("Input your ID and your PassWord.\n")
		fmt.Scanf("%d %d\n", &ID, &password)
		routerTest.GetUserSelfTest(ID, password)
		pause()
	} else if option == 2 {
		var ID int
		fmt.Printf("Input the Car ID for search\n")
		fmt.Scanf("%d\n", &ID)
		routerTest.GetCarTestNormal(ID)
		pause()
	} else {
		return false
	}
	return true
}

func clean() {
	cmd := exec.Command("cmd.exe", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func pause() {
	fmt.Scanf("\n")
}

func main() {
	routerTest.Init()
	var opt int
	for true {
		clean()
		fmt.Printf("Select Your Option.\n")
		fmt.Printf("1.User Option.\n")
		fmt.Printf("2.Admin Option\n")
		fmt.Printf("3.Exit\n")
		fmt.Scanf("%d\n", &opt)
		if opt == 1 {
			for true {
				if !opt1() {
					break
				}
			}
		} else if opt == 2 {
			for true {
				if !opt2() {
					break
				}
			}
		} else if opt == 3 {
			fmt.Printf("Have a good day\n")
			pause()
			break
		}

	}

	/*
		routerTest.NewUserTest("Ming", true, 1)
		routerTest.GetUserTest(1, "")
		routerTest.GetUserTest(-1, "Ming")
		routerTest.DeleteUserTest(3)
	*/
	/*
		routerTest.NewCarTest(12)
		routerTest.GetCarTest(1)
		routerTest.GetCarInRangeTest(1, 2)
		routerTest.DeleteCarTest(2)
	*/
	/*
		routerTest.NewRentTest(1, 5)
		routerTest.GetRentTest(1)
		routerTest.GetPayBackTest(1, 5)
		routerTest.DeleteRentTest(1, 5)
	*/
	/*
		routerTest.GetRentStatusTest(7)
		routerTest.SetRentStatusTest(7, 21, 23, 0)
	*/
	/*
		routerTest.NewFixLogTest(1, 10)
		routerTest.GetFixLogInRangeTest(1, 20)
		routerTest.DeleteFixLogTest(10)
	*/
	/*
		routerTest.ReportTest("year")
		routerTest.ReportTest("quarter")
		routerTest.ReportTest("month")
		routerTest.ReportTest("day")
	*/
	/*
		routerTest.GetUserSelfTest(1, 1)
		routerTest.GetCarTestNormal(1)
	*/

}
