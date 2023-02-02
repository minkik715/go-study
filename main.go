package main

import (
	"fmt"
	"learngo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{}
	dictionary["fuck2"] = "you"
	result, err := dictionary.Search("fuck")
	if result == "" {
		fmt.Println(err)
	}
	dictionary.Add("first", "two")
	dictionary.Add("first", "two")

	fmt.Println(dictionary)
}
