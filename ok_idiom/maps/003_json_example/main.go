package main

import (
	"encoding/json"
	"log"
	"strconv"
)

func main() {

	jsonBytes := []byte(`{
		"name":"John", 
		"age":"17", 
		"car":null
	}`)

	checkAge(jsonBytes)

}

func checkAge(jsonBytes []byte) {

	var jsonMap map[string]interface{}

	err := json.Unmarshal(jsonBytes, &jsonMap)
	if err != nil {
		log.Fatalln(err)
	}

	age := jsonMap["age"]
	if age == nil {
		log.Println("age is nil")
		return
	}

	ageAsFloat, ok := age.(float64)
	if ok {
		if ageAsFloat > 18 {
			log.Println("age is greater than 18")
		} else {
			log.Println("age is less than 18")
		}
		return
	}
	if ageAsString, ok := age.(string); ok {
		// convert ageAsString to float64
		stringsAsFloat, err := strconv.ParseFloat(ageAsString, 64)
		if err != nil {
			log.Println("age is string but not a number")
			return
		}
		if stringsAsFloat > 18 {
			log.Println("age is greater than 18")
		} else {
			log.Println("age is less than 18")
		}
	} else {
		log.Println("age is not float64 neither string")
		return
	}
}
