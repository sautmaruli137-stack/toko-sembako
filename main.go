package main

import (
	"html/template"
	"log"
	"net/http"
)

// Struct untuk data yang akan dikirim ke HTML
type PageData struct {
	Title   string
	Message string
}

func main() {
	// Handler untuk file statis (CSS, JS, gambar)
	// URL /static/ akan diarahkan ke folder ./static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handler untuk halaman utama
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Template tidak ditemukan", http.StatusInternalServerError)
			log.Println("Error parsing template:", err)
			return
		}

		data := PageData{
			Title:   "Halaman Go + HTML + CSS",
			Message: "Halo dari Golang dengan CSS!",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Gagal merender template", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	})

	log.Println("Server berjalan di http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server gagal:", err)
	}
}