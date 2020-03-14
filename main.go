package main

import (
	"fmt"

	"github.com/changhoi/learngo/mydict"
)

func main() {
	dic := mydict.Dictionary{}
	base := "hello"
	dic.Add(base, "first")
	dic.Update(base, "second")
	word, _ := dic.Search(base)
	fmt.Println(word)

	dic.Delete(base)
	_, err := dic.Search(base)
	if err != nil {
		fmt.Println(err)
	}
}
