package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/rmatsuoka/myvideos/controller"
	"github.com/rmatsuoka/myvideos/model"
)

func main() {
	t := template.Must(template.ParseFiles("template/watch.html"))
	m := model.NewModel()
	h := controller.Handler{M: m, T: t}

	http.HandleFunc("/watch/", h.WatchHandler)
	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.FS(m.FS()))))
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("."))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
