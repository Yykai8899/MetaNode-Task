package main

import (
	"task-go/task-go/go-base_4/structs"
)

func main() {
	r := structs.SetupRouter()

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
