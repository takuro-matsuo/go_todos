package main

import (
	"fmt"
	"go_todos/app/models"
)

func main() {
	// controllers.StartMainServer()

	user, _ := models.GetUserByEmail("test@gmail.com")
	fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)
}
