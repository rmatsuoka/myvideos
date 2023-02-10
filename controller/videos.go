package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"path"

	"github.com/rmatsuoka/myvideos/model"
)

type Video struct {
	model.VideoModel
}

func (v Video) Info(w http.ResponseWriter, req *http.Request) {
	id := path.Base(req.URL.Path)
	switch req.Method {
	case "GET":
		info, err := v.VideoModel.Info(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(info)
	case "PUT":
		info := new(model.VideoInfo)
		defer req.Body.Close()
		err := json.NewDecoder(req.Body).Decode(info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = v.UpdateInfo(id, info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}

func (v Video) SetInfo(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	info := new(model.VideoInfo)
	defer req.Body.Close()
	err := json.NewDecoder(req.Body).Decode(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := v.VideoModel.SetInfo(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s := struct{ ID string }{id}
	json.NewEncoder(w).Encode(s)
}

func (v Video) Description(w http.ResponseWriter, req *http.Request) {
	id := path.Base(req.URL.Path)
	switch req.Method {
	case "GET":
		desc, err := v.VideoModel.Description(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Write(desc)
	case "PUT":
		defer req.Body.Close()
		desc, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = v.SetDescription(id, desc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}
