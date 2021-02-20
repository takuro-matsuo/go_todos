package main

import (
	"fmt"
	"go_todo/todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.PassWord = "testtest"
	fmt.Println(u)
	u.CreateUser()
}
