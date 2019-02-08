package main

import (
	"fmt"
)

type User struct {
	ID int `json:"user_id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	City string `json:"city"`
	Phone string `json:"phone"`
}

type Users []*User

func (u *User) String() string {
	return fmt.Sprintf("%d, %s, %s, %s, %s", u.ID, u.Username, u.Fullname, u.City, u.Phone)
}

func (us *Users) String() string {
	var rows string
	for _, u := range *us {
		fmt.Printf("%#v", u)
		rows += fmt.Sprintf("%d, %s, %s, %s, %s\n", u.ID, u.Username, u.Fullname, u.City, u.Phone)
	}
	fmt.Println(rows)
	return rows
}