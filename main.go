package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Eg: ", os.Args[0], "'{a:1,b:2}'")
		os.Exit(1)
	}

	var js interface{}
	json.Unmarshal([]byte(os.Args[1]), &js)

	indent, err := json.MarshalIndent(js, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(indent))
}
