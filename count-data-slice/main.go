package main

import "fmt"

func howManyElements(data []any) int {
	return len(data)
}

func main() {
	input := []interface{}{"Edo", "Budi", "Joko", "Tono", "Edo", "Budi", "Joko", "Tono"}
	output := howManyElements(input)
	fmt.Println(output)
}
