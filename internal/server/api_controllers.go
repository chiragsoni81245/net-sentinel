package server

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/chiragsoni81245/net-sentinel/internal/types"
	"github.com/chiragsoni81245/net-sentinel/internal/utils"
)


type APIControllers struct {
    Server *types.Server
}

func (apiC *APIControllers) Login(w http.ResponseWriter, r *http.Request) {
    if strings.ToUpper(r.Method) != "POST" {
        http.Error(w, "Unsupported method", http.StatusNotFound)
        return
    }


    err := r.ParseForm()
    if err != nil {
        log.Println(err)
        http.Error(w, "Invalid data", http.StatusBadRequest)
    }

    username := r.FormValue("username")
    password := r.FormValue("password")

    var user struct {
        id int
        password string
    };
    row := apiC.Server.DB.QueryRow(`SELECT id, password FROM user WHERE username=$1`, username)
    err = row.Scan(&user)
    
    if user.password != password {
        http.Error(w, "Unauthorized user", http.StatusUnauthorized)
    }

    token, err  := utils.GenerateJWTToken(user.id, apiC.Server.Config)
    if err != nil {
        log.Println(err)
        http.Error(w, "Something went wrong", http.StatusInternalServerError)
    }

    http.SetCookie(w, &http.Cookie{
        Name: "token",
        Value: token,
        Path: "/",
        Expires: time.Now().Add(time.Duration(apiC.Server.Config.Server.TokenExpiration) * time.Hour),
        HttpOnly: true,
        Secure: true,
        SameSite: http.SameSiteStrictMode,
    })
    
    http.Redirect(w, r, "/", 302)
}

