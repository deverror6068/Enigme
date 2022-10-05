package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	file, err := ioutil.ReadFile("C:/Users/canel/Documents/Ynov/Tp/Enigme/Enigme/enigme.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	fmt.Println(text)
}
