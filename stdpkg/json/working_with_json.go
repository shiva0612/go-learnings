package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// encoding_decoding()
	rawmsg()
}

func encoding_decoding() {
	f, _ := os.Create("encode.json")
	a := map[int]int{1: 1, 2: 2}
	err := json.NewEncoder(f).Encode(a)
	if err != nil {
		log.Println(err.Error())
	}

	ff, _ := os.Open("encode.json")
	aa := map[string]int{}
	json.NewDecoder(ff).Decode(&aa)
	fmt.Println(aa)
}

type MyData struct {
	ID int `json:"id"`
	/*
		you do not know if it is map[string]string, map[string]int
		it may have {name:shiva, age:20} or {name:shiva, city:knr}
		in these situations, u can use jaw.rawMsg -> get the bytes -> store in DB if required -> process it later
	*/
	Payload json.RawMessage `json:"payload"`
}

func rawmsg() {
	jsonStr := `{"id": 123, "payload": {"name": "John", "age": 30}}`

	var data MyData
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("ID:", data.ID)
	fmt.Println("Raw Payload:", string(data.Payload))

	// Access specific fields within the payload
	var payloadData struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	err = json.Unmarshal(data.Payload, &payloadData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Name:", payloadData.Name)
	fmt.Println("Age:", payloadData.Age)
}
