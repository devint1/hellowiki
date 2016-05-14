package controller

import (
	"github.com/devint1/hellowiki/model"
	"github.com/devint1/hellowiki/util"
	"net/http"
	"path/filepath"
	"strings"
)

func ViewController(w http.ResponseWriter, r *http.Request) {
	m := util.GetWikiFile(w, r)
	if m == nil {
		return
	}
	title := m[2]

	// Show a list of wikis to view
	if title == "" {
		files, _ := filepath.Glob("wiki/*.txt");
		for index, file := range files {
			file = strings.Replace(file, ".txt", "", -1)
			files[index] = strings.Replace(file, "wiki/", "", -1)
		}
		util.RenderTemplate(w, "view-list", files);
		return
	}

	p, err := model.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	util.RenderTemplate(w, "view", p)
}

