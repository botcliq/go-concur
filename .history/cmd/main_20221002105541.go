package main

import (
	"encoding/json"
	"fmt"
	"go-concur/internal/request"
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

	fmt.Println(config)

	tasks := []*Task{
		NewTask(func() error { return nil }),
		NewTask(func() error { return nil }),
		NewTask(func() error { return nil }),
	}

	p := worker.NewPool(tasks, config.Concurrency)
	p.Run()

	var numErrors int
	for _, task := range p.Tasks {
		if task.Err != nil {
			log.Error(task.Err)
			numErrors++
		}
		if numErrors >= 10 {
			log.Error("Too many errors.")
			break
		}
	}
}
