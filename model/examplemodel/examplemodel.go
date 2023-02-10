package examplemodel

import (
	"embed"
	"io/fs"

	"github.com/rmatsuoka/myvideos/model"
)

//go:embed videos/* thumb/*
var exampleFS embed.FS

type exampleModel int

// FS implements model.Model
func (exampleModel) FS() fs.FS {
	return exampleFS
}

// SearchModel implements model.Model
func (exampleModel) Search() model.SearchModel {
	return searchModel(0)
}

// Tag implements model.Model
func (exampleModel) Tag() model.TagModel {
	return tagMap
}

// VideoInfo implements model.Model
func (exampleModel) VideoInfo() model.VideoModel {
	return videoInfoModel(0)
}

var Model model.Model = exampleModel(0)
