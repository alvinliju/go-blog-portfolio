package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/war-logs", warlogs)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))

	fmt.Println("🚀 Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", "layout.html", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", "layout.html", nil)
}

func warlogs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "war-logs.md")
}

func renderTemplate(w http.ResponseWriter, contentTempl string, layoutTmpl string, data interface{}) {
	//parsing the specified template file passed as input
	t, err := template.ParseFiles("templates/"+layoutTmpl, "templates/"+contentTempl)
	if err != nil {
		http.Error(w, "Error parsing document", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
