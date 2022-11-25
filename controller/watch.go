package controller

import (
	"net/http"
	"path"

	"github.com/rmatsuoka/myvideos/model"
)

type WatchPage struct {
	Info          *model.VideoInfo
	RelatedVideos []*model.VideoInfo
}

func (h *Handler) WatchHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)

	info, err := h.M.VideoInfo().Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	rv, err := h.M.Search().RelatedTo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	wp := &WatchPage{Info: info, RelatedVideos: rv}
	err = h.T.ExecuteTemplate(w, "watch.html", wp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
