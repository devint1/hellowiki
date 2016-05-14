package controller

import (
	"github.com/devint1/hellowiki/model"
	"github.com/devint1/hellowiki/util"
	"net/http"
)

func EditController(w http.ResponseWriter, r *http.Request) {
	m := util.GetWikiFile(w, r)
	if m == nil {
		return
	}
	title := m[2]
	p, err := model.LoadPage(title)
	if err != nil {
		p = &model.Page{Title: title}
	}
	util.RenderTemplate(w, "edit", p)
}

