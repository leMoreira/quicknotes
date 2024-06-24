package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"views/templates/base.html",
		"views/templates/pages/home.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", nil)
}

func noteView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Nota não encontrada", http.StatusNotFound)
		return
	}

	files := []string{
		"views/templates/home.html",
		"views/templates/pages/note-view.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "base", nil)

}

func noteNew(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/note-new.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
	}
	t.ExecuteTemplate(w, "base", nil)
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Rejeitar a requisição

		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Método não permitido")
		return
	}
	fmt.Fprint(w, "Criando uma nova Nota")

}

func main() {
	fmt.Println("Servidor rodando na porta 8080 ")

	mux := http.NewServeMux()

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/new", noteNew)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":8080", mux)

}
