package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	if position == "left" {
		return arr[1:]
	} else if position == "right" {
		return arr[:len(arr)-1]
	}
	return arr
}

func main() {
	arr := []interface{}{"Edo", "Budi", "Joko", "Tono"}
	newArr := removeLeftRight(arr, "left")
	fmt.Println(newArr)

	fmt.Println()

	arr = []interface{}{"Edo", "Budi", "Joko", "Tono"}
	newArr = removeLeftRight(arr, "right")
	fmt.Println(newArr)
}
