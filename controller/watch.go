package controller

import (
	"net/http"
	"path"
)

func (h *Handler) WatchHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)

	info, err := h.M.VideoInfo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = h.T.ExecuteTemplate(w, "watch.html", info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
