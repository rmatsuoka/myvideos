package controller

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/rmatsuoka/myvideos/model"
)

func (h *Handler) VideoInfo(w http.ResponseWriter, req *http.Request) {
	id := path.Base(req.URL.Path)
	switch req.Method {
	case "GET":
		info, err := h.M.VideoInfo(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(info)
	case "POST":
		info := new(model.VideoInfo)
		err := json.NewDecoder(req.Body).Decode(info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = h.M.UpdateVideoInfo(id, info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}

func (h Handler) CreateVideoInfo(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	info := new(model.VideoInfo)
	err := json.NewDecoder(req.Body).Decode(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.M.CreateVideoInfo(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s := map[string]string{"id": id}
	json.NewEncoder(w).Encode(s)
}
