package controller

import (
	"encoding/json"
	"net/http"
)

type addTagBody struct {
	VideoID string
	TagName string
}

func (h *Handler) AddTag(w http.ResponseWriter, r *http.Request) {
	var b addTagBody
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tag, err := h.M.Tag().Add(b.VideoID, b.TagName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type deleteTagBody struct {
	VideoID string
	TagID   string
}

func (h *Handler) DeleteTag(w http.ResponseWriter, r *http.Request) {
	var b deleteTagBody
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.M.Tag().Delete(b.VideoID, b.TagID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
