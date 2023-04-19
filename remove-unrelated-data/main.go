package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	delete(dataMap, key)
	return dataMap
}

func main() {
	dataMap := map[string]interface{}{
		"name":    "Edo",
		"age":     20,
		"address": "Jakarta",
	}

	newMap := removeUnrelated(dataMap, "address")

	fmt.Println(newMap)
}
