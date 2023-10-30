package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{
		Name: "Ram",
		Age:  24,
	}
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error encoding json:", err)
		return
	}
	fmt.Println("JSON Data:", string(jsonData))

	var p2 Person
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		fmt.Println("Error Decoding JSON:", err)
		return
	}
	fmt.Println("Decode Person:", p2)
}
