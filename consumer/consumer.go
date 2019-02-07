package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit Fruits
	Veggies Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

func main() {
	url := "http://localhost:8080/users"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	var pay Payload

	err = json.Unmarshal(body, &pay)
	if err != nil {
		panic(err)
	}

	for i, veggies := range pay.Stuff.Veggies {
		fmt.Printf("%v: %v\n", i, veggies)
	}
}
