package main

var list Users

func loadInitialData() {
	crealhex := &User{
		ID:       1,
		Username: "crealhex",
		Fullname: "Luis Enco",
		Age:      20,
		Email:    "encoluis@hotmail.com",
		City:     "Lima",
		Phone:    "952736356",
	}
	list = append(list, crealhex)
}

func getAll() *Users {
	var users Users
	var friends Users
	for i, user := range list {
		for _, friend := range list[i].Friends {
			curr := &User{
				ID:       friend.ID,
				Username: friend.Username,
				Fullname: friend.Fullname,
				Age:      friend.Age,
				Email:    friend.Email,
				City:     friend.City,
				Phone:    friend.Phone,
			}
			friends = append(friends, curr)
		}
		curr := &User{
			ID:       user.ID,
			Username: user.Username,
			Fullname: user.Fullname,
			Age:      user.Age,
			Email:    user.Email,
			City:     user.City,
			Phone:    user.Phone,
			Friends:  friends,
		}
		friends = nil
		users = append(users, curr)
	}
	return &users
}

func getByID(id int) *User {
	for _, v := range list {
		if v.ID == id {
			return v
		}
	}
	return nil
}

func add(c *User) {
	c.ID = genNewID()
	list = append(list, c)
}

func genNewID() int {
	size := len(list)
	if size > 0 {
		return list[size-1].ID + 1
	}
	return 1
}

func update(id int, data *User) {
	user := findUser(id)
	user.Username = data.Username
	user.Fullname = data.Fullname
	user.Age = data.Age
	user.Email = data.Email
	user.City = data.City
	user.Phone = data.Phone
}

func findUser(id int) *User {
	for _, user := range list {
		if user.ID == id {
			return user
		}
	}
	return nil
}

func deleteu(id int) {
	list = append(list[:id-1], list[id:]...)
}

func addFriend(one, two int) {
	user, other := &User{}, &User{}
	for _, v := range list {
		if v.ID == one {
			user = v
		}
		if v.ID == two {
			other = v
		}
	}
	user.Friends = append(user.Friends, other)
	other.Friends = append(other.Friends, user)
}