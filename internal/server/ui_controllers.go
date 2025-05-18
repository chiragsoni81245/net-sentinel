package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/chiragsoni81245/net-sentinel/internal/types"
)

type UIControllers struct {
    Server *types.Server
}

type HTTPError struct {
    StatusCode int
    Message string
}


func (uc *UIControllers) Dashboard(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("internal/server/templates/index.html")
    if err != nil {
        http.Error(w, "Something went wrong", http.StatusInternalServerError)
        log.Println(err)
        return
    }
    tmpl.ExecuteTemplate(w, "index.html", nil)
}
