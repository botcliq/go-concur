package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fileContent, err := os.Open("config.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("The File is opened successfully...")
	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)
	var config request.Config

	json.Unmarshal(byteResult, &config)
	var res map[string]interface{}
	json.Unmarshal([]byte(byteResult), &res)

	fmt.Println(res["users"])
}
