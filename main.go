package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		handleErr("no action provided")
	}
	runCmd(args[0], args[1:]...)
}

func runCmd(action string, args...string) {
	switch action {
	case "add":
		if len(args) == 0 {
			handleErr("add: no description provided")
		}
		add(args[0])
	case "update":
	case "delete":
	case "mark":
	case "list":
	case "help":
	default: handleErr("invalid action provided")
	}
}


func add(desc string) {
	tasks := getTasks()
	tasks.Elements = append(tasks.Elements, Task{
			tasks.NextId(),
			desc,
			"todo",
			time.Now(),
			time.Now(),
		})
	jsonBytes := Unwrap(json.Marshal(tasks))
	Expect(os.WriteFile("./tasks.json", jsonBytes, 0644))
}

type Task struct {
	Id uint `json:"id"`
	Desc string `json:"desc"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Tasks struct {
	Elements []Task `json:"tasks"`
}

func (self Tasks) NextId() uint {
	if len(self.Elements) == 0 {
		return 1
	} 
	highestId := self.Elements[0].Id

	for _, task := range self.Elements[1:] {
		if task.Id > highestId {
			highestId = task.Id
		}
	}	

	return highestId + 1
}

func getTasks() Tasks {
	var tasks Tasks
	bytes, err := os.ReadFile("./tasks.json")
	if err != nil {
		fmt.Printf(
			"%s\n",
			Italic("task.json not found, creating..."),
		)
		Unwrap(os.Create("tasks.json"))
		jsonBytes := Unwrap(json.Marshal(new(Tasks)))
		Expect(os.WriteFile("./tasks.json", jsonBytes, 0644))
		bytes = Unwrap(os.ReadFile("./tasks.json"))
	}

	Expect(json.Unmarshal(bytes, &tasks))
	return tasks
}

