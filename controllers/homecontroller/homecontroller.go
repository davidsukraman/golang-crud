package homecontroller

import (
	"html/template"
	"log"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("views/home/index.html")

	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	temp.Execute(w, nil)
}
