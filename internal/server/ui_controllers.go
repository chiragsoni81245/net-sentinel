package server

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/chiragsoni81245/net-sentinel/internal/packets"
	"github.com/chiragsoni81245/net-sentinel/internal/types"
	"github.com/chiragsoni81245/net-sentinel/internal/utils"
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
        http.Redirect(w, r, "/error", 302)
        log.Println(err)
        return
    }
    tmpl.ExecuteTemplate(w, "index.html", nil)
}

func (uc *UIControllers) Error(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("internal/server/templates/error.html")
    if err != nil {
        http.Redirect(w, r, "/error", 302)
        log.Println(err)
        return
    }
    tmpl.ExecuteTemplate(w, "error.html", nil)
}

func (uc *UIControllers) Login(w http.ResponseWriter, r *http.Request) {
    if r.Context().Value("userId") != nil {
        http.Redirect(w, r, "/", 302)
        return
    }
    tmpl, err := template.ParseFiles("internal/server/templates/login.html")
    if err != nil {
        http.Redirect(w, r, "/error", 302)
        log.Println(err)
        return
    }
    toasts := utils.ParseToasts(w, r)
    var payload struct {
        Toasts []types.Toast
    }
    payload.Toasts = *toasts
    tmpl.ExecuteTemplate(w, "login.html", payload)
}

func (uc *UIControllers) Logout(w http.ResponseWriter, r *http.Request) {
    if r.Context().Value("userId") != nil {
        http.SetCookie(w, &http.Cookie{
            Name: "token",
            Value: "",
            Expires: time.Now(),
        })
        http.Redirect(w, r, "/login", 302)
        return
    }
}

func (uc *UIControllers) Devices(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("internal/server/templates/devices.html")
    if err != nil {
        http.Redirect(w, r, "/error", 302)
        log.Println(err)
        return
    }

    
    devices, err := packets.GetAllDevices()
    if err != nil {
        http.Redirect(w, r, "/error", 302)
        log.Println(err)
    }

    tmpl.ExecuteTemplate(w, "devices.html", devices)
}
