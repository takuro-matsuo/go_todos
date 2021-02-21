package controllers

import (
	"go_todo/todo_app/config"
	"net/http"
)

// 127.0.0.1 ループバックアドレス（自分を指すIPアドレス）
func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
