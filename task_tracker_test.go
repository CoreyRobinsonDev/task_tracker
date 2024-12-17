package main

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestAddNoJson(t *testing.T) {
	cleanup()
	var ans Tasks
	expect := Tasks{
		[]Task{{
			Id: 1,
			Desc: "test",
			Status: "todo",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}},
	}
	add("test")
	bytes := Unwrap(os.ReadFile("./tasks.json"))
	Expect(json.Unmarshal(bytes, &ans))

	if expect.Elements[0].Id != ans.Elements[0].Id ||
	expect.Elements[0].Desc != ans.Elements[0].Desc ||
	expect.Elements[0].Status != ans.Elements[0].Status {
		t.Errorf(
			"add(test) =\n{\n\tId: %d,\n\tDesc: %s,\n\tStatus: %s,\n}\nexpected =\n{\n\tId: 1,\n\tDesc: test,\n\tStatus: todo,\n} ",
			ans.Elements[0].Id,
			ans.Elements[0].Desc,
			ans.Elements[0].Status,
		)
	}
	cleanup()
}

func TestAddJson(t *testing.T) {
	cleanup()
	var ans Tasks
	expect := Tasks{
		[]Task{{
			Id: 2,
			Desc: "test2",
			Status: "todo",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}},
	}
	add("test1")
	add("test2")
	bytes := Unwrap(os.ReadFile("./tasks.json"))
	Expect(json.Unmarshal(bytes, &ans))
	if expect.Elements[0].Id != ans.Elements[1].Id ||
	expect.Elements[0].Desc != ans.Elements[1].Desc ||
	expect.Elements[0].Status != ans.Elements[1].Status {
		t.Errorf(
			"add(test) =\n{\n\tId: %d,\n\tDesc: %s,\n\tStatus: %s,\n}\nexpected =\n{\n\tId: 2,\n\tDesc: test2,\n\tStatus: todo,\n} ",
			ans.Elements[1].Id,
			ans.Elements[1].Desc,
			ans.Elements[1].Status,
		)
	}
	cleanup()
}


func cleanup() {
	_, err := os.ReadFile("./tasks.json")
	if err == nil {
		Expect(os.Remove("./tasks.json"))
	}
}

