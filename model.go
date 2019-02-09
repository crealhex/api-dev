package main

type User struct {
	ID int `json:"user_id"`
	Username string `json:"username,omitempty"`
	Fullname string `json:"fullname,omitempty"`
	Age uint8 `json:"age,omitempty"`
	Email string `json:"email,omitempty"`
	City string `json:"city,omitempty"`
	Phone string `json:"phone,omitempty"`
	Friends []User `json:"friends,omitempty"`
}

type Users []*User
/*
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
*/