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

	tasks := []*worker.Task{}
	fmt.Println(config)
	for i := 0; i < config.Total; i++ {
		for j := 0; j < len(config.Apis); j++ {
			tasks = append(tasks, worker.NewTask(request.Get, config.Apis[j], fmt.Sprintf("user%d:api %d", i, j)))

		}

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
