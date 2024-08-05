package main

import (
	"fmt"
	"task_manager/router"
)

func main() {
	fmt.Println("This is api for task management")
	R := router.NewRouter()
	R.Run()
}
