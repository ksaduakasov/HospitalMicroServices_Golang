package main

import (
	"fmt"
	"github.com/golangcollege/sessions"
	"log"
	"net/http"
	"runtime/debug"
)

type Application struct {
	ErrorLog log.Logger
	InfoLog  log.Logger
	Session  *sessions.Session
}
func (a *Application) ServerError(w http.ResponseWriter, err error)  {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	a.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError)
}
func (a *Application) BadRequestError(w http.ResponseWriter, err error)  {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	a.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusBadRequest),
		http.StatusBadRequest)
}
func (a *Application) MethodNotAllowedError(w http.ResponseWriter, err error)  {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	a.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed),
		http.StatusMethodNotAllowed)
}
func (a *Application) UnauthorizedError(w http.ResponseWriter, err error)  {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	a.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusUnauthorized),
		http.StatusUnauthorized)
}
func (a *Application) ForbiddenError(w http.ResponseWriter, err error)  {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	a.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusForbidden),
		http.StatusForbidden)
}
func (a *Application) NotFoundError(w http.ResponseWriter, err error)  {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	a.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusNotFound),
		http.StatusNotFound)
}