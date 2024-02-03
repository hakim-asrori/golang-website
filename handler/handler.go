package handler

import (
	"golangweb/entity"
	"html/template"
	"net/http"
	"path"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Halaman Index"))

	temp, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":   "Home",
		"content": "Lorem ipsum dolor sit amet.",
	}

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")

	// idNum, err := strconv.Atoi(id)
	// if err != nil || idNum < 1 {
	// 	http.NotFound(w, r)
	// 	return
	// }

	temp, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// data := entity.Product{ID: 1, Name: "Mobilio", Price: 20000, Stock: 2}
	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 20000, Stock: 6},
		{ID: 2, Name: "Xpander", Price: 20000, Stock: 2},
		{ID: 3, Name: "Pajero Sport", Price: 20000, Stock: 7},
	}

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
