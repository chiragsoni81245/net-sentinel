package server

import (
	"database/sql"
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
    if r.Context().Value("userId") != nil {
        http.Redirect(w, r, "/", 302)
        return
    }

    err := r.ParseForm()
    if err != nil {
        log.Println(err)
        utils.SendJSON(w, `{"error": "Invalid data"}`, http.StatusBadRequest)
        return
    }

    username := r.FormValue("username")
    password := r.FormValue("password")


    var user struct {
        id int
        password string
    };

    invalidCredentialsToast := types.Toast{
        Type: "error",
        Text: "Invalid Credentials",
    }
    somethingWentWrongToast := types.Toast{
        Type: "error",
        Text: "Something went wrong",
    }

    row := apiC.Server.DB.QueryRow(`SELECT id, password FROM user WHERE username=$1`, username)
    err = row.Scan(&user.id, &user.password)
    if err != nil {
        if err == sql.ErrNoRows {
            utils.SetToasts(w, &[]types.Toast{invalidCredentialsToast})
            http.Redirect(w, r, "/login", 302)
            return
        }
        log.Println(err)
        utils.SetToasts(w, &[]types.Toast{somethingWentWrongToast})
        http.Redirect(w, r, "/login", 302)
        return
    }
    
    if user.password != password {
        utils.SetToasts(w, &[]types.Toast{invalidCredentialsToast})
        http.Redirect(w, r, "/login", 302)
    }

    token, err  := utils.GenerateJWTToken(user.id, apiC.Server.Config)
    if err != nil {
        log.Println(err)
        utils.SetToasts(w, &[]types.Toast{somethingWentWrongToast})
        http.Redirect(w, r, "/login", 302)
        return
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

