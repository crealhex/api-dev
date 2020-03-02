package main

type User struct {
	ID 			int 		`json:"user_id"`
	Username 	string 		`json:"username,omitempty"`
	Fullname 	string 		`json:"fullname,omitempty"`
	Age 		uint8 		`json:"age,omitempty"`
	Email 		string 		`json:"email,omitempty"`
	City 		string 		`json:"city,omitempty"`
	Phone 		string 		`json:"phone,omitempty"`
	Friends 	[]*User 	`json:"friends,omitempty"`
}

type Friend struct {
	ID 			int 		`json:"user_id"`
	Username 	string 		`json:"username,omitempty"`
	Fullname 	string 		`json:"fullname,omitempty"`
	Age 		uint8 		`json:"age,omitempty"`
	Email 		string 		`json:"email,omitempty"`
	City 		string 		`json:"city,omitempty"`
	Phone 		string 		`json:"phone,omitempty"`
}

type Users []*User