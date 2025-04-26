package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"course-name"`
	Price int 	
	Platform string `json:"platform"`
	Password string `json:"-"` 
	Tags []string  `json:"tags,omitempty"` // omitempty - if the field is nill ignore that field.
}

func main() {
	fmt.Println("Wellcome to JSON")
	EncodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "1234", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "1234", []string{"web-dev", "js"}},
		{"GOLANG Bootcamp", 299, "LearnCodeOnline.in", "1234", nil},
	}

	// package this data as JSON data

	// finalJson, err := json.Marshal(lcoCourses)

	/*
		json.MarshalIndent()
		params : interface, prefix, indentation
	*/
	finalJson, err := json.MarshalIndent(lcoCourses, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))

}