package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

func Home(w http.ResponseWriter, r *http.Request) {

	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Fatal("Something inside home handler", err)

	}
}

func renderPage(w http.ResponseWriter, html string, data jet.VarMap) error {

	view, err := views.GetTemplate(html)
	if err != nil {
		log.Fatal("Something wrong while rendering", err)
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Fatal("Something went wrong while executing the page")
		return err
	}
	return nil
}
