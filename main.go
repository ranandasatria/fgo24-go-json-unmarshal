package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var baseURL = "https://jsonplaceholder.typicode.com"

type Company struct {
	Name string
	CatchPhrase string
	Bs string
}

type User struct {
	ID    int
	Name  string
	Email string
	Company Company
}

func fetchUsers() ([]User, error) {
	var err error
	var client = &http.Client{}
	var data[]User

	request, err := http.NewRequest("GET", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	var users, err = fetchUsers()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	for _, each := range users {
		fmt.Printf("%v\n Name: %v \n Email: %v \n Company: \n  \t Name: %v \n \t Catchphrase: %v \n \t Bs: %v \n", each.ID, each.Name, strings.ToLower(each.Email), each.Company.Name, each.Company.CatchPhrase, each.Company.Bs)
	}

}

