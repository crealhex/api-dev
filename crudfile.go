package main

var list Users

func loadInitialData() {
	u1 := User{
		ID: 1,
		Username: "crealhex",
		Fullname: "Luis Enco",
		City: "Lima",
		Phone: "952736356",
	}

	u2 := User{
		ID: 2,
		Username: "crealhex",
		Fullname: "Luis Enco",
		City: "Lima",
		Phone: "952736356",
	}

	u3 := User{
		ID: 3,
		Username: "crealhex",
		Fullname: "Luis Enco",
		City: "Lima",
		Phone: "952736356",
	}
	list = append(list, &u1, &u2, &u3)
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

func update(id int, c *User) {
	for _, v := range list {
		if v.ID == id {
			v = c
			return
		}
	}
}

func delete(id int) {
	list[len(list) - 1], list[id] = list[id], list[len(list) - 1]
	list = list[:len(list) - 1]
}