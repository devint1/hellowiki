package main

import (
	"github.com/devint1/hellowiki/controller"
	"net/http"
)

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/view/", http.StatusFound)
			return
		}
		if fn == nil {
			http.ServeFile(w, r, "public" + r.URL.Path)
			return
		}
		fn(w, r)
	}
}

func main() {
	http.HandleFunc("/", makeHandler(nil))
	http.HandleFunc("/view/", makeHandler(controller.ViewController))
	http.HandleFunc("/edit/", makeHandler(controller.EditController))
	http.HandleFunc("/save/", makeHandler(controller.SaveController))

	http.ListenAndServe(":8080", nil)
}

