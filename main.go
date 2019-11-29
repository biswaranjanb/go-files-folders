package main

import (
	"io/ioutil"
	"strings"
	"log"
	"fmt"
)

var filee = "platformbrigade.txt"

func main(){
	//Read all in the string
	r := strings.NewReader("We are a part of Platform brigade.")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

	//Read all in the directory
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	//Read File 
	data, err := ioutil.ReadFile(filee)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", data)

	//Write File
	message := []byte("Hello, Gophers!")
	err1 := ioutil.WriteFile(filee, message, 0644)
	if err != nil {
		log.Fatal(err1)
	}

}	