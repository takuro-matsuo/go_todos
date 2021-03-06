package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"go_todos/config"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top) //top
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
