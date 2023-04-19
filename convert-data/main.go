package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	parts := strings.Split(data, ",")
	age, _ := strconv.Atoi(parts[1])

	return User{
		Name:    parts[0],
		Age:     age,
		Address: parts[2],
	}
}

func main() {
	data := "Alexander Abraham,23,Galaxy"
	user := ConvertData(data)

	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Age: %d\n", user.Age)
	fmt.Printf("Address: %s\n", user.Address)
}
