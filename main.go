package main

import (
	"os"
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

}

func getTaskFile() 

