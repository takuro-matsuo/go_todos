package main

import (
	"fmt"

	"go_todos/app/controllers"
	"go_todos/app/models"
)

func TestConnection() {

}

func main() {
	fmt.Println(models.Db)
	go controllers.StartMainServer()

	for {

	}
}
