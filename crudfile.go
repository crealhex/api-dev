package main

var list Users

func loadInitialData() {
	crealhex := User{
		ID: 1,
		Username: "crealhex",
		Fullname: "Luis Enco",
		Age: 20,
		Email: "encoluis@hotmail.com",
		City: "Lima",
		Phone: "952736356",
	}
	list = append(list, &crealhex)
}

func getAll() Users {
	return list
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
		return list[size - 1].ID + 1
	}
	return 1
}

func update(id int, u *User) {
	for i, v := range list {
		if v.ID == id {
			list[i] = u
			return
		}
	}
}

func deleteu(id int) {
	list = append(list[:id-1], list[id:]...)
}