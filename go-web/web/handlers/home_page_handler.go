package handlers

import (
	"go-web/internal/app/web"
	"net/http"
)

type pageContent struct {
	Count int
}

var content = pageContent{
	Count: 0,
}

func HomePageHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl := web.NewTemplate("home", "web/templates/home.html")
	content.Count += 1
	tmpl.Render(w, content)
}
