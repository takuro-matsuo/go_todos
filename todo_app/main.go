package main

import (
	"fmt"
	"go_todo/todo_app/app/models"
)

func main() {
	// fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.PassWord = "testtest"
	fmt.Println(u)
	u.CreateUser()

	u2, _ := models.GetUser(1)
	fmt.Println(u2)
	u2.Name = "Test2"
	u2.Email = "test2@example.com"
	u2.UpdateUser()
	u2, _ = models.GetUser(1)
	fmt.Println(u2)
}
