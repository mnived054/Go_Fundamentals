package main

import (
	"net/http"
	"text/template"
)

type ContactDetails struct {
	Name    string
	Email   string
	Subject string
	Message string
}

// this is mux handler

func main() {
	tmpl := template.Must(template.ParseFiles("D:/Go data types/forms/form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Name:    r.FormValue("name"),
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8088", nil)
}
