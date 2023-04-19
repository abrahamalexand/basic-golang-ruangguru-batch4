package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		newElement := []int{newData}
		newElement = append(newElement, data...)
		return newElement
	} else if position == "down" {
		newElement := []int{newData}
		data = append(data, newElement...)
		return data
	}

	return nil
}

func main() {
	data := []int{1, 2, 3, 4}
	newData := 5
	position := "up"

	result := AddElement(data, newData, position)
	fmt.Println(result)

	fmt.Println()

	data = []int{3, 4, 5, 7}
	newData = 9
	position = "down"

	result = AddElement(data, newData, position)
	fmt.Println(result)
}
