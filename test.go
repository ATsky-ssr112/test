package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
	}
	templates := template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/base.html",
		"templates/view.html"))
	dat := struct {
		Title string
		Name  string
		User string
		Pass string
		Value string
	}{
		Title: "テスト",
		Name:  "まっさん",
		User: "Massan",
		Pass: "Massan111",
		Value: "コンスト",
	}
	err := templates.ExecuteTemplate(w, "base", dat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


