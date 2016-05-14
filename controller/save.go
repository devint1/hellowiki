package controller

import (
	"github.com/devint1/hellowiki/model"
	"github.com/devint1/hellowiki/util"
	"net/http"
)

func SaveController(w http.ResponseWriter, r *http.Request) {
	m := util.GetWikiFile(w, r)
	if m == nil {
		return
	}
	title := m[2]
	body := r.FormValue("body")
	p := &model.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

