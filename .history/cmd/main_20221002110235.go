package main

import (
	"encoding/json"
	"fmt"
	"go-concur/internal/request"
	"go-concur/internal/worker"
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

	tasks := []*worker.Task{
		worker.NewTask(request.Get, config.Apis[0].Url, 0),
	}

	p := worker.NewPool(tasks, config.Concurrency)
	p.Run()

	var numErrors int
	for _, task := range p.Tasks {
		if task.Err != nil {
			log.Println(task.Err)
			numErrors++
		}
		if numErrors >= 10 {
			log.Println("Too many errors.")
			break
		}
	}
}
