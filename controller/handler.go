package controller

import (
	"html/template"

	"github.com/rmatsuoka/myvideos/model"
)

type Handler struct {
	M *model.Model
	T *template.Template
}
