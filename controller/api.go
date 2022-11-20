package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rmatsuoka/myvideos/model"
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
	tag, err := h.M.AddTag(b.VideoID, b.TagName)
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
	err = h.M.DeleteTag(b.VideoID, b.TagID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type updateVideoInfoBody struct {
	VideoID   string
	VideoInfo model.VideoInfo
}

func (h *Handler) UpdateVideoInfo(w http.ResponseWriter, r *http.Request) {
	var b updateVideoInfoBody
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.M.UpdateVideoInfo(b.VideoID, &b.VideoInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type VideoIDBody struct {
	VideoID string
}

type responseIncrementLikebody struct {
	Likes int
}

func (h *Handler) IncrementLike(w http.ResponseWriter, r *http.Request) {
	var b VideoIDBody
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	i, err := h.M.IncrementLike(b.VideoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&responseIncrementLikebody{i})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) TagsWithVideo(w http.ResponseWriter, r *http.Request) {
	var b VideoIDBody
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t, err := h.M.TagsWithVideo(b.VideoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
