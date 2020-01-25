package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"changhoi", "eunjin"}
	for _, person := range people {
		go isGood(person, c)
	}
	result := <-c
	fmt.Println(result)
}

func isGood(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true
}
