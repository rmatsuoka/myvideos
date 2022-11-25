package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/rmatsuoka/myvideos/controller"
	"github.com/rmatsuoka/myvideos/model/examplemodel"
)

func main() {
	t := template.Must(template.ParseFiles("template/watch.html"))
	m := examplemodel.Model
	h := controller.Handler{M: m, T: t}

	http.HandleFunc("/watch/", h.WatchHandler)
	http.Handle("/model/", http.StripPrefix("/model/", http.FileServer(http.FS(m.FS()))))
	http.Handle("/statics/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/api/addtag", h.AddTag)
	http.HandleFunc("/api/deletetag", h.DeleteTag)
	http.HandleFunc("/api/tagswithvideo", h.TagsWithVideo)
	http.HandleFunc("/api/updatevideoinfo", h.UpdateVideoInfo)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
