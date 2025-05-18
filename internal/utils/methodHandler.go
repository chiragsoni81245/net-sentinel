package utils

import (
	"net/http"
	"strings"
)


func NewMethodHandler() *MethodHandler {
    return &MethodHandler{
        methodMapFunc: make(map[string]http.HandlerFunc),
    }
}

type MethodHandler struct {
    methodMapFunc map[string]http.HandlerFunc
}

func (mh *MethodHandler) Get(handler http.HandlerFunc) {
    mh.methodMapFunc["GET"] = handler
}

func (mh *MethodHandler) Post(handler http.HandlerFunc) {
    mh.methodMapFunc["POST"] = handler
}

func (mh *MethodHandler) Put(handler http.HandlerFunc) {
    mh.methodMapFunc["PUT"] = handler
}

func (mh *MethodHandler) Delete(handler http.HandlerFunc) {
    mh.methodMapFunc["DELETE"] = handler
}

func (mh *MethodHandler) Handler(w http.ResponseWriter, r *http.Request) {
    handler, ok := mh.methodMapFunc[strings.ToUpper(r.Method)] 
    if !ok {
        http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
        return
    }
    handler(w, r)
}
