package main

import (
	"fmt"

	"github.com/changhoi/learngo/mydict"
)

func main() {

	dic := mydict.Dictionary{"first": "First word"}

	err := dic.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}

	def, err := dic.Search("hello")
	fmt.Println(def)

	err = dic.Add("hello", "greeting")
	if err != nil {
		fmt.Println(err)
	}
}
