package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/chiragsoni81245/net-sentinel/internal/types"
)

func SetToasts(w http.ResponseWriter, toasts *[]types.Toast) {
    toastsString := ""
    for _, toast := range *toasts {
        if toastsString == "" {
            toastsString = fmt.Sprintf("%s=%s", toast.Type, toast.Text)
        } else {
            toastsString = fmt.Sprintf("%s,%s=%s", toastsString, toast.Type, toast.Text)
        }
    }

    http.SetCookie(w, &http.Cookie{
        Name: "toasts",
        Value: toastsString,
    })
}

func ParseToasts(w http.ResponseWriter, r *http.Request) *[]types.Toast {
    var toasts []types.Toast
    toasts_cookie, err := r.Cookie("toasts")
    if err == nil {
        for _, toast := range strings.Split(toasts_cookie.Value, ",") {
            toast_split := strings.Split(toast, "=")
            if len(toast_split) < 2 {
                continue
            }
            toasts = append(toasts, types.Toast{Type: toast_split[0], Text: toast_split[1]}) 
        }
    }

    http.SetCookie(w, &http.Cookie{
        Name: "toasts",
        Value: "",
        Expires: time.Now(),
    })

    return &toasts
}
