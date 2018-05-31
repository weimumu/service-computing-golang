package entities

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type userDB struct {
	db *gorm.DB
}

var mydb *userDB

func GetDao() *userDB {
	if mydb != nil {
		return mydb
	}
	db, err := gorm.Open("mysql", "test:test2017@tcp(139.199.174.146:3306)/mytest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println(db.HasTable(&User{}))
	if (!db.HasTable(&User{})) {
		db.CreateTable(&User{})
	}
	mydb = &userDB{db}
	return mydb
}

func (userDB *userDB) StoreUser(u *User) error {
	if userDB.db.Create(u).Error != nil {
		fmt.Println("error")
		return userDB.db.Error
	}
	return nil
}

func (userDB *userDB) GetAllUsers() ([]*User, error) {
	users := make([]*User, 0)
	if userDB.db.Find(&users).Error != nil {
		return nil, userDB.db.Error
	}
	return users, nil
}

func (userDB *userDB) GetUser(uid uint64) (*User, error) {
	var u User
	if userDB.db.Find(&u, uid).Error != nil {
		return nil, userDB.db.Error
	}
	if u.Username != "" {
		return &u, nil
	}
	return nil, nil
}
