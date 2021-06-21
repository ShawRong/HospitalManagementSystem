package datasource

import (
	"testing"
)

func TestDBConnect(t *testing.T) {
	if err := DBConnect(); err != nil {
		t.Errorf("%v\n", err)
	}
}

/*
func TestCreateUser(t *testing.T) {
	test := User{"Wang", true, 1, 10}
	if err := CreateUser(CMSdb, test); err != nil {
		t.Errorf("%v/n", err)
	}
}
PASS
*/

func TestFindUserbyID(t *testing.T) {
	var user User
	var err error
	if user, err = FindUserbyID(CMSdb, 1); err != nil {
		t.Errorf("%v/n", err)
	}
	if user.Username != "Wang" {
		t.Errorf("%s %d", user.Username, user.Honesty)
	}
}
